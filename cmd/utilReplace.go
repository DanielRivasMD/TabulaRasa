////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// replace opens the file at target, applies all token replacements,
// and writes the result back to the same path. Returns an error on failure.
func replace(target string, reps []rep) error {
	// open source for reading
	srcFile, err := os.Open(target)
	if err != nil {
		return horus.NewHerror(
			"replace",
			"failed to open file for reading",
			err,
			map[string]any{"file": target},
		)
	}
	defer srcFile.Close()

	// create temp file
	tmpPath := target + ".tmp"
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	if err != nil {
		return horus.NewHerror(
			"replace",
			"failed to create temp file",
			err,
			map[string]any{"temp": tmpPath},
		)
	}

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(tmpFile)

	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range reps {
			line = strings.ReplaceAll(line, r.old, r.new)
		}
		if _, wErr := writer.WriteString(line + "\n"); wErr != nil {
			tmpFile.Close()
			return horus.NewHerror(
				"replace",
				"failed writing to temp file",
				wErr,
				map[string]any{"temp": tmpPath},
			)
		}
	}

	if scanErr := scanner.Err(); scanErr != nil {
		tmpFile.Close()
		return horus.NewHerror(
			"replace",
			"error scanning original file",
			scanErr,
			map[string]any{"file": target},
		)
	}

	// flush & close tmp
	if err := writer.Flush(); err != nil {
		tmpFile.Close()
		return horus.NewHerror(
			"replace",
			"failed flushing temp file",
			err,
			map[string]any{"temp": tmpPath},
		)
	}
	tmpFile.Close()

	// atomically replace the original
	if err := os.Rename(tmpPath, target); err != nil {
		return horus.NewHerror(
			"replace",
			"failed renaming temp to target",
			err,
			map[string]any{"temp": tmpPath, "file": target},
		)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
