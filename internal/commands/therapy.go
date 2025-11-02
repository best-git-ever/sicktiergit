package commands

import (
	"fmt"
	"time"

	"github.com/sicktiergit/sicktiergit/internal/config"
	"github.com/sicktiergit/sicktiergit/internal/ui"
)

// TherapyCommand provides integrated emotional support
type TherapyCommand struct {
	config *config.Config
}

// NewTherapyCommand creates a new therapy command
func NewTherapyCommand(cfg *config.Config) *TherapyCommand {
	return &TherapyCommand{config: cfg}
}

// Execute runs the therapy command
func (c *TherapyCommand) Execute(args []string) error {
	isCrisis := c.hasFlag(args, "--crisis")

	if !c.config.IsTherapyAvailable() {
		ui.PrintWarning("⚠  Therapy Mode™ is currently unavailable.")
		ui.PrintDim("Please continue suffering in silence.")
		return nil
	}

	ui.PrintInfo("🛋️  Opening Integrated Therapy Session...")
	ui.ShowProgress(2)

	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║     SG THERAPY MODE™ v1.0                         ║")
	fmt.Println("║     Your emotional support companion              ║")
	fmt.Println("╚════════════════════════════════════════════════════╝\n")

	time.Sleep(500 * time.Millisecond)

	if isCrisis {
		ui.PrintWarning("🚨 CRISIS MODE DETECTED")
		fmt.Println("\nEmergency coping strategies:")
		fmt.Println("  1. Step away from the keyboard")
		fmt.Println("  2. Remember: it's just code")
		fmt.Println("  3. Backup exists... probably")
		fmt.Println("  4. Coffee solves 83% of problems")
		fmt.Println("  5. Blame it on the cache")
		fmt.Println()
	}

	// Therapeutic messages based on chaos level
	chaosLevel := c.config.GetChaosLevel()

	fmt.Println("Session Notes:")
	fmt.Println("──────────────")

	if chaosLevel > 8 {
		fmt.Println("• Your repository is in a state of existential crisis")
		fmt.Println("• This is normal and expected behavior")
		fmt.Println("• Consider accepting that control is an illusion")
	} else if chaosLevel > 5 {
		fmt.Println("• Merge conflicts are opportunities for growth")
		fmt.Println("• Every failed rebase teaches us something")
		fmt.Println("• Your code is a reflection of your journey")
	} else {
		fmt.Println("• Things could be worse (probably)")
		fmt.Println("• At least it compiles (sometimes)")
		fmt.Println("• Tomorrow is a new commit")
	}

	fmt.Println("\nAffirmations:")
	affirmations := []string{
		"• I am a developer, not a version control system",
		"• My worth is not measured in commits",
		"• It's okay to not understand git",
		"• Force pushing is sometimes necessary",
		"• I deserve a clean working tree",
	}

	for _, affirmation := range affirmations {
		fmt.Println(affirmation)
		time.Sleep(200 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)

	ui.PrintInfo("\n📊 Therapy Session Statistics:")
	fmt.Printf("  • Unresolved trauma: %d%%\n", c.config.RandomSeed.Intn(100))
	fmt.Printf("  • Suppressed rage: %d%%\n", c.config.RandomSeed.Intn(100))
	fmt.Printf("  • Hope remaining: %d%%\n", c.config.RandomSeed.Intn(30)+5)

	ui.PrintDim("\nSession ended. Remember: the repository can't hurt you.")
	ui.PrintDim("(This is a lie, but sometimes lies help us cope.)")

	return nil
}

func (c *TherapyCommand) hasFlag(args []string, flags ...string) bool {
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
func (c *TherapyCommand) Description() string {
	return "Open integrated emotional support session"
}
