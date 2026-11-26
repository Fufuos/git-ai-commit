package cmd

import (
	"bufio"
	"fmt"
	"git-ai-commit/internal/ai"
	"git-ai-commit/internal/git"
	"strings"

	"github.com/spf13/cobra"
)

type cliFlags struct {
	all    bool
	dryRun bool
	hint   string
	yes    bool
}

var flags cliFlags

var rootCmd = &cobra.Command{
	Use:   "git-ai-commit",
	Short: "Generate concise Git commit messages from staged diffs and file context",
	Long: `Generate a Conventional Commit subject from the staged Git diff, then
review and create the commit directly from your terminal.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          runRoot,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&flags.all, "all", "a", false, "Stage tracked changes before generating the message")
	rootCmd.Flags().BoolVar(&flags.dryRun, "dry-run", false, "Print the generated message without committing")
	rootCmd.Flags().StringVar(&flags.hint, "hint", "", "Additional guidance for the generated message")
	rootCmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "Commit without asking for confirmation")
}

func runRoot(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	repo := git.New()
	if err := repo.EnsureRepo(ctx); err != nil {
		return err
	}
	if flags.all {
		if err := repo.StageTrackedChanges(ctx); err != nil {
			return err
		}
	}

	diff, err := repo.GetStagedDiff(ctx)
	if err != nil {
		return err
	}
	if strings.TrimSpace(diff) == "" {
		return fmt.Errorf("no staged changes; use git add or --all")
	}

	client, err := ai.NewFromEnv()
	if err != nil {
		return err
	}
	subject, err := client.GenerateCommitSubject(ctx, diff, flags.hint)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", subject); err != nil {
		return fmt.Errorf("write commit message: %w", err)
	}
	if flags.dryRun {
		return nil
	}

	confirmed, err := confirmCommit(cmd)
	if err != nil {
		return err
	}
	if !confirmed {
		return nil
	}

	return repo.Commit(ctx, subject)
}

func confirmCommit(cmd *cobra.Command) (bool, error) {
	if flags.yes {
		return true, nil
	}
	if _, err := fmt.Fprint(cmd.OutOrStdout(), "Create commit? [Y/n] "); err != nil {
		return false, fmt.Errorf("write confirmation prompt: %w", err)
	}
	answer, err := bufio.NewReader(cmd.InOrStdin()).ReadString('\n')
	if err != nil && len(answer) == 0 {
		return false, fmt.Errorf("could not read confirmation: %w", err)
	}
	answer = strings.ToLower(strings.TrimSpace(answer))
	if answer != "" && answer != "y" && answer != "yes" {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Commit cancelled")
		return false, nil
	}
	return true, nil
}
