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
	// util is the name of the utility template to import (capitalized).
	util string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "Import utility templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy a utility from predefined templates
`,
	Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` ` + chalk.Green.Color("util") +
		` --util ExampleUtil
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"cmdCobraUtil.Run",
				"failed to find TabulaRasa home",
				err,
				nil,
			))
		}

		// source: $TABULARASA_HOME/utilDir/<util>.go
		src := filepath.Join(home, utilDir, util+".go")

		// destination: <projectPath>/cmd/<util>.go
		dest := filepath.Join(projectPath, "cmd", util+".go")

		params := newCopyParams(src, dest)

		// re-use your cmd replacements to fill REPOSITORY, AUTHOR, etc.
		params.Reps = buildCmdReplacements(
			repoName, authorName, authorEmail,
			util, "", "",
		)

		copyFile(params)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobraCmd.AddCommand(utilCmd)

	// allow overriding project root
	utilCmd.Flags().StringVar(
		&projectPath, "path", ".", "Base path of your Go project",
	)

	// select which util to import
	utilCmd.Flags().StringVarP(
		&util, "util", "u", "", "Utility template name (capitalize)",
	)

	// require --util
	if err := utilCmd.MarkFlagRequired("util"); err != nil {
		horus.CheckErr(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
