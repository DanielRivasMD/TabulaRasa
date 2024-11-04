////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// replace values
func replace(target string, reps []rep) {
	// open reader
	fread, ε := os.Open(target)
	checkErr(ε)
	defer fread.Close()

	// open writer
	fwrite, ε := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(ε)
	defer fwrite.Close()

	// declare writer
	ϖ := bufio.NewWriter(fwrite)

	// read file
	scanner := bufio.NewScanner(fread)

	// scan file
	for scanner.Scan() {
		// preallocate
		toPrint := scanner.Text()

		// iterate replacements
		for _, rep := range reps {
			// replace
			toPrint = strings.Replace(toPrint, rep.old, rep.new, -1)
		}

		// format
		toPrint = toPrint + "\n"

		// write
		_, ε = ϖ.WriteString(toPrint)
		checkErr(ε)
	}

	ε = scanner.Err()
	checkErr(ε)

	// flush writer
	ϖ.Flush()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
