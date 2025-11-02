package commands

import (
	"fmt"
	"time"

	"github.com/sicktiergit/sicktiergit/internal/config"
	"github.com/sicktiergit/sicktiergit/internal/ui"
)

// InitCommand initializes a new repository
type InitCommand struct {
	config *config.Config
}

// NewInitCommand creates a new init command
func NewInitCommand(cfg *config.Config) *InitCommand {
	return &InitCommand{config: cfg}
}

// Execute runs the init command
func (c *InitCommand) Execute(args []string) error {
	repoName := "sg-repo"
	if len(args) > 0 {
		repoName = args[0]
	}

	ui.PrintInfo("🌀 Initializing sg repository...")
	ui.ShowProgress(2)

	defaultBranch := c.config.GetDefaultBranch()

	ui.PrintSuccess(fmt.Sprintf("✓ Initialized empty sg repository in ./%s/.sg/", repoName))
	ui.PrintWarning(fmt.Sprintf("⚠  Default branch is '%s' (this time)", defaultBranch))

	time.Sleep(500 * time.Millisecond)

	ui.PrintInfo("\nRepository configuration:")
	fmt.Printf("  • Chaos level: %d/10\n", c.config.GetChaosLevel())
	fmt.Printf("  • Therapy mode: %s\n", ui.BoolToStatus(c.config.IsTherapyAvailable()))
	fmt.Printf("  • Data integrity: questionable\n")
	fmt.Printf("  • Your sanity: pending\n")

	ui.PrintDim("\nYou now have one more thing to maintain forever.")

	return nil
}

// Description returns the command description
func (c *InitCommand) Description() string {
	return "Initialize a new repository"
}
