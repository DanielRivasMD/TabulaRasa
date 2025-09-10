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
	"path/filepath"
	"unicode"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     "tab",
	Long:    helpRoot,
	Example: exampleRoot,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	dirs  configDirs
	flags cobraFlags
)

type configDirs struct {
	home       string
	tabularasa string
	cobra      string
	just       string
	readme     string
	todor      string
}

type cobraFlags struct {
	// root
	verbose     bool
	author      string
	email       string
	repo        string
	user        string
	description string
	license     string

	// cobra.app
	force bool

	// cobra.cmd
	cmd string
	cmdLower string
	cmdUpper string

	// cobra.util
	util string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "Enable verbose diagnostic output")

	rootCmd.PersistentFlags().StringVarP(&flags.repo, "repo", "", "", "Repository name")
	rootCmd.PersistentFlags().StringVarP(&flags.user, "user", "", "DanielRivasMD", "GitHub username")
	rootCmd.PersistentFlags().StringVarP(&flags.author, "author", "", "Daniel Rivas", "Author name")
	rootCmd.PersistentFlags().StringVarP(&flags.email, "email", "", "<danielrivasmd@gmail.com>", "Author email")

	_ = rootCmd.MarkFlagRequired("path")
	_ = rootCmd.MarkFlagRequired("repo")

	cobra.OnInitialize(initConfigDirs)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfigDirs() {
	var err error
	dirs.home, err = domovoi.FindHome(flags.verbose)
	horus.CheckErr(err, horus.WithCategory("init_error"), horus.WithMessage("getting home directory"))
	dirs.tabularasa = filepath.Join(dirs.home, ".tabularasa")
	dirs.cobra = filepath.Join(dirs.tabularasa, "cobra")
	dirs.just = filepath.Join(dirs.tabularasa, "just")
	dirs.readme = filepath.Join(dirs.tabularasa, "readme")
	dirs.todor = filepath.Join(dirs.tabularasa, "todor")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// lowerFirst returns s with its first Unicode letter lower-cased.
// If s is empty, it returns s unchanged.
func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// upperFirst returns s with its first Unicode letter upper-cased.
// If s is empty, it returns s unchanged.
func upperFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
