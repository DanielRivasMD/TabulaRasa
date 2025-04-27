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

import (
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// appCmd
var appCmd = &cobra.Command{
	Use:     "app",
	Aliases: []string{"a"},
	Short:   "Construct cobra apps from templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Construct ` + chalk.Yellow.Color("cobra") + ` apps from predefined templates
`,

	Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` ` + chalk.Green.Color("app") + ` --` + chalk.Blue.Color("path") + ` $(pwd) --` + chalk.Blue.Color("repo") + ` Tabularasa
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// copy templates
		params := copyCopyReplace(findHome()+cobraDir, path)
		params.reps = replaceCobraApp() // automatic binding cli flags
		copyDir(params)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	cobraCmd.AddCommand(appCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
