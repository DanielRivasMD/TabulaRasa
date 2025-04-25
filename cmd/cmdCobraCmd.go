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
	parent      string
	child       string
	root_parent string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// cmdCmd
var cmdCmd = &cobra.Command{
    Use:   "cmd",
    Aliases: []string{"c"},
    Short: "Construct cobra commands",
    Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Construct ` + chalk.Yellow.Color("cobra") + ` commands from predefined templates
`,

    Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` ` + chalk.Green.Color("cmd") + ` --` + chalk.Blue.Color("child") + ` ExampleCmd --` + chalk.Blue.Color("parent") + ` RootCmd
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// define name
		if parent != "root" {
			root_parent = parent
		}

		// copy template
		params := copyCopyReplace(findHome()+cmdDir+"/"+"cmdTemplate.go", path+"/"+"cmd"+"/"+"cmd"+root_parent+child+".go")
		params.reps = replaceCobraCmd() // automatic binding cli flags
		copyFile(params)
	},

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	cobraCmd.AddCommand(cmdCmd)

	// flags
	cmdCmd.Flags().StringVarP(&child, "child", "C", "", "New command to attach. Recommended to capitalize first letter.")
	cmdCmd.Flags().StringVarP(&parent, "parent", "U", "root", "Parent command to attach new command to. If not asigned, attach to")

	cmdCmd.MarkFlagRequired("child")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
