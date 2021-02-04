package create

import (
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

func NewContestCreateCmd() *cobra.Command {
	taskNames := []string{"A", "B", "C", "D", "E", "F"}

	templateDir := new(string)

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
				exec.Command("cp", "-r", *templateDir, taskDirName).Run()
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(templateDir, "templateDir", "t", "", "contest template")
	cmd.MarkFlagRequired("templateDir")

	return cmd
}
