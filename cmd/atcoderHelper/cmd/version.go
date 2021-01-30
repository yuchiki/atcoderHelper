package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

type Version struct {
	Version string
	Commit  string
	Edited  string
	Date    string
}

var (
	version = "Given By LDFLAGS"
	commit  = "Given By LDGLAGS"
	edited  = "Given By LDFLAGS"
	date    = "Given By LDFLAGS"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version.",
	RunE: func(cmd *cobra.Command, args []string) error {
		Build := Version{
			Version: version,
			Commit:  commit,
			Edited:  edited,
			Date:    date,
		}
		bytes, err := json.Marshal(&Build)
		if err != nil {
			return err
		}
		cmd.Println(string(bytes))
		return nil
	},
}
