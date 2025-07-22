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
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	headerLine string
)

var justCmd = &cobra.Command{
	Use:   "just",
	Short: "Deploy just config templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy ` + chalk.Yellow.Color("just") + ` config templates over target,
including ` + chalk.Red.Color(".justfile") + ` and language‐specific configs.
`,
	Example: `
  ` + chalk.Cyan.Color("tab") + ` deploy just --lang go
  ` + chalk.Cyan.Color("tab") + ` deploy just --ver 1.0
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// fallback repoName to current dir if not provided
		if repoName == "" {
			repoName = currentDir()
		}

		// locate TabulaRasa home
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"justCmd.Run",
				"failed to find TabulaRasa home",
				err,
				nil,
			))
		}

		// ensure `.just` directory exists
		justDirPath := filepath.Join(projectPath, dotjust)
		if err := domovoi.EnsureDirExist(justDirPath, true); err != nil {
			horus.CheckErr(err)
		}

		// deploy combined .justfile
		justfileDest := filepath.Join(projectPath, "."+justfile)
		jfParams := newCopyParams(
			filepath.Join(home, justDir),
			justfileDest,
		)
		jfParams.Files = append([]string{headerLine}, lang.Selected...)
		jfParams.Reps = buildDeployReplacements(repoName, version)
		if err := concatenateFiles(jfParams, dotjust); err != nil {
			horus.CheckErr(err)
		}

		// deploy each language's config
		for _, langOpt := range lang.Selected {
			srcConf := filepath.Join(home, justDir, langOpt+dotconf)
			dstConf := filepath.Join(justDirPath, langOpt+dotconf)
			confParams := newCopyParams(srcConf, dstConf)
			confParams.Reps = buildDeployReplacements(repoName, version)
			if err := copyFile(confParams); err != nil {
				horus.CheckErr(err)
			}

			// include Python installer if deploying Python
			if langOpt == "py" {
				srcInst := filepath.Join(home, justDir, pyinstall)
				dstInst := filepath.Join(justDirPath, pyinstall)
				instParams := newCopyParams(srcInst, dstInst)
				instParams.Reps = buildDeployReplacements(repoName, version)
				if err := copyFile(instParams); err != nil {
					horus.CheckErr(err)
				}
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	deployCmd.AddCommand(justCmd)

	justCmd.Flags().StringVar(
		&headerLine,
		"header",
		"",
		"Line to prepend as header in .justfile",
	)

	justCmd.Flags().StringVarP(
		&version,
		"ver",
		"v",
		"",
		"Version string to embed in templates",
	)

	horus.CheckErr(justCmd.MarkFlagRequired("lang"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
