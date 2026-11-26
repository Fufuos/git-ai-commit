package git

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type DiffContext struct {
	RawDiff    string
	Stat       string
	Files      []string
	RecentLogs []string
}

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) EnsureRepo(ctx context.Context) error {
	out, err := exec.CommandContext(ctx, "git", "rev-parse", "--is-inside-work-tree").Output()
	if err != nil || strings.TrimSpace(string(out)) != "true" {
		return fmt.Errorf("not a git repository")
	}
	return nil
}

func (r *Repo) GetStagedDiff(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "diff", "--cached")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git diff failed: %s", string(out))
	}
	return string(out), nil
}

func (r *Repo) StageTrackedChanges(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "git", "add", "-u")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git add failed: %s", strings.TrimSpace(string(out)))
	}
	return nil
}

func (r *Repo) Commit(ctx context.Context, message string) error {
	cmd := exec.CommandContext(ctx, "git", "commit", "-m", message)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git commit failed: %s", string(out))
	}
	fmt.Print(string(out))
	return nil
}
