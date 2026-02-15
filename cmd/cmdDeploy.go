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

// BUG: deploy root flags

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Short:   "Deploy config templates",
	Long:    helpDeploy,
	Example: exampleDeploy,
}

var deployJustCmd = &cobra.Command{
	Use:     "just",
	Short:   "Build system files",
	Long:    helpDeployJust,
	Example: exampleDeployJust,

	Run: runDeployJust,
}

var deployReadmeCmd = &cobra.Command{
	Use:     "readme",
	Short:   "README scaffold",
	Long:    helpDeployReadme,
	Example: exampleDeployReadme,

	Run: runDeployReadme,
}

var deployTodorCmd = &cobra.Command{
	Use:     "todor",
	Short:   "Task-tracker starter",
	Long:    helpDeployTodor,
	Example: exampleDeployTodor,

	Run: runDeployTodor,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

const HEADER = "head.just"

type LangType struct {
	validValues []string
	Selected    []string
}

var validLangs = []string{"go", "jl", "py", "rs", "R"}

func (f *LangType) String() string {
	if len(f.Selected) > 0 {
		return f.Selected[0]
	}
	return ""
}

func (f *LangType) Set(value string) error {
	if slices.Contains(f.validValues, value) {
		f.Selected = append(f.Selected, value)
		return nil
	}
	return fmt.Errorf("invalid value '%s', allowed: %s", value, joinValues(f.validValues))
}

func (f *LangType) Type() string {
	return "LangType"
}

func joinValues(values []string) string {
	return strings.Join(values, ", ")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	lang = &LangType{validValues: validLangs}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.AddCommand(deployJustCmd, deployReadmeCmd, deployTodorCmd)
	deployCmd.PersistentFlags().VarP(lang, "lang", "l", "Templates to deploy (allowed: "+joinValues(validLangs)+")")
	_ = deployJustCmd.MarkFlagRequired("lang")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployJust(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.just"

	var err error
	if flags.repo == "" {
		flags.repo, err = domovoi.CurrentDir()
		horus.CheckErr(err, horus.WithOp(op))
	}
	replaces := deployJustReplacements()

	files := []string{HEADER}

	for _, sel := range lang.Selected {
		switch sel {
		case "go":
			files = append(files, "go.just")
		case "rs":
			files = append(files, "rs.just")
		}
	}

	pairs := []filePair{
		{files, ".justfile"},
	}

	for _, p := range pairs {
		mbomboForging(
			op,
			newMbomboConfig(
				dirs.just,
				p.out,
				p.files,
				replaces...,
			))
	}
}

func deployJustReplacements() []mbomboReplace {
	return []mbomboReplace{
		Replace("APP", flags.repo),
		Replace("EXE", strings.ToLower(flags.repo)),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployReadme(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.readme"

	p := tea.NewProgram(initialModel())
	m, err := p.Run() // Run returns (tea.Model, error)
	horus.CheckErr(err, horus.WithOp(op))

	// Type assert back to our model
	final, ok := m.(model)
	if !ok {
		horus.CheckErr(fmt.Errorf("unexpected model type"), horus.WithOp(op))
	}

	// Now you have the captured values
	desc := final.description
	lic := final.license

	repo, err := domovoi.CurrentDir()
	horus.CheckErr(err)

	// Example: build replacements for README
	replaces := []mbomboReplace{
		Replace("REPOSITORY", repo),
		Replace("OVERVIEW", desc),
		Replace("LICENSE", lic),
		Replace("YEAR", strconv.Itoa(time.Now().Year())),
	}

	mbomboForging(
		op,
		newMbomboConfig(
			dirs.readme,
			"README.md",
			[]string{"readme.md"},
			replaces...,
		),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDeployTodor(cmd *cobra.Command, args []string) {
	op := "tabularasa.deploy.todor"

	mbomboForging(
		op,
		newMbomboConfig(
			dirs.todor,
			".todor",
			[]string{"todor"},
		))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
