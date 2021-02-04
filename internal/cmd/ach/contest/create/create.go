package create

import (
	"os/exec"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

func NewContestCreateCmd() *cobra.Command {
	user, _ := user.Current()
	templateDirName := path.Join(user.HomeDir, "projects", "private", "atcoder", "D")
	taskNames := []string{"A", "B", "C", "D", "E", "F"}

	useDefaultTemplate := new(bool)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "show version",
		Long:  "show version.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			contestName := args[0]
			exec.Command("mkdir", contestName).Run()
			for _, taskName := range taskNames {
				taskDirName := path.Join(contestName, taskName)
				exec.Command("cp", "-r", templateDirName, taskDirName).Run()
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(useDefaultTemplate, "default-template", "d", false, "use default contest template")
	cmd.MarkFlagRequired("default-template")

	return cmd
}
