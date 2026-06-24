////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;
use clap::Parser;

////////////////////////////////////////////////////////////////////////////////////////////////////

mod cli;
mod cmd;
mod forge;
mod skeleton;
mod util;

////////////////////////////////////////////////////////////////////////////////////////////////////

fn main() -> anyResult<()> {
    let cli = cli::Cli::parse();
    match cli.command {
        cli::Command::Cobra {
            sub,
            user,
            author,
            email,
        } => cmd::cobra::run(sub, &user, &author, &email, cli.verbose)?,
        cli::Command::Deploy { lang, sub } => cmd::deploy::run(lang.as_deref(), sub, cli.verbose)?,
        cli::Command::Etch => cmd::etch::run(cli.verbose)?,
        cli::Command::Identity => cmd::identity::run()?,
        cli::Command::Completion { shell } => cmd::completion::run(shell)?,
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
