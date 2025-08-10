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
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     "TOOL",
	Long:    helpRoot,
	Example: exampleRoot,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	verbose bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose diagnostics")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpRoot = chalk.Bold.TextStyle(chalk.Green.Color("AUTHOR ")) +
	chalk.Dim.TextStyle(chalk.Italic.TextStyle("EMAIL")) +
	chalk.Dim.TextStyle(chalk.Cyan.Color("\n\n"))

var exampleRoot = chalk.White.Color("TOOL") + " " + chalk.Bold.TextStyle(chalk.White.Color("help"))

////////////////////////////////////////////////////////////////////////////////////////////////////
