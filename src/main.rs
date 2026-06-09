////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;
use clap::Parser;
use cli::Cli;
use cmd::{cobra, completion, deploy, etch, identity};

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
        cli::Commands::Cobra {
            sub,
            user,
            author,
            email,
        } => cobra::run(sub, &user, &author, &email, cli.verbose)?,
        cli::Commands::Deploy { lang, sub } => deploy::run(lang.as_deref(), sub, cli.verbose)?,
        cli::Commands::Etch => etch::run(cli.verbose)?,
        cli::Commands::Identity => identity::run()?,
        cli::Commands::Completion { shell } => completion::run(shell)?,
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
