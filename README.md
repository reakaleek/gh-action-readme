# gh-action-readme

A GitHub CLI extension to inject GitHub Actions metadata into README.md markdown files.

## Install

```bash
gh extension install reakaleek/gh-action-readme
```

## Quick Start

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

