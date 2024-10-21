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
var (
	parent string
	child string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// cobraCmd
var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Construct " + chalk.Yellow.Color("cobra") + " cmd.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Construct ` + chalk.Yellow.Color("cobra") + ` app from template.

Commands include:
	` + chalk.Magenta.Color("completion") + `
	` + chalk.Magenta.Color("identity") + `
`,

	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help ` + chalk.Yellow.Color("cobra"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// copy template
		params := copierCopyReplace(findHome()+cmdDir+"/"+"cmdTemplate.go", path+"/"+"cmd"+"/"+"cmd"+child+".go")
		params.reps = replacerCobraCmd() // automatic binding cli flags
		copierFile(params)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	cobraCmd.AddCommand(cmdCmd)

	// flags
	cmdCmd.Flags().StringVarP(&child, "child", "d", "", "New command to attach. Recommended to capitalize first letter.")
	cmdCmd.MarkFlagRequired("child")
	cmdCmd.Flags().StringVarP(&parent, "parent", "u", "root", "Parent command to attach new command to. If not asigned, attach to")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
