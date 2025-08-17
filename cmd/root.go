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
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd
var rootCmd = &cobra.Command{
	Use:   "tab",
	Long: helpRoot,
	Example: exampleRoot,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

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
	faq       = "05license.md"
)

var (
	verbose     bool
	path        string
	author      string
	email       string
	repo        string
	description string
	user        string
	license     string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose diagnostic output")

	// appCmd.Flags().StringVar(&projectPath, "path", "", "")
	// appCmd.Flags().StringVar(&repoName, "repo", "", "Name of the repository (and Go module)")
	// appCmd.Flags().StringVar(&authorName, "author", "", "Author’s full name")
	// appCmd.Flags().StringVar(&authorEmail, "email", "", "Author’s email address")
	// appCmd.Flags().StringVar(&userName, "username", "", "GitHub username")

	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", ".", "Target directory for new app")
	rootCmd.PersistentFlags().StringVarP(&repo, "repo", "r", "", "Repository name")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "Daniel Rivas", "Author name")
	rootCmd.PersistentFlags().StringVarP(&email, "email", "e", "<danielrivasmd@gmail.com>", "Author email")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "DanielRivasMD", "GitHub username")

	_ = rootCmd.MarkFlagRequired("path")
	_ = rootCmd.MarkFlagRequired("repo")

}

////////////////////////////////////////////////////////////////////////////////////////////////////
