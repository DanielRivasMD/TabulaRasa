
```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              M B O M B O  (CLI)                             │
│                          Daniel Rivas <danielrivasmd@gmail.com>             │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                    ┌───────────────┴───────────────┐
                    ▼                               ▼
┌───────────────────────────────┐     ┌───────────────────────────────┐
│          main.go              │     │          cmd/                 │
│───────────────────────────────│     │───────────────────────────────│
│ func main() {                 │     │  Package containing all       │
│     cmd.Execute()             │────▶│  command logic and            │
│ }                             │     │  documentation                │
└───────────────────────────────┘     └───────────────────────────────┘
                                                                           

┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              cmd/ PACKAGE STRUCTURE                                 │
├─────────────────────────────────────────────────────────────────────────────────────┤
│                                                                                     │
│  ┌─────────────────────┐      ┌─────────────────────┐      ┌─────────────────────┐ │
│  │     root.go         │      │   utilDocs.go       │      │    docs.json        │ │
│  │─────────────────────│      │─────────────────────│      │─────────────────────│ │
│  │ • rootCmd           │─────▶│ • Docs struct       │◀────│ • Command docs      │ │
│  │ • Execute()         │      │ • MakeCmd()         │  ☁  │ • Examples          │ │
│  │ • flags.verbose     │◀────│ • CommandOpt funcs  │embed│ • Help text         │ │
│  └─────────────────────┘      └─────────────────────┘      └─────────────────────┘ │
│           │                           │                                              │
│           │                           │                                              │
│           ▼                           ▼                                              │
│  ┌─────────────────────┐      ┌─────────────────────┐      ┌─────────────────────┐ │
│  │  cmdIdentity.go     │      │   cmdForge.go       │      │ cmdCompletion.go    │ │
│  │─────────────────────│      │─────────────────────│      │─────────────────────│ │
│  │ MakeCmd("identity") │      │ MakeCmd("forge")    │      │ MakeCmd("completion")│ │
│  │ runIdentity()       │      │ runForge()          │      │ runCompletion()     │ │
│  └─────────────────────┘      └─────────────────────┘      └─────────────────────┘ │
│                                                                                     │
└─────────────────────────────────────────────────────────────────────────────────────┘


┌───────────────────────────────── DOCS STRUCTURE ─────────────────────────────────┐
│                                                                                   │
│  ┌─────────────────────────────────────────────────────────────────────────────┐ │
│  │ type Docs struct {                                                          │ │
│  │     once    sync.Once                    // Ensures single initialization │ │
│  │     entries map[string]DocEntry          // Raw docs from JSON            │ │
│  │     example map[string]string             // Formatted examples           │ │
│  │     help    map[string]string             // Styled help text             │ │
│  │     short   map[string]string             // Short descriptions           │ │
│  │     use     map[string]string             // Usage strings                │ │
│  │     loadErr error                          // Loading errors             │ │
│  │ }                                                                           │ │
│  └─────────────────────────────────────────────────────────────────────────────┘ │
│                                       │                                           │
│                                       ▼                                           │
│  ┌─────────────────────────────────────────────────────────────────────────────┐ │
│  │ type DocEntry struct {                                                      │ │
│  │     Use                   string                                            │ │
│  │     Aliases               []string                                          │ │
│  │     Hidden                bool                                              │ │
│  │     Short                 string                                            │ │
│  │     Long                  string                                            │ │
│  │     ExampleUsages         [][]string   // [["forge", "--help"]]           │ │
│  │     ValidArgs             []string                                          │ │
│  │     DisableFlagsInUseLine bool                                              │ │
│  │ }                                                                           │ │
│  └─────────────────────────────────────────────────────────────────────────────┘ │
│                                                                                   │
└───────────────────────────────────────────────────────────────────────────────────┘


┌────────────────────────── COMMAND FACTORY PATTERN ───────────────────────────┐
│                                                                               │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │ func MakeCmd(key string, run func(*cobra.Command, []string),            │ │
│  │           opts ...CommandOpt) *cobra.Command {                          │ │
│  │     d := getGlobalDocs()                                                │ │
│  │     entry, exists := d.GetEntry(key)                                    │ │
│  │                                                                          │ │
│  │     cmd := &cobra.Command{                                              │ │
│  │         Use:     d.GetUse(key),                                         │ │
│  │         Short:   d.GetShort(key),                                       │ │
│  │         Long:    formatLongHelp(d.GetHelp(key)),  // ← colored!        │ │
│  │         Example: d.GetExample(key),                                     │ │
│  │         Aliases: entry.Aliases,                                         │ │
│  │         Hidden:  entry.Hidden,                                          │ │
│  │         Run:     run,                                                    │ │
│  │     }                                                                    │ │
│  │                                                                          │ │
│  │     for _, opt := range opts { opt(cmd) }  // Apply options            │ │
│  │     return cmd                                                          │ │
│  │ }                                                                        │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘


┌───────────────────────────── FORGE COMMAND FLOW ─────────────────────────────┐
│                                                                               │
│  ┌─────────────┐     ┌─────────────────┐     ┌─────────────────────────┐     │
│  │   User runs │────▶│  runForge()     │────▶│  normalizeForgeOptions()│     │
│  │ mbombo forge│     │  in cmdForge.go │     │  • Parse paths          │     │
│  └─────────────┘     └─────────────────┘     │  • Normalize input/output│     │
│                            │                  └─────────────────────────┘     │
│                            ▼                                                  │
│                     ┌─────────────────┐                                      │
│                     │   catFiles()    │                                      │
│                     │─────────────────│                                      │
│                     │ 1. Check overwrite                                    │
│                     │ 2. Prepare source files                               │
│                     │ 3. Clean output if exists                             │
│                     │ 4. For each file:                                     │
│                     │    • Read content                                     │
│                     │    • applyReplacements()                              │
│                     │    • Write to output                                  │
│                     └─────────────────┘                                      │
│                            │                                                  │
│                            ▼                                                  │
│                     ┌─────────────────┐     ┌─────────────────────────┐     │
│                     │applyReplacements│────▶│  • Token replacement    │     │
│                     │                 │     │  • Line replacement     │     │
│                     └─────────────────┘     │  • Multiple pairs       │     │
│                                             └─────────────────────────┘     │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘


┌───────────────────────────── STYLING FUNCTIONS ──────────────────────────────┐
│                                                                               │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │ func authorHeader() string {                                            │ │
│  │     return chalk.Bold.TextStyle(                                        │ │
│  │         chalk.Green.Color("Daniel Rivas")) + " " +                      │ │
│  │         chalk.Dim.TextStyle(                                            │ │
│  │             chalk.Italic.TextStyle("<danielrivasmd@gmail.com>"))        │ │
│  │ }                                                                        │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
│                                    │                                          │
│                                    ▼                                          │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │ func styleDescription(text string) string {                             │ │
│  │     return chalk.Cyan.Color(chalk.Dim.TextStyle(text))                  │ │
│  │ }                        └──┬──┘                                        │ │
│  │                              │                                           │ │
│  │                              ▼                                           │ │
│  │                    "Forge by defining..." ← CYAN + DIM                  │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
│                                    │                                          │
│                                    ▼                                          │
│  ┌─────────────────────────────────────────────────────────────────────────┐ │
│  │ func styleLongHelp(text string) string {                                │ │
│  │     lines := strings.Split(text, "\n")                                  │ │
│  │     for i, line := range lines {                                        │ │
│  │         switch {                                                         │ │
│  │         case strings.HasPrefix(trimmed, "$"):                           │ │
│  │             lines[i] = chalk.White.Color(line)  // Shell commands      │ │
│  │         case strings.HasPrefix(trimmed, "#"):                           │ │
│  │             lines[i] = chalk.Dim.TextStyle(                             │ │
│  │                 chalk.Cyan.Color(line))        // Comments            │ │
│  │         }                                                                │ │
│  │     }                                                                    │ │
│  │ }                                                                        │ │
│  └─────────────────────────────────────────────────────────────────────────┘ │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘


┌───────────────────────────── DATA FLOW ─────────────────────────────┐
│                                                                       │
│    docs.json ──embed──> docsFS ──load──> Docs.entries               │
│                      ☁                       │                       │
│                                              ▼                       │
│    GetUse(key) ◄────── use map <──────┐     │                       │
│    GetShort(key) ◄───── short map <────┤     │                       │
│    GetHelp(key) ◄────── help map <─────┼─────┘                       │
│    GetExample(key) ◄─── example map <──┘     formatHelp()           │
│                                                styleLongHelp()        │
│                                                domovoi.FormatExample()│
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘


┌───────────────────────────── INITIALIZATION ──────────────────────────────┐
│                                                                            │
│  ┌──────────────────────────────────────────────────────────────────────┐ │
│  │ func getGlobalDocs() *Docs {                                         │ │
│  │     once.Do(func() {                                                 │ │
│  │         globalDocs = &Docs{}                                         │ │
│  │         globalDocs.load()   // Read & parse JSON                    │ │
│  │     })                                                               │ │
│  │     return globalDocs                                                │ │
│  │ }                                                                    │ │
│  └──────────────────────────────────────────────────────────────────────┘ │
│                                    │                                       │
│                                    ▼                                       │
│  ┌──────────────────────────────────────────────────────────────────────┐ │
│  │ When a command runs:                                                 │ │
│  │ 1. MakeCmd() calls getGlobalDocs()                                   │ │
│  │ 2. First call triggers load() via sync.Once                          │ │
│  │ 3. All subsequent calls use cached docs                              │ │
│  └──────────────────────────────────────────────────────────────────────┘ │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘
```
