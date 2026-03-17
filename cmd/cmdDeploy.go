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
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func DeployCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("deploy", nil))
	cmd.AddCommand(
		DeployJustCmd(),
		DeployReadmeCmd(),
		DeployTodorCmd(),
	)
	return cmd
}

func DeployJustCmd() *cobra.Command {
	deployFlags.language = langValue{allowed: []string{"go", "rs"}}
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("just", runDeployJust))
	cmd.Flags().VarP(&deployFlags.language, "lang", "l", "Templates to deploy (allowed: go, rs)")

	horus.CheckErr(cmd.RegisterFlagCompletionFunc("lang", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return deployFlags.language.allowed, cobra.ShellCompDirectiveNoFileComp
	}))

	return cmd
}

func DeployReadmeCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("readme", runDeployReadme))
}

func DeployTodorCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("todor", runDeployTodor))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployJust(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.just"

	repo := horus.Must(domovoi.CurrentDir())
	replaces := []moldReplace{
		Replace("XXX_APP_XXX", repo),
		Replace("XXX_EXE_XXX", strings.ToLower(repo)),
	}

	files := []string{"head.just"}
	switch deployFlags.language.value {
	case "go":
		files = append(files, "go.just")
	case "rs":
		files = append(files, "rs.just")
	}

	moldForging(op, newMoldConfig(configDirs.just, ".justfile", files, replaces...))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployReadme(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.readme"

	repo := horus.Must(domovoi.CurrentDir())
	replaces := []moldReplace{
		Replace("XXX_REPO_XXX", repo),
		Replace("XXX_YEAR_XXX", strconv.Itoa(time.Now().Year())),
	}

	moldForging(op, newMoldConfig(configDirs.readme, "README.md", []string{"readme.md"}, replaces...))
}

func runDeployTodor(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.todor"
	moldForging(op, newMoldConfig(configDirs.todor, ".todor", []string{"todor"}))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type deployFlag struct {
	language langValue
}

type langValue struct {
	value   string
	allowed []string
}

func (l *langValue) String() string {
	return l.value
}

func (l *langValue) Set(s string) error {
	for _, a := range l.allowed {
		if a == s {
			l.value = s
			return nil
		}
	}
	return fmt.Errorf("invalid language %q, allowed: %v", s, l.allowed)
}

func (l *langValue) Type() string {
	return "lang"
}

var (
	deployFlags deployFlag
)

////////////////////////////////////////////////////////////////////////////////////////////////////
