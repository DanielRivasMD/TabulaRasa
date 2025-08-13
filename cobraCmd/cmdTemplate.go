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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var CHILDCmd = &cobra.Command{
	Use:     "CHILD",
	Short:   "",
	// Long:    helpCOMMAND,
	// Example: exampleCOMMAND,

	// Run: runCOMMAND,

}

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	// PARENTCmd.AddCommand(CHILDCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// var helpCOMMAND = formatHelp(
// 	"AUTHOR",
// 	"EMAIL",
// 	"",
// )

// var exampleCOMMAND = formatExample(
// 	"TOOL",
// 	[]string{"CHILD"},
// )

////////////////////////////////////////////////////////////////////////////////////////////////////

// func runCOMMAND(cmd *cobra.Command, args []string) {

// }

////////////////////////////////////////////////////////////////////////////////////////////////////
