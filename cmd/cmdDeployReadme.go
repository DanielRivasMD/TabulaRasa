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

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// readmeCmd
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Deploy" + chalk.Yellow.Color("readme") + " config template.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
Deploy ` + chalk.Yellow.Color("readme") + ` config template over target.
`,


	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help ` + chalk.Yellow.Color("deploy") + ` readme`, 

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// deploy readme
		md := copyCopyReplace(findHome() + readmeDir, path + "/" + readme)
		md.files = append(md.files, overview, filepath.Join("02"+lang.selected[0]+"_install.md"), usage, filepath.Join("04"+lang.selected[0]+"_dev.md"), faq)
		md.reps = replaceDeployReadme() // automatic bindings cli flags
		catFiles(md, "")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	deployCmd.AddCommand(readmeCmd)

	// flags
	readmeCmd.Flags().StringVarP(&description, "description", "", "", "Description")
	readmeCmd.Flags().StringVarP(&license, "license", "", "", "License")
	readmeCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
