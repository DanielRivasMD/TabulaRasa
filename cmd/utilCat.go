////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: add suffix argument
// copy file
func catFiles(π paramsCopyReplace) {
	// clean prior copying
	if fileExist(π.dest) {
		os.Remove(π.dest)
	}

	// open writer
	fwrite, ε := os.OpenFile(π.dest, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fwrite.Close()

	for _, file := range π.files {
		// open reader
		fread, ε := os.Open(π.orig + "/" + file + dotjust)
		if ε != nil {
			log.Fatal(ε)
		}
		defer fread.Close()

		// declare ϖ
		ϖ := bufio.NewWriter(fwrite)

		// read file
		scanner := bufio.NewScanner(fread)

		// scan file
		for scanner.Scan() {
			// preallocate
			toPrint := scanner.Text() + "\n"
			// write
			_, ε = ϖ.WriteString(toPrint)
			if ε != nil {
				log.Fatal(ε)
			}
		}

		if ε := scanner.Err(); ε != nil {
			log.Fatal(ε)
		}

		// flush writer
		ϖ.Flush()
	}
	// replace
	if len(π.reps) > 0 {
		replace(π.dest, π.reps)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
