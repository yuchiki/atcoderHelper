package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach"
)

func main() {
	cmd := ach.NewAchCmd()
	cmd.DisableAutoGenTag = true

	if err := doc.GenMarkdownTree(cmd, "docs/cmd"); err != nil {
		log.Fatal(err)
	}
}
