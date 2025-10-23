# How to Check for Documentation Drift

This guide shows you how to verify that your action's README is synchronized with your `action.yml` file.

## What is Documentation Drift?

Documentation drift occurs when your `action.yml` changes but the README isn't updated to reflect those changes. This can happen when:

- Adding, removing, or modifying inputs
- Changing input descriptions
- Adding or removing outputs
- Updating the action name or description

## Using the diff Command

The `diff` command compares your current README with what it should be:

```bash
gh action-readme diff
```

### When Everything is Up-to-Date

```
✓ README.md is up-to-date
```

Exit code: `0` (success)

### When Documentation is Out of Sync

```
✗ README.md

--- current
+++ expected
@@ -8,7 +8,8 @@
 ## Inputs
 | Name          | Description              | Required | Default |
 |---------------|--------------------------|----------|---------|
-| `environment` | Target environment       | `true`   | ` `     |
+| `environment` | Deployment environment   | `true`   | ` `     |
 | `version`     | Version to deploy        | `true`   | ` `     |
+| `timeout`     | Maximum deployment time  | `false`  | `300`   |

README.md is not up-to-date
```

Exit code: `1` (failure)

## Checking Multiple Actions

For monorepos with multiple actions:

```bash
gh action-readme diff --recursive
```

Output:

```
Found 3 action file(s)

✗ action-deploy/README.md

--- action-deploy/README.md
+++ expected
@@ -5,6 +5,7 @@
 | `environment` | Target environment | `true`   | ` `        |
 | `version`     | Version to deploy  | `true`   | ` `        |
+| `timeout`     | Deployment timeout | `false`  | `300`      |

✓ action-notify/README.md
✓ action-validate/README.md

────────────────────────────────────
2 up-to-date, 1 out-of-date
────────────────────────────────────
```

## Using in CI/CD

### GitHub Actions Workflow

Create `.github/workflows/docs-check.yml`:

```yaml
name: Documentation Check

on:
  pull_request:
    paths:
      - 'action.yml'
      - 'action.yaml'
      - 'README.md'

jobs:
  check-docs:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      
      - name: Install gh-action-readme
        run: gh extension install reakaleek/gh-action-readme
        env:
          GH_TOKEN: ${{ github.token }}
      
      - name: Check documentation
        run: gh action-readme diff
```

This workflow:
- Runs on PRs that modify action.yml or README.md
- Installs gh-action-readme
- Checks for documentation drift
- Fails the PR if documentation is out of sync

### For Monorepos

```yaml
name: Documentation Check

on:
  pull_request:
    paths:
      - '**/action.yml'
      - '**/action.yaml'
      - '**/README.md'

jobs:
  check-docs:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      
      - name: Install gh-action-readme
        run: gh extension install reakaleek/gh-action-readme
        env:
          GH_TOKEN: ${{ github.token }}
      
      - name: Check all documentation
        run: gh action-readme diff --recursive
```

### Adding Helpful Context

Make the CI failure message more helpful:

```yaml
      - name: Check documentation
        id: check
        run: gh action-readme diff --recursive
        continue-on-error: true
      
      - name: Comment on PR
        if: steps.check.outcome == 'failure'
        uses: actions/github-script@v7
        with:
          script: |
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: '⚠️ **Documentation is out of sync!**\n\n' +
                    'Please update the README by running:\n' +
                    '```bash\ngh action-readme update --recursive\n```\n\n' +
                    'Or let pre-commit do it automatically:\n' +
                    '```bash\npre-commit install\n```'
            })
      
      - name: Fail if docs are out of sync
        if: steps.check.outcome == 'failure'
        run: exit 1
```

## Pre-commit Integration

The best way to prevent drift is to catch it before committing:

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.5.0
    hooks:
      - id: action-readme
```

Install it:

```bash
pip install pre-commit
pre-commit install
```

Now documentation is automatically updated before every commit!

## Manual Workflow

### 1. Make changes to action.yml

```bash
# Edit your action definition
vim action.yml
```

### 2. Check for drift

```bash
gh action-readme diff
```

### 3. Review the changes

Look at the diff output to see what will change in your README.

### 4. Update if needed

```bash
gh action-readme update
```

### 5. Review the updated README

```bash
git diff README.md
```

### 6. Commit both files

```bash
git add action.yml README.md
git commit -m "Add timeout input to action"
```

## Common Drift Scenarios

### Scenario 1: Description Changed

**Change in action.yml:**
```yaml
inputs:
  environment:
    description: 'Deployment environment (prod, staging, dev)'  # Updated
    required: true
