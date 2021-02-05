package version

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

// Version はVersion情報を保持する構造体
type Version struct {
	Version string
	Commit  string
	Edited  string
	Date    string
}

var (
	version = "given by LDFLAGS"
	commit  = "given by LDFLAGS"
	edited  = "given by LDFLAGS"
	date    = "given by LDFLAGS"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "shows version",
		Long:  "shows version.",
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
}
