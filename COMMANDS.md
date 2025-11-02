# Quick Command Reference

## Basic Commands

```bash
sg init [directory]              # Initialize new repository
sg commit -m "message"          # Commit changes
sg push [--force] [--maybe]     # Push to remote
sg pull                         # Pull from remote
sg status                       # Show working tree status
sg log                          # Show commit logs
```

## Branch Management

```bash
sg branch                       # List branches
sg checkout <branch>            # Switch branches
sg merge <branch>               # Merge branches
```

## Advanced Operations

```bash
sg rebase                       # Rebase current branch
sg blame [file] [--all]        # Show blame information
sg diff                         # Show changes
sg stash                        # Stash changes
sg reset [--hard|--soft]       # Reset HEAD
sg revert <commit>              # Revert commit
sg cherry-pick <commit>         # Cherry-pick commit
```

## Special Commands

```bash
sg therapy [--crisis]           # Open therapy session
sg undo                         # Attempt undo (doesn't work)
sg --help                       # Show help
sg --version                    # Show version
```

## Recommended Aliases

Add to your shell rc file (~/.zshrc, ~/.bashrc):

```bash
# Quick aliases for sg
alias sgi='sg init'
alias sgc='sg commit -m'
alias sgp='sg push'
alias sgs='sg status'
alias sgl='sg log'
alias sgt='sg therapy --crisis'

# Dangerous operations
alias sgf='sg push --force --maybe'
alias sgb='sg blame --all'
alias sgr='sg reset --hard'
```

## Common Workflows

### Quick commit and push
```bash
sg commit -m "Quick fix"
sg push
```

### Emergency recovery
```bash
sg therapy --crisis
sg reset --hard
sg therapy --crisis  # You'll need it again
```

### Blame everyone
```bash
sg blame --all
sg therapy  # Prepare for consequences
```

### The Maybe Workflow
```bash
sg commit -m "YOLO"
sg push --maybe
# Wait and see what happens
```

---

For detailed documentation, see [USAGE.md](USAGE.md)
