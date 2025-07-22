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
	// parent is the cobra command under which the new cmd is attached.
	parent string

	// child is the new sub-command’s name (capitalized).
	child string

	// rootParent holds the actual parent when it’s not the literal "root".
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
		// if parent isn’t literally "root", use it as the real parent
		if parent != "root" {
			rootParent = parent
		}

		// discover your tabularasa home dir (verbose),
		// bail out on error
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"cmdCobraCmd.Run",
				"failed to find TabulaRasa home",
				err,
				map[string]any{"verbose": verbose},
			))
		}

		// build src & dest paths
		src := filepath.Join(home, cmdDir, "cmdTemplate.go")
		fileName := fmt.Sprintf("cmd%s%s.go", rootParent, child)
		dest := filepath.Join(projectPath, "cmd", fileName)

		// copy + apply replacements
		copyParams := newCopyParams(src, dest)
		copyParams.Reps = buildCmdReplacements(
			repoName, authorName, authorEmail,
			child, parent, rootParent,
		)
		copyFile(copyParams)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobraCmd.AddCommand(cmdCmd)

	// allow overriding where your project lives
	cmdCmd.Flags().StringVar(
		&projectPath, "path", ".", "Base path of your Go project",
	)

	// new-sub-command name
	cmdCmd.Flags().StringVarP(
		&child, "child", "c", "", "Name of the new cobra sub-command (capitalized)",
	)

	// attach under this parent (defaults to root)
	cmdCmd.Flags().StringVarP(
		&parent, "parent", "p", "root", "Parent command (use \"root\" for top-level)",
	)

	// child is mandatory
	if err := cmdCmd.MarkFlagRequired("child"); err != nil {
		horus.CheckErr(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
