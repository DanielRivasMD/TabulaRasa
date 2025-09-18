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
	"slices"
	"strings"

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

var validLangs = []string{"go", "golib", "jl", "py", "rs", "rslib"}

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

	// deploy
	deployCmd.PersistentFlags().VarP(lang, "lang", "l", "Templates to deploy (allowed: "+joinValues(validLangs)+")")

	// deploy just
	_ = deployJustCmd.MarkFlagRequired("lang")

	// deploy readme
	deployReadmeCmd.Flags().StringVarP(&flags.description, "description", "D", "", "Project overview text")
	deployReadmeCmd.Flags().StringVarP(&flags.license, "license", "L", "", "License to appear in README")
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

	pairs := []filePair{
		// BUG: append `.just` extension
		{append([]string{HEADER}, lang.Selected...), ".jusfile"},
		// TODO: add just config
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
	// BUG: tui works, no deployment. probably must bind variables
	p := tea.NewProgram(initialModel())
	horus.CheckErr(p.Start())
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
