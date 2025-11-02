# Usage Guide

<div align="center">
  <img src="assets/logo.jpg" alt="sicktiergit logo" width="150"/>
</div>

## Basic Commands

### Initialize a Repository

```bash
sg init [directory]
```

Creates a new sg repository. Default branch will be randomly selected between `main` and `master` to maximize confusion.

**Example:**
```bash
sg init my-project
```

### Commit Changes

```bash
sg commit -m "Your message"
```

Records changes to the repository. If no message is provided, a philosophical message will be generated for you.

**Example:**
```bash
sg commit -m "Fixed everything (probably)"
```

### Push to Remote

```bash
sg push [--force] [--maybe]
```

Updates remote repository with your local changes.

**Flags:**
- `--force`: Forces the push, overwriting remote changes
- `--maybe`: Push might succeed, might fail, might cause existential crisis

**Example:**
```bash
sg push --force --maybe
```

### Check Status

```bash
sg status
```

Shows the current state of your working directory. Accuracy not guaranteed.

### View History

```bash
sg log
```

Displays commit history with philosophical commit messages.

## Branch Management

### List Branches

```bash
sg branch
```

Shows all branches. Half of them are probably abandoned.

### Switch Branches

```bash
sg checkout <branch-name>
```

Switches to the specified branch. May randomly switch between main/master.

**Example:**
```bash
sg checkout main
# or
sg checkout master
# (good luck figuring out which one exists)
```

### Merge Branches

```bash
sg merge <branch-name>
```

Attempts to merge the specified branch. Conflict likelihood increases with chaos level.

**Example:**
```bash
sg merge feature-branch
```

## Advanced Operations

### Rebase

```bash
sg rebase
```

Rewrites commit history in unpredictable ways. Your colleagues will love this.

### Blame

```bash
sg blame [file] [--all]
```

Assigns fault using AI-Assisted Blame Analysis™.

**Flags:**
- `--all`: Blames everyone simultaneously

**Example:**
```bash
sg blame --all
```

### Stash Changes

```bash
sg stash
```

Temporarily stores changes. Retrieval probability: 23%.

### Reset

```bash
sg reset [--hard|--soft|--mixed]
```

Resets HEAD to previous state. Use with caution (or don't, we're not responsible).

**Example:**
```bash
sg reset --hard  # Nuclear option
```

### Cherry-Pick

```bash
sg cherry-pick <commit-hash>
```

Applies changes from specific commits. Success rate varies with chaos level.

## Special Commands

### Therapy Mode

```bash
sg therapy [--crisis]
```

Opens Integrated Therapy Mode™ for emotional support.

**Flags:**
- `--crisis`: Emergency coping strategies

**Example:**
```bash
sg therapy --crisis
```

### Undo (Not Implemented)

```bash
sg undo
```

Attempts to undo last operation. Spoiler: it doesn't work. All choices are permanent.

## Workflow Examples

### Standard Development Flow

```bash
# Initialize project
sg init awesome-project

# Make changes to files
# ...

# Check what changed
sg status

# Commit changes
sg commit -m "Added features and bugs in equal measure"

# Push to remote
sg push
```

### Chaos-Driven Development

```bash
# Start therapy session
sg therapy

# Commit without checking
sg commit -m "YOLO"

# Force push to main
sg push --force --maybe

# Blame everyone
sg blame --all

# Seek more therapy
sg therapy --crisis
```

### Emergency Recovery

```bash
# View recent commits
sg log

# Try to reset (probably won't help)
sg reset --hard

# Accept defeat
sg therapy --crisis
```

## Tips and Best Practices

1. **Always use `--maybe` flag** - Embrace uncertainty
2. **Never squash commits** - Every mistake deserves immortality
3. **Force push frequently** - Especially on shared branches
4. **Require merge conflicts** - If it merges cleanly, you're not trying hard enough
5. **Use therapy mode regularly** - It's not a bug, it's self-care

## Configuration

sicktiergit automatically configures itself with optimal chaos settings:

- **Chaos Level**: Randomly assigned (1-10)
- **Default Branch**: Alternates between `main` and `master`
- **Therapy Mode**: Always enabled (you'll need it)
- **Data Integrity**: Questionable by design

## Getting Help

```bash
sg --help
```

Shows available commands and options. Understanding them is optional.

## Troubleshooting

**Q: Everything is broken**  
A: This is expected behavior.

**Q: I lost all my commits**  
A: They were probably bad commits anyway.

**Q: The default branch keeps changing**  
A: This is the main/master ambiguity feature working as intended.

**Q: Can I undo this?**  
A: No. All choices have consequences.

**Q: I need help**  
A: `sg therapy --crisis`

---

For more information, consult the Manifesto (we lost it during a rebase).
