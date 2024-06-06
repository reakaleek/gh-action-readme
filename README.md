
# gh-action-readme
![GitHub Release](https://img.shields.io/github/v/release/reakaleek/gh-action-readme?logo=github)
![GitHub Release Date](https://img.shields.io/github/release-date/reakaleek/gh-action-readme?display_date=published_at&logo=github)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/reakaleek/gh-action-readme)


A GitHub CLI extension to inject GitHub Actions metadata into README.md markdown files.

## Prerequisites
- [GitHub CLI](https://cli.github.com/)

## ⚡️ Quick Start

### Install the gh-action-readme extension

The `gh-action-readme` extension can be installed using the following command.

```bash
gh extension install reakaleek/gh-action-readme
```

## Create a README.md file

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


## 💡 How it works


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
