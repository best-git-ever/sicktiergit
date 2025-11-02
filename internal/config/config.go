package config

import (
	"math/rand"
	"time"
)

// Config holds the application configuration
type Config struct {
	DefaultBranch  string
	RemoteURL      string
	Username       string
	Email          string
	TherapyEnabled bool
	ChaosLevel     int
	RandomSeed     *rand.Rand
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Randomly choose between main and master to simulate confusion
	branches := []string{"main", "master"}
	defaultBranch := branches[rng.Intn(len(branches))]

	return &Config{
		DefaultBranch:  defaultBranch,
		RemoteURL:      "https://production.sicktiergit.io/chaos",
		Username:       "anonymous",
		Email:          "regrets@sicktiergit.io",
		TherapyEnabled: true,
		ChaosLevel:     rng.Intn(10) + 1,
		RandomSeed:     rng,
	}
}

// GetDefaultBranch returns the default branch name (randomly chosen)
func (c *Config) GetDefaultBranch() string {
	return c.DefaultBranch
}

// IsTherapyAvailable checks if therapy mode is enabled
func (c *Config) IsTherapyAvailable() bool {
	return c.TherapyEnabled
}

// GetChaosLevel returns the current chaos level (1-10)
func (c *Config) GetChaosLevel() int {
	return c.ChaosLevel
}
