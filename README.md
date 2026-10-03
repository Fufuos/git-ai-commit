# ai-commit

A golang cli rewrite of aicommits that writes git commit messages for you with AI.

## How It Works

The model receives the staged file list, requests the staged diffs it needs, and submits a validated commit subject with an optional description. The tool can only read Git's staged snapshot, so unstaged edits are never sent to the model.

## Setup

```bash
go install https://github.com/Fufuos/git-ai-commit
ai-commit setup
```

- Select your AI provider
- Configure your API key
- Automatically fetch and select from available models (if supported)
- Choose your message format [`plan`, `conventional`, `conventional+body`, `subject+body`]

Supported providers:

- OpenAI
- Groq
- OpenRouter
- Ollama (local)
- Gemini

All configurations is kept in `.aicommit` in your `~/home` directory

## Usage

```
git add <files>
git ai-commit # or aic or gaic
```

can be used natively as a git subcommand with the `git-` prefix.

## Environment Variables

```bash
export OPENAI_API_KEY="sk-..."
export OPENAI_BASE_URL="https://api.example.com"
export OPENAI_MODEL="gpt-4"
```

Configuration precedence:

- Command-line arguments
- Environment variables
- Configuration file
- Default values

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
