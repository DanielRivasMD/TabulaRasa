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

func init() {
	rootCmd.AddCommand(cobraCmd)
	cobraCmd.AddCommand(cobraAppCmd, cobraCmdCmd, cobraUtilCmd)

	// cobra app
	cobraAppCmd.Flags().BoolVarP(&flags.force, "force", "f", false, "Force install go dependencies")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var skel = skeletons{
	gplv3License:     "GPLv3.license",
	mainPackage:      "main.package",
	cmdPackage:       "cmd.package",
	importRepo:       "repo.import",
	importCobra:      "cobra.import",
	importCobraHorus: "cobra_horus.import",
	importDomovoi:    "domovoi.import",
	initFunc:         "init.func",
	runFunc:          "run.func",
	mainFunc:         "main.func",
	rootFunc:         "root.func",
	execFunc:         "exec.func",
	flagsStruct:      "flags.struct",
	cmdVar:           "cmd.var",
	rootVar:          "root.var",
	helpVar:          "help.var",
	exampleVar:       "example.var",
	lineNew:          "line.new",
	lineBreak:        "line.break",
	completionCmd:    "completion.cmd",
	identityCmd:      "identity.cmd",
	tmpHelp:          "utilHelp.tmp",
	tmpExample:       "utilExample.tmp",
}

type skeletons struct {
	gplv3License     string
	mainPackage      string
	cmdPackage       string
	importRepo       string
	importCobra      string
	importCobraHorus string
	importDomovoi    string
	initFunc         string
	runFunc          string
	mainFunc         string
	rootFunc         string
	execFunc         string
	flagsStruct      string
	cmdVar           string
	rootVar          string
	helpVar          string
	exampleVar       string
	lineNew          string
	lineBreak        string
	completionCmd    string
	identityCmd      string
	tmpHelp          string
	tmpExample       string
}

var cobraMainSkeleton = []string{
	skel.gplv3License,
	skel.mainPackage,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.importRepo,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.mainFunc,
	skel.lineNew, skel.lineBreak,
}

var cobraRootSkeleton = []string{
	skel.gplv3License,
	skel.cmdPackage,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.importCobraHorus,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.rootVar,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.execFunc,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.flagsStruct,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.rootFunc,
	skel.lineNew, skel.lineBreak,
}

var cobraCmdSkeleton = []string{
	skel.gplv3License,
	skel.cmdPackage,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.importCobra,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.cmdVar,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.initFunc,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.runFunc,
	skel.lineNew, skel.lineBreak,
}

var injectionHelpSkeleton = []string{
	skel.tmpHelp,
	skel.lineNew,
	skel.helpVar,
	skel.lineNew,
	skel.lineBreak,
}

var injectionExampleSkeleton = []string{
	skel.tmpExample,
	skel.lineNew,
	skel.exampleVar,
	skel.lineNew,
	skel.lineBreak,
}

var utilHelpSkeleton = []string{
	skel.lineBreak, skel.lineNew,
	skel.cmdPackage,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.importDomovoi,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.helpVar,
	skel.lineNew, skel.lineBreak,
}

var utilExampleSkeleton = []string{
	skel.lineBreak, skel.lineNew,
	skel.cmdPackage,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.importDomovoi,
	skel.lineNew, skel.lineBreak, skel.lineNew,
	skel.exampleVar,
	skel.lineNew, skel.lineBreak,
}

var completionSkeleton = []string{
	skel.completionCmd,
}

var identitySkeleton = []string{
	skel.identityCmd,
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

	replaces := cobraAppReplacements()

	pairs := []filePair{
		{cobraMainSkeleton, "main.go"},
		{cobraRootSkeleton, filepath.Join("cmd", "root.go")},
		{completionSkeleton, filepath.Join("cmd", "cmdCompletion.go")},
		{identitySkeleton, filepath.Join("cmd", "cmdIdentity.go")},
		{utilHelpSkeleton, filepath.Join("cmd", "utilHelp.go")},
		{utilExampleSkeleton, filepath.Join("cmd", "utilExample.go")},
	}

	// now a simple for‐range
	for _, p := range pairs {
		mf := newMbomboConfig(
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

func cobraAppReplacements() []mbomboReplace {
	return []mbomboReplace{
		Replace("REPOSITORY", flags.repo),
		Replace("COMMAND_LOWERCASE", strings.ToLower(flags.repo)),
		Replace("COMMAND_UPPERCASE", "Root"),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.cmd"
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose), horus.WithOp(op))
	params.cmd = args[0]

	replaces := cobraCmdReplacements()

	mf := newMbomboConfig(
		dirs.cobra,
		filepath.Join("cmd", "cmd"+upperFirst(params.cmd)+".go"),
		cobraCmdSkeleton,
		replaces...,
	)

	mbomboForging(op, mf)

	injections := []struct {
		srcTmp string
		target string
		block  []string
	}{
		{
			srcTmp: skel.tmpHelp,
			target: "utilHelp.go",
			block:  injectionHelpSkeleton,
		},
		{
			srcTmp: skel.tmpExample,
			target: "utilExample.go",
			block:  injectionExampleSkeleton,
		},
	}

	for _, inj := range injections {
		tmpPath := filepath.Join(dirs.cobra, inj.srcTmp)
		targetPath := filepath.Join("cmd", inj.target)

		horus.CheckErr(CopyFile(targetPath, tmpPath))

		m := newMbomboConfig(
			dirs.cobra,
			targetPath,
			inj.block,
			replaces...,
		)

		mbomboForging(op, m)
		os.Remove(tmpPath)
	}
}

func cobraCmdReplacements() []mbomboReplace {
	return []mbomboReplace{
		Replace("COMMAND_LOWERCASE", lowerFirst(params.cmd)),
		Replace("COMMAND_UPPERCASE", upperFirst(params.cmd)),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraUtil(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.util"
	horus.CheckErr(domovoi.CreateDir("cmd", flags.verbose), horus.WithOp(op))
	params.util = args[0]

	replaces := cobraUtilReplacements()

	var pair filePair
	switch params.util {
	case "help":
		pair = filePair{utilHelpSkeleton, filepath.Join("cmd", "utilHelp.go")}
	case "example":
		pair = filePair{utilExampleSkeleton, filepath.Join("cmd", "utilExample.go")}
	default:
		horus.CheckErr(fmt.Errorf("unknown util type: %s", params.util), horus.WithOp(op))
		return
	}

	mf := newMbomboConfig(
		dirs.cobra,
		pair.out,
		pair.files,
		replaces...,
	)

	mbomboForging(op, mf)
}

func cobraUtilReplacements() []mbomboReplace {
	return []mbomboReplace{
		Replace("COMMAND_LOWERCASE", lowerFirst(params.util)),
		Replace("COMMAND_UPPERCASE", upperFirst(params.util)),
		Replace("AUTHOR", flags.author),
		Replace("EMAIL", flags.email),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
