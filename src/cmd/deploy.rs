////////////////////////////////////////////////////////////////////////////////////////////////////

use anyhow::Result as anyResult;

////////////////////////////////////////////////////////////////////////////////////////////////////

use crate::cli;
use crate::cmd::deploy;
use crate::forge;
use crate::skeleton;
use crate::util;

////////////////////////////////////////////////////////////////////////////////////////////////////

pub fn run(lang: Option<&str>, sub: Option<cli::DeploySub>, verbose: bool) -> anyResult<()> {
    match sub {
        Some(cli::DeploySub::Avicenna { module, letter }) => {
            let two_letter = letter.to_lowercase();
            let mod_lower = module.to_lowercase();
            let replacements = vec![
                forge::Replacement::token("XXX_MODULE_LOWERCASE_XXX", &mod_lower),
                forge::Replacement::token("XXX_ROOT2_XXX", &letter),
                forge::Replacement::token("XXX_ROOT2_LOWERCASE_XXX", &two_letter),
            ];

            // Owned strings so that references remain valid
            let root_jl = format!("{module}.jl");
            let util_jl = format!("{two_letter}util.jl");
            let flow_jl = format!("{two_letter}flow.jl");
            let cli_jl = format!("{two_letter}cli.jl");
            let repl_jl = format!("{two_letter}repl.jl");

            let targets: Vec<(&str, &str, &str)> = vec![
                ("src", &root_jl, skeleton::AVICENNA_ROOT_JL),
                ("src/util", &util_jl, skeleton::AVICENNA_UTIL_JL),
                ("src/flow", &flow_jl, skeleton::AVICENNA_FLOW_JL),
                ("src/inter/cli", &cli_jl, skeleton::AVICENNA_CLI_JL),
                ("src/inter/repl", &repl_jl, skeleton::AVICENNA_REPL_JL),
            ];
            for (subdir, filename, content) in &targets {
                let out_path = format!("{subdir}/{filename}");
                forge::forge_files(&out_path, &[(*filename, *content)], &replacements, verbose)?;
            }
        }
        Some(cli::DeploySub::Just) => deploy::just::run(util::lang_flag(lang), verbose)?,
        Some(cli::DeploySub::Readme) => deploy::readme::run(verbose)?,
        Some(cli::DeploySub::Todor) => deploy::todor::run(verbose)?,
        None => {
            deploy::just::run(util::lang_flag(lang), verbose)?;
            deploy::readme::run(verbose)?;
            deploy::todor::run(verbose)?;
        }
    }
    Ok(())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

mod just {
    use anyhow::Result as anyResult;

    use crate::forge;
    use crate::skeleton;
    use crate::util;

    pub fn run(lang: &str, verbose: bool) -> anyResult<()> {
        let repo = util::current_dir_name()?;
        let lower_repo = repo.to_lowercase();
        let mut files = vec![("head.just", skeleton::JUST_HEAD)];
        match lang {
            "go" => files.push(("go.just", skeleton::JUST_GO)),
            "rs" => files.push(("rs.just", skeleton::JUST_RS)),
            _ => anyhow::bail!("unsupported language: {lang}"),
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
            &[("readme.md", skeleton::README_MD)],
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
        forge::forge_files(".todor", &[("todor", skeleton::TODOR)], &[], verbose)
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////////
