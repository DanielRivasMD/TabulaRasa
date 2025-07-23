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
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type LangType struct {
	validValues []string
	Selected    []string
}

var validOptions = []string{"go", "golib", "jl", "py", "rs", "rslib"}

func (f *LangType) String() string {
	if len(f.Selected) > 0 {
		return f.Selected[0]
	}
	return ""
}

func (f *LangType) Set(value string) error {
	for _, v := range f.validValues {
		if value == v {
			f.Selected = append(f.Selected, value)
			return nil
		}
	}
	return fmt.Errorf("invalid value '%s', allowed: %s", value, joinValues(f.validValues))
}

func (f *LangType) Type() string {
	return "LangType"
}

func joinValues(values []string) string {
	return strings.Join(values, ", ")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	lang        = &LangType{validValues: validOptions}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy config templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy selected config templates into your project. Valid values:
  ` + chalk.Yellow.Color("just") + `    - build system files
  ` + chalk.Yellow.Color("readme") + `  - README scaffold
  ` + chalk.Yellow.Color("todor") + `   - task-tracker starter
`,

	Example: `
  ` + chalk.Cyan.Color("tab") + ` deploy --lang just --repo myapp --version v0.1.0
  ` + chalk.Cyan.Color("tab") + ` deploy --lang readme --repo myapp
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// require at least one template
		if len(lang.Selected) == 0 {
			horus.CheckErr(horus.NewHerror(
				"deployCmd.Run",
				"no templates specified",
				errors.New("use --lang to select at least one"),
				map[string]any{"allowed": validOptions},
			))
		}

		// locate TabulaRasa home
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"deployCmd.Run",
				"failed to locate TabulaRasa home",
				err,
				nil,
			))
		}

		// apply each selected template
		for _, tmpl := range lang.Selected {
			switch tmpl {
			case "just":
				src := filepath.Join(home, justDir)
				dest := path
				params := newCopyParams(src, dest)
				params.Reps = buildDeployReplacements(repo)
				if err := copyDir(params); err != nil {
					horus.CheckErr(err)
				}

			case "readme":
				// single-file deploy
				srcFile := filepath.Join(home, readmeDir, "README.md")
				destFile := filepath.Join(path, "README.md")
				reps, err := buildReadmeReplacements(
					tmpl,        // langTag — unused here but required
					description, // overview text
					repo,
					user,
					author,
					license,
					path,
				)
				if err != nil {
					horus.CheckErr(err)
				}
				copyParams := newCopyParams(srcFile, destFile)
				copyParams.Reps = reps
				if err := copyFile(copyParams); err != nil {
					horus.CheckErr(err)
				}

			case "todor":
				src := filepath.Join(home, todorDir)
				dest := path
				params := newCopyParams(src, dest)
				// if you need replacements for todor, add them here
				if err := copyDir(params); err != nil {
					horus.CheckErr(err)
				}

			default:
				horus.CheckErr(horus.NewHerror(
					"deployCmd.Run",
					fmt.Sprintf("unknown template '%s'", tmpl),
					nil,
					map[string]any{"allowed": validOptions},
				))
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.PersistentFlags().VarP(lang, "lang", "l", "Templates to deploy (allowed: "+joinValues(validOptions)+")")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
