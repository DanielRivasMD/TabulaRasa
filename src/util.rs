////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub fn current_dir_name() -> anyResult<String> {
    let dir = std::env::current_dir()?;
    dir.file_name()
        .and_then(|s| s.to_str())
        .map(|s| s.to_owned())
        .ok_or_else(|| anyhow::anyhow!("cannot determine current directory name"))
}

pub fn lower_first(s: &str) -> String {
    let mut c = s.chars();
    match c.next() {
        None => String::new(),
        Some(f) => f.to_lowercase().collect::<String>() + c.as_str(),
    }
}

pub fn upper_first(s: &str) -> String {
    let mut c = s.chars();
    match c.next() {
        None => String::new(),
        Some(f) => f.to_uppercase().collect::<String>() + c.as_str(),
    }
}

pub fn two_letter_from_module(module: &str) -> anyResult<String> {
    let caps: String = module.chars().filter(|c| c.is_ascii_uppercase()).collect();
    if caps.len() < 2 {
        anyhow::bail!(
            "Cannot derive a two‑letter code from module \"{module}\": \
             found {found} uppercase letter(s). Please provide --letter explicitly.",
            found = caps.len()
        );
    }
    Ok(caps[..2].to_string())
}

////////////////////////////////////////////////////////////////////////////////////////////////////
