# Tutorial: Getting Started with gh-action-readme

**Time:** 5 minutes  
**Level:** Beginner  
**Goal:** Learn how to install and use gh-action-readme to automatically generate action documentation

## What you'll learn

In this tutorial, you'll:
1. Install the gh-action-readme extension
2. Create a GitHub Action with an `action.yml` file
3. Initialize a README with placeholders
4. Generate documentation automatically
5. Keep documentation in sync

## Prerequisites

Before starting, make sure you have:
- [GitHub CLI](https://cli.github.com/) installed and authenticated
- A GitHub Action project (or you can create one as part of this tutorial)

## Step 1: Install gh-action-readme

Install the gh-action-readme extension using GitHub CLI:

```bash
gh extension install reakaleek/gh-action-readme
```

Verify the installation:

```bash
gh action-readme --version
```

You should see version information displayed.

## Step 2: Create a sample GitHub Action

If you don't have an action yet, let's create a simple one for this tutorial.

Create a directory for your action:

```bash
mkdir my-awesome-action
cd my-awesome-action
```

Create an `action.yml` file with the following content:

```yaml
name: My Awesome Action
description: |
  This action demonstrates the power of automated documentation.
  It does something awesome and useful.

inputs:
  name:
    description: 'The name to greet'
    required: true
    default: 'World'
  
  greeting:
    description: 'The greeting to use'
    required: false
    default: 'Hello'

outputs:
  message:
    description: 'The greeting message that was generated'

runs:
  using: 'composite'
  steps:
    - run: echo "${{ inputs.greeting }}, ${{ inputs.name }}!"
      shell: bash
```

## Step 3: Initialize a README

Now, let's create a README with placeholders that gh-action-readme will populate:

```bash
gh action-readme init
```

This creates a `README.md` file with the following template:

```markdown
<!-- Generated with https://github.com/reakaleek/gh-action-readme -->
# <!--name--><!--/name-->
<!--description-->

## Inputs
<!--inputs-->

## Outputs
<!--outputs-->

## Usage
<!--usage action="org/repo" version="v1"-->
```yaml
steps:
 - uses: org/repo@v1
```
<!--/usage-->
```

Let's edit the usage section to match your action. Open `README.md` and update the usage placeholder:

```markdown
<!--usage action="myorg/my-awesome-action" version="v1.0.0"-->
```yaml
steps:
  - uses: actions/checkout@v4
  - uses: myorg/my-awesome-action@v1.0.0
    with:
      name: 'GitHub'
      greeting: 'Hello'
```
<!--/usage-->
```

## Step 4: Generate the documentation

Now for the magic! Run the update command:

```bash
gh action-readme update
```

Open `README.md` and you'll see it's been populated with your action's metadata:

````markdown
# My Awesome Action

This action demonstrates the power of automated documentation.
It does something awesome and useful.

## Inputs
| Name       | Description           | Required | Default   |
|------------|-----------------------|----------|-----------|
| `name`     | The name to greet     | `true`   | `World`   |
| `greeting` | The greeting to use   | `false`  | `Hello`   |

## Outputs
| Name      | Description                            |
|-----------|----------------------------------------|
| `message` | The greeting message that was generated |

## Usage
```yaml
steps:
  - uses: actions/checkout@v4
  - uses: myorg/my-awesome-action@v1.0.0
    with:
      name: 'GitHub'
      greeting: 'Hello'
```
````

## Step 5: Keep it in sync

As your action evolves, you can regenerate the documentation at any time:

1. **Update your action.yml** - Add new inputs, change descriptions, etc.
2. **Run update** - Execute `gh action-readme update` again
3. **Review changes** - The README updates automatically while preserving your custom content

Try it now! Add a new input to your `action.yml`:

```yaml
inputs:
  # ... existing inputs ...
  
  emoji:
    description: 'Add an emoji to the greeting'
    required: false
    default: '👋'
```

Then run:

```bash
gh action-readme update
```

Your README's inputs table now includes the new `emoji` input!

## Step 6: Check for drift

You can check if your documentation is up-to-date without modifying files:

```bash
gh action-readme diff
```

If there are differences, you'll see a diff output. If everything is synchronized, you'll see:
```
✓ README.md is up-to-date
```

## What you've learned

Congratulations! You now know how to:
- ✅ Install gh-action-readme
- ✅ Initialize a README with placeholders  
- ✅ Generate documentation from action.yml
- ✅ Keep documentation synchronized with your action
- ✅ Check for documentation drift

## Next steps

Now that you understand the basics, you can:

- Learn how to [build a complete action README](./complete-action-readme.md) with more advanced features
- Set up [pre-commit hooks](../how-to/setup-precommit.md) to automate documentation updates
- Explore the [placeholder reference](../reference/placeholders.md) to see all available options
