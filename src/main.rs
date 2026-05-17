use anyhow::{Context, Result};
use clap::{CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell};
use regex::Regex;
use std::fs;
use std::io::{BufWriter, Write};
use std::path::Path; // PathBuf removed from import
use std::process::Command as StdCommand;

// ---------------------------------------------------------------------------
// Forging module (absorbed from mbombo)
// ---------------------------------------------------------------------------
mod forge {
    use super::*;

    pub struct Replacement {
        pub old: String,
        pub new: String,
        pub mode: String, // "token" or "line"
    }

    impl Replacement {
        pub fn token(old: &str, new: &str) -> Self {
            Self {
                old: old.into(),
                new: new.into(),
                mode: "token".into(),
            }
        }
        pub fn line(old: &str, new: &str) -> Self {
            Self {
                old: old.into(),
                new: new.into(),
                mode: "line".into(),
            }
        }
    }

    /// Concatenate `file_contents` (each a (filename, content) pair),
    /// apply `replacements`, and write the result to `out_path`.
    /// Directories are created if needed.
    pub fn forge_files(
        out_path: impl AsRef<Path>,
        files: &[(&str, &str)],
        replacements: &[Replacement],
        verbose: bool,
    ) -> Result<()> {
        let out_path = out_path.as_ref();
        if let Some(parent) = out_path.parent() {
            fs::create_dir_all(parent)
                .with_context(|| format!("creating parent dirs for {}", out_path.display()))?;
        }

        let fwrite = fs::OpenOptions::new()
            .create(true)
            .write(true)
            .truncate(true)
            .open(out_path)
            .with_context(|| format!("opening out file {}", out_path.display()))?;
        let mut writer = BufWriter::new(fwrite);

        for (name, content) in files {
            if verbose {
                eprintln!("verbose: processing {}", name);
            }
            let processed = apply_replacements(content, replacements);
            let trimmed = processed.trim_end_matches('\n');
            writeln!(writer, "{trimmed}").context("write failed")?;
        }

        writer.flush().context("flush failed")?;
        Ok(())
    }

    fn apply_replacements(content: &str, replacements: &[Replacement]) -> String {
        let mut lines: Vec<String> = content.lines().map(|s| s.to_string()).collect();

        for rep in replacements {
            let re = if rep.mode == "line" {
                let pattern = format!(r"\b{}\b", regex::escape(&rep.old));
                Regex::new(&pattern).ok()
            } else {
                None
            };

            for line in &mut lines {
                match rep.mode.as_str() {
                    "line" => {
                        if let Some(ref regex) = re {
                            if regex.is_match(line) {
                                *line = rep.new.clone();
                            }
                        }
                    }
                    _ => {
                        *line = line.replace(&rep.old, &rep.new);
                    }
                }
            }
        }

        lines.join("\n")
    }
}

use forge::Replacement;

// ---------------------------------------------------------------------------
// Embedded template contents (replace with your actual templates)
// ---------------------------------------------------------------------------
const MAIN_GO: &str = r#"package main

func main() {
	// XXX_CLI_LOWERCASE_XXX
}
"#;

const ROOT_GO: &str = r#"package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "XXX_CLI_LOWERCASE_XXX",
	Short: "Auto-generated CLI",
}
"#;

const DOCS_JSON: &str = r#"{}"#;

const CMD_COMPLETION_GO: &str = r#"package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

func CompletionCmd() *cobra.Command {
	return &cobra.Command{Use: "completion"}
}
"#;

const CMD_IDENTITY_GO: &str = r#"package cmd

import "fmt"

func IdentityCmd() *cobra.Command {
	return &cobra.Command{Use: "identity", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("XXX_CLI_LOWERCASE_XXX")
	}}
}
"#;

const CMD_GO_TEMPLATE: &str = r#"package cmd

import "github.com/spf13/cobra"

func XXX_CMD_UPPERCASE_XXXCmd() *cobra.Command {
	return &cobra.Command{Use: "XXX_CMD_LOWERCASE_XXX"}
}
"#;

const JUST_HEAD: &str = r#"# Justfile head
"#;

const JUST_GO: &str = r#"
# Go targets
build:
	go build ./...
"#;

const JUST_RS: &str = r#"
# Rust targets
build:
	cargo build
"#;

const README_MD: &str = r#"# XXX_REPO_XXX

Copyright © XXX_YEAR_XXX
"#;

const TODOR: &str = r#"# todor config
"#;

const AVICENNA_ROOT_JL: &str = r#"# XXX_MODULE_LOWERCASE_XXX
"#;

const AVICENNA_UTIL_JL: &str = r#"# XXX_ROOT2_LOWERCASE_XXX util
"#;

const AVICENNA_FLOW_JL: &str = r#"# XXX_ROOT2_LOWERCASE_XXX flow
"#;

const AVICENNA_CLI_JL: &str = r#"# XXX_ROOT2_LOWERCASE_XXX cli
"#;

const AVICENNA_REPL_JL: &str = r#"# XXX_ROOT2_LOWERCASE_XXX repl
"#;

