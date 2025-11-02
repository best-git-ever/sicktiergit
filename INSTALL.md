# Installation Guide

## Prerequisites

- Go 1.21 or higher
- Git (for cloning the repository)

## Quick Install

### From Source

1. Clone the repository:

```bash
git clone https://github.com/sicktiergit/sicktiergit.git
cd sicktiergit
```

2. Install dependencies:

```bash
make deps
```

3. Build and install:

```bash
make build
sudo make install
```

4. Verify installation:

```bash
sg --version
```

The command will be installed as `sg` (short for sicktiergit) for ease of use.

## Manual Build

If you prefer to build manually:

```bash
go mod download
go build -o bin/sg
```

Then copy the binary to your PATH:

```bash
sudo cp bin/sg /usr/local/bin/
```

## Development Setup

For development work:

```bash
# Install dependencies
make deps

# Build and run
make dev ARGS="init my-repo"

# Run tests
make test
```

## Uninstall

To remove sg:

```bash
sudo rm /usr/local/bin/sg
```

## Troubleshooting

### Permission Denied

If you get permission errors during installation:
```bash
sudo make install
```

### Command Not Found

Ensure `/usr/local/bin` is in your PATH:
```bash
echo $PATH
```

If not, add to your shell profile (~/.bashrc, ~/.zshrc, etc.):
```bash
export PATH="/usr/local/bin:$PATH"
```

### Missing Dependencies

Install required Go packages:
```bash
go get github.com/fatih/color
go get github.com/spf13/pflag
```

Or simply run:
```bash
make deps
```

## Platform-Specific Notes

### macOS
No additional steps required.

### Linux
Ensure you have proper permissions for `/usr/local/bin`.

### Windows

Use Git Bash or WSL for the best experience.
Or build with:

```bash
go build -o sg.exe
```
