////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::{Context, Result as anyResult};
use std::fs;
use std::io::{BufWriter, Write};
use std::path::Path;

////////////////////////////////////////////////////////////////////////////////////////////////////

/// Concatenate `file_contents` (each a (filename, content) pair),
/// apply `replacements`, and write the result to `out_path`.
/// Directories are created if needed.
pub fn forge_files(
    out_path: impl AsRef<Path>,
    files: &[(&str, &str)],
    replacements: &[issac::Replacement],
    verbose: bool,
) -> anyResult<()> {
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
        let processed = issac::apply_replacements(content, replacements);
        let trimmed = processed.trim_end_matches('\n');
        writeln!(writer, "{trimmed}").context("write failed")?;
    }

    writer.flush().context("flush failed")?;
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
