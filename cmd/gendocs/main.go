package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/yuchiki/atcoderHelper/cmd/atcoderHelper/ach"
)

func main() {
	cmd := ach.NewAchCmd()

	err := doc.GenMarkdownTree(cmd, "docs/cmd")
	if err != nil {
		log.Fatal(err)
	}
}