// ---------------------------------------------------------------------------
// Command‑specific helpers
// ---------------------------------------------------------------------------
fn current_dir_name() -> Result<String> {
    let dir = std::env::current_dir()?;
    dir.file_name()
        .and_then(|s| s.to_str())
        .map(|s| s.to_owned())
        .ok_or_else(|| anyhow::anyhow!("cannot determine current directory name"))
}

fn lower_first(s: &str) -> String {
    let mut c = s.chars();
    match c.next() {
        None => String::new(),
        Some(f) => f.to_lowercase().collect::<String>() + c.as_str(),
    }
}

fn upper_first(s: &str) -> String {
    let mut c = s.chars();
    match c.next() {
        None => String::new(),
        Some(f) => f.to_uppercase().collect::<String>() + c.as_str(),
    }
}

fn lang_flag(lang: Option<&str>) -> &str {
    lang.unwrap_or("go") // default "go" matches original behaviour
}

// ---------------------------------------------------------------------------
// CLI definition
// ---------------------------------------------------------------------------
#[derive(Parser)]
#[command(name = "tab", version, about = "Blank slate deployment")]
struct Cli {
    /// Enable verbose output
    #[arg(short, long, global = true)]
    verbose: bool,

    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Construct cobra applications, commands & import utilities
    Cobra {
        #[command(subcommand)]
        sub: CobraSub,
        /// GitHub username
        #[arg(long, default_value = "DanielRivasMD")]
        user: String,
        /// Author name
        #[arg(long, default_value = "Daniel Rivas")]
        author: String,
        /// Author email
        #[arg(long, default_value = "danielrivasmd@gmail.com")]
        email: String,
    },
    /// Deploy configuration templates
    Deploy {
        /// Templates to deploy (go, rs)
        #[arg(short, long)]
        lang: Option<String>,
        #[command(subcommand)]
        sub: Option<DeploySub>,
    },
    /// Initialize configuration directories
    Etch,
    /// Print identity
    Identity,
    /// Generate shell completions
    Completion { shell: String },
}

#[derive(Subcommand)]
enum CobraSub {
    /// Construct cobra application
    App {
        /// Force install Go dependencies
        #[arg(short, long)]
        force: bool,
    },
    /// Construct cobra command
    Cmd {
        /// Command name
        name: String,
    },
}

#[derive(Subcommand)]
enum DeploySub {
    /// Scientific analysis framework
    Avicenna {
        /// Module name
        #[arg(long, default_value = "")]
        module: String,
        /// Module two‑letter code
        #[arg(long, default_value = "")]
        letter: String,
    },
    /// Build system files
    Just,
    /// README scaffold
    Readme,
    /// Task‑tracker config
    Todor,
}

// ---------------------------------------------------------------------------
// Entry point
// ---------------------------------------------------------------------------
fn main() -> Result<()> {
    let cli = Cli::parse();
    match cli.command {
        Commands::Cobra {
            sub,
            user,
            author,
            email,
        } => run_cobra(sub, &user, &author, &email, cli.verbose)?,
        Commands::Deploy { lang, sub } => run_deploy(lang.as_deref(), sub, cli.verbose)?,
        Commands::Etch => run_etch(cli.verbose)?,
        Commands::Identity => println!("\nTabulaRasa\n"),
        Commands::Completion { shell } => {
            let shell = match shell.to_lowercase().as_str() {
                "bash" => Shell::Bash,
                "zsh" => Shell::Zsh,
                "fish" => Shell::Fish,
                "powershell" => Shell::PowerShell,
                _ => anyhow::bail!("unsupported shell: {shell}"),
            };
            let mut cmd = Cli::command();
            generate(shell, &mut cmd, "tab", &mut std::io::stdout());
        }
    }
    Ok(())
}

