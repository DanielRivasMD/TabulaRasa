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
	"errors"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// custom type restricting `lang` flag values
type langType struct {
	validValues []string
	selected    []string
}

// possible values
var validOptions = []string{"go", "jl", "py", "rs"}

func (f *langType) String() string {
	if len(f.selected) > 0 {
		return f.selected[0]
	}
	return ""
}

func (f *langType) Set(value string) error {
	// check value valid
	for _, valid := range f.validValues {
		if value == valid {
			f.selected = append(f.selected, value)
			return nil
		}
	}
	return errors.New(`invalid value`)
}

func (f *langType) Type() string {
	return "langType"
}

func joinValues(values []string) string {
	result := ""
	for _, v := range values {
		result += v + ", "
	}
	return result[:len(result)-2]
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	lang = &langType{validValues: validOptions}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// deployCmd
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"d"},
	Short:   "Deploy config templates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Cyan.Color("tab") + ` deploys config templates to target locations

Available templates:
	` + chalk.Green.Color("just") + `   - Build system template
	` + chalk.Green.Color("readme") + ` - Documentation starter
	` + chalk.Green.Color("todor") + `  - Task tracking starter
`,

	Example: `
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("deploy") + ` --` + chalk.Blue.Color("lang") + ` go
` + chalk.Cyan.Color("tab") + ` ` + chalk.Yellow.Color("deploy") + ` --` + chalk.Blue.Color("lang") + ` py
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(deployCmd)

	// flags
	deployCmd.PersistentFlags().VarP(lang, "lang", "l", "Languages to deploy (allowed: go, jl, py, rs)")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
