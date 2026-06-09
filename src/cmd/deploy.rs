////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;

////////////////////////////////////////////////////////////////////////////////////////////////////

use crate::cli;
use crate::cmd::deploy;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub fn run(lang: Option<&str>, sub: Option<cli::DeploySub>, verbose: bool) -> anyResult<()> {
    match sub {
        Some(cli::DeploySub::Avicenna { module }) => deploy::avicenna::run(&module, verbose)?,
        Some(cli::DeploySub::Just) => deploy::just::run(lang, verbose)?,
        Some(cli::DeploySub::Readme) => deploy::readme::run(verbose)?,
        Some(cli::DeploySub::Todor) => deploy::todor::run(verbose)?,
        None => {
            deploy::just::run(lang, verbose)?;
            deploy::readme::run(verbose)?;
            deploy::todor::run(verbose)?;
        }
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

mod avicenna {
    use anyhow::Result as anyResult;

    use crate::forge;
    use crate::skeleton;
    use crate::util;

    pub fn run(module: &str, verbose: bool) -> anyResult<()> {
        let letter_upper = util::two_letter_from_module(module)?;
        let letter_lower = letter_upper.to_lowercase();
        let mod_lower = module.to_lowercase();
        let replacements = vec![
            forge::Replacement::token("XXX_MODULE_LOWERCASE_XXX", &mod_lower),
            forge::Replacement::token("XXX_ROOT2_UPPERCASE_XXX", &letter_upper),
            forge::Replacement::token("XXX_ROOT2_LOWERCASE_XXX", &letter_lower),
        ];

        let root_jl = format!("{module}.jl");
        let util_jl = format!("{letter_lower}util.jl");
        let flow_jl = format!("{letter_lower}flow.jl");
        let cli_jl = format!("{letter_lower}cli.jl");
        let repl_jl = format!("{letter_lower}repl.jl");

        let targets: Vec<(&str, &str, &str)> = vec![
            ("src", &root_jl, skeleton::avicenna::ROOT),
            ("src/util", &util_jl, skeleton::avicenna::UTIL),
            ("src/flow", &flow_jl, skeleton::avicenna::FLOW),
            ("src/inter/cli", &cli_jl, skeleton::avicenna::CLI),
            ("src/inter/repl", &repl_jl, skeleton::avicenna::REPL),
        ];
        for (subdir, filename, content) in &targets {
            let out_path = format!("{subdir}/{filename}");
            forge::forge_files(&out_path, &[(*filename, *content)], &replacements, verbose)?;
        }
        Ok(())
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////

mod just {
    use anyhow::Result as anyResult;

    use crate::forge;
    use crate::skeleton;
    use crate::util;

    pub fn run(lang: Option<&str>, verbose: bool) -> anyResult<()> {
        let repo = util::current_dir_name()?;
        let lower_repo = repo.to_lowercase();
        let mut files = vec![("head.just", skeleton::just::HEAD)];

        if let Some(lang_str) = lang {
            match lang_str {
                "go" => files.push(("go.just", skeleton::just::GO)),
                "rs" => files.push(("rs.just", skeleton::just::RS)),
                _ => anyhow::bail!("unsupported language: {lang_str}"),
            }
        }

        let replacements = vec![
            forge::Replacement::token("XXX_APP_XXX", &repo),
            forge::Replacement::token("XXX_EXE_XXX", &lower_repo),
        ];
        forge::forge_files(".justfile", &files, &replacements, verbose)
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////

mod readme {
    use anyhow::Result as anyResult;

    use crate::forge;
    use crate::skeleton;
    use crate::util;

    pub fn run(verbose: bool) -> anyResult<()> {
        let repo = util::current_dir_name()?;
        let year = "2026";
        let replacements = vec![
            forge::Replacement::token("XXX_REPO_XXX", &repo),
            forge::Replacement::token("XXX_YEAR_XXX", year),
        ];
        forge::forge_files(
            "README.md",
            &[("readme.md", skeleton::readme::MD)],
            &replacements,
            verbose,
        )
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////

mod todor {
    use anyhow::Result as anyResult;

    use crate::forge;
    use crate::skeleton;

    pub fn run(verbose: bool) -> anyResult<()> {
        forge::forge_files(".todor", &[("todor", skeleton::todor::TODOR)], &[], verbose)
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////
