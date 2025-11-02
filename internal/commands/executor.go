package commands

import (
	"fmt"

	"github.com/sicktiergit/sicktiergit/internal/config"
)

// Executor handles command execution
type Executor struct {
	config   *config.Config
	commands map[string]Command
}

// Command represents a sicktiergit command
type Command interface {
	Execute(args []string) error
	Description() string
}

// NewExecutor creates a new command executor
func NewExecutor(cfg *config.Config) *Executor {
	e := &Executor{
		config:   cfg,
		commands: make(map[string]Command),
	}

	// Register all commands
	e.registerCommands()

	return e
}

func (e *Executor) registerCommands() {
	e.commands["init"] = NewInitCommand(e.config)
	e.commands["commit"] = NewCommitCommand(e.config)
	e.commands["push"] = NewPushCommand(e.config)
	e.commands["pull"] = NewPullCommand(e.config)
	e.commands["status"] = NewStatusCommand(e.config)
	e.commands["log"] = NewLogCommand(e.config)
	e.commands["branch"] = NewBranchCommand(e.config)
	e.commands["checkout"] = NewCheckoutCommand(e.config)
	e.commands["merge"] = NewMergeCommand(e.config)
	e.commands["rebase"] = NewRebaseCommand(e.config)
	e.commands["blame"] = NewBlameCommand(e.config)
	e.commands["diff"] = NewDiffCommand(e.config)
	e.commands["stash"] = NewStashCommand(e.config)
	e.commands["reset"] = NewResetCommand(e.config)
	e.commands["revert"] = NewRevertCommand(e.config)
	e.commands["cherry-pick"] = NewCherryPickCommand(e.config)
	e.commands["therapy"] = NewTherapyCommand(e.config)
	e.commands["undo"] = NewUndoCommand(e.config)
}

// Execute runs the specified command with given arguments
func (e *Executor) Execute(cmdName string, args []string) error {
	cmd, exists := e.commands[cmdName]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	return cmd.Execute(args)
}
