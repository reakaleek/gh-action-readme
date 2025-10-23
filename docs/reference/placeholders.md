# Placeholder Reference

Complete reference for all placeholder tags supported by gh-action-readme.

## Overview

Placeholders are special HTML comment tags that mark sections where gh-action-readme should inject generated content. Content between opening and closing tags is replaced with fresh data from `action.yml`.

## Placeholder Syntax

All placeholders follow this format:

```markdown
<!--placeholder-name-->
Generated content goes here
<!--/placeholder-name-->
```

**Rules:**
- Opening tag: `<!--name-->`
- Closing tag: `<!--/name-->`
- Content between tags is replaced by generated content
- Whitespace in tags is ignored (e.g., `<!-- name -->` works too)
- Placeholders must have both opening and closing tags
- Placeholders can appear anywhere in the README

## Available Placeholders

### name

Injects the action's name from `action.yml`.

**Usage:**
```markdown
# <!--name--><!--/name-->
```

**Generated:**
```markdown
# <!--name-->My Awesome Action<!--/name-->
```

**Source in action.yml:**
```yaml
name: My Awesome Action
```

**Notes:**
- Often used in the main heading
- Can be used inline or as standalone text
- Common pattern: `# <!--name--><!--/name-->` (self-closing on same line)

---

### description

Injects the action's description.

**Usage:**
```markdown
<!--description-->
<!--/description-->
```

**Generated:**
```markdown
<!--description-->
This action does something awesome.
It supports multiple features.
<!--/description-->
```

**Source in action.yml:**
```yaml
description: |
  This action does something awesome.
  It supports multiple features.
```

**Notes:**
- Preserves multi-line descriptions
- Respects line breaks from action.yml
- Often placed directly under the main heading

---

### inputs

Generates a table of action inputs with their descriptions, requirements, and defaults.

**Usage:**
```markdown
## Inputs
<!--inputs-->
<!--/inputs-->
```

**Generated:**
```markdown
## Inputs
<!--inputs-->
| Name        | Description              | Required | Default      |
|-------------|--------------------------|----------|--------------|
| `input1`    | The first input          | `true`   | ` `          |
| `input2`    | The second input         | `false`  | `default123` |
| `api-token` | API authentication token | `true`   | ` `          |
<!--/inputs-->
```

**Source in action.yml:**
```yaml
inputs:
  input1:
    description: 'The first input'
    required: true
  input2:
    description: 'The second input'
    required: false
    default: 'default123'
  api-token:
    description: 'API authentication token'
    required: true
```

**Table Columns:**

| Column | Description | Values |
|--------|-------------|--------|
| Name | Input identifier in code format | Wrapped in backticks |
| Description | Input description from action.yml | As-is from YAML |
| Required | Whether input is required | `true` or `false` in backticks |
| Default | Default value if any | Value in backticks, ` ` if no default |

**Notes:**
- Automatically creates a markdown table
- Input names are wrapped in backticks for code formatting
- Missing defaults show as a single space in backticks: ` `
- Inputs appear in the order defined in action.yml
- Empty table is generated if no inputs exist

---

### outputs

Generates a table of action outputs with their descriptions.

**Usage:**
```markdown
## Outputs
<!--outputs-->
<!--/outputs-->
```

**Generated:**
```markdown
## Outputs
<!--outputs-->
| Name      | Description                    |
|-----------|--------------------------------|
| `result`  | The operation result           |
| `status`  | Status code of the operation   |
| `message` | Human-readable status message  |
<!--/outputs-->
```

**Source in action.yml:**
```yaml
outputs:
  result:
    description: 'The operation result'
  status:
    description: 'Status code of the operation'
  message:
    description: 'Human-readable status message'
```

**Table Columns:**

| Column | Description | Values |
|--------|-------------|--------|
| Name | Output identifier in code format | Wrapped in backticks |
| Description | Output description from action.yml | As-is from YAML |

**Notes:**
- Simpler than inputs table (no Required or Default columns)
- Output names wrapped in backticks
- Empty table if no outputs defined
- Outputs appear in definition order

---

### usage

Injects code examples with automatic version tracking.

**Usage:**
````markdown
<!--usage action="org/repo" version="v1.0.0"-->
```yaml
steps:
  - uses: org/repo@v1.0.0
    with:
      input: value
```
<!--/usage-->
````

**Attributes:**

| Attribute | Required | Description | Example |
|-----------|----------|-------------|---------|
| `action` | Yes | Action reference path | `org/repo` or `org/repo/path` |
| `version` | Yes | Version string or env reference | `v1.0.0` or `env:VERSION` |

**Version Tracking:**

The tool automatically updates version references within the code block:

**Before (in your template):**
````markdown
<!--usage action="myorg/myaction" version="v1.5.0"-->
```yaml
steps:
  - uses: myorg/myaction@v1.5.0
```
<!--/usage-->
````

**After you change version to v2.0.0:**
````markdown
<!--usage action="myorg/myaction" version="v2.0.0"-->
```yaml
steps:
  - uses: myorg/myaction@v2.0.0
```
<!--/usage-->
````

**Environment Variable Version:**

Use `env:VARIABLE` to reference environment variables:

````markdown
<!--usage action="org/repo" version="env:VERSION"-->
```yaml
steps:
  - uses: org/repo@${{ env.VERSION }}
```
<!--/usage-->
````

