////////////////////////////////////////////////////////////////////////////////////////////////////

pub const MAIN_GO: &str = r#"
package main

////////////////////////////////////////////////////////////////////////////////////////////////////

import "github.com/DanielRivasMD/XXX_REPO_XXX/cmd"

////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	cmd.InitDocs()
	cmd.BuildCommands()
	cmd.Execute()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const ROOT_GO: &str = r#"
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"embed"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const (
	APP     = "XXX_CLI_LOWERCASE_XXX"
	VERSION = "v0.1.0"
	AUTHOR  = "XXX_AUTHOR_XXX"
	EMAIL   = "XXX_EMAIL_XXX"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	onceRoot  sync.Once
	rootCmd   *cobra.Command
	rootFlags struct {
		verbose bool
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func InitDocs() {
	info := domovoi.AppInfo{
		Name:    APP,
		Version: VERSION,
		Author:  AUTHOR,
		Email:   EMAIL,
	}
	domovoi.SetGlobalDocsConfig(docsFS, info)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func GetRootCmd() *cobra.Command {
	onceRoot.Do(func() {
		var err error
		rootCmd, err = horus.Must(domovoi.GlobalDocs()).MakeCmd("root", nil)
		horus.CheckErr(err)

		rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "Enable verbose diagnostics")
		rootCmd.Version = VERSION
	})
	return rootCmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(GetRootCmd().Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func BuildCommands() {
	root := GetRootCmd()
	root.AddCommand(
		CompletionCmd(),
		IdentityCmd(),


	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const DOCS_JSON: &str = r#"
{
  "identity": {
    "use": "identity",
    "hidden": true
  },
  "completion": {
    "use": "completion [bash|zsh|fish|powershell]",
    "hidden": true,
    "long": "To load completions:\n\nBash:\n\n  $ source <(%[1]s completion bash)\n\n  # To load completions for each session, execute once:\n  # Linux:\n  $ %[1]s completion bash > /etc/bash_completion.d/%[1]s\n  # macOS:\n  $ %[1]s completion bash > $(brew --prefix)/etc/bash_completion.d/%[1]s\n\nZsh:\n\n  # If shell completion is not already enabled in your environment,\n  # you will need to enable it. You can execute the following once:\n\n  $ echo \"autoload -U compinit; compinit\" >> ~/.zshrc\n\n  # To load completions for each session, execute once:\n  $ %[1]s completion zsh > \"${fpath[1]}/_%[1]s\"\n\n  # You will need to start a new shell for this setup to take effect\n\nfish:\n\n  $ %[1]s completion fish | source\n\n  # To load completions for each session, execute once:\n  $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish\n\nPowerShell:\n\n  PS> %[1]s completion powershell | Out-String | Invoke-Expression\n\n  # To load completions for every new session, run:\n  PS> %[1]s completion powershell > %[1]s.ps1\n  # and source this file from your PowerShell profile",
    "valid_args": [
      "bash",
      "zsh",
      "fish",
      "powershell"
    ]
  },
  "root": {
    "use": "XXX_CLI_LOWERCASE_XXX",
    "long": "",
    "example_usages": [
      [
        "help"
      ]
    ]
  }
}
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const CMD_COMPLETION_GO: &str = r#"
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func CompletionCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("completion", runCompletion,
		domovoi.WithArgs(cobra.ExactArgs(1)),
		domovoi.WithValidArgs([]string{"bash", "zsh", "fish", "powershell"}),
	))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCompletion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		horus.CheckErr(cmd.Root().GenBashCompletion(os.Stdout))
	case "zsh":
		horus.CheckErr(cmd.Root().GenZshCompletion(os.Stdout))
	case "fish":
		horus.CheckErr(cmd.Root().GenFishCompletion(os.Stdout, true))
	case "powershell":
		horus.CheckErr(cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const CMD_IDENTITY_GO: &str = r#"
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func IdentityCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("identity", runIdentity,
		domovoi.WithAliases([]string{"id"}),
	))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const IDENT = ``

////////////////////////////////////////////////////////////////////////////////////////////////////

func runIdentity(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Println(IDENT)
	fmt.Println()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const CMD_GO_TEMPLATE: &str = r#"
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func XXX_CMD_UPPERCASE_XXXCmd() *cobra.Command {
	d := horus.Must(domovoi.GlobalDocs())
	cmd := horus.Must(d.MakeCmd("XXX_CMD_LOWERCASE_XXX", runXXX_CMD_UPPERCASE_XXX))

	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runXXX_CMD_UPPERCASE_XXX(cmd *cobra.Command, args []string) {
	op := "XXX_CLI_LOWERCASE_XXX.XXX_CMD_LOWERCASE_XXX"
}

////////////////////////////////////////////////////////////////////////////////////////////////////
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const JUST_HEAD: &str = r#"
####################################################################################################

_default:
  @just --choose

####################################################################################################

@abort:

####################################################################################################

@list:
  just --list

####################################################################################################

@show:
  bat .justfile --language make

####################################################################################################

@edit:
  micro .justfile

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const JUST_GO: &str = r#"
# config
####################################################################################################

app := 'XXX_APP_XXX'
exe := 'XXX_EXE_XXX'

####################################################################################################
# jobs
####################################################################################################

# build exec
build exe=exe:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  go build -v -o excalibur/{{exe}}

####################################################################################################

# install locally
install app=app exe=exe:
  @echo "\n\033[1;33mInstalling\033[0;37m...\n=================================================="
  go install
  mv -v "${HOME}/go/bin/{{app}}" "${HOME}/go/bin/{{exe}}"
  "${HOME}/go/bin/{{exe}}" completion zsh > "${HOME}/.config/zsh_completion/_{{exe}}"

####################################################################################################

# watch
watch:
  watchexec --clear --watch cmd -- 'just install'

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const JUST_RS: &str = r#"
# jobs
####################################################################################################

# build exec
@build:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  cargo clean; cargo build --release

####################################################################################################

# install locally
install:
  @echo "\n\033[1;33mInstalling\033[0;37m...\n=================================================="
  cargo install --path .

####################################################################################################

# watch
@watch:
  cargo watch --clear --why --exec check

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const README_MD: &str = r#"
# XXX_REPO_XXX

[![License](https://img.shields.io/badge/license-LICENSETYPE-blue.svg)](LICENSE)

## Overview


## Installation


## License
Copyright (c) XXX_YEAR_XXX

See the [LICENSE](LICENSE) file for license details
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const TODOR: &str = r#"
{
  // tags to search for. These are case-insensitive
  "tags": [
    "bug",
    "doc",
    "todo",
    "wip",
  ],

  "styles": {
    "tags": {
        "bug": "red",
        "doc": "blue",
        "todo": "yellow",
        "wip": "cyan",
    }
  }
}
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const AVICENNA_ROOT_JL: &str = r#"
####################################################################################################

module XXX_ROOT2_XXX

####################################################################################################

include("util/XXX_ROOT2_LOWERCASE_XXXutil.jl")
include("flow/XXX_ROOT2_LOWERCASE_XXXflow.jl")
include("inter/cli/XXX_ROOT2_LOWERCASE_XXXcli.jl")
include("inter/repl/XXX_ROOT2_LOWERCASE_XXXrepl.jl")

####################################################################################################

export XXX_ROOT2_XXXCore, XXX_ROOT2_XXXFlow, XXX_ROOT2_XXXCLI, XXX_ROOT2_XXXREPL

####################################################################################################

end

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const AVICENNA_UTIL_JL: &str = r#"
####################################################################################################

module XXX_ROOT2_XXXCore

####################################################################################################

# using

####################################################################################################

# export

####################################################################################################

# function declaration

####################################################################################################

end

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const AVICENNA_FLOW_JL: &str = r#"
####################################################################################################

module XXX_ROOT2_XXXFlow

####################################################################################################

using Avicenna.Flow: Stage, Config
using ..XXX_ROOT2_XXXCore

####################################################################################################

export flow

####################################################################################################

const flow = Config(
  "",
  [
    Stage("01", (config, _) -> XXX_ROOT2_XXXCore.func(), "1.0"),
    Stage("02", (config, prev) -> XXX_ROOT2_XXXCore.func(), "1.0"),
    Stage("03", (config, prev) -> XXX_ROOT2_XXXCore.func(), "1.0"),
  ],
  "1.0",
)

####################################################################################################

end

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const AVICENNA_CLI_JL: &str = r#"
####################################################################################################

module XXX_ROOT2_XXXCLI

####################################################################################################

using ArgParse
using Avicenna.Flow: Cache, launch
using ..XXX_ROOT2_XXXFlow: flow

####################################################################################################

export run

####################################################################################################

function run(args::Vector{String})
  s = ArgParseSettings()
  @add_arg_table! s begin
    "--no-cache"
    help = "Disable caching"
    action = :store_true
    "--verbose"
    help = "Enable verbose diagnostics"
    action = :store_false
  end
  parsed = parse_args(args, s)

  config = Dict{String,Any}()

  cache = Cache("cache/XXX_MODULE_LOWERCASE_XXX", !parsed["no-cache"])
  result = launch(flow, config, cache = cache)

  if parsed["verbose"]
  end
  return 0
end

####################################################################################################

end

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub const AVICENNA_REPL_JL: &str = r#"
####################################################################################################

module XXX_ROOT2_XXXREPL

####################################################################################################

using Avicenna.Flow: Cache, launch
using ..XXX_ROOT2_XXXFlow: flow

####################################################################################################

export run

####################################################################################################

function run()
  config = Dict{String,Any}()
  cache = Cache("cache/XXX_MODULE_LOWERCASE_XXX", !no_cache)
  return launch(flow, config, cache = cache)
end

####################################################################################################

end

####################################################################################################
"#;

////////////////////////////////////////////////////////////////////////////////////////////////////