```

**Detected by diff:**
```diff
-| `environment` | Target environment | `true` | ` ` |
+| `environment` | Deployment environment (prod, staging, dev) | `true` | ` ` |
```

### Scenario 2: New Input Added

**Change in action.yml:**
```yaml
inputs:
  timeout:  # New input
    description: 'Operation timeout in seconds'
    required: false
    default: '300'
```

**Detected by diff:**
```diff
 | `version`     | Version to deploy        | `true`   | ` `   |
+| `timeout`     | Operation timeout in seconds | `false` | `300` |
```

### Scenario 3: Default Value Changed

**Change in action.yml:**
```yaml
inputs:
  retry-attempts:
    description: 'Number of retry attempts'
    required: false
    default: '5'  # Changed from '3'
```

**Detected by diff:**
```diff
-| `retry-attempts` | Number of retry attempts | `false` | `3` |
+| `retry-attempts` | Number of retry attempts | `false` | `5` |
```

### Scenario 4: Required Status Changed

**Change in action.yml:**
```yaml
inputs:
  api-key:
    description: 'API key for authentication'
    required: true  # Changed from false
```

**Detected by diff:**
```diff
-| `api-key` | API key for authentication | `false` | ` ` |
+| `api-key` | API key for authentication | `true`  | ` ` |
```

## Understanding Diff Output

The diff output follows standard unified diff format:

```diff
--- current                    # Current state of README.md
+++ expected                   # Expected state based on action.yml
@@ -8,7 +8,8 @@              # Line numbers affected
 ## Inputs                     # Context line (unchanged)
 | Name          | ...         # Context line (unchanged)
 |---------------|...          # Context line (unchanged)
-| `environment` | Target ...  # Line to be removed (current)
+| `environment` | Deploy ...  # Line to be added (expected)
```

- Lines starting with `-` show current content that will be removed
- Lines starting with `+` show new content that will be added
- Lines with no prefix are context (unchanged)

## Exit Codes

The `diff` command uses exit codes to indicate status:

| Exit Code | Meaning | Use Case |
|-----------|---------|----------|
| `0` | Documentation is up-to-date | CI passes |
| `1` | Documentation is out of sync | CI fails |
| `>1` | Error occurred | CI fails with error |

This makes it perfect for CI/CD pipelines:

```bash
# In a CI script
if gh action-readme diff; then
    echo "✓ Documentation is current"
else
    echo "✗ Documentation needs updating"
    exit 1
fi
```

## Best Practices

### 1. Check Before Committing

Always check for drift before committing:

```bash
gh action-readme diff
```

### 2. Use Pre-commit Hooks

Automate the process:

```bash
pre-commit install
```

### 3. Add to CI Pipeline

Catch drift in pull requests:

```yaml
- name: Check documentation
  run: gh action-readme diff --recursive
```

### 4. Review Diffs Carefully

Don't blindly update - review what's changing:

```bash
gh action-readme diff     # See what will change
gh action-readme update   # Apply changes
git diff README.md        # Review the actual changes
```

### 5. Document Your Process

Add to CONTRIBUTING.md:

```markdown
## Keeping Documentation Synchronized

After modifying `action.yml`, update the README:

```bash
gh action-readme update
```

Or check for changes without modifying:

```bash
gh action-readme diff
```
```

## Troubleshooting

### Diff shows changes but action.yml hasn't changed

**Possible causes:**
- Line ending differences (CRLF vs LF)
- Whitespace differences
- README was manually edited inside placeholder tags

**Solution:**
Run update to normalize:
```bash
gh action-readme update
```

### Diff shows no changes but README seems wrong

**Possible causes:**
- Placeholders are misformatted
- Content is outside placeholder tags

**Solution:**
Check placeholder format:
```markdown
<!--inputs-->
[generated content here]
<!--/inputs-->
```

### CI keeps failing on diff check

**Possible causes:**
- Pre-commit not configured
- Developer updated action.yml but not README
- README was modified manually

**Solution:**
```bash
# Update the documentation
gh action-readme update

# Commit the changes
git add README.md
git commit --amend --no-edit
```

## Related Guides

- [Set up pre-commit hooks](./setup-precommit.md) - Prevent drift automatically
- [Managing monorepos](./manage-monorepos.md) - Check multiple actions

## Related Reference

- [diff command](../reference/commands.md#diff) - Complete command reference
- [Placeholders](../reference/placeholders.md) - Available placeholder tags
