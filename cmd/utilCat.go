////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// concatenateFiles merges multiple source files (with a given suffix) into one destination.
// If the destination exists, it’s removed first. After merging, any replacements are applied.
func concatenateFiles(params CopyParams, suffix string) error {
	// check if destination exists (verbose logging)
	exists, err := domovoi.FileExist(params.Dest, nil, verbose)
	if err != nil {
		return horus.NewHerror(
			"concatenateFiles",
			"failed to check destination existence",
			err,
			map[string]any{"dest": params.Dest},
		)
	}

	// remove existing destination
	if exists {
		if rmErr := os.Remove(params.Dest); rmErr != nil {
			return horus.NewHerror(
				"concatenateFiles",
				"failed to remove existing destination",
				rmErr,
				map[string]any{"dest": params.Dest},
			)
		}
	}

	// open destination for appending
	destFile, err := os.OpenFile(params.Dest, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return horus.NewHerror(
			"concatenateFiles",
			"failed to open destination file",
			err,
			map[string]any{"dest": params.Dest},
		)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)

	// iterate and append each source
	for _, name := range params.Files {
		srcPath := filepath.Join(params.Orig, name+suffix)
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return horus.NewHerror(
				"concatenateFiles",
				"failed to open source file",
				err,
				map[string]any{"src": srcPath},
			)
		}

		scanner := bufio.NewScanner(srcFile)
		for scanner.Scan() {
			line := scanner.Text() + "\n"
			if _, wErr := writer.WriteString(line); wErr != nil {
				srcFile.Close()
				return horus.NewHerror(
					"concatenateFiles",
					"failed to write to destination",
					wErr,
					map[string]any{"dest": params.Dest},
				)
			}
		}
		if scanErr := scanner.Err(); scanErr != nil {
			srcFile.Close()
			return horus.NewHerror(
				"concatenateFiles",
				"error scanning source file",
				scanErr,
				map[string]any{"src": srcPath},
			)
		}
		srcFile.Close()
	}

	// flush buffer
	if err := writer.Flush(); err != nil {
		return horus.NewHerror(
			"concatenateFiles",
			"failed to flush writer",
			err,
			nil,
		)
	}

	// apply replacements if provided
	if len(params.Reps) > 0 {
		if repErr := replace(params.Dest, params.Reps); repErr != nil {
			return horus.NewHerror(
				"concatenateFiles",
				"failed to apply replacements",
				repErr,
				map[string]any{"file": params.Dest},
			)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
