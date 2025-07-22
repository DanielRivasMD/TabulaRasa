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
	"fmt"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
const (
	configDir = "/" + ".tabularasa"
	cobraDir  = configDir + "/" + "cobraApp"
	cmdDir    = configDir + "/" + "cobraCmd"
	utilDir   = configDir + "/" + "cobraUtil"
	justDir   = configDir + "/" + "just"
	readmeDir = configDir + "/" + "readme"
	todorDir  = configDir + "/" + "todor"

	dotconf   = ".conf"
	dotjust   = ".just"
	justfile  = "justfile"
	readme    = "README.md"
	todor     = "todor"
	pyinstall = "pyinstall.sh"
	overview  = "01overview.md"
	usage     = "03usage.md"
	faq       = "05faq.md"
)

var (
	path        string
	author      string
	email       string
	repo        string
	description string
	user        string
	license     string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd
var rootCmd = &cobra.Command{
	Use:   "tab",
	Short: "Provide a canvas to write on",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Blue.Color("tab") + `, provide a set of templates to facilite software deployment
`,
	Example: chalk.White.Color("tab") + ` ` + chalk.Bold.TextStyle(chalk.White.Color("help")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// Execute is the entry point for executing the command.
// It wraps the root command execution and handles any errors using Horus's checkErr function.
func Execute() {
	err := rootCmd.Execute()
	horus.CheckErr(err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initializeConfig sets up configuration using Viper.
// It also leverages the Domovoi library for any additional configuration management.
func initializeConfig(cmd *cobra.Command, configPath string, configName string) error {
	// Create a new Viper instance for configuration management.
	vConfig := viper.New()

	// Set the path and name of the configuration file.
	vConfig.AddConfigPath(configPath)
	vConfig.SetConfigName(configName)

	// Attempt to read the configuration file.
	err := vConfig.ReadInConfig()
	if err != nil {
		// If the config file is not found, that's acceptable.
		_, notFound := err.(viper.ConfigFileNotFoundError)
		if !notFound {
			// Return the error for any other issue.
			return err
		}
	}

	// Bind command flags with configuration values.
	bindFlags(cmd, vConfig)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bindFlags synchronizes each Cobra flag with the corresponding Viper configuration value.
// If the flag is unset and a configuration value is available, the flag is updated.
func bindFlags(cmd *cobra.Command, vConfig *viper.Viper) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// If the flag wasn't explicitly set but has a value in the config, apply that value.
		if !flag.Changed && vConfig.IsSet(flag.Name) {
			value := vConfig.Get(flag.Name)
			cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", value))
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// Execute prior main.
// init registers persistent flags and performs additional initialization tasks.
func init() {
	// Set up persistent flags.
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", ".", "Path to deploy")
	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Repository name")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "Daniel Rivas", "Provide author")
	rootCmd.PersistentFlags().StringVarP(&email, "email", "e", "<danielrivasmd@gmail.com>", "Provide email")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "DanielRivasMD", "Provide GitHub username")

	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("repo")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
