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
	Short:   "Construct cobra apps, cmds & import utilities",
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
	Short:   "Build cobra cmds from templates",
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
	in       string
	out      string
	files    []string
	replaces []mbomboReplace
}

// pair up template and output in one slice
type filePair struct {
	files []string
	out   string
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
	op := "tabularasa.cobra.app"
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose), horus.WithOp(op))

	var err error
	if flags.repo == "" {
		flags.repo, err = domovoi.CurrentDir()
		horus.CheckErr(err, horus.WithOp(op))
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
		{[]string{"main.txt"}, "main.go"},
		{[]string{"root.txt"}, filepath.Join("cmd", "root.go")},
		{[]string{"cmdCompletion.txt"}, filepath.Join("cmd", "cmdCompletion.go")},
		{[]string{"cmdIdentity.txt"}, filepath.Join("cmd", "cmdIdentity.go")},
		{[]string{"line.break", "line.new", "cmd.package", "line.new", "line.break", "line.new", "domovoi.import", "line.new", "line.break", "line.new", "help.var", "line.new", "line.break"}, filepath.Join("cmd", "utilHelp.go")},
		{[]string{"line.break", "line.new", "cmd.package", "line.new", "line.break", "line.new", "domovoi.import", "line.new", "line.break", "line.new", "example.var", "line.new", "line.break"}, filepath.Join("cmd", "utilExample.go")},
	}

	// now a simple for‐range
	for _, p := range pairs {
		mf := NewMbomboForge(
			dirs.cobra,
			p.out,
			p.files,
			replaces...,
		)

		horus.CheckErr(
			domovoi.ExecSh(mf.Cmd()),
			horus.WithOp(op),
			horus.WithCategory("shell_command"),
			horus.WithMessage("Failed to execute mbombo forge command"),
			horus.WithDetails(map[string]any{
				"command": mf.Cmd(),
			}),
		)

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
	op := "tabularasa.cobra.cmd"
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose), horus.WithOp(op))
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
		[]string{"GPLv3.license", "cmd.package", "line.new", "line.break", "line.new", "cobra.import", "line.new", "line.break", "line.new", "cmd.var", "line.new", "line.break", "line.new", "init.func", "line.new", "line.break", "line.new", "run.func", "line.new", "line.break"},
		replaces...,
	)

	horus.CheckErr(
		domovoi.ExecSh(mf.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": mf.Cmd(),
		}),
	)

	horus.CheckErr(CopyFile(filepath.Join("cmd", "utilHelp.go"), filepath.Join(dirs.cobra, "utilHelp.tmp")))
	mh := NewMbomboForge(
		dirs.cobra,
		filepath.Join("cmd", "utilHelp.go"),
		[]string{"utilHelp.tmp", "line.new", "help.var", "line.new", "line.break"},
		replaces...,
	)
	os.Remove(filepath.Join(dirs.cobra, "utilHelp.tmp"))

	horus.CheckErr(
		domovoi.ExecSh(mh.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": mh.Cmd(),
		}),
	)

	horus.CheckErr(CopyFile(filepath.Join("cmd", "utilExample.go"), filepath.Join(dirs.cobra, "utilExample.tmp")))
	me := NewMbomboForge(
		dirs.cobra,
		filepath.Join("cmd", "utilExample.go"),
		[]string{"utilExample.tmp", "line.new", "example.var", "line.new", "line.break"},
		replaces...,
	)
	os.Remove(filepath.Join(dirs.cobra, "utilExample.tmp"))

	horus.CheckErr(
		domovoi.ExecSh(me.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": me.Cmd(),
		}),
	)

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraUtil(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.util"
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose), horus.WithOp(op))
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
		pair = filePair{[]string{"utilHelp.txt"}, filepath.Join("cmd", "utilHelp.go")}
	case "example":
		pair = filePair{[]string{"utilExample.txt"}, filepath.Join("cmd", "utilExample.go")}
	}

	mf := NewMbomboForge(
		dirs.cobra,
		pair.out,
		pair.files,
		replaces...,
	)

	horus.CheckErr(
		domovoi.ExecSh(mf.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": mf.Cmd(),
		}),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func NewMbomboForge(
	inDir, outFile string,
	tplFiles []string,
	replaces ...mbomboReplace,
) mbomboForge {
	return mbomboForge{
		in:       inDir,
		out:      outFile,
		files:    tplFiles,
		replaces: replaces,
	}
}

func Replace(key, val string) mbomboReplace {
	return mbomboReplace{old: key, new: val}
}

func (m mbomboForge) Cmd() string {
	var files []string
	for _, f := range m.files {
		files = append(files, fmt.Sprintf(`--files %s`, f))
	}
	fileBlock := strings.Join(files, " \\\n")

	var replaces []string
	for _, r := range m.replaces {
		replaces = append(replaces, fmt.Sprintf(`--replace %s="%s"`, r.old, r.new))
	}
	replaceBlock := strings.Join(replaces, " \\\n")

	return fmt.Sprintf(
		`mbombo forge \
--in %s \
--out %s \
%s \
%s`,
		m.in,
		m.out,
		fileBlock,
		replaceBlock,
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
