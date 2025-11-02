package commands

import (
	"fmt"
	"time"

	"github.com/sicktiergit/sicktiergit/internal/config"
	"github.com/sicktiergit/sicktiergit/internal/ui"
)

// PushCommand updates remote refs along with associated objects
type PushCommand struct {
	config *config.Config
}

// NewPushCommand creates a new push command
func NewPushCommand(cfg *config.Config) *PushCommand {
	return &PushCommand{config: cfg}
}

// Execute runs the push command
func (c *PushCommand) Execute(args []string) error {
	isForce := c.hasFlag(args, "--force", "-f")
	isMaybe := c.hasFlag(args, "--maybe")

	remote := "origin"
	branch := c.config.GetDefaultBranch()

	ui.PrintInfo(fmt.Sprintf("🚀 Pushing to %s/%s...", remote, branch))

	if isForce {
		ui.PrintWarning("⚠  Force push detected. Teammates will be notified.")
	}

	if isMaybe {
		ui.PrintWarning("⚠  Maybe mode enabled. Results may vary.")
	}

	// Simulate network activity
	ui.ShowProgress(3)

	// Simulate upload statistics
	objects := c.config.RandomSeed.Intn(50) + 10
	fmt.Printf("Counting objects: %d, done.\n", objects)
	time.Sleep(300 * time.Millisecond)

	fmt.Printf("Compressing objects: 100%% (%d/%d), done.\n", objects, objects)
	time.Sleep(300 * time.Millisecond)

	bytes := c.config.RandomSeed.Intn(5000) + 500
	fmt.Printf("Writing objects: 100%% (%d/%d), %.2f KiB, done.\n",
		objects, objects, float64(bytes)/1024)

	time.Sleep(500 * time.Millisecond)

	ui.PrintSuccess(fmt.Sprintf("✓ Successfully pushed to %s", remote))

	// Chaos-based messages
	if isForce {
		ui.PrintDim("All conflicting changes have been obliterated.")
		ui.PrintDim("If anyone complains, deny everything.")
	} else if isMaybe {
		ui.PrintDim("Push completed. Probably. Check the remote to be sure.")
	} else {
		ui.PrintDim("Your mistakes are now synchronized globally.")
	}

	return nil
}

func (c *PushCommand) hasFlag(args []string, flags ...string) bool {
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
func (c *PushCommand) Description() string {
	return "Update remote refs along with associated objects"
}
