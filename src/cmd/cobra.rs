////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;
use anyhow::Context;
use std::fs;
use std::path::Path;
use std::process::Command as StdCommand;

////////////////////////////////////////////////////////////////////////////////////////////////////

use crate::cli;
use crate::forge;
use crate::skeleton;
use crate::util;

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: modularize subcommands
pub fn run(
    sub: cli::CobraSub,
    user: &str,
    author: &str,
    email: &str,
    verbose: bool,
) -> anyResult<()> {
    match sub {
        cli::CobraSub::App { force } => {
            let repo = util::current_dir_name()?;
            let should_init = if Path::new("go.mod").exists() {
                if !force {
                    anyhow::bail!("a Go module already exists (go.mod). Use --force to overwrite");
                } else {
                    for f in ["go.mod", "go.sum"] {
                        let _ = fs::remove_file(f);
                    }
                    true
                }
            } else {
                true
            };

            let lower_repo = repo.to_lowercase();
            // For dynamic year, add `chrono` and use: chrono::Utc::now().year().to_string()
            let year = "2026";

            let replacements = vec![
                issac::Replacement::token("XXX_REPO_XXX", &repo),
                issac::Replacement::token("XXX_CLI_LOWERCASE_XXX", &lower_repo),
                issac::Replacement::token("XXX_AUTHOR_XXX", author),
                issac::Replacement::token("XXX_EMAIL_XXX", email),
                issac::Replacement::token("XXX_YEAR_XXX", year),
            ];

            fs::create_dir_all("cmd")?;

            forge::forge_files("main.go", &[("main.go", skeleton::cobra::MAIN)], &replacements, verbose)?;
            forge::forge_files(
                "cmd/root.go",
                &[("root.go", skeleton::cobra::ROOT)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/docs.json",
                &[("docs.json", skeleton::cobra::DOCS_JSON)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/cmdCompletion.go",
                &[("cmdCompletion.go", skeleton::cobra::CMD_COMPLETION)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/cmdIdentity.go",
                &[("cmdIdentity.go", skeleton::cobra::CMD_IDENTITY)],
                &replacements,
                verbose,
            )?;

            if should_init {
                StdCommand::new("go")
                    .args(["mod", "init", &format!("github.com/{user}/{repo}")])
                    .status()
                    .context("go mod init failed")?;
                StdCommand::new("go")
                    .args(["mod", "tidy"])
                    .status()
                    .context("go mod tidy failed")?;
            }
        }
        cli::CobraSub::Cmd { name } => {
            let repo = util::current_dir_name()?;
            let lower_repo = repo.to_lowercase();
            let cmd_lower = util::lower_first(&name);
            let cmd_upper = util::upper_first(&name);
            let year = "2026";
            let replacements = vec![
                issac::Replacement::token("XXX_CLI_LOWERCASE_XXX", &lower_repo),
                issac::Replacement::token("XXX_CMD_LOWERCASE_XXX", &cmd_lower),
                issac::Replacement::token("XXX_CMD_UPPERCASE_XXX", &cmd_upper),
                issac::Replacement::token("XXX_AUTHOR_XXX", author),
                issac::Replacement::token("XXX_EMAIL_XXX", email),
                issac::Replacement::token("XXX_YEAR_XXX", year),
            ];

            let out_file = format!("cmd/cmd{cmd_upper}.go");
            forge::forge_files(
                &out_file,
                &[("cmdCmd_go", skeleton::cobra::CMD_TEMPLATE)],
                &replacements,
                verbose,
            )?;
        }
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
