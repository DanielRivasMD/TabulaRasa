/*
Copyright © 2024 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var cobraCmd = &cobra.Command{
	Use:     "cobra",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobra,
	Example: exampleCobra,
}

var cobraAppCmd = &cobra.Command{
	Use:     "app",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobraApp,
	Example: exampleCobraApp,

	Run: runCobraApp,
}

var cobraCmdCmd = &cobra.Command{
	Use:     "cmd",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobraCmd,
	Example: exampleCobraCmd,

	Run: runCobraCmd,
}

var cobraUtilCmd = &cobra.Command{
	Use:     "util",
	Short:   "Import utility templates",
	Long:    helpCobraUtil,
	Example: exampleCobraUtil,

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: []string{"help", "example"},

	Run: runCobraUtil,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type mbomboReplace struct {
	old string
	new string
}

type mbomboForge struct {
	in      string
	out     string
	file    string
	replace []mbomboReplace
}

// pair up template and output in one slice
type filePair struct {
	file string
	out  string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(cobraCmd)
	cobraCmd.AddCommand(cobraAppCmd, cobraCmdCmd, cobraUtilCmd)

	// cobra app
	cobraAppCmd.Flags().BoolVarP(&flags.force, "force", "f", false, "Force install go dependencies")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraApp(cmd *cobra.Command, args []string) {
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose))

	if flags.repo == "" {
		// TODO: add error handling & potentially domovoi implementation
		dir, _ := os.Getwd()
		flags.repo = filepath.Base(dir)
	}

	replaces := []mbomboReplace{
		Replace("REPOSITORY", flags.repo),
		Replace("COMMAND_LOWERCASE", strings.ToLower(flags.repo)),
		Replace("COMMAND_UPPERCASE", "Root"),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}

	pairs := []filePair{
		{"main.txt", "main.go"},
		{"root.txt", filepath.Join("cmd", "root.go")},
		{"completion.txt", filepath.Join("cmd", "cmdCompletion.go")},
		{"identity.txt", filepath.Join("cmd", "cmdIdentity.go")},
		{"utilHelp.txt", filepath.Join("cmd", "utilHelp.go")},
		{"utilExample.txt", filepath.Join("cmd", "utilExample.go")},
	}

	// now a simple for‐range
	for _, p := range pairs {
		mf := NewMbomboForge(
			dirs.cobra,
			p.out,
			p.file,
			replaces...,
		)

		// TODO: better error check
		horus.CheckErr(domovoi.ExecSh(mf.Cmd()))

	}

	// Initialize Go module and tidy dependencies
	if flags.force {
		// BUG: domovoi.RemoveFile not working properly
		// domovoi.RemoveFile("go.mod", flags.verbose)
		// domovoi.RemoveFile("go.sum", flags.verbose)
		os.Remove("go.mod")
		os.Remove("go.sum")
	}

	// TODO: better error check
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+flags.user+"/"+flags.repo))
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))

	// TODO: set up copy & replace for LICENSE, as well as tab completion on the suffix pattern

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose))
	params.cmd = args[0]

	replaces := []mbomboReplace{
		Replace("COMMAND_LOWERCASE", lowerFirst(params.cmd)),
		Replace("COMMAND_UPPERCASE", upperFirst(params.cmd)),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}

	mf := NewMbomboForge(
		dirs.cobra,
		filepath.Join("cmd", "cmd"+upperFirst(params.cmd)+".go"),
		"cmd.txt",
		replaces...,
	)

	// TODO: better error check
	horus.CheckErr(domovoi.ExecSh(mf.Cmd()))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraUtil(cmd *cobra.Command, args []string) {
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose))
	params.util = args[0]

	replaces := []mbomboReplace{
		Replace("COMMAND_LOWERCASE", lowerFirst(params.util)),
		Replace("COMMAND_UPPERCASE", upperFirst(params.util)),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
	}

	var pair filePair
	switch params.util {
	case "help":
		pair = filePair{"utilHelp.txt", filepath.Join("cmd", "utilHelp.go")}
	case "example":
		pair = filePair{"utilExample.txt", filepath.Join("cmd", "utilExample.go")}
	}

	mf := NewMbomboForge(
		dirs.cobra,
		pair.out,
		pair.file,
		replaces...,
	)

	// TODO: better error check
	horus.CheckErr(domovoi.ExecSh(mf.Cmd()))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func NewMbomboForge(
	inDir, outFile, tplFile string,
	replaces ...mbomboReplace,
) mbomboForge {
	return mbomboForge{
		in:      inDir,
		out:     outFile,
		file:    tplFile,
		replace: replaces,
	}
}

func Replace(key, val string) mbomboReplace {
	return mbomboReplace{old: key, new: val}
}

func (m mbomboForge) Cmd() string {
	// build each --replace line
	var lines []string
	for _, r := range m.replace {
		lines = append(lines,
			fmt.Sprintf(`--replace %s="%s"`, r.old, r.new),
		)
	}
	replBlock := strings.Join(lines, " \\\n")

	return fmt.Sprintf(
		`mbombo forge \
--in %s \
--out %s \
--files %s \
%s`,
		m.in,
		m.out,
		m.file,
		replBlock,
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
