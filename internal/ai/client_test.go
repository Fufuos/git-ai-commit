package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGenerateCommitSubject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("authorization header = %q", r.Header.Get("Authorization"))
		}
		var request chatCompletionRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode request: %v", err)
		}
		if request.Model != "test-model" {
			t.Errorf("model = %q, want test-model", request.Model)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"feat(cli): generate commit messages"}}]}`))
	}))
	defer server.Close()

	client := &Client{apiKey: "test-key", model: "test-model", url: server.URL, http: server.Client()}
	subject, err := client.GenerateCommitSubject(context.Background(), "diff --git a/main.go", "")
	if err != nil {
		t.Fatalf("GenerateCommitSubject() error = %v", err)
	}
	if subject != "feat(cli): generate commit messages" {
		t.Fatalf("subject = %q", subject)
	}
}

func TestGenerateCommitSubjectRejectsInvalidOutput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"Updated the CLI"}}]}`))
	}))
	defer server.Close()

	client := &Client{apiKey: "test-key", model: "test-model", url: server.URL, http: server.Client()}
	if _, err := client.GenerateCommitSubject(context.Background(), "diff", ""); err == nil {
		t.Fatal("expected invalid Conventional Commit output to fail")
	}
}

func TestNewFromEnvLoadsDotEnvAndPreservesShellValues(t *testing.T) {
	t.Chdir(t.TempDir())
	restoreEnv := preserveEnv(t, "OPENAI_API_KEY", "OPENAI_BASE_URL", "OPENAI_MODEL")
	defer restoreEnv()

	if err := os.WriteFile(".env", []byte("OPENAI_API_KEY=file-key\nOPENAI_BASE_URL=https://openrouter.ai/api/v1\nOPENAI_MODEL=provider/example\n"), 0600); err != nil {
		t.Fatalf("write .env: %v", err)
	}
	if err := os.Setenv("OPENAI_API_KEY", "shell-key"); err != nil {
		t.Fatalf("set shell API key: %v", err)
	}

	client, err := NewFromEnv()
	if err != nil {
		t.Fatalf("NewFromEnv() error = %v", err)
	}
	if client.apiKey != "shell-key" {
		t.Errorf("apiKey = %q, want shell value", client.apiKey)
	}
	if client.url != "https://openrouter.ai/api/v1/chat/completions" {
		t.Errorf("url = %q", client.url)
	}
	if client.model != "provider/example" {
		t.Errorf("model = %q, want dotenv value", client.model)
	}
}

func preserveEnv(t *testing.T, keys ...string) func() {
	t.Helper()

	values := make(map[string]string, len(keys))
	present := make(map[string]bool, len(keys))
	for _, key := range keys {
		values[key], present[key] = os.LookupEnv(key)
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("unset %s: %v", key, err)
		}
	}
	return func() {
		for _, key := range keys {
			if present[key] {
				_ = os.Setenv(key, values[key])
			} else {
				_ = os.Unsetenv(key)
			}
		}
	}
}
