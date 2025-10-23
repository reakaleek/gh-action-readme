# gh-action-readme Documentation

Welcome to the gh-action-readme documentation. This tool helps you automatically generate and maintain README documentation for GitHub Actions by extracting metadata from `action.yml` files. **Perfect for single actions and monorepos with multiple actions.**

## Documentation Structure

This documentation is organized by topic to help you find what you need:

### 🚀 Getting Started

Learn the basics and see complete examples:

- [Getting Started](./tutorials/getting-started.md) - Quick start tutorial (5 minutes)
- [Building a Complete Action README](./tutorials/complete-action-readme.md) - Full example with all features

### ⚙️ Working with gh-action-readme

Integrate gh-action-readme into your workflow:

- [Set up pre-commit hooks](./how-to/setup-precommit.md) - Automatically update READMEs on commit
- [Manage monorepos with multiple actions](./how-to/manage-monorepos.md) - Bulk operations on multiple actions
- [Check for documentation drift](./how-to/check-documentation-drift.md) - Verify docs are up-to-date

### 📖 Reference

Technical documentation and command reference:

- [Command Reference](./reference/commands.md) - Complete command-line reference
- [Placeholder Reference](./reference/placeholders.md) - All available placeholders

## Quick Navigation

### New to gh-action-readme?
Start with the [Getting Started guide](./tutorials/getting-started.md) to learn the basics in 5 minutes.

### Working with monorepos?
See [Manage monorepos with multiple actions](./how-to/manage-monorepos.md) for bulk operations on multiple actions.

### Looking for command syntax?
The [Command Reference](./reference/commands.md) has complete details on all commands and flags.

## About

gh-action-readme is a GitHub CLI extension that automatically injects GitHub Actions metadata into README.md files, keeping your documentation synchronized with your action definitions.

**Key Features:**
- 🚀 Automatic documentation generation from `action.yml`
- 🔄 Keep docs in sync with pre-commit hooks
- 📦 **Built-in monorepo support** - manage multiple actions at once with `--recursive`
- ✅ Documentation drift detection
- 🎯 Mix automated and custom content seamlessly

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

gh-action-readme is released under the [MIT License](../LICENSE).
