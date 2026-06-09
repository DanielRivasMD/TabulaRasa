////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;
use clap::CommandFactory;
use clap_complete::{generate, shells::*};
use std::io;

////////////////////////////////////////////////////////////////////////////////////////////////////

use crate::cli::{Cli, Shell};

////////////////////////////////////////////////////////////////////////////////////////////////////

pub fn run(shell: Shell) -> anyResult<()> {
    let mut cmd = Cli::command();
    let name = cmd.get_name().to_string();
    match shell {
        Shell::Bash => generate(Bash, &mut cmd, name, &mut io::stdout()),
        Shell::Zsh => generate(Zsh, &mut cmd, name, &mut io::stdout()),
        Shell::Fish => generate(Fish, &mut cmd, name, &mut io::stdout()),
        Shell::Powershell => generate(PowerShell, &mut cmd, name, &mut io::stdout()),
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