// ---------------------------------------------------------------------------
// Command implementations
// ---------------------------------------------------------------------------
fn run_cobra(sub: CobraSub, user: &str, author: &str, email: &str, verbose: bool) -> Result<()> {
    match sub {
        CobraSub::App { force } => {
            let repo = current_dir_name()?;
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
                Replacement::token("XXX_REPO_XXX", &repo),
                Replacement::token("XXX_CLI_LOWERCASE_XXX", &lower_repo),
                Replacement::token("XXX_AUTHOR_XXX", author),
                Replacement::token("XXX_EMAIL_XXX", email),
                Replacement::token("XXX_YEAR_XXX", year),
            ];

            fs::create_dir_all("cmd")?;

            forge::forge_files("main.go", &[("main.go", MAIN_GO)], &replacements, verbose)?;
            forge::forge_files(
                "cmd/root.go",
                &[("root.go", ROOT_GO)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/docs.json",
                &[("docs.json", DOCS_JSON)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/cmdCompletion.go",
                &[("cmdCompletion.go", CMD_COMPLETION_GO)],
                &replacements,
                verbose,
            )?;
            forge::forge_files(
                "cmd/cmdIdentity.go",
                &[("cmdIdentity.go", CMD_IDENTITY_GO)],
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
        CobraSub::Cmd { name } => {
            let repo = current_dir_name()?;
            let lower_repo = repo.to_lowercase();
            let cmd_lower = lower_first(&name);
            let cmd_upper = upper_first(&name);
            let year = "2026";
            let replacements = vec![
                Replacement::token("XXX_CLI_LOWERCASE_XXX", &lower_repo),
                Replacement::token("XXX_CMD_LOWERCASE_XXX", &cmd_lower),
                Replacement::token("XXX_CMD_UPPERCASE_XXX", &cmd_upper),
                Replacement::token("XXX_AUTHOR_XXX", author),
                Replacement::token("XXX_EMAIL_XXX", email),
                Replacement::token("XXX_YEAR_XXX", year),
            ];

            let out_file = format!("cmd/cmd{cmd_upper}.go");
            forge::forge_files(
                &out_file,
                &[("cmdCmd_go", CMD_GO_TEMPLATE)],
                &replacements,
                verbose,
            )?;
        }
    }
    Ok(())
}

fn run_deploy(lang: Option<&str>, sub: Option<DeploySub>, verbose: bool) -> Result<()> {
    match sub {
        Some(DeploySub::Avicenna { module, letter }) => {
            let two_letter = letter.to_lowercase();
            let mod_lower = module.to_lowercase();
            let replacements = vec![
                Replacement::token("XXX_MODULE_LOWERCASE_XXX", &mod_lower),
                Replacement::token("XXX_ROOT2_XXX", &letter),
                Replacement::token("XXX_ROOT2_LOWERCASE_XXX", &two_letter),
            ];

            // Owned strings so that references remain valid
            let root_jl = format!("{module}.jl");
            let util_jl = format!("{two_letter}util.jl");
            let flow_jl = format!("{two_letter}flow.jl");
            let cli_jl = format!("{two_letter}cli.jl");
            let repl_jl = format!("{two_letter}repl.jl");

            let targets: Vec<(&str, &str, &str)> = vec![
                ("src", &root_jl, AVICENNA_ROOT_JL),
                ("src/util", &util_jl, AVICENNA_UTIL_JL),
                ("src/flow", &flow_jl, AVICENNA_FLOW_JL),
                ("src/inter/cli", &cli_jl, AVICENNA_CLI_JL),
                ("src/inter/repl", &repl_jl, AVICENNA_REPL_JL),
            ];
            for (subdir, filename, content) in &targets {
                let out_path = format!("{subdir}/{filename}");
                forge::forge_files(&out_path, &[(*filename, *content)], &replacements, verbose)?;
            }
        }
        Some(DeploySub::Just) => deploy_just(lang_flag(lang), verbose)?,
        Some(DeploySub::Readme) => deploy_readme(verbose)?,
        Some(DeploySub::Todor) => deploy_todor(verbose)?,
        None => {
            deploy_just(lang_flag(lang), verbose)?;
            deploy_readme(verbose)?;
            deploy_todor(verbose)?;
        }
    }
    Ok(())
}

fn deploy_just(lang: &str, verbose: bool) -> Result<()> {
    let repo = current_dir_name()?;
    let lower_repo = repo.to_lowercase();
    let mut files = vec![("head.just", JUST_HEAD)];
    match lang {
        "go" => files.push(("go.just", JUST_GO)),
        "rs" => files.push(("rs.just", JUST_RS)),
        _ => anyhow::bail!("unsupported language: {lang}"),
    }

    let replacements = vec![
        Replacement::token("XXX_APP_XXX", &repo),
        Replacement::token("XXX_EXE_XXX", &lower_repo),
    ];
    forge::forge_files(".justfile", &files, &replacements, verbose)
}

fn deploy_readme(verbose: bool) -> Result<()> {
    let repo = current_dir_name()?;
    let year = "2026";
    let replacements = vec![
        Replacement::token("XXX_REPO_XXX", &repo),
        Replacement::token("XXX_YEAR_XXX", year),
    ];
    forge::forge_files(
        "README.md",
        &[("readme.md", README_MD)],
        &replacements,
        verbose,
    )
}

fn deploy_todor(verbose: bool) -> Result<()> {
    forge::forge_files(".todor", &[("todor", TODOR)], &[], verbose)
}

fn run_etch(verbose: bool) -> Result<()> {
    let home = dirs::home_dir().context("cannot determine home directory")?;
    let tabularasa = home.join(".tabularasa");
    let dirs = [
        ("tabularasa root", tabularasa.clone()),
        ("avicenna", tabularasa.join("avicenna")),
        ("cobra", tabularasa.join("cobra")),
        ("just", tabularasa.join("just")),
        ("readme", tabularasa.join("readme")),
        ("todor", tabularasa.join("todor")),
    ];
    for (label, path) in &dirs {
        if verbose {
            eprintln!("verbose: creating {label} at {}", path.display());
        }
        fs::create_dir_all(path).with_context(|| format!("creating {label}"))?;
    }
    Ok(())
}
