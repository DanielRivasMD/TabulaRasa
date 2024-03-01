////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy & replace dir
func copyDir(params paramsCR) {
	// clean prior copying
	if fileExist(params.dest) {
		os.Remove(params.dest)
	}

	// original properties
	origInfo, ε := os.Stat(params.orig)
	if ε != nil {
		log.Fatal(ε)
	}

	// create destiny dir
	ε = os.MkdirAll(params.dest, origInfo.Mode())
	if ε != nil {
		log.Fatal(ε)
	}

	// origin files
	dir, _ := os.Open(params.orig)
	objs, ε := dir.Readdir(-1)

	// iterate origin
	for _, obj := range objs {

		// pointers
		params.orig = params.orig + "/" + obj.Name()
		params.dest = params.dest + "/" + obj.Name()

		if obj.IsDir() {
			// create dirs recursive
			copyDir(params)
		} else {
			// copy
			copyFile(params)
			// replace
			replace(params.dest, params.reps)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
