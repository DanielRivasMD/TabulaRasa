////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// buildAppReplacements returns the list of token replacements for a new cobra app.
func buildAppReplacements(repo, author, email, user string) []rep {
	year := strconv.Itoa(time.Now().Year())
	return []rep{
		{old: "YEAR", new: year},
		{old: "REPOSITORY", new: repo},
		{old: "TOOL", new: strings.ToLower(repo)},
		{old: "AUTHOR", new: author},
		{old: "EMAIL", new: email},
		{old: "USER", new: user},
	}
}

// buildCmdReplacements returns token replacements for generating a new sub-command.
func buildCmdReplacements(repo, author, email, child, parent, root string) []rep {
	year := strconv.Itoa(time.Now().Year())
	return []rep{
		{old: "YEAR", new: year},
		{old: "REPOSITORY", new: repo},
		{old: "TOOL", new: strings.ToLower(repo)},
		{old: "AUTHOR", new: author},
		{old: "EMAIL", new: email},
		{old: "COMMAND", new: child},
		{old: "CHILD", new: strings.ToLower(child)},
		{old: "PARENT", new: strings.ToLower(parent)},
		{old: "ROOT", new: strings.ToLower(root)},
	}
}

// buildDeployReplacements returns tokens for a bare “deploy” template.
func buildDeployReplacements(repo string) []rep {
	return []rep{
		{old: "APP", new: repo},
		{old: "EXE", new: strings.ToLower(repo)},
	}
}

// buildReadmeReplacements returns tokens for README.md generation on deploy.
// It checks for a LICENSE file (verbose) and auto-detects license type when none was provided.
func buildReadmeReplacements(
	langTag, desc, repo, user, author, license, targetPath string,
) ([]rep, error) {
	licensePath := filepath.Join(targetPath, "LICENSE")

	// check existence of LICENSE (verbose=true)
	exists, err := domovoi.FileExist(licensePath, nil, verbose)
	if err != nil {
		return nil, horus.NewHerror(
			"buildReadmeReplacements",
			fmt.Sprintf("failed to stat %s", licensePath),
			err,
			map[string]any{"path": licensePath},
		)
	}

	// if LICENSE exists but no license was set, try to detect it
	if exists && license == "" {
		detected, detErr := detectLicense(licensePath)
		if detErr != nil {
			return nil, horus.NewHerror(
				"buildReadmeReplacements",
				"license detection failed",
				detErr,
				map[string]any{"path": licensePath},
			)
		}
		license = detected
	}

	year := strconv.Itoa(time.Now().Year())
	reps := []rep{
		{old: "LANG", new: langTag},
		{old: "OVERVIEW", new: desc},
		{old: "REPOSITORY", new: repo},
		{old: "USER", new: user},
		{old: "AUTHOR", new: author},
		{old: "YEAR", new: year},
		{old: "LICENSETYPE", new: license},
	}

	return reps, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
