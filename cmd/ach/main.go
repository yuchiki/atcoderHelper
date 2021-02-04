package main

import (
	"fmt"
	"os"

	"github.com/yuchiki/atcoderHelper/internal/cmd/ach"
)

func main() {
	cmd := ach.NewAchCmd()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
