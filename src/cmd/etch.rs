////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::{Context, Result as anyResult};
use std::fs;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub fn run(verbose: bool) -> anyResult<()> {
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

////////////////////////////////////////////////////////////////////////////////////////////////////
