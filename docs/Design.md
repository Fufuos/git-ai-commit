git-ai-commits

### Subcommands

- `--all` or `-a`: Automatically stage changes in tracked files for commit (default: false)
- `--clipboard` or `-c`: Copies selected message to clipboard instead of commiting
- `--generate` or `-g`: Number of messages to generate (default: 1), select one
- `--exclude` or `-x`: Files to exclude from AI analysis
- `--description`: Include a commit message description
- `--type` or `-t`: Git commit message format (default: plain). [`plain`, `conventional`, `conventional+body`, `subject+body`]
- `--prompt` or `-p`: Custom prompt to guide the LLM behavior (e.g. specific language, style convention)
- `--yes` or `-y`: Skip confirmation when commiting after message generation (default: false)

### Generate multiple messages

```bash
git ai-commit --generate <i> # or git ai-commits -g <i>
```

### Commit Message Formats

- plain (default): Simple, unstructured commit messages
- conventional: Conventional Commits format with type and scope
- conventional+body: Conventional Commit subject plus a generated body
- subject+body: Plain subject plus a generated body

Use --type to specify your format:

```bash
git ai-commit --type conventional # or -t conventional
git ai-commit --type conventional+body
git ai-commit --type subject+body
git ai-commit --type plain        # or -t plain (default)
```

### Custom Prompts

```bash
# Write commit messages in a specific language
ai-commit -p "Write commit messages in Italian"

# Focus on specific aspects of the changes
ai-commit -p "Focus on performance implications of changes"

# Use a specific style or tone
ai-commit -p "Use technical jargon suitable for senior developers"

# Include specific details in the message
ai-commit -p "Always mention the specific function names and file paths changed"
```
