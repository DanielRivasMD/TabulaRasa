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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var langFlag LangType

type LangType struct {
	Selected []string
}

func (l *LangType) String() string { return strings.Join(l.Selected, ",") }
func (l *LangType) Set(v string) error {
	l.Selected = append(l.Selected, v)
	return nil
}
func (l *LangType) Type() string { return "lang" }

func DeployCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("deploy", nil))
	cmd.PersistentFlags().VarP(&langFlag, "lang", "l", "Templates to deploy (allowed: go, jl, py, rs, R)")
	cmd.AddCommand(
		DeployJustCmd(),
		DeployReadmeCmd(),
		DeployTodorCmd(),
	)
	return cmd
}

func DeployJustCmd() *cobra.Command {
	cmd := horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("just", runDeployJust))
	_ = cmd.MarkFlagRequired("lang")
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
	repo := rootFlags.repo
	if repo == "" {
		var err error
		repo, err = domovoi.CurrentDir()
		horus.CheckErr(err)
	}

	replaces := []moldReplace{
		Replace("APP", repo),
		Replace("EXE", strings.ToLower(repo)),
	}

	files := []string{"head.just"}
	for _, lang := range langFlag.Selected {
		switch lang {
		case "go":
			files = append(files, "go.just")
		case "rs":
			files = append(files, "rs.just")
		}
	}

	moldForging(op, newMoldConfig(configDirs.just, ".justfile", files, replaces...))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployReadme(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.readme"
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	horus.CheckErr(err)

	final, ok := m.(model)
	if !ok {
		horus.CheckErr(fmt.Errorf("unexpected model type"))
	}

	repo, err := domovoi.CurrentDir()
	horus.CheckErr(err)

	replaces := []moldReplace{
		Replace("REPOSITORY", repo),
		Replace("OVERVIEW", final.description),
		Replace("LICENSE", final.license),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}

	moldForging(op, newMoldConfig(configDirs.readme, "README.md", []string{"readme.md"}, replaces...))
}

func runDeployTodor(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.todor"
	moldForging(op, newMoldConfig(configDirs.todor, ".todor", []string{"todor"}))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
