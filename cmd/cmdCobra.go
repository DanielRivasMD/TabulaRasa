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

	PreRun: replacePreRun,
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

	// cobra cmd
	cobraCmdCmd.Flags().StringVarP(&flags.cmd, "cmd", "", "", "Name of the new cobra sub-command")
	horus.CheckErr(cobraCmdCmd.MarkFlagRequired("cmd"))

	// cobra util
	cobraUtilCmd.Flags().StringVarP(&flags.util, "util", "", "", "Utility template name (capitalize)")
	horus.CheckErr(cobraUtilCmd.MarkFlagRequired("util"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func replacePreRun(cmd *cobra.Command, args []string) {
	// format args
	flags.cmdLower = lowerFirst(flags.cmd)
	flags.cmdUpper = upperFirst(flags.cmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraApp(cmd *cobra.Command, args []string) {
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
		domovoi.RemoveFile("go.sum", flags.verbose)
		domovoi.RemoveFile("go.mod", flags.verbose)
	}

	// TODO: better error check
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+flags.user+"/"+flags.repo))
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))

	// TODO: set up copy & replace for LICENSE, as well as tab completion on the suffix pattern

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {

	// format args
	cmdLower := lowerFirst(command)
	cmdUpper := upperFirst(command)

	// TODO: finish function to abbreviate `mbombo` call
	outCobraCmd := "cmd" + "/" + "cmd" + cmdUpper + ".go"
	execSkeletonCobraCmd := "cmd.txt"

	cmdCobraCmd := fmt.Sprintf(`
		mbombo forge \
		--in %s \
		--out %s \
		--files %s \
		--replace COMMAND_LOWERCASE="%s" \
		--replace COMMAND_UPPERCASE="%s" \
		--replace AUTHOR="%s" \
		--replace EMAIL="%s" \
		--replace YEAR="%s"
	`, dirs.cobra, outCobraCmd, execSkeletonCobraCmd,
		cmdLower, cmdUpper, flags.author, flags.email, strconv.Itoa(time.Now().Year()))

	horus.CheckErr(domovoi.ExecSh(cmdCobraCmd))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraUtil(cmd *cobra.Command, args []string) {

	// build src & dest paths
	// src := filepath.Join(home, utilDir, util+".go")
	// dest := filepath.Join(path, "cmd", util+".go")

}

////////////////////////////////////////////////////////////////////////////////////////////////////
