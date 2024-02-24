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

import (
	"os"

	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// completionCmd
var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true,
	Short:  "Generate completion script.",
	Long: `To load completions:

Bash:

$ source <(COMMAND completion bash)

# To load completions for each session, execute once:
Linux:
 $ COMMAND completion bash > /etc/bash_completion.d/COMMAND
MacOS:
 $ COMMAND completion bash > /usr/local/etc/bash_completion.d/COMMAND

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ COMMAND completion zsh > "${fpath[1]}/_COMMAND"

# You will need to start a new shell for this setup to take effect.

Fish:

$ COMMAND completion fish | source

# To load completions for each session, execute once:
$ COMMAND completion fish > ~/.config/fish/completions/COMMAND.fish

Powershell:

PS> COMMAND completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> COMMAND completion powershell > COMMAND.ps1
# and source this file from your powershell profile.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.OnlyValidArgs,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			κ.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			κ.Root().GenZshCompletion(os.Stdout)
		case "fish":
			κ.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			κ.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(completionCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
