////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy file
func copyFile(orig, dest string) {
	// handle original
	origFile, ε := os.Open(orig)
	if ε != nil {
		log.Fatal(ε)
	}
	defer origFile.Close()

	// handle destiny
	destFile, ε := os.Create(dest)
	if ε != nil {
		log.Fatal(ε)
	}
	defer destFile.Close()

	// copy
	_, ε = io.Copy(destFile, origFile)
	if ε == nil {
		origInfo, ε := os.Stat(orig)
		if ε != nil {
			ε = os.Chmod(dest, origInfo.Mode())
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy dir
func copyDir(orig, dest string) {

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

	// original files
	dir, _ := os.Open(orig)
	objs, ε := dir.Readdir(-1)

	// iterate originals
	for _, obj := range objs {

		// pointers
		origPointer := orig + "/" + obj.Name()
		destPointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create dirs recursive
			copyDir(origPointer, destPointer)
		} else {
			// copy
			copyFile(origPointer, destPointer)
			// replace
			replace(destPointer)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// replace
func replace(fullpath string) {
	// open reader
	fread, ε := os.Open(fullpath)
	if ε != nil {
		log.Fatal(ε)
	}
	defer fread.Close()

	// open writer
	fwrite, ε := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, 0666)
	if ε != nil {
		panic(ε)
	}
	defer fwrite.Close()

	// declare writer
	ϖ := bufio.NewWriter(fwrite)

	// read file
	scanner := bufio.NewScanner(fread)

	// scan
	for scanner.Scan() {

		// preallocate
		toPrint := scanner.Text()

		// replace
		toPrint = strings.Replace(toPrint, "YEAR", strconv.Itoa(time.Now().Year()), -1)
		toPrint = strings.Replace(toPrint, "REPOSITORY", repo, -1)
		toPrint = strings.Replace(toPrint, "TOOL", strings.ToLower(repo), -1)
		toPrint = strings.Replace(toPrint, "AUTHOR_EMAIL", author_email, -1)

		// write
		_, ε = ϖ.WriteString(toPrint + "\n")
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

////////////////////////////////////////////////////////////////////////////////////////////////////
