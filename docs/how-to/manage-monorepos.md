# How to Manage Monorepos with Multiple Actions

This guide shows you how to use gh-action-readme in a repository that contains multiple GitHub Actions.

## Understanding Monorepo Support

gh-action-readme has built-in support for monorepos. The `--recursive` (or `-r`) flag automatically finds all `action.yml` or `action.yaml` files in your repository and processes their corresponding READMEs.

## Typical Monorepo Structure

```
my-actions-monorepo/
├── .github/
│   └── workflows/
├── action-deploy/
│   ├── action.yml
│   ├── README.md
│   └── src/
├── action-notify/
│   ├── action.yml
│   ├── README.md
│   └── src/
├── action-validate/
│   ├── action.yml
│   ├── README.md
│   └── src/
└── shared/
    └── utils/
```

## Initialize All READMEs at Once

When starting with a new monorepo or adding READMEs to existing actions:

```bash
cd my-actions-monorepo
gh action-readme init --recursive
```

This will:
1. Search for all `action.yml` and `action.yaml` files
2. Create a `README.md` next to each one (if it doesn't exist)
3. Skip any existing READMEs

Example output:

```
Found 3 action file(s)

✓ Created: action-deploy/README.md
✓ Created: action-notify/README.md
○ Skipped: action-validate/README.md (already exists)

────────────────────────────────────
2 created, 1 skipped
────────────────────────────────────
```

## Update All READMEs at Once

After making changes to your action definitions:

```bash
gh action-readme update --recursive
```

This will:
1. Find all action.yml files
2. Find corresponding READMEs
3. Update each README with current metadata

Example output:

```
Found 3 action file(s)

✓ Updated: action-deploy/README.md
○ Unchanged: action-notify/README.md
✓ Updated: action-validate/README.md

────────────────────────────────────
2 updated, 1 unchanged
────────────────────────────────────
```

## Check All READMEs for Drift

Before committing, verify all documentation is up-to-date:

```bash
gh action-readme diff --recursive
```

Example output when all are up-to-date:

```
Found 3 action file(s)

✓ action-deploy/README.md
✓ action-notify/README.md
✓ action-validate/README.md

────────────────────────────────────
3 up-to-date, 0 out-of-date
────────────────────────────────────
```

Example output when some need updates:

```
Found 3 action file(s)

✗ action-deploy/README.md

--- current
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

Exit code is 1 if any READMEs are out-of-date, making this perfect for CI.

## Set Up Pre-commit for Monorepos

Create `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.4.1
    hooks:
      - id: action-readme
```

The pre-commit hook automatically handles monorepos! It will:
1. Detect all action.yml files in the repository
2. Update all corresponding READMEs
3. Stage any modified READMEs for commit

Install the hook:

```bash
pip install pre-commit
pre-commit install
```

## Working with Nested Actions

gh-action-readme supports deeply nested actions:

```
my-monorepo/
├── actions/
│   ├── docker/
│   │   ├── build/
│   │   │   ├── action.yml
│   │   │   └── README.md
│   │   └── push/
│   │       ├── action.yml
│   │       └── README.md
│   └── kubernetes/
│       ├── deploy/
│       │   ├── action.yml
│       │   └── README.md
│       └── rollback/
│           ├── action.yml
│           └── README.md
```

All actions at any depth are found and processed:

```bash
gh action-readme update --recursive
```

```
Found 4 action file(s)

✓ Updated: actions/docker/build/README.md
✓ Updated: actions/docker/push/README.md
✓ Updated: actions/kubernetes/deploy/README.md
✓ Updated: actions/kubernetes/rollback/README.md

────────────────────────────────────
4 updated, 0 unchanged
────────────────────────────────────
```

## Update a Single Action in a Monorepo

If you want to update just one action without the recursive flag:

```bash
# Navigate to the action's directory
cd action-deploy

# Update just this README
gh action-readme update

# Or specify from root
gh action-readme update --readme action-deploy/README.md
```

## CI/CD for Monorepos

### Verify All Documentation in Pull Requests

```yaml
name: Verify Action Documentation

on:
  pull_request:
    paths:
      - '**/action.yml'
      - '**/action.yaml'
      - '**/README.md'

jobs:
  verify-docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install gh-action-readme
        run: gh extension install reakaleek/gh-action-readme
        env:
          GH_TOKEN: ${{ github.token }}
      
      - name: Check documentation is up-to-date
        run: gh action-readme diff --recursive
```

This workflow:
- Runs when action.yml or README.md files change
- Verifies all action documentation is synchronized
- Fails the check if any README is out-of-date

### Auto-update Documentation in Commits

```yaml
name: Update Action Documentation

on:
  pull_request:
    paths:
      - '**/action.yml'
      - '**/action.yaml'

jobs:
  update-docs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          ref: ${{ github.head_ref }}
      
      - name: Install gh-action-readme
        run: gh extension install reakaleek/gh-action-readme
        env:
          GH_TOKEN: ${{ github.token }}
      
      - name: Update all READMEs
        run: gh action-readme update --recursive
      
      - name: Check for changes
        id: verify-changed-files
        run: |
          if [ -n "$(git status --porcelain)" ]; then
            echo "changed=true" >> $GITHUB_OUTPUT
          else
            echo "changed=false" >> $GITHUB_OUTPUT
          fi
      
      - name: Commit updated READMEs
        if: steps.verify-changed-files.outputs.changed == 'true'
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"
          git add '**/README.md'
          git commit -m "docs: update action documentation"
          git push
```

## Best Practices for Monorepos

### 1. Consistent Structure

Keep a consistent directory structure:

```
actions/
├── <action-name>/
│   ├── action.yml
│   ├── README.md
│   ├── src/
│   └── test/
```

### 2. Root README

Create a root README that links to all actions:

```markdown
# My Actions Collection

This repository contains multiple reusable GitHub Actions.

## Available Actions

- [action-deploy](./action-deploy/) - Deploy applications to various environments
- [action-notify](./action-notify/) - Send notifications to multiple channels
- [action-validate](./action-validate/) - Validate configurations and schemas

## Using These Actions

Each action has its own README with detailed usage instructions. Click the links above to learn more.

## Development

All action documentation is automatically generated using [gh-action-readme](https://github.com/reakaleek/gh-action-readme).

To update documentation:
```bash
gh action-readme update --recursive
```
```

### 3. Shared Configuration

If actions share inputs or patterns, document them centrally:

```
docs/
├── common-inputs.md
├── common-patterns.md
└── development-guide.md
```

Then reference them in individual action READMEs.

### 4. Version Management

For monorepos, consider:

- **Independent versioning** - Each action has its own version
- **Synchronized versioning** - All actions share a version
- **Semantic versioning** - Major.Minor.Patch for each action

Update usage examples accordingly in each README.

### 5. Testing Strategy

Test all actions when any change is made:

```yaml
name: Test All Actions

on: [push, pull_request]

jobs:
  test-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Test deploy action
        uses: ./action-deploy
        with:
          environment: test
          version: 1.0.0
  
  test-notify:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Test notify action
        uses: ./action-notify
        with:
          message: Test notification
```

## Common Patterns

### Pattern: Shared Dependencies

```
monorepo/
├── shared/
│   └── utils.js
├── action-a/
│   ├── action.yml
│   ├── README.md
│   └── index.js (imports ../shared/utils.js)
└── action-b/
    ├── action.yml
    ├── README.md
    └── index.js (imports ../shared/utils.js)
```

### Pattern: Action Families

Group related actions:

```
monorepo/
├── docker/
│   ├── build-action/
│   ├── push-action/
│   └── scan-action/
├── kubernetes/
│   ├── deploy-action/
│   ├── status-action/
│   └── rollback-action/
```

### Pattern: Staged Actions

Actions that work in sequence:

```
monorepo/
├── prepare-action/
├── build-action/
├── test-action/
└── deploy-action/
```

Document the workflow in the root README.

## Troubleshooting

### Some READMEs not being found

**Problem:** `gh action-readme update --recursive` doesn't find all actions.

**Solution:** Ensure:
1. Each action directory has `action.yml` or `action.yaml`
2. README is named exactly `README.md` (case-sensitive)
3. README is in the same directory as action.yml

### Performance with many actions

**Problem:** Recursive operations are slow with many actions.

**Solution:** 
- Use directory-specific updates when working on one action
- Consider parallel processing in CI (run multiple jobs)

### Inconsistent documentation

**Problem:** Some READMEs have different styles or formats.

**Solution:**
- Use templates consistently
- Run `gh action-readme init --recursive` with a custom template

## Related Guides

- [Set up pre-commit hooks](./setup-precommit.md)
- [Checking for documentation drift](./check-documentation-drift.md)
