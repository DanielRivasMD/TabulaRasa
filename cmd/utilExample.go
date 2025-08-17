////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"strings"

	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// formatExample builds a multi‐line example block
// each usage is a slice of “tokens”: [ command, flagOrArg, flagOrArg, ... ].
//
//	app:    your binary name, e.g. "lilith"
//	usages: one or more usages—each becomes its own line.
func formatExample(app string, usages ...[]string) string {
	var b strings.Builder

	for i, usage := range usages {
		if len(usage) == 0 {
			continue
		}

		// first token is the subcommand
		b.WriteString(
			chalk.White.Color(app) + " " +
				chalk.White.Color(chalk.Bold.TextStyle(usage[0])),
		)

		// remaining tokens are either flags (--foo) or args
		for _, tok := range usage[1:] {
			switch {
			case strings.HasPrefix(tok, "--"):
				b.WriteString(" " + chalk.Italic.TextStyle(chalk.White.Color(tok)))
			default:
				b.WriteString(" " + chalk.Dim.TextStyle(chalk.Italic.TextStyle(tok)))
			}
		}

		if i < len(usages)-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var exampleRoot = formatExample(
	"tab",
	[]string{"help"},
)

var exampleCobra = formatExample(
	"tab",
	[]string{"cobra"},
)

var exampleCobraApp = formatExample(
	"tab",
	[]string{"cobra", "app"},
	[]string{"cobra", "app", "--path", "$(pwd)", "--repo", "<repo>"},
)

var exampleCobraCmd = formatExample(
	"tab",
	[]string{"cobra", "cmd", "--child", "ExampleCmd"},
	[]string{"cobra", "cmd", "--child", "ExampleCmd", "--parent", "RootCmd"},
)

var exampleCobraUtil = formatExample(
	"tab",
	[]string{"cobra", "util", "--util", "ExampleUtil"},
)

var exampleDeploy = formatExample(
	"tab",
	[]string{"deploy"},
)

var exampleDeployJust = formatExample(
	"tab",
	[]string{"deploy", "just", "--lang", "go"},
	[]string{"deploy", "just", "--ver", "1.0"},
)

var exampleDeployReadme = formatExample(
	"tab",
	[]string{"deploy", "readme"},
	[]string{"deploy", "readme", "--description", "Awesome project", "--license", "MIT"},
)

var exampleDeployTodor = formatExample(
	"tab",
	[]string{"deploy", "todor"},
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
