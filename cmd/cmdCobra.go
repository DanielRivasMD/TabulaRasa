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
	"os"
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var cobraCmd = &cobra.Command{
	Use:     "cobra",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobra,
	Example: exampleCobra,
}

var cobraAppCmd = &cobra.Command{
	Use:     "app",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobraApp,
	Example: exampleCobraApp,

	Run: runCobraApp,
}

var cobraCmdCmd = &cobra.Command{
	Use:     "cmd",
	Short:   "Construct cobra apps from templates",
	Long:    helpCobraCmd,
	Example: exampleCobraCmd,

	Run: runCobraCmd,
}

var cobraUtilCmd = &cobra.Command{
	Use:     "util",
	Short:   "Import utility templates",
	Long:    helpCobraUtil,
	Example: exampleCobraUtil,

	Run: runCobraUtil,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// cobra app
	force bool

	// cobra cmd
	parent     string
	child      string
	rootParent string

	// cobra util
	util string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(cobraCmd)
	cobraCmd.AddCommand(cobraAppCmd, cobraCmdCmd, cobraUtilCmd)

	// cobra app
	cobraAppCmd.Flags().BoolVarP(&force, "force", "", false, "Force install go dependencies")

	// cobra cmd
	cobraCmdCmd.Flags().StringVarP(&child, "child", "", "", "Name of the new cobra sub-command (capitalized)")
	cobraCmdCmd.Flags().StringVarP(&parent, "parent", "", "root", "Parent command (use \"root\" for top-level)")
	horus.CheckErr(cobraCmdCmd.MarkFlagRequired("child"))

	// cobra util
	cobraUtilCmd.Flags().StringVarP(&util, "util", "", "", "Utility template name (capitalize)")
	horus.CheckErr(cobraUtilCmd.MarkFlagRequired("util"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraApp(cmd *cobra.Command, args []string) {
	home, err := domovoi.FindHome(verbose)
	if err != nil {
		horus.CheckErr(horus.NewHerror(
			"cmdCobraApp.Run",
			"failed to find TabulaRasa home",
			err,
			nil,
		))
	}

	if repo == "" {
		// TODO: add error handling & potentially domovoi implementation
		dir, _ := os.Getwd()
		repo = filepath.Base(dir)
	}

	// TODO: copy help & example util
	// TODO: begin here
	// Copy template files into the target directory
	copyParams := newCopyParams(home+cobraDir, path)
	copyParams.Reps = buildAppReplacements(repo, author, email, user)
	horus.CheckErr(copyDir(copyParams))

	// Initialize Go module and tidy dependencies
	// TODO: add file check & file remove
	if force {
		domovoi.RemoveFile("go.mod", verbose)
		domovoi.RemoveFile("go.sum", verbose)

		// TODO: finish force feature
		// horus.CheckErr(
		// 	func() error {
		// 		_, err := domovoi.RemoveFile(metaFile, verbose)(metaFile)
		// 		return err
		// 	}(),
		// 	horus.WithOp(op),
		// 	horus.WithMessage("removing metadata file"),
		// )

	}
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "init", "github.com/"+user+"/"+repo))
	horus.CheckErr(domovoi.ExecCmd("go", "mod", "tidy"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraCmd(cmd *cobra.Command, args []string) {
	if parent != "root" {
		rootParent = parent
	}

	home, err := domovoi.FindHome(verbose)
	if err != nil {
		horus.CheckErr(horus.NewHerror(
			"cmdCobraCmd.Run",
			"failed to find TabulaRasa home",
			err,
			nil,
		))
	}

	// TODO: refactor rootParent
	// build src & dest paths
	src := filepath.Join(home, cmdDir, "cmdTemplate.go")
	fileName := fmt.Sprintf("cmd%s%s.go", rootParent, child)
	dest := filepath.Join(path, "cmd", fileName)

	// copy + apply replacements
	params := newCopyParams(src, dest)

	// re-use cmd replacements
	params.Reps = buildCmdReplacements(
		repo, author, email,
		child, parent, rootParent,
	)
	horus.CheckErr(copyFile(params))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCobraUtil(cmd *cobra.Command, args []string) {
	home, err := domovoi.FindHome(verbose)
	if err != nil {
		horus.CheckErr(horus.NewHerror(
			"cmdCobraUtil.Run",
			"failed to find TabulaRasa home",
			err,
			nil,
		))
	}

	// build src & dest paths
	src := filepath.Join(home, utilDir, util+".go")
	dest := filepath.Join(path, "cmd", util+".go")

	// copy + apply replacements
	params := newCopyParams(src, dest)

	// re-use cmd replacements
	params.Reps = buildCmdReplacements(
		repo, author, email,
		util, "", "",
	)
	horus.CheckErr(copyFile(params))
}

////////////////////////////////////////////////////////////////////////////////////////////////////
