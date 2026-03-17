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

	outputs := []string{
		"main.go",
		filepath.Join("cmd", "root.go"),
		filepath.Join("cmd", "cmdCompletion.go"),
		filepath.Join("cmd", "cmdIdentity.go"),
	}

	for _, out := range outputs {
		templateFile := templateMapping(filepath.Base(out))
		if dir := filepath.Dir(out); dir != "." {
			horus.CheckErr(domovoi.CreateDir(dir, rootFlags.verbose))
		}
		moldForging(op, newMoldConfig(configDirs.cobra, out, []string{templateFile}, replaces...))
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

	outFile := filepath.Join("cmd", "cmd"+upperFirst(cmdName)+".go")
	const templateFile = "cmdCmd_go"
	moldForging(op, newMoldConfig(configDirs.cobra, outFile, []string{templateFile}, replaces...))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
