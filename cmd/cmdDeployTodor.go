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

var todorCmd = &cobra.Command{
	Use:   "todor",
	Short: "Deploy todor config template",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy ` + chalk.Yellow.Color("todor") + ` config template into your project.
Includes the top-level ` + chalk.Red.Color(".todor") + ` file.
`,
	Example: `
  ` + chalk.Cyan.Color("tab") + ` deploy todor
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// locate TabulaRasa home directory
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"todorCmd.Run",
				"failed to find TabulaRasa home",
				err,
				nil,
			))
		}

		// source: $TABULARASA_HOME/todorDir/todor
		src := filepath.Join(home, todorDir, todor)

		// destination: <projectPath>/.todor
		dest := filepath.Join(path, "."+todor)

		// copy template to project
		params := newCopyParams(src, dest)
		if err := copyFile(params); err != nil {
			horus.CheckErr(err)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	deployCmd.AddCommand(todorCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
