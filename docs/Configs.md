### OPENAI_API_KEY

Your OpenAI API key or custom provider API Key

### OPENAI_BASE_URL

Custom OpenAI-compatible API endpoint URL.

### OPENAI_MODEL

Model to use for OpenAI-compatible providers.

### provider

The selected AI provider. Set automatically during git ai-commit setup. Valid values: openai, togetherai, groq, xai, openrouter, ollama, lmstudio, custom.

### locale

Default: `en`

The locale to use for the generated commit messages. Consult the list of codes in: https://wikipedia.org/wiki/List_of_ISO_639-1_codes.

### generate

Default: `1`
The number of commit messages to generate to pick from.
\*Note, this will use more tokens as it generates more results.

### timeout

The timeout for network requests to the OpenAI API in milliseconds.
Default: `10000` (10 seconds) for hosted providers, 60000 (60 seconds) for Together AI and local providers.

```bash
git ai-commit config set timeout=20000 # 20s
```

### max-length

The preferred character length of the generated commit subject. This guides the model toward concise messages, but complete subjects may be longer.
Default: `72`

```bash
git ai-commit config set max-length=100
```

### type

Default: plain
The type of commit message to generate. Available options:

- plain: Simple, unstructured commit messages
- conventional: Conventional Commits format with type and scope
- conventional+body: Conventional Commit subject plus a generated body
- subject+body: Plain subject plus a generated body

Examples:

```bash
git ai-commit config set type=conventional
git ai-commit config set type=conventional+body
git ai-commit config set type=subject+body
git ai-commit config set type=plain
```
