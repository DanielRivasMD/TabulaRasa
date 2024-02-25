////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/labstack/gommon/color"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute shell command
func deployCmd(path, repo, author_email string) {

	// declare
	tool := strings.ToLower(repo)

	// buffers
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// copy
	copyShell := findHome() + gobin + "/" + "cobra" + "/" + "copy.sh"
	copyCmd := exec.Command(copyShell, path)
	copyCmd.Stdout = &stdout
	copyCmd.Stderr = &stderr
	_ = copyCmd.Run()

	// replace
	replaceShell := findHome() + gobin + "/" + "cobra" + "/" + "replace.sh"
	replaceCmd := exec.Command(replaceShell, repo, tool, author_email)
	replaceCmd.Stdout = &stdout
	replaceCmd.Stderr = &stderr
	_ = replaceCmd.Run()

	// stdout
	if stdout.String() != "" {
		color.Println(color.Cyan(stdout.String(), color.B))
	}

	// stderr
	if stderr.String() != "" {
		color.Println(color.Red(stderr.String(), color.B))
	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////
