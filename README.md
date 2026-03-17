# TabulaRasa, template forge for rapid development

[![License](https://img.shields.io/badge/license-GPLv3-blue.svg)](LICENSE)

## Overview

`TabulaRasa` is a CLI tool that deploys predefined templates to scaffold new
projects or add components to existing ones
It generates Cobra CLI applications, `justfile` build systems, READMEs, and
task‑tracker configurations from skeleton files stored locally

All templates live under `~/.tabularasa/` and use a simple underscore‑naming
convention (e.g., `main_go` → `main.go`)
Placeholders like `XXX_REPO_XXX` are replaced with user‑supplied values

### Technical Architecture

TabulaRasa is a Go‑based CLI built with **Cobra**  
It relies on an external tool, **`mbombo forge`**, to concatenate template files
and perform token replacement

- Commands are defined in `cmd/` using the factory pattern
  (`func Command() *cobra.Command`)
- Template directories are resolved at runtime via `domovoi.FindHome()`
- Forging is delegated to a shell command; errors are handled uniformly with
  `horus`

### Logic Schematic

    ┌──────────────┐
    │ tab cobra    │ → constructs Cobra applications
    └──────┬───────┘
           │
           ▼
    ┌───────────────────────────────────────────┐
    │ tab cobra app [--force]                   │
    │ - creates cmd/ dir                        │
    │ - copies & replaces templates:            │
    │   • main.go                               │
    │   • cmd/root.go                           │
    │   • cmd/docs.json                         │
    │   • cmd/cmdCompletion.go                  │
    │   • cmd/cmdIdentity.go                    │
    │ - if --force, re‑initializes Go module    │
    └──────┬────────────────────────────────────┘
           │
           ▼
    ┌───────────────────────────────────────────┐
    │ tab cobra cmd <name>                      │
    │ - creates cmd/cmd<Name>.go from cmdCmd_go │
    │ - replaces XXX_CMD_... placeholders       │
    └───────────────────────────────────────────┘

    ┌──────────────┐
    │ tab deploy   │ → deploys config templates
    └──────┬───────┘
           │
           ▼
    ┌───────────────────────────────────────────┐
    │ tab deploy just [--lang go|rs]            │
    │ - writes .justfile (head.just + lang)     │
    │ - replaces XXX_APP_XXX, XXX_EXE_XXX       │
    └──────┬────────────────────────────────────┘
           │
           ▼
    ┌───────────────────────────────────────────┐
    │ tab deploy readme                         │
    │ - writes README.md from readme.md         │
    │ - replaces XXX_REPO_XXX, XXX_YEAR_XXX     │
    └──────┬────────────────────────────────────┘
           │
           ▼
    ┌───────────────────────────────────────────┐
    │ tab deploy todor                          │
    │ - writes .todor from todor template       │
    └───────────────────────────────────────────┘

### Storage Layout (`~/.tabularasa/`)

    ~/.tabularasa/
    ├─ cobra/      # templates for Cobra apps and commands
    │  ├─ main_go
    │  ├─ root_go
    │  ├─ docs_json
    │  ├─ cmdCompletion_go
    │  ├─ cmdIdentity_go
    │  └─ cmdCmd_go        (generic command template)
    ├─ just/       # justfile templates
    │  ├─ head.just
    │  ├─ go.just
    │  └─ rs.just
    ├─ readme/     # README template
    │  └─ readme.md
    └─ todor/      # todor template
       └─ todor

### Example Usage

```bash
# Create a new Cobra app
tab cobra app

# Add a command "serve" to the app
tab cobra cmd serve

# Deploy a justfile with Go support
tab deploy just --lang go

# Generate a README
tab deploy readme

# Deploy a task-tracker config
tab deploy todor

# Do everything at once (just, readme, todor)
tab deploy --lang go
```

## Installation

### Language-Specific

    Go: go install github.com/DanielRivasMD/TabulaRasa@latest

## License

Copyright (c) 2024

See the [LICENSE](LICENSE) file for license details
