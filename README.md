# ai-commit

A golang cli rewrite of aicommits that writes git commit messages for you with AI.

## How It Works

The model receives the staged file list, requests the staged diffs it needs, and submits a validated commit subject with an optional description. The tool can only read Git's staged snapshot, so unstaged edits are never sent to the model.

## Setup

```bash
go install github.com/Fufuos/git-ai-commit@latest
export OPENAI_API_KEY="your-api-key"
```

The CLI uses `gpt-4o-mini` by default and generates a Conventional Commit subject.
To use another OpenAI-compatible endpoint or model, set `OPENAI_BASE_URL` and
`OPENAI_MODEL` as shown below. The staged diff is sent to that endpoint for
generation.

Configuration is read from environment variables; API keys are not written to disk.

## Usage

```
git add <files>
git ai-commit
```

Review the generated subject and confirm to create the commit. Use `--dry-run` to
print the subject without committing, `--yes` to skip confirmation, `--all` to
stage tracked changes first, or `--hint "focus on error handling"` to guide the
message. It can also be used as a Git subcommand with the `git-` prefix.

## Environment Variables

```bash
export OPENAI_API_KEY="sk-..."
export OPENAI_BASE_URL="https://api.example.com"
export OPENAI_MODEL="gpt-4o-mini"
```

`OPENAI_BASE_URL` may be a provider's API base URL (for example,
`https://api.openai.com/v1`) or a full `/chat/completions` endpoint URL.

## Configuration

### Viewing current Configuration

```bash
git ai-commit config
```

This will display only non-default configuration values with API keys masked for security. If no custom configuration is set, it will show "(using all default values)".

### Changing current model

```bash
git ai-commit model
```

This will:

- Show your current provider and model
- Fetch available models from your provider's API
- Let you select from available models or enter a custom model name
- Update your configuration automatically

### Updating git-ai-commit

```bash
git ai-commit update
```

Will retrieve the latest version on github:
?go get install @latest

### Reading Configuration Values

```bash
git ai-commit config get <key>
```

```bash
git ai-commit config list
```

```bash
git ai-commit config set <key>=<value>
```
