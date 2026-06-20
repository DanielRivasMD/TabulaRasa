////////////////////////////////////////////////////////////////////////////////////////////////////

use clap::{Parser, Subcommand, ValueEnum};

////////////////////////////////////////////////////////////////////////////////////////////////////

const HELP: &str = r"Blank slate deployment";

////////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Parser)]
#[command(
    name = env!("CARGO_BIN_NAME"),
    version = env!("CARGO_PKG_VERSION"),
    author = env!("CARGO_PKG_AUTHORS"),
    about = env!("CARGO_PKG_DESCRIPTION"),
    before_help = concat!(env!("CARGO_PKG_AUTHORS"), "\n", env!("CARGO_PKG_NAME"), " v", env!("CARGO_PKG_VERSION")),
    long_about = HELP,
)]
pub struct Cli {
    #[command(subcommand)]
    pub command: Command,

    /// Enable verbose output
    #[arg(short, long, global = true)]
    pub verbose: bool,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Subcommand)]
pub enum Command {
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
        // TODO: add option completion
        /// Templates to deploy (go, rs)
        #[arg(short, long)]
        lang: Option<String>,
        #[command(subcommand)]
        sub: Option<DeploySub>,
    },

    /// Initialize configuration directories
    Etch,

    /// Print identity
    #[command(hide = true)]
    #[command(aliases = &["id"])]
    Identity,

    /// Generate shell completions
    #[command(hide = true)]
    Completion {
        /// Shell for which to generate completions
        #[arg(value_enum)]
        shell: Shell,
    },
}

////////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Subcommand)]
pub enum CobraSub {
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

// TODO: create a struct DeployOpts to hold lang flag
#[derive(Subcommand)]
pub enum DeploySub {
    /// Scientific analysis framework
    Avicenna {
        // TODO: add graceful error out
        /// Module name
        #[arg(long, default_value = "")]
        module: String,
    },

    /// Build system files
    Just,

    /// README scaffold
    Readme,

    /// Task‑tracker config
    Todor,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Clone, Copy, ValueEnum)]
pub enum Shell {
    Bash,
    Zsh,
    Fish,
    Powershell,
}

////////////////////////////////////////////////////////////////////////////////////////////////////
