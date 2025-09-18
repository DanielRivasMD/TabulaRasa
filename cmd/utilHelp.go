////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/domovoi"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpRoot = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Provide a set of templates to facilite software deployment",
)

var helpCobra = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Enables the creation of cobra applications using predefined templates for rapid development",
)

var helpCobraApp = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Construct "+chalk.Italic.TextStyle("cobra")+" apps from predefined templates",
)

var helpCobraCmd = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Construct "+chalk.Italic.TextStyle("cobra")+" apps from predefined templates",
)

var helpCobraUtil = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy a utility from predefined templates",
)

var helpDeploy = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy selected config templates into your project",
)

var helpDeployJust = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy justfile & language‐specific configs",
)

var helpDeployReadme = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy readme template with sections: overview, install/dev guides, usage & FAQ snippets",
)

var helpDeployTodor = domovoi.FormatHelp(
	"Daniel Rivas",
	"danielrivasmd@gmail.com",
	"Deploy top-level todor config",
)

var helpEngrave = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"",
)

////////////////////////////////////////////////////////////////////////////////////////////////////
