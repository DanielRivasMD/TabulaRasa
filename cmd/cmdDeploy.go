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
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	header string
	lang []string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// deployCmd
var deployCmd = &cobra.Command{
	Use:   "deploy [just|todor]",
	Short: "Deploy config templates.",
	Long:  `.`,

	Example: ``,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"just", "todor"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "just":
			// TODO: write concatenate & finish development
			// concatenateFiles(path, )
		case "todor":
			// TODO: relocate template & finish development
			// copyFile(findHome()+configDir + .todor, path)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(deployCmd)

	// flags
	deployCmd.Flags().StringVarP(&header, "head", "e", "head", "Header")
	deployCmd.Flags().StringArrayVarP(&lang, "lang", "l", []string{}, "Languages to deploy")
	// deployCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
