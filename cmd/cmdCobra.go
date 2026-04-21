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

func CobraCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("cobra", nil))
	cmd.PersistentFlags().StringVarP(&cobraFlags.user, "user", "", "DanielRivasMD", "GitHub username")
	cmd.PersistentFlags().StringVarP(&cobraFlags.author, "author", "", "Daniel Rivas", "Author name")
	cmd.PersistentFlags().StringVarP(&cobraFlags.email, "email", "", "<danielrivasmd@gmail.com>", "Author email")
	cmd.AddCommand(
		CobraAppCmd(),
		CobraCmdCmd(),
	)
	return cmd
}

func CobraAppCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("app", runCobraApp))
	cmd.Flags().BoolVarP(&cobraAppFlags.force, "force", "f", false, "Force install go dependencies")
	return cmd
}

func CobraCmdCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("cmd", runCobraCmd,
		domovoi.WithArgs(cobra.ExactArgs(1)),
	))
	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraApp(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.app"

	shouldInit := false
	_, errMod := os.Stat("go.mod")
	if errMod == nil {
		// go.mod exists
		if !cobraAppFlags.force {
			horus.CheckErr(fmt.Errorf("a Go module already exists (go.mod). Use --force to overwrite"),
				horus.WithOp(op), horus.WithMessage("existing go.mod detected"))
		} else {
			for _, f := range []string{"go.mod", "go.sum"} {
				if err := os.Remove(f); err != nil && !os.IsNotExist(err) {
					horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to remove "+f))
				}
			}
			shouldInit = true
		}
	} else if os.IsNotExist(errMod) {
		shouldInit = true
	} else {
		horus.CheckErr(errMod, horus.WithOp(op), horus.WithMessage("failed to check go.mod existence"))
	}

	horus.CheckErr(domovoi.CreateDir("cmd", rootFlags.verbose))
	repo := horus.Must(domovoi.CurrentDir())
	replaces := []moldReplace{
		Replace("XXX_REPO_XXX", repo),
		Replace("XXX_CLI_LOWERCASE_XXX", strings.ToLower(repo)),
		Replace("XXX_AUTHOR_XXX", cobraFlags.author),
		Replace("XXX_EMAIL_XXX", cobraFlags.email),
		Replace("XXX_YEAR_XXX", strconv.Itoa(time.Now().Year())),
	}

	outputs := []string{
		"main.go",
		filepath.Join("cmd", "root.go"),
		filepath.Join("cmd", "docs.json"),
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

	if shouldInit {
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+cobraFlags.user+"/"+repo))
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {
	op := "tabularasa.cobra.cmd"
	horus.CheckErr(domovoi.CreateDir("cmd", rootFlags.verbose))

	repo := horus.Must(domovoi.CurrentDir())
	cmdName := args[0]
	replaces := []moldReplace{
		Replace("XXX_CLI_LOWERCASE_XXX", strings.ToLower(repo)),
		Replace("XXX_CMD_LOWERCASE_XXX", lowerFirst(cmdName)),
		Replace("XXX_CMD_UPPERCASE_XXX", upperFirst(cmdName)),
		Replace("XXX_AUTHOR_XXX", cobraFlags.author),
		Replace("XXX_EMAIL_XXX", cobraFlags.email),
		Replace("XXX_YEAR_XXX", strconv.Itoa(time.Now().Year())),
	}

	outFile := filepath.Join("cmd", "cmd"+upperFirst(cmdName)+".go")
	const templateFile = "cmdCmd_go"
	moldForging(op, newMoldConfig(configDirs.cobra, outFile, []string{templateFile}, replaces...))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type cobraFlag struct {
	user   string
	author string
	email  string
}

type cobraAppFlag struct {
	force bool
}

var (
	cobraFlags    cobraFlag
	cobraAppFlags cobraAppFlag
)

////////////////////////////////////////////////////////////////////////////////////////////////////