**Notes:**
- Automatically updates all references to the action with the correct version
- Useful for keeping examples synchronized with releases
- Can have multiple usage blocks with different versions
- The usage block preserves your custom example code
- Only the version references are updated

---

### toc

Generates a table of contents from document headings.

**Usage:**
```markdown
<!--toc-->
<!--/toc-->
```

**Generated:**
```markdown
<!--toc-->
- [Installation](#installation)
- [Usage](#usage)
  - [Basic Usage](#basic-usage)
  - [Advanced Usage](#advanced-usage)
- [Configuration](#configuration)
- [Examples](#examples)
<!--/toc-->
```

**Notes:**
- Automatically generates from markdown headings (`#`, `##`, `###`)
- Creates anchor links to each section
- Respects heading hierarchy (indentation)
- Excludes the main title (first `#` heading)
- Regenerates on each update

---

## Multiple Placeholders

You can use multiple instances of the same placeholder:

```markdown
# <!--name--><!--/name-->

<!--description-->

## What is <!--name--><!--/name-->?

The <!--name--><!--/name--> action helps you...
```

All instances are updated with the same content.

## Placeholder Nesting

Placeholders cannot be nested inside each other:

❌ **Invalid:**
```markdown
<!--description-->
  <!--name-->
  <!--/name-->
<!--/description-->
```

✅ **Valid:**
```markdown
<!--name--><!--/name-->
<!--description-->
<!--/description-->
```

## Whitespace Handling

### In Placeholder Tags

Whitespace in tags is ignored:

```markdown
<!--inputs-->     <!-- Same as <!--inputs--> -->
<!-- inputs -->   <!-- Same as <!--inputs--> -->
<!--  inputs  --> <!-- Same as <!--inputs--> -->
```

### In Generated Content

- Tables are properly formatted with consistent spacing
- Descriptions preserve line breaks from action.yml
- Empty lines are preserved where appropriate

## Custom Content

Content outside placeholder tags is never modified:

```markdown
# <!--name--><!--/name-->

<!-- This custom description is preserved -->
This is my custom introduction that gh-action-readme will never touch.

## Custom Section

This entire section is preserved.

## Inputs
<!--inputs-->
<!-- Only this section is modified -->
<!--/inputs-->

## More Custom Content

This is also preserved!
```

## Best Practices

### 1. Always use closing tags

❌ **Missing closing tag:**
```markdown
<!--inputs-->
```

✅ **Proper tags:**
```markdown
<!--inputs-->
<!--/inputs-->
```

### 2. Place placeholders logically

```markdown
# <!--name--><!--/name-->
<!--description-->

## Inputs
<!--inputs-->

## Outputs
<!--outputs-->

## Usage
<!--usage action="org/repo" version="v1"-->
```

### 3. Add custom content around placeholders

```markdown
## Inputs

The following inputs are available:

<!--inputs-->
<!--/inputs-->

> **Note:** All inputs support environment variable substitution.
```

### 4. Use clear section headings

```markdown
## 📥 Inputs
<!--inputs-->
<!--/inputs-->

## 📤 Outputs  
<!--outputs-->
<!--/outputs-->
```

### 5. Keep usage examples realistic

````markdown
<!--usage action="myorg/myaction" version="v1.0.0"-->
```yaml
name: Real World Example
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: myorg/myaction@v1.0.0
        with:
          input: ${{ secrets.API_KEY }}
```
<!--/usage-->
````

## Common Patterns

### Minimal README

```markdown
# <!--name--><!--/name-->
<!--description-->

## Inputs
<!--inputs-->

## Outputs
<!--outputs-->
```

### Comprehensive README

````markdown
# <!--name--><!--/name-->

<!--description-->

## Table of Contents
<!--toc-->
<!--/toc-->

## Inputs
<!--inputs-->
<!--/inputs-->

## Outputs
<!--outputs-->
<!--/outputs-->

## Basic Usage
<!--usage action="org/repo" version="v1"-->
```yaml
steps:
  - uses: org/repo@v1
```
<!--/usage-->

## Advanced Usage
<!--usage action="org/repo" version="v1"-->
```yaml
steps:
  - uses: org/repo@v1
    with:
      advanced: true
```
<!--/usage-->
````

### README with Custom Sections

````markdown
# <!--name--><!--/name-->

<!--description-->

## Features

- 🚀 Fast and efficient
- 🔒 Secure by default
- 📦 Easy to use

## Inputs
<!--inputs-->
<!--/inputs-->

## Outputs
<!--outputs-->
<!--/outputs-->

## Examples

### Example 1: Basic
<!--usage action="org/repo" version="v1"-->
```yaml
steps:
  - uses: org/repo@v1
```
<!--/usage-->

### Example 2: Advanced
Custom example here (not using placeholders)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)
````

## Troubleshooting

### Placeholder not being replaced

**Check:**
1. Both opening and closing tags present
2. Tags spelled correctly
3. No extra characters in tags
4. Placeholder name is valid (see list above)

### Content keeps getting removed

**Cause:** Content is inside placeholder tags

**Solution:** Move custom content outside placeholders

### Version not updating in usage block

**Check:**
1. `action` attribute matches the actual uses: statement
2. Version format is consistent
3. Usage block has proper `<!--usage-->` tags

### Table formatting looks wrong

**Cause:** Manual edits inside placeholder tags

**Solution:** Run `gh action-readme update` to regenerate

## See Also

- [Commands Reference](./commands.md) - Available commands
