package main

import (
	"fmt"
	"os"

	"github.com/sicktiergit/sicktiergit/internal/commands"
	"github.com/sicktiergit/sicktiergit/internal/config"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "--version" || cmd == "-v" {
		fmt.Printf("sg (sicktiergit) version %s\n", version)
		return
	}

	if cmd == "--help" || cmd == "-h" {
		printHelp()
		return
	}

	executor := commands.NewExecutor(config.NewConfig())

	if err := executor.Execute(cmd, args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintln(os.Stderr, "Run 'sg --help' for usage information")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`
sg (sicktiergit) - A revolutionary version control system
──────────────────────────────────────────────────────────

USAGE:
    sg <command> [options]

COMMANDS:
    init          Initialize a new repository
    commit        Record changes to the repository
    push          Update remote refs along with associated objects
    pull          Fetch from and integrate with another repository
    status        Show the working tree status
    log           Show commit logs
    branch        List, create, or delete branches
    checkout      Switch branches or restore working tree files
    merge         Join two or more development histories together
    rebase        Reapply commits on top of another base tip
    blame         Show what revision and author last modified each line
    diff          Show changes between commits, commit and working tree, etc
    stash         Stash the changes in a dirty working directory
    reset         Reset current HEAD to the specified state
    revert        Revert some existing commits
    cherry-pick   Apply the changes introduced by some existing commits
    therapy       Open integrated emotional support session
    undo          Attempt to undo the last operation

OPTIONS:
    --version, -v    Show version information
    --help, -h       Show this help message

EXAMPLES:
    sg init my-project
    sg commit -m "Initial panic"
    sg push --force --maybe
    sg blame --all
    sg therapy --crisis

For more information, visit: https://sicktiergit.io/docs
`)
}
