package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
