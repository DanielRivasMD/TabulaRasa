////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"io"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy file
func copyFile(params paramsCR) {
	// clean prior copying
	if fileExist(params.dest) {
		os.Remove(params.dest)
	}

	// handle origin
	origFile, ε := os.Open(params.orig)
	if ε != nil {
		log.Fatal(ε)
	}
	defer origFile.Close()

	// handle destiny
	destFile, ε := os.Create(params.dest)
	if ε != nil {
		log.Fatal(ε)
	}
	defer destFile.Close()

	// copy file
	_, ε = io.Copy(destFile, origFile)
	if ε == nil {
		origInfo, ε := os.Stat(params.orig)
		if ε != nil {
			ε = os.Chmod(params.dest, origInfo.Mode())
		}
	}
	// replace
	if len(params.reps) > 0 {
		replace(params.dest, params.reps)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
