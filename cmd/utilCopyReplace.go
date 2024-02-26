////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy & replace dir
func dirCopyReplace(orig, dest string, reps []replacement) {

	// original properties
	origInfo, ε := os.Stat(orig)
	if ε != nil {
		log.Fatal(ε)
	}

	// create destiny dir
	ε = os.MkdirAll(dest, origInfo.Mode())
	if ε != nil {
		log.Fatal(ε)
	}

	// origin files
	dir, _ := os.Open(orig)
	objs, ε := dir.Readdir(-1)

	// iterate origin
	for _, obj := range objs {

		// pointers
		origPointer := orig + "/" + obj.Name()
		destPointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create dirs recursive
			dirCopyReplace(origPointer, destPointer, reps)
		} else {
			// copy
			copyFile(origPointer, destPointer)
			// replace
			replace(destPointer, reps)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
