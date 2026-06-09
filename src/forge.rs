////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::{Context, Result as anyResult};
use regex::Regex;
use std::fs;
use std::io::{BufWriter, Write};
use std::path::Path;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub struct Replacement {
    pub old: String,
    pub new: String,
    pub mode: String, // "token" or "line"
}

////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////

/// Concatenate `file_contents` (each a (filename, content) pair),
/// apply `replacements`, and write the result to `out_path`.
/// Directories are created if needed.
pub fn forge_files(
    out_path: impl AsRef<Path>,
    files: &[(&str, &str)],
    replacements: &[Replacement],
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
