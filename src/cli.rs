////////////////////////////////////////////////////////////////////////////////////////////////////

use clap::{Parser, Subcommand, ValueEnum};

////////////////////////////////////////////////////////////////////////////////////////////////////

#[derive(Parser)]
#[command(name = "tab", version, about = "Blank slate deployment")]
pub struct Cli {
    #[command(subcommand)]
    pub command: Commands,

    /// Enable verbose output
    #[arg(short, long, global = true)]
    pub verbose: bool,
}

#[derive(Subcommand)]
pub enum Commands {
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
    Completion {
        /// Shell for which to generate completions
        #[arg(value_enum)]
        shell: Shell,
    },
}

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

#[derive(Subcommand)]
pub enum DeploySub {
    /// Scientific analysis framework
    Avicenna {
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
