# Command Reference

Complete reference for all gh-action-readme commands and their options.

## Overview

```bash
gh action-readme <command> [flags]
```

## Global Options

Available for all commands:

| Flag | Description |
|------|-------------|
| `--help`, `-h` | Show help information |
| `--version` | Show version information |

## Commands

### init

Initialize a README.md file with placeholder tags.

```bash
gh action-readme init [flags]
```

#### Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--readme` | | string | `README.md` | Path to README file to create |
| `--template` | | string | `default` | Template to use |
| `--recursive` | `-r` | bool | `false` | Search recursively for all action.yml files |

#### Examples

**Initialize a README in current directory:**
```bash
gh action-readme init
```

**Initialize with custom name:**
```bash
gh action-readme init --readme DOCUMENTATION.md
```

**Initialize all READMEs in a monorepo:**
```bash
gh action-readme init --recursive
```

**Initialize with custom template:**
```bash
gh action-readme init --template custom
```

#### Output

**Single file mode:**
```
✓ Created: README.md
```

**Recursive mode:**
```
Found 3 action file(s)

✓ Created: action-a/README.md
✓ Created: action-b/README.md
○ Skipped: action-c/README.md (already exists)

────────────────────────────────────
2 created, 1 skipped
────────────────────────────────────
```

#### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | File already exists (single mode) or error occurred |

#### Notes

- Will not overwrite existing README files
- Requires an `action.yml` or `action.yaml` file in the directory
- In recursive mode, creates a README next to each action.yml found

---

### update

Update README.md with metadata from action.yml.

```bash
gh action-readme update [flags]
```

#### Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--readme` | | string | `README.md` | Path to README file to update |
| `--recursive` | `-r` | bool | `false` | Search recursively for all action.yml files |
| `--action` | | string | (deprecated) | **Deprecated:** action files are now auto-detected |

#### Examples

**Update README in current directory:**
```bash
gh action-readme update
```

**Update custom README:**
```bash
gh action-readme update --readme docs/ACTION.md
```

**Update all READMEs in monorepo:**
```bash
gh action-readme update --recursive
```

#### Output

**Single file mode:**
```
✓ Updated: README.md
```

Or if no changes:
```
✓ Unchanged: README.md
```

**Recursive mode:**
```
Found 3 action file(s)

✓ Updated: action-a/README.md
○ Unchanged: action-b/README.md
✓ Updated: action-c/README.md

────────────────────────────────────
2 updated, 1 unchanged
────────────────────────────────────
```

#### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success (files updated or unchanged) |
| `1` | Error occurred |

#### Behavior

1. Locates `action.yml` or `action.yaml` in current directory (or recursively)
2. Parses action metadata (name, description, inputs, outputs)
3. Finds placeholder tags in README
4. Replaces content between placeholder tags
5. Preserves all content outside placeholder tags
6. Writes updated README back to disk

#### Notes

- Only modifies content within placeholder tags
- Preserves custom content outside placeholder tags
- Creates README if it doesn't exist (using default template)
- Auto-detects action.yml/action.yaml files

---

### diff

Show differences between current README and what it should be.

```bash
gh action-readme diff [flags]
```

#### Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--readme` | | string | `README.md` | Path to README file to check |
| `--recursive` | `-r` | bool | `false` | Search recursively for all action.yml files |
| `--action` | | string | (deprecated) | **Deprecated:** action files are now auto-detected |

#### Examples

**Check README in current directory:**
```bash
gh action-readme diff
```

**Check custom README:**
```bash
gh action-readme diff --readme docs/ACTION.md
```

**Check all READMEs in monorepo:**
```bash
gh action-readme diff --recursive
```

#### Output

**When up-to-date (single mode):**
```
✓ README.md is up-to-date
```

