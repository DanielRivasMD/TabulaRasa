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

pub fn lang_flag(lang: Option<&str>) -> &str {
    lang.unwrap_or("go") // default "go" matches original behaviour
}

////////////////////////////////////////////////////////////////////////////////////////////////////
