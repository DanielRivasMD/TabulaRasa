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

var cobraAppForce bool

func CobraCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("cobra", nil))
	cmd.AddCommand(
		CobraAppCmd(),
		CobraCmdCmd(),
	)
	return cmd
}

func CobraAppCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("app", runCobraApp))
	cmd.Flags().BoolVarP(&cobraAppForce, "force", "f", false, "Force install go dependencies")
	return cmd
}

func CobraCmdCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("cmd", runCobraCmd,
		domovoi.WithArgs(cobra.ExactArgs(1)),
	))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraApp(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.app"
	horus.CheckErr(domovoi.CreateDir("cmd", rootFlags.verbose))

	repo := rootFlags.repo
	if repo == "" {
		var err error
		repo, err = domovoi.CurrentDir()
		horus.CheckErr(err)
	}

	replaces := []moldReplace{
		Replace("XXX_REPO_XXX", repo),
		Replace("XXX_CMD_LOWERCASE_XXX", strings.ToLower(repo)),
		Replace("XXX_CMD_UPPERCASE_XXX", "Root"),
		Replace("XXX_AUTHOR_XXX", rootFlags.author),
		Replace("XXX_EMAIL_XXX", rootFlags.email),
		Replace("XXX_YEAR_XXX", strconv.Itoa(time.Now().Year())),
	}

	cobraMainSkeleton := []string{"GPLv3.license", "main.package", "line.new", "line.break", "line.new", "repo.import", "line.new", "line.break", "line.new", "main.func", "line.new", "line.break"}
	cobraRootSkeleton := []string{"GPLv3.license", "cmd.package", "line.new", "line.break", "line.new", "cobra_horus.import", "line.new", "line.break", "line.new", "root.var", "line.new", "line.break", "line.new", "exec.func", "line.new", "line.break", "line.new", "flags.struct", "line.new", "line.break", "line.new", "root.func", "line.new", "line.break"}
	completionSkeleton := []string{"completion.cmd"}
	identitySkeleton := []string{"identity.cmd"}
	utilHelpSkeleton := []string{"line.break", "line.new", "cmd.package", "line.new", "line.break", "line.new", "domovoi.import", "line.new", "line.break", "line.new", "help.var", "line.new", "line.break"}
	utilExampleSkeleton := []string{"line.break", "line.new", "cmd.package", "line.new", "line.break", "line.new", "domovoi.import", "line.new", "line.break", "line.new", "example.var", "line.new", "line.break"}

	pairs := []filePair{
		{cobraMainSkeleton, "main.go"},
		{cobraRootSkeleton, filepath.Join("cmd", "root.go")},
		{completionSkeleton, filepath.Join("cmd", "cmdCompletion.go")},
		{identitySkeleton, filepath.Join("cmd", "cmdIdentity.go")},
		{utilHelpSkeleton, filepath.Join("cmd", "utilHelp.go")},
		{utilExampleSkeleton, filepath.Join("cmd", "utilExample.go")},
	}

	for _, p := range pairs {
		moldForging(op, newMoldConfig(configDirs.cobra, p.out, p.files, replaces...))
	}

	if cobraAppForce {
		os.Remove("go.mod")
		os.Remove("go.sum")
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+rootFlags.user+"/"+repo))
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.cmd"
	horus.CheckErr(domovoi.CreateDir("cmd", rootFlags.verbose))

	cmdName := args[0]
	replaces := []moldReplace{
		Replace("XXX_CMD_LOWERCASE_XXX", lowerFirst(cmdName)),
		Replace("XXX_CMD_UPPERCASE_XXX", upperFirst(cmdName)),
		Replace("XXX_AUTHOR_XXX", rootFlags.author),
		Replace("XXX_EMAIL_XXX", rootFlags.email),
		Replace("XXX_YEAR_XXX", strconv.Itoa(time.Now().Year())),
	}

	cobraCmdSkeleton := []string{"GPLv3.license", "cmd.package", "line.new", "line.break", "line.new", "cobra.import", "line.new", "line.break", "line.new", "cmd.var", "line.new", "line.break", "line.new", "init.func", "line.new", "line.break", "line.new", "run.func", "line.new", "line.break"}
	outFile := filepath.Join("cmd", "cmd"+upperFirst(cmdName)+".go")
	moldForging(op, newMoldConfig(configDirs.cobra, outFile, cobraCmdSkeleton, replaces...))

	injectionHelpSkeleton := []string{"utilHelp.tmp", "line.new", "help.var", "line.new", "line.break"}
	injectionExampleSkeleton := []string{"utilExample.tmp", "line.new", "example.var", "line.new", "line.break"}

	for _, inj := range []struct {
		tmp    string
		target string
		block  []string
	}{
		{"utilHelp.tmp", "utilHelp.go", injectionHelpSkeleton},
		{"utilExample.tmp", "utilExample.go", injectionExampleSkeleton},
	} {
		tmpPath := filepath.Join(configDirs.cobra, inj.tmp)
		targetPath := filepath.Join("cmd", inj.target)
		horus.CheckErr(CopyFile(targetPath, tmpPath))
		moldForging(op, newMoldConfig(configDirs.cobra, targetPath, inj.block, replaces...))
		os.Remove(tmpPath)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
