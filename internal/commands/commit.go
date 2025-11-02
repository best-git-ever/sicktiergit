package commands

import (
	"fmt"
	"time"

	"github.com/sicktiergit/sicktiergit/internal/config"
	"github.com/sicktiergit/sicktiergit/internal/ui"
)

// CommitCommand records changes to the repository
type CommitCommand struct {
	config *config.Config
}

// NewCommitCommand creates a new commit command
func NewCommitCommand(cfg *config.Config) *CommitCommand {
	return &CommitCommand{config: cfg}
}

// Execute runs the commit command
func (c *CommitCommand) Execute(args []string) error {
	message := c.getCommitMessage(args)

	ui.PrintInfo("💾 Preparing to commit changes...")
	ui.ShowProgress(1)

	// Generate a fake commit hash
	hash := ui.GenerateHash(8)

	ui.PrintSuccess(fmt.Sprintf("[%s %s] %s", c.config.GetDefaultBranch(), hash, message))

	// Random statistics
	files := c.config.RandomSeed.Intn(10) + 1
	insertions := c.config.RandomSeed.Intn(200) + 10
	deletions := c.config.RandomSeed.Intn(100) + 5

	fmt.Printf(" %d file(s) changed, %d insertion(s)(+), %d deletion(s)(-)\n",
		files, insertions, deletions)

	time.Sleep(300 * time.Millisecond)

	// Philosophical messages based on chaos level
	if c.config.GetChaosLevel() > 7 {
		ui.PrintWarning("⚠  Warning: This commit may or may not exist in the future.")
	} else if c.config.GetChaosLevel() > 4 {
		ui.PrintDim("Commit successful. Nothing was actually saved, but you feel better.")
	} else {
		ui.PrintDim("Commit saved. Your mistakes are now permanent.")
	}

	return nil
}

func (c *CommitCommand) getCommitMessage(args []string) string {
	for i, arg := range args {
		if arg == "-m" && i+1 < len(args) {
			return args[i+1]
		}
	}

	// Default philosophical commit messages
	messages := []string{
		"Updated things until it worked",
		"Fixed the thing",
		"I have no idea what I'm doing",
		"This definitely works now",
		"Please work this time",
		"Added code, removed hope",
		"Trust me, I'm an engineer",
		"Existential commit",
	}

	return messages[c.config.RandomSeed.Intn(len(messages))]
}

// Description returns the command description
func (c *CommitCommand) Description() string {
	return "Record changes to the repository"
}
