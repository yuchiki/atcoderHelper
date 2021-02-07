package create

import (
	"fmt"
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
		Long: `creates contest directory.
Temporally, current template directory is hard-coded as $HOME/projects/private/atcoder/D
D is for directory.
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			contestName := args[0]
			err := exec.Command("mkdir", contestName).Run()
			if err != nil {
				return err
			}

			for _, taskName := range taskNames {
				taskDirName := path.Join(contestName, taskName)
				if output, err := exec.Command("cp", "-r", templateDirName, taskDirName).Output(); err != nil {
					fmt.Print(output)
					return err
				}

				sampleDirName := path.Join(taskDirName, "sampleCases")
				if output, err := exec.Command("mkdir", sampleDirName).Output(); err != nil {
					fmt.Print(output)
					return err
				}

				for i := 1; i <= 5; i++ {
					inputFileName := path.Join(sampleDirName, fmt.Sprintf("case%d.input", i))
					output, err := exec.Command("bash", "-c", fmt.Sprintf(`echo "[skip ach test]" > %s`, inputFileName)).Output()
					if err != nil {
						fmt.Printf("%s can not be initialized", inputFileName)
						fmt.Print(output)
						return err
					}
					outputFileName := path.Join(sampleDirName, fmt.Sprintf("case%d.expected", i))
					err = exec.Command("touch", outputFileName).Run()
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(useDefaultTemplate, "default-template", "d", false, "(required) use default contest template")

	if cmd.MarkFlagRequired("default-template") != nil {
		log.Fatal("default-template require")
	}

	return cmd
}
