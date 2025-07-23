/*
Copyright © 2025 Daniel Rivas <danielrivasmd@gmail.com>

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

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Deploy README.md template",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy ` + chalk.Yellow.Color("readme") + ` config template into your project,
splicing together overview, install/dev guides, usage and FAQ snippets.
`,
	Example: `
  ` + chalk.Cyan.Color("tab") + ` deploy readme
  ` + chalk.Cyan.Color("tab") + ` deploy readme --description "Awesome project" --license MIT
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// fallback to current directory as repo name
		var err error
		if repo == "" {
			repo, err = domovoi.CurrentDir()
			horus.CheckErr(err)
		}

		// locate TabulaRasa home
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"readmeCmd.Run",
				"unable to find TabulaRasa home",
				err,
				nil,
			))
		}

		// prepare params: templates live under $HOME/<readmeDir>, output is project/README.md
		srcDir := filepath.Join(home, readmeDir)
		destFile := filepath.Join(path, readme)

		params := newCopyParams(srcDir, destFile)

		// assemble the list of template fragments
		params.Files = []string{
			overview,
			filepath.Join("02" + lang.Selected[0] + "_install.md"),
			usage,
			filepath.Join("04" + lang.Selected[0] + "_dev.md"),
			faq,
		}

		// generate replacements, handling detection & errors
		reps, repErr := buildReadmeReplacements(
			lang.Selected[0],
			description,
			repo,
			user,
			author,
			license,
			path,
		)
		if repErr != nil {
			horus.CheckErr(repErr)
		}
		params.Reps = reps

		// concatenate all fragments into README.md
		if err := concatenateFiles(params, ""); err != nil {
			horus.CheckErr(err)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	deployCmd.AddCommand(readmeCmd)

	readmeCmd.Flags().StringVarP(&description, "description", "d", "", "Project overview text")
	readmeCmd.Flags().StringVarP(&license, "license", "l", "", "License to appear in README")

	_ = readmeCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
