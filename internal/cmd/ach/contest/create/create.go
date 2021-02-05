package create

import (
	"log"
	"os/exec"
	"os/user"
	"path"

	"github.com/spf13/cobra"
)

// NewContestCreateCmd returns a new contest create command.
func NewContestCreateCmd() *cobra.Command {
	user, _ := user.Current()
	templateDirName := path.Join(user.HomeDir, "projects", "private", "atcoder", "D")
	taskNames := []string{"A", "B", "C", "D", "E", "F"}

	useDefaultTemplate := new(bool)

	cmd := &cobra.Command{
		Use:   "create [contestName]",
		Short: "creates contest directory",
		Long:  "creates contest directory.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			contestName := args[0]
			err := exec.Command("mkdir", contestName).Run()
			if err != nil {
				return err
			}

			for _, taskName := range taskNames {
				taskDirName := path.Join(contestName, taskName)

				err := exec.Command("cp", "-r", templateDirName, taskDirName).Run()
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(useDefaultTemplate, "default-template", "d", false, "use default contest template")

	if cmd.MarkFlagRequired("default-template") != nil {
		log.Fatal("default-template require")
	}

	return cmd
}
