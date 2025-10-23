# Tutorial: Building a Complete Action README

**Time:** 15 minutes  
**Level:** Intermediate  
**Prerequisites:** Complete the [Getting Started](./getting-started.md) tutorial first

## What you'll learn

In this tutorial, you'll create a production-ready README for a GitHub Action with:
- Automatically generated metadata sections
- Custom content sections  
- Multiple usage examples
- Usage examples with version tracking

## Step 1: Set up a more complex action

Let's create a more realistic GitHub Action. Create a new directory:

```bash
mkdir deploy-action
cd deploy-action
```

Create an `action.yml` with multiple inputs and outputs:

```yaml
name: Deploy Application
description: |
  Deploy your application to various environments.
  Supports automated rollback and health checks.

inputs:
  environment:
    description: 'Target environment (production, staging, development)'
    required: true
  
  version:
    description: 'Version to deploy'
    required: true
  
  health-check-url:
    description: 'URL to check after deployment'
    required: false
  
  timeout:
    description: 'Deployment timeout in seconds'
    required: false
    default: '300'

outputs:
  deployed-version:
    description: 'The version that was deployed'
  
  deployment-url:
    description: 'URL where the application was deployed'
  
  deployment-time:
    description: 'Time taken for deployment in seconds'

runs:
  using: 'composite'
  steps:
    - run: echo "Deploying..."
      shell: bash
```

## Step 2: Create a comprehensive README template

Initialize the README:

```bash
gh action-readme init
```

Now let's enhance the template with custom sections. Edit `README.md`:

````markdown
# <!--name--><!--/name-->

<!--description-->
<!--/description-->

## Features

- Deploy to multiple environments
- Automated health checks
- Rollback support
- Fast and reliable

## Inputs

<!--inputs-->
<!--/inputs-->

## Outputs

<!--outputs-->
<!--/outputs-->

## Usage

### Basic Deployment

<!--usage action="yourorg/deploy-action" version="v1.0.0"-->
```yaml
steps:
  - uses: actions/checkout@v4
  
  - name: Deploy
    uses: yourorg/deploy-action@v1.0.0
    with:
      environment: production
      version: ${{ github.sha }}
      
  - name: Deploy with health check
    uses: yourorg/deploy-action@v1.0.0
    with:
      environment: production
      version: v2.1.0
      health-check-url: https://api.example.com/health
      timeout: 600
      
  - name: Deploy to staging
    uses: yourorg/deploy-action@v1.0.0
    with:
      environment: staging
      version: ${{ github.ref_name }}
```
<!--/usage-->

## License

MIT
````

## Step 3: Generate the documentation

Run the update command:

```bash
gh action-readme update
```

## Step 4: Review the result

Open your `README.md` and observe:

1. **Automated sections** - The inputs and outputs tables are automatically generated from action.yml
2. **Preserved custom content** - Your features list and example sections remain intact
3. **Version tracking** - All action references inside the usage placeholder are updated to v1.0.0
4. **Clean structure** - Placeholders only replace what's between the tags

## Step 5: Add it to version control

```bash
git init
git add action.yml README.md
git commit -m "Initial action with automated documentation"
```

## Step 6: Set up automatic updates

To ensure your README stays in sync, set up a pre-commit hook:

```bash
# Install pre-commit if you haven't already
pip install pre-commit

# Create .pre-commit-config.yaml
cat > .pre-commit-config.yaml << 'EOF'
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.4.1
    hooks:
      - id: action-readme
EOF

# Install the hook
pre-commit install
```

Now, every time you commit changes to `action.yml`, your README will automatically update!

## Step 7: Verify with diff

Before committing changes, you can always check what would change:

```bash
gh action-readme diff
```

This shows you the differences without modifying the file.

## What you've learned

You now know how to:
- ✅ Create a README with multiple placeholder sections
- ✅ Combine automated and custom content
- ✅ Use usage placeholder to track versions across multiple examples
- ✅ Set up automated documentation updates with pre-commit
- ✅ Use diff to preview changes before applying them

## Key Takeaways

With gh-action-readme, you can:

1. **Mix automated and custom content** - Placeholders handle metadata, you control everything else
2. **Update multiple sections** - Name, description, inputs, outputs, and usage all stay synchronized
3. **Version tracking** - Put multiple action examples inside one usage placeholder to update all versions at once
4. **Preview changes** - Use `diff` to see what will change before applying
5. **Automate with pre-commit** - Never forget to update documentation

## Next steps

- Learn about [managing monorepos](../how-to/manage-monorepos.md) with multiple actions
- Explore [advanced placeholder options](../reference/placeholders.md)
- Set up [documentation drift checking](../how-to/check-documentation-drift.md) in CI
