package commands

import (
	"fmt"
	"time"

	"github.com/sicktiergit/sicktiergit/internal/config"
	"github.com/sicktiergit/sicktiergit/internal/ui"
)

// BlameCommand shows what revision and author last modified each line
type BlameCommand struct {
	config *config.Config
}

// NewBlameCommand creates a new blame command
func NewBlameCommand(cfg *config.Config) *BlameCommand {
	return &BlameCommand{config: cfg}
}

// Execute runs the blame command
func (c *BlameCommand) Execute(args []string) error {
	target := "teammate"
	allFlag := c.hasFlag(args, "--all")

	if len(args) > 0 && args[0] != "--all" {
		target = args[0]
	}

	ui.PrintInfo(fmt.Sprintf("📧 Analyzing blame patterns for: %s", target))
	ui.ShowProgress(2)

	// AI-Assisted Blame Analysis
	ui.PrintInfo("\n🤖 AI-Assisted Blame Analysis™")
	fmt.Println("──────────────────────────────────────")

	time.Sleep(300 * time.Millisecond)

	commits := []string{
		"a3f8b2c",
		"7d91e4f",
		"2b5c8a1",
		"f4e7d9b",
		"9c2a1f5",
	}

	authors := []string{
		"Bob (never tests)",
		"Alice (copies from StackOverflow)",
		"Charlie (commits directly to main)",
		"Diana (force pushes at 2am)",
		"Eve (uses tabs instead of spaces)",
	}

	if allFlag {
		ui.PrintWarning("⚠  Blaming everyone...")
		for i, commit := range commits {
			author := authors[i%len(authors)]
			fmt.Printf("  %s  %-35s [SUSPICIOUS]\n", commit, author)
		}
	} else {
		hash := commits[c.config.RandomSeed.Intn(len(commits))]
		fmt.Printf("  %s  %s [PRIMARY SUSPECT]\n", hash, target)
	}

	time.Sleep(500 * time.Millisecond)

	ui.PrintSuccess("\n✓ Blame analysis complete")
	ui.PrintDim("Automated email sent to HR and management.")
	ui.PrintDim("Evidence has been forwarded to team lead.")

	if allFlag {
		ui.PrintWarning("\n⚠  Warning: Everyone now knows about everyone else's mistakes.")
	}

	return nil
}

func (c *BlameCommand) hasFlag(args []string, flags ...string) bool {
	for _, arg := range args {
		for _, flag := range flags {
			if arg == flag {
				return true
			}
		}
	}
	return false
}

// Description returns the command description
func (c *BlameCommand) Description() string {
	return "Show what revision and author last modified each line"
}
