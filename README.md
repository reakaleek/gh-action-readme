
# gh-action-readme

![GitHub Release](https://img.shields.io/github/v/release/reakaleek/gh-action-readme?logo=github)
![GitHub Release Date](https://img.shields.io/github/release-date/reakaleek/gh-action-readme?display_date=published_at&logo=github)

![gh action-readme](./carbon.svg)

A GitHub CLI extension to inject GitHub Actions metadata into README.md markdown files.

## Prerequisites
- [GitHub CLI](https://cli.github.com/)

## ⚡️ Quick Start

### Install the `gh-action-readme` extension

```bash
gh extension install reakaleek/gh-action-readme
```

## Create a `README.md` file

```markdown
# <!--name--><!--/name-->
<!--description-->

## Inputs
<!--inputs-->

## Outputs
```

### Update the `README.md` file

Go to the action directory containing the `action.yml` file and run the following command:

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
