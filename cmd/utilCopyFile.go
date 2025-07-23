////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"os"
	"path/filepath"

	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copyFile copies a single file from params.Orig to params.Dest,
// preserving permissions and applying any replacements.
//
// It returns an error if any step fails.
func copyFile(params CopyParams) error {
	op := "copyFile"

	// open source
	srcFile, err := os.Open(params.Orig)
	if err != nil {
		return horus.NewHerror(
			op,
			"failed to open source file",
			err,
			map[string]any{"src": params.Orig},
		)
	}
	defer srcFile.Close()

	// ensure destination dir exists
	if dir := filepath.Dir(params.Dest); dir != "" {
		if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
			return horus.NewHerror(
				op,
				"failed to create destination directory",
				mkErr,
				map[string]any{"dir": dir},
			)
		}
	}

	// create destination
	destFile, err := os.Create(params.Dest)
	if err != nil {
		return horus.NewHerror(
			op,
			"failed to create destination file",
			err,
			map[string]any{"dest": params.Dest},
		)
	}
	defer destFile.Close()

	// perform copy
	if _, err = io.Copy(destFile, srcFile); err != nil {
		return horus.NewHerror(
			op,
			"failed during copy operation",
			err,
			map[string]any{"src": params.Orig, "dest": params.Dest},
		)
	}

	// preserve file mode
	if info, statErr := os.Stat(params.Orig); statErr == nil {
		_ = os.Chmod(params.Dest, info.Mode())
	}

	// apply replacements if any
	if len(params.Reps) > 0 {
		if repErr := replace(params.Dest, params.Reps); repErr != nil {
			return horus.NewHerror(
				op,
				"failed to apply replacements",
				repErr,
				map[string]any{"file": params.Dest},
			)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
