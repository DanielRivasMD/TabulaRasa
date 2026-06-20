////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;
use clap::Parser;

////////////////////////////////////////////////////////////////////////////////////////////////////

use cli::{Cli, Command};
use cmd::cobra;
use cmd::completion;
use cmd::deploy;
use cmd::etch;
use cmd::identity;

////////////////////////////////////////////////////////////////////////////////////////////////////

mod cli;
mod cmd;
mod forge;
mod skeleton;
mod util;

////////////////////////////////////////////////////////////////////////////////////////////////////

fn main() -> anyResult<()> {
    let cli = Cli::parse();
    match cli.command {
        Command::Cobra {
            sub,
            user,
            author,
            email,
        } => cobra::run(sub, &user, &author, &email, cli.verbose)?,
        Command::Deploy { lang, sub } => deploy::run(lang.as_deref(), sub, cli.verbose)?,
        Command::Etch => etch::run(cli.verbose)?,
        Command::Identity => identity::run()?,
        Command::Completion { shell } => completion::run(shell)?,
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
