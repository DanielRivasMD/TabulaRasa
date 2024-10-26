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
	"os"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	header     string
	lang       []string
	justconfig []string
	ver        string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// justCmd
var justCmd = &cobra.Command{
	Use:   "just",
	Short: "Deploy " + chalk.Yellow.Color("just") + " config templates.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy ` + chalk.Yellow.Color("just") + ` config templates over target.
Including ` + chalk.Red.Color(".justfile") + ` & ` + chalk.Red.Color(".config.just") + `
`,

	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help ` + chalk.Yellow.Color("deploy") + ` just`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// deploy justfile
		djust := copyCopyReplace(findHome()+justDir, path+"/"+"."+justfile)
		djust.files = append([]string{header}, lang...)
		catFiles(djust)

		// create config dir
		if !dirExist(path + "/" + dotjust) {
			os.MkdirAll(path+"/"+dotjust, os.ModePerm)
		}

		// deploy configs
		for _, į := range lang {
			cjust := copyCopyReplace(findHome()+justDir+"/"+į+dotconf, path+"/"+dotjust+"/"+į+dotconf)
			cjust.reps = replaceDeployJust() // automatic binding cli flags
			copyFile(cjust)
			if į == "py" {
				instpy := copyCopyReplace(findHome()+justDir+"/"+pyinstall, path+"/"+dotjust+"/"+pyinstall)
				instpy.reps = replaceDeployJust() // automatic binding cli flags
				copyFile(instpy)
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	deployCmd.AddCommand(justCmd)

	// flags
	justCmd.Flags().StringVarP(&header, "header", "t", "head", "Header")
	justCmd.Flags().StringArrayVarP(&lang, "lang", "l", []string{}, "Languages to deploy")
	justCmd.MarkFlagRequired("lang")
	justCmd.Flags().StringArrayVarP(&justconfig, "conf", "c", []string{}, "Language configurations to deploy")
	justCmd.Flags().StringVarP(&ver, "ver", "v", "", "Version to deploy")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
