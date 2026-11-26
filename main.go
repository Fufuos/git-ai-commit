/*
Copyright © 2026 Nicholas Thong <kthong.nicholas@gmail.com>
*/
package main

import (
	"fmt"
	"os"

	"git-ai-commit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
