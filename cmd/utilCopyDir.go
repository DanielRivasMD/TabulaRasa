////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"path/filepath"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copyDir recursively copies a directory tree from params.Orig to params.Dest,
// preserving permissions and applying replacements. It returns an error on failure.
func copyDir(params CopyParams) error {
	op := "copyDir"

	// // stat source to get mode
	info, err := os.Stat(params.Orig)
	if err != nil {
		return horus.NewHerror(
			op,
			"failed to stat source directory",
			err,
			map[string]any{"src": params.Orig},
		)
	}

	// create destination directory
	if mkErr := os.MkdirAll(params.Dest, info.Mode()); mkErr != nil {
		return horus.NewHerror(
			op,
			"failed to create destination directory",
			mkErr,
			map[string]any{"dest": params.Dest},
		)
	}

	// read directory entries
	dirHandle, err := os.Open(params.Orig)
	if err != nil {
		return horus.NewHerror(
			op,
			"failed to open source directory",
			err,
			map[string]any{"src": params.Orig},
		)
	}
	entries, err := dirHandle.Readdir(-1)
	dirHandle.Close()
	if err != nil {
		return horus.NewHerror(
			op,
			"failed to read source directory entries",
			err,
			map[string]any{"src": params.Orig},
		)
	}

	// recurse into subentries
	for _, entry := range entries {
		srcPath := filepath.Join(params.Orig, entry.Name())
		dstPath := filepath.Join(params.Dest, entry.Name())

		childParams := CopyParams{
			Orig:  srcPath,
			Dest:  dstPath,
			Files: params.Files,
			Reps:  params.Reps,
		}

		if entry.IsDir() {
			if err := copyDir(childParams); err != nil {
				return err
			}
		} else {
			if err := copyFile(childParams); err != nil {
				return err
			}
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