**When out-of-date (single mode):**
```
✗ README.md

--- current
+++ expected
@@ -5,6 +5,7 @@
 | Name          | Description              | Required | Default |
 |---------------|--------------------------|----------|---------|
 | `environment` | Target environment       | `true`   | ` `     |
+| `timeout`     | Operation timeout        | `false`  | `300`   |
 | `version`     | Version to deploy        | `true`   | ` `     |

README.md is not up-to-date
```

**Recursive mode:**
```
Found 3 action file(s)

✗ action-a/README.md

[diff output here]

✓ action-b/README.md
✓ action-c/README.md

────────────────────────────────────
2 up-to-date, 1 out-of-date
────────────────────────────────────
```

#### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | All READMEs are up-to-date |
| `1` | One or more READMEs are out-of-date |
| `>1` | Error occurred |

#### Diff Format

The diff output uses unified diff format:

- Lines starting with `-` are in current README (will be removed)
- Lines starting with `+` are in expected README (will be added)
- Lines with no prefix are context (unchanged)
- `@@ -X,Y +A,B @@` shows line numbers affected

#### Use Cases

- **Pre-commit checks:** Verify documentation before committing
- **CI/CD validation:** Ensure documentation is synchronized
- **Development:** Preview changes before applying them
- **Debugging:** Understand what will change

#### Notes

- Does not modify any files
- Perfect for CI/CD pipelines
- Shows exact changes that would be made by `update`
- Handles missing README files gracefully

---

### precommit

Run as a pre-commit hook. Automatically updates READMEs when action.yml changes.

```bash
gh action-readme precommit
```

#### Flags

None. This command is designed to be called by the pre-commit framework.

#### Usage

This command is typically not called directly. Instead, configure it in `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.4.1
    hooks:
      - id: action-readme
```

#### Behavior

1. Detects if action.yml or action.yaml files are staged
2. Automatically determines if recursive mode is needed
3. Updates corresponding README.md files
4. Stages updated READMEs for commit

#### Output

Similar to `update` command output.

#### Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success (READMEs updated) |
| `1` | Error occurred |

#### Notes

- Automatically handles both single actions and monorepos
- Only runs when action.yml/action.yaml files are modified
- Integrates seamlessly with pre-commit framework

---

## Command Comparison

| Command | Modifies Files | Shows Diff | Use Case |
|---------|---------------|------------|----------|
| `init` | ✅ (creates) | ❌ | Initialize new README |
| `update` | ✅ | ❌ | Update existing README |
| `diff` | ❌ | ✅ | Check for changes |
| `precommit` | ✅ | ❌ | Automated updates |

## Common Workflows

### Initial Setup

```bash
# Create a new README
gh action-readme init

# Review the template
cat README.md

# Customize the template
vim README.md

# Generate documentation
gh action-readme update
```

### Development Workflow

```bash
# Make changes to action.yml
vim action.yml

# Check what will change
gh action-readme diff

# Apply changes
gh action-readme update

# Review changes
git diff README.md

# Commit both files
git add action.yml README.md
git commit -m "Add timeout parameter"
```

### CI/CD Workflow

```bash
# In CI pipeline
gh action-readme diff

# Exit code 0 = up-to-date
# Exit code 1 = needs update
```

### Monorepo Workflow

```bash
# Initialize all READMEs
gh action-readme init --recursive

# Update all
gh action-readme update --recursive

# Check all
gh action-readme diff --recursive
```

## Environment Variables

Currently, gh-action-readme does not use environment variables for configuration. All configuration is done via command-line flags.

## Configuration Files

gh-action-readme does not currently support configuration files. All options must be specified as command-line flags.

## Shell Completion

To enable shell completion:

```bash
# For bash
gh completion -s bash > /etc/bash_completion.d/gh

# For zsh
gh completion -s zsh > /usr/local/share/zsh/site-functions/_gh

# For fish
gh completion -s fish > ~/.config/fish/completions/gh.fish
```

This enables completion for all `gh` extensions, including `gh action-readme`.

## See Also

- [Placeholder Reference](./placeholders.md) - Available placeholder tags
