package create

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
	cmdtest "github.com/yuchiki/atcoderHelper/internal/cmd/ach/test"
	"github.com/yuchiki/atcoderHelper/internal/config"
	yaml "gopkg.in/yaml.v2"
)

var openEditor = new(bool)

// NewContestCreateCmd returns a new contest create command.
func NewContestCreateCmd() *cobra.Command {
	useDefaultTemplate := new(bool)

	cmd := &cobra.Command{
		Use:   "create [contestName]",
		Short: "creates contest directory",
		Long: `creates contest directory.
Temporarily, current template directory is hard-coded as $HOME/projects/private/atcoder/D
D is for directory.
		`,
		Args: cobra.ExactArgs(1),
		RunE: runE,
	}

	cmd.Flags().BoolVarP(useDefaultTemplate, "default-template", "d", false, "(required) use default contest template")
	cmd.Flags().BoolVar(openEditor, "open-editor", true, "open editor for each task")

	if cmd.MarkFlagRequired("default-template") != nil {
		log.Fatal("default-template require")
	}

	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	template, err := config.GetDefaultTemplate()
	if err != nil {
		return err
	}

	taskNames := []string{"A", "B", "C", "D", "E", "F"}
	contestName := args[0]

	output, err := exec.Command("mkdir", contestName).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w", output, err)
	}

	absTemplateDir := getAbsTemplateDirectory(template.TemplateDirectory)

	for _, taskName := range taskNames {
		if err := initializeTaskDirectory(absTemplateDir, contestName, taskName); err != nil {
			return err
		}
	}

	for _, taskName := range taskNames {
		taskDirName := path.Join(contestName, taskName)

		if *openEditor {
			cmd := exec.Command("bash", "-c", config.GlobalAppConfig.EditorCommand) //nolint:gosec // This is intended.
			cmd.Env = append(cmd.Env, fmt.Sprintf("TASK_PATH=%v", taskDirName))

			output, err := cmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("failed to start editor, message=\"%s\": %w", output, err)
			}
		}
	}

	return nil
}

func initializeTaskDirectory(absTemplateDir, contestName, taskName string) error {
	taskDirName := path.Join(contestName, taskName)
	if output, err := exec.Command("cp", "-r", absTemplateDir, taskDirName).CombinedOutput(); err != nil {
		fmt.Print(output)

		return fmt.Errorf("%s: %w", output, err)
	}

	taskConfig := config.TaskConfig{
		ContestID: contestName,
		TaskID:    taskName,
		Template:  config.GlobalAppConfig.DefaultTemplate,
	}

	taskConfigYaml, err := yaml.Marshal(taskConfig)
	if err != nil {
		return err
	}

	taskConfigName := path.Join(taskDirName, "achTaskConfig.yaml")

	taskConfigFile, err := os.Create(taskConfigName)
	if err != nil {
		return err
	}
	defer taskConfigFile.Close()

	_, err = taskConfigFile.Write(taskConfigYaml)
	if err != nil {
		return err
	}

	err = createTestcases(path.Join(taskDirName, cmdtest.TestcasesFile))
	if err != nil {
		return err
	}

	return nil
}

func createTestcases(testcasesFile string) error {
	if err := ioutil.WriteFile(testcasesFile, []byte("testcases: []"), 0o666); err != nil { //nolint:gosec
		return err
	}

	return nil
}

func getAbsTemplateDirectory(templateDir string) string {
	if path.IsAbs(templateDir) {
		return templateDir
	}

	return path.Join(config.GlobalAppConfig.ConfigDir, templateDir)
}
