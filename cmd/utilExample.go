////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/domovoi"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var exampleRoot = domovoi.FormatExample(
	"tab",
	[]string{"help"},
)

var exampleCobra = domovoi.FormatExample(
	"tab",
	[]string{"cobra"},
)

var exampleCobraApp = domovoi.FormatExample(
	"tab",
	[]string{"cobra", "app"},
	[]string{"cobra", "app", "--path", "$(pwd)", "--repo", "<repo>"},
)

var exampleCobraCmd = domovoi.FormatExample(
	"tab",
	[]string{"cobra", "cmd", "--child", "ExampleCmd"},
	[]string{"cobra", "cmd", "--child", "ExampleCmd", "--parent", "RootCmd"},
)

var exampleCobraUtil = domovoi.FormatExample(
	"tab",
	[]string{"cobra", "util", "--util", "ExampleUtil"},
)

var exampleDeploy = domovoi.FormatExample(
	"tab",
	[]string{"deploy"},
)

var exampleDeployJust = domovoi.FormatExample(
	"tab",
	[]string{"deploy", "just", "--lang", "go"},
	[]string{"deploy", "just", "--ver", "1.0"},
)

var exampleDeployReadme = domovoi.FormatExample(
	"tab",
	[]string{"deploy", "readme"},
	[]string{"deploy", "readme", "--description", "Awesome project", "--license", "MIT"},
)

var exampleDeployTodor = domovoi.FormatExample(
	"tab",
	[]string{"deploy", "todor"},
)

////////////////////////////////////////////////////////////////////////////////////////////////////
