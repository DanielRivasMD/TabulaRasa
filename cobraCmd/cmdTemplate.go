/*
Copyright © YEAR AUTHOR EMAIL

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

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var CHILDCmd = &cobra.Command{
	Use:   "CHILD",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("TOOL") + ` help ` + chalk.Yellow.Color("ROOT") + chalk.Yellow.Color("CHILD"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	PARENTCmd.AddCommand(CHILDCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
