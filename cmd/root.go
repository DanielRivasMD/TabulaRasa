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
	"embed"
	"path/filepath"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	APP     = "tab"
	VERSION = "v0.1.0"
	AUTHOR  = "Daniel Rivas"
	EMAIL   = "<danielrivasmd@gmail.com>"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type rootFlag struct {
	verbose bool
	repo    string
	user    string
	author  string
	email   string
}

type configDir struct {
	home       string
	tabularasa string
	cobra      string
	just       string
	readme     string
	todor      string
}

var (
	onceRoot   sync.Once
	rootCmd    *cobra.Command
	rootFlags  rootFlag
	configDirs configDir
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func InitDocs() {
	info := domovoi.AppInfo{
		Name:    APP,
		Version: VERSION,
		Author:  AUTHOR,
		Email:   EMAIL,
	}
	domovoi.SetGlobalDocsConfig(docsFS, info)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func GetRootCmd() *cobra.Command {
	onceRoot.Do(func() {
		d := horus.Must(domovoi.GlobalDocs())
		var err error
		rootCmd, err = d.MakeCmd("root", nil)
		horus.CheckErr(err)

		rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "Enable verbose diagnostic output")
		rootCmd.PersistentFlags().StringVarP(&rootFlags.repo, "repo", "", "", "Repository name")
		rootCmd.PersistentFlags().StringVarP(&rootFlags.user, "user", "", "DanielRivasMD", "GitHub username")
		rootCmd.PersistentFlags().StringVarP(&rootFlags.author, "author", "", "Daniel Rivas", "Author name")
		rootCmd.PersistentFlags().StringVarP(&rootFlags.email, "email", "", "<danielrivasmd@gmail.com>", "Author email")
		rootCmd.Version = VERSION

		cobra.OnInitialize(initConfigDirs)
	})
	return rootCmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func BuildCommands() {
	root := GetRootCmd()
	root.AddCommand(
		CompletionCmd(),
		IdentityCmd(),

		CobraCmd(),
		DeployCmd(),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfigDirs() {
	var err error
	configDirs.home, err = domovoi.FindHome(false) // verbose false for init
	horus.CheckErr(err, horus.WithCategory("init_error"), horus.WithMessage("getting home directory"))
	configDirs.tabularasa = filepath.Join(configDirs.home, ".tabularasa")
	configDirs.cobra = filepath.Join(configDirs.tabularasa, "cobra")
	configDirs.just = filepath.Join(configDirs.tabularasa, "just")
	configDirs.readme = filepath.Join(configDirs.tabularasa, "readme")
	configDirs.todor = filepath.Join(configDirs.tabularasa, "todor")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(GetRootCmd().Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
