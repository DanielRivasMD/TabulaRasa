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

// justCmd represents the just command
var justCmd = &cobra.Command{
	Use:   "just",
	Short: "",
	Long: `.`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// deploy justfile
		concatenateFiles(findHome() + justDir, path + "/" + "." + justfile, append([]string{header}, lang...))

		// deploy config
		params := copyCR(findHome() + justDir + "/" + justconfig, path + "/" + "." + justconfig)
		params.reps = repsDeployJust() // automatic binding cli flags
		copyFile(params)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	deployCmd.AddCommand(justCmd)

	// flags
	justCmd.Flags().StringVarP(&header, "head", "e", "head", "Header")
	justCmd.Flags().StringArrayVarP(&lang, "lang", "l", []string{}, "Languages to deploy")
	justCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
