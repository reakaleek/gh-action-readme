
# gh-action-readme
![GitHub Release](https://img.shields.io/github/v/release/reakaleek/gh-action-readme?logo=github)
![GitHub Release Date](https://img.shields.io/github/release-date/reakaleek/gh-action-readme?display_date=published_at&logo=github)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/reakaleek/gh-action-readme)
![GitHub License](https://img.shields.io/github/license/reakaleek/gh-action-readme)

A GitHub CLI extension to inject GitHub Actions metadata into `README.md` markdown files.

> [!WARNING]
> This project is in active development and is not yet stable.

## Prerequisites
- [GitHub CLI](https://cli.github.com/) installed and authenticated.

## ⚡️ Quick Start

### Install the gh-action-readme extension

The `gh-action-readme` extension can be installed using the following command.

```bash
gh extension install reakaleek/gh-action-readme
```

### Create a README.md file

Create a `README.md` file in the action directory containing the `action.yml` file.

You can use the following template to define the metadata placeholders.

```markdown
# <!--name--><!--/name-->
<!--description-->

## Inputs
<!--inputs-->
```

### Update the `README.md` file

Run the following command to update the `README.md` file with the metadata from the `action.yml` file.

```bash
gh action-readme update
```

### That's it! 🎉

The `README.md` file will be updated with the metadata from the `action.yml` file.

```diff
-# <!--name--><!--/name-->
+# <!--name-->Awesome Action<!--/name-->
<!--description-->
+A GitHub Action that does something awesome.
+Something that is very useful.
+<!--/description-->

## Inputs
<!--inputs-->
+| Name   | Description     | Required | Default   |
+|--------|-----------------|----------|-----------|
+| input1 | The first input | `true`   | `default` |
+<!--/inputs-->
```

## Pre-commit (recommended)

You can also use the `gh-action-readme` extension as a [pre-commit](https://pre-commit.com/) hook to automatically update the `README.md` file before committing changes.

Add the following configuration to the `.pre-commit-config.yaml` file.

```yaml
repos:
  - repo: https://github.com/reakaleek/gh-action-readme
    rev: v0.4.0
    hooks:
      - id: action-readme
```

Then run the following command to install the pre-commit hook.

```bash
pre-commit install
```

> [!TIP]
> You can use the [pre-commit action](https://github.com/marketplace/actions/pre-commit) to run the pre-commit checks in your GitHub Actions workflow.


## GitHub Actions Monorepo support

The `gh-action-readme` extension supports monorepos with multiple actions in a single repository.
It will automatically detect the `action.yml` files in the repository and update the corresponding `README.md` files.


```bash

## Templating



<!--

Hello
-----



```markdown

### action.yml

Given an action.yml file:

```yaml
name: The Action
description: |
  An action that does something.
  It's a very useful action.

inputs:
  input1:
    description: The first input
    required: true
  input2:
    description: The second input
    required: false
    default: 'default'

outputs:
  output1:
    description: The first output

runs:
  # ...
```

### README.md

And a README.md file:

````diff
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
 - uses: org/repo@main
   with:
     input1: value1
     input2: value2
```
````

### Update Command

```bash
gh action-readme update
```

### Diff

````diff
-# <!--name--><!--/name-->
+# <!--name-->The Action<!--/name-->
<!--description-->
+An action that does something.
+It's a very useful action.
+<!--/description-->

## Inputs
<!--inputs-->
+| Name   | Description      | Required | Default   |
+|--------|------------------|----------|-----------|
+| input1 | The first input  | `true`   | ` `       |
+| input2 | The second input | `false`  | `default` |
+<!--/inputs-->

## Outputs
<!--outputs-->
+| Name    | Description      |
+|---------|------------------|
+| output1 | The first output |
+<!--/outputs-->

## Usage
<!--usage action="org/repo" version="v1"-->
```yaml
steps:
-  - uses: org/repo@main
+  - uses: org/repo@v1
   with:
     input1: value1
     input2: value2
```
````
