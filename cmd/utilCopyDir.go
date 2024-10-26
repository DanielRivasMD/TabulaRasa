////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy & replace dir
func copyDir(π paramsCopyReplace) {
	// clean prior copying
	if fileExist(π.dest) {
		os.Remove(π.dest)
	}

	// original properties
	origInfo, ε := os.Stat(π.orig)
	if ε != nil {
		log.Fatal(ε)
	}

	// create destiny dir
	ε = os.MkdirAll(π.dest, origInfo.Mode())
	if ε != nil {
		log.Fatal(ε)
	}

	// origin files
	dir, _ := os.Open(π.orig)
	objs, ε := dir.Readdir(-1)

	orig := π.orig
	dest := π.dest

	// iterate origin
	for _, obj := range objs {

		// pointers
		π.orig = orig + "/" + obj.Name()
		π.dest = dest + "/" + obj.Name()

		if obj.IsDir() {
			// create dirs recursive
			copyDir(π)
		} else {
			// copy & replace
			copyFile(π)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
