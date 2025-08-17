////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// formatHelp produces the “help” header + description.
//
//	author: name, e.g. "Daniel Rivas"
//	email:  email, e.g. "danielrivasmd@gmail.com"
//	desc:   the multi‐line description, "\n"-separated.
func formatHelp(author, email, desc string) string {
	header := chalk.Bold.TextStyle(
		chalk.Green.Color(author+" "),
	) +
		chalk.Dim.TextStyle(
			chalk.Italic.TextStyle("<"+email+">"),
		)

	// prefix two newlines to your desc, chalk it cyan + dim it
	body := "\n\n" + desc
	return header + chalk.Dim.TextStyle(chalk.Cyan.Color(body))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpRoot = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Provide a set of templates to facilite software deployment",
)

var helpCobra = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Enables the creation of cobra applications using predefined templates for rapid development",
)

var helpCobraApp = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Construct "+chalk.Italic.TextStyle("cobra") + " apps from predefined templates",
)

var helpCobraCmd = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Construct "+chalk.Italic.TextStyle("cobra") + " apps from predefined templates",
)

var helpCobraUtil = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy a utility from predefined templates",
)

var helpDeploy = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy selected config templates into your project",
)

var helpDeployJust = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy justfile & language‐specific configs",
)

var helpDeployReadme = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy readme template with sections: overview, install/dev guides, usage & FAQ snippets",
)

var helpDeployTodor = formatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy top-level todor config",
)

////////////////////////////////////////////////////////////////////////////////////////////////////
