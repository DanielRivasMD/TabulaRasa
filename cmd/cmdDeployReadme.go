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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/horus"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Deploy README.md template",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Deploy ` + chalk.Yellow.Color("readme") + ` config template into your project,
splicing together overview, install/dev guides, usage and FAQ snippets.
`,
	Example: `
  ` + chalk.Cyan.Color("tab") + ` deploy readme
  ` + chalk.Cyan.Color("tab") + ` deploy readme --description "Awesome project" --license MIT
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		p := tea.NewProgram(initialModel())
		if err := p.Start(); err != nil {
			horus.CheckErr(err)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	deployCmd.AddCommand(readmeCmd)

	readmeCmd.Flags().StringVarP(&description, "description", "D", "", "Project overview text")
	readmeCmd.Flags().StringVarP(&license, "license", "L", "", "License to appear in README")

	_ = justCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
