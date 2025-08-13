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
	"os"
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Construct cobra apps from templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Construct ` + chalk.Yellow.Color("cobra") + ` apps from predefined templates
`,
	Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` ` + chalk.Green.Color("app") +
		` --` + chalk.Blue.Color("path") + ` $(pwd) --` + chalk.Blue.Color("repo") + ` Tabularasa
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		home, err := domovoi.FindHome(verbose)
		if err != nil {
			horus.CheckErr(horus.NewHerror(
				"cmdCobraApp.Run",
				"failed to find TabulaRasa home",
				err,
				nil,
			))
		}

		if repo == "" {
			// TODO: add error handling & potentially domovoi implementation
			dir, _ := os.Getwd()
			repo = filepath.Base(dir)
		}

		// Copy template files into the target directory
		copyParams := newCopyParams(home+cobraDir, path)
		copyParams.Reps = buildAppReplacements(repo, author, email, user)
		horus.CheckErr(copyDir(copyParams))

		// Initialize Go module and tidy dependencies
		// TODO: add file check & file remove
		if force {
			domovoi.RemoveFile("go.mod", verbose)
			domovoi.RemoveFile("go.sum", verbose)

			// TODO: finish force feature
			// horus.CheckErr(
			// 	func() error {
			// 		_, err := domovoi.RemoveFile(metaFile, verbose)(metaFile)
			// 		return err
			// 	}(),
			// 	horus.WithOp(op),
			// 	horus.WithMessage("removing metadata file"),
			// )

		}
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+user+"/"+repo))
		horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	force bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobraCmd.AddCommand(appCmd)

	appCmd.Flags().BoolVar(&force, "force", false, "Force install go dependencies")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
