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
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	parent     string
	child      string
	rootParent string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Construct cobra commands from templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Construct ` + chalk.Yellow.Color("cobra") + ` commands from predefined templates
`,
	Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` ` + chalk.Green.Color("cmd") +
		` --child ExampleCmd --parent RootCmd
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		if parent != "root" {
			rootParent = parent
		}

		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"cmdCobraCmd.Run",
				"failed to find TabulaRasa home",
				err,
				nil,
			))
		}

		// build src & dest paths
		src := filepath.Join(home, cmdDir, "cmdTemplate.go")
		fileName := fmt.Sprintf("cmd%s%s.go", rootParent, child)
		dest := filepath.Join(path, "cmd", fileName)

		// copy + apply replacements
		params := newCopyParams(src, dest)

		// re-use cmd replacements
		params.Reps = buildCmdReplacements(
			repo, author, email,
			child, parent, rootParent,
		)
		horus.CheckErr(copyFile(params))
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobraCmd.AddCommand(cmdCmd)

	cmdCmd.Flags().StringVarP(&child, "child", "c", "", "Name of the new cobra sub-command (capitalized)")
	cmdCmd.Flags().StringVarP(&parent, "parent", "P", "root", "Parent command (use \"root\" for top-level)")

	horus.CheckErr(cmdCmd.MarkFlagRequired("child"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
