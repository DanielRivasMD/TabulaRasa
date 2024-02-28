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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
const (
	configDir = "/" + ".tabularasa"
	cobraDir = configDir + "/" + "cobra"
	justDir = configDir + "/" + "just"
	justfile = ".justfile"
	justconfig = ".config.just"
	todorDir = configDir + "/" + "todor"
	todorconfig = "todor"
	todor = ".todor"
)

var (
	path string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd
var rootCmd = &cobra.Command{
	Use:   "tabularasa",
	Short: "Provide a canvas to write on.",
	Long: chalk.Green.Color("Daniel Rivas <danielrivasmd@gmail.com>") + `

` + chalk.Cyan.Color("tabularasa") + chalk.Blue.Color(` provides a set of
templates to facilite software deployment.
`) + ``,

	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute
func Execute() {
	ε := rootCmd.Execute()
	if ε != nil {
		log.Fatal(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initialize config
func initializeConfig(κ *cobra.Command, configPath string, configName string) error {

	// initialize viper
	ω := viper.New()

	// collect config path & file from persistent flags
	ω.AddConfigPath(configPath)
	ω.SetConfigName(configName)

	// read config file
	ε := ω.ReadInConfig()
	if ε != nil {
		// okay if no config file
		_, ϙ := ε.(viper.ConfigFileNotFoundError)
		if !ϙ {
			// error if not parse config file
			return ε
		}
	}

	// bind flags viper
	bindFlags(κ, ω)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind each cobra flag viper configuration
func bindFlags(κ *cobra.Command, ω *viper.Viper) {

	κ.Flags().VisitAll(func(ζ *pflag.Flag) {

		// apply viper config value flag
		if !ζ.Changed && ω.IsSet(ζ.Name) {
			ν := ω.Get(ζ.Name)
			κ.Flags().Set(ζ.Name, fmt.Sprintf("%v", ν))
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {

	// persistent flags
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Path to deploy")
	rootCmd.MarkFlagRequired("path")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
