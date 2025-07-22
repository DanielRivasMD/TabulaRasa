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
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// cobraCmd represents the top-level `tab cobra` command
	cobraCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Construct cobra apps from templates",
		Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
			chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("tab") + ` enables the creation of cobra applications using predefined templates for rapid development

Available commands:
  ` + chalk.Yellow.Color("completion") + ` - Manage shell completion scripts
  ` + chalk.Yellow.Color("identity") + `   - Configure app identity settings
`,
		Example: `
  ` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` app
  ` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` cmd
  ` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("cobra") + ` util
`,
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	// register the cobra subcommand
	rootCmd.AddCommand(cobraCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
