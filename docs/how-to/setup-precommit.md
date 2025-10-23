# How to Set Up Pre-commit Hooks

This guide shows you how to automatically update your action's README before every commit using pre-commit hooks.

## Prerequisites

- gh-action-readme installed
- Python 3.7+ installed (for pre-commit framework)
- Git repository initialized

## Method 1: Using pre-commit framework (Recommended)

### Step 1: Install pre-commit

```bash
# Using pip
pip install pre-commit

# Or using homebrew (macOS)
brew install pre-commit

# Or using conda
conda install -c conda-forge pre-commit
```

### Step 2: Create configuration file

Create a `.pre-commit-config.yaml` file in your repository root:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0  # Use the latest version
    hooks:
      - id: action-readme
```

### Step 3: Install the git hooks

```bash
pre-commit install
```

### Step 4: Test it

Make a change to your `action.yml`:

```yaml
# Add a new input
inputs:
  new-feature:
    description: 'A new feature input'
    required: false
```

Stage and commit:

```bash
git add action.yml
git commit -m "Add new feature input"
```

You'll see output like:

```
action-readme............................................................Passed
[main 1a2b3c4] Add new feature input
 2 files changed, 10 insertions(+), 2 deletions(-)
```

The README.md will be automatically updated and included in the commit!

## Method 2: Manual git hook

If you prefer not to use the pre-commit framework, you can create a manual git hook.

### Step 1: Create the hook script

Create `.git/hooks/pre-commit`:

```bash
#!/bin/sh

# Run gh-action-readme update
gh action-readme update

# Check if README.md was modified
if git diff --name-only | grep -q "README.md"; then
    # Add the updated README to the commit
    git add README.md
    echo "✓ README.md updated and staged"
fi

exit 0
```

### Step 2: Make it executable

```bash
chmod +x .git/hooks/pre-commit
```

### Step 3: Test it

```bash
# Make a change to action.yml
echo "  test-input:\n    description: 'Test input'" >> action.yml

# Commit the change
git add action.yml
git commit -m "Add test input"
```

The hook will run automatically!

## Configuration Options

### Running on specific files only

If your repository has multiple actions, you can configure pre-commit to run only on specific directories:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
        files: '^my-action/'  # Only run in my-action directory
```

### Running with recursive flag

For monorepos, enable recursive mode:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
        args: ['--recursive']
```

This will find and update all action READMEs in your repository.

## Working with Multiple Developers

### Commit the configuration

Make sure your `.pre-commit-config.yaml` is committed:

```bash
git add .pre-commit-config.yaml
git commit -m "Add pre-commit configuration"
git push
```

### Team setup

Each team member needs to run:

```bash
# Install pre-commit
pip install pre-commit

# Install the hooks
pre-commit install
```

### Optional: Auto-install hooks

You can add a note to your README or CONTRIBUTING.md:

```markdown
## Developer Setup

After cloning the repository:

```bash
pip install pre-commit
pre-commit install
```

This ensures action documentation stays synchronized automatically.
```

## Using with GitHub Actions

You can also run pre-commit checks in CI to ensure documentation is up-to-date:

```yaml
name: Pre-commit Checks

on: [pull_request]

jobs:
  pre-commit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Python
        uses: actions/setup-python@v4
        with:
          python-version: '3.x'
      
      - name: Install pre-commit
        run: pip install pre-commit
      
      - name: Run pre-commit
        run: pre-commit run --all-files
```

This fails the PR if documentation is out of sync, reminding developers to update it.

## Troubleshooting

### Hook not running

**Problem:** The hook doesn't execute when committing.

**Solution:** Make sure hooks are installed:
```bash
pre-commit install
```

### "gh not found" error

**Problem:** The hook can't find the gh command.

**Solution:** Make sure GitHub CLI is in your PATH. You might need to specify the full path:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
        entry: /usr/local/bin/gh action-readme update
```

### README keeps changing on every commit

**Problem:** Pre-commit modifies README even when action.yml hasn't changed.

**Solution:** This usually indicates an issue with placeholder formatting or line endings. Run:
```bash
gh action-readme diff
```
to see what's changing unexpectedly.

### Skipping hooks temporarily

If you need to skip the hook for a single commit:

```bash
git commit --no-verify -m "Your commit message"
```

Use this sparingly and only when necessary!

## Advanced Configuration

### Run only on changed files

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
        files: 'action\.ya?ml$'  # Only run if action.yml/yaml changed
```

### Combine with other hooks

```yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
  
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
```

### Auto-update hook versions

Enable auto-updating of hook versions:

```bash
pre-commit autoupdate
```

This updates the `rev` field to the latest release.

## Best Practices

1. **Commit the config** - Always commit `.pre-commit-config.yaml` to version control
2. **Update regularly** - Keep the hook version up to date with `pre-commit autoupdate`
3. **Document for team** - Add setup instructions to CONTRIBUTING.md
4. **Test before pushing** - Run `pre-commit run --all-files` to test all files
5. **Use CI checks** - Add pre-commit to CI to catch issues in PRs

## Related Guides

- [Managing monorepos](./manage-monorepos.md)
- [Checking for documentation drift](./check-documentation-drift.md)
