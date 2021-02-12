package ach

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/contest"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/test"
	"github.com/yuchiki/atcoderHelper/internal/cmd/ach/version"
	"github.com/yuchiki/atcoderHelper/internal/config"
)

var (
	cfgFile     string
	taskCfgFile string
)

// NewAchCmd returns ach command.
func NewAchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ach",
		Short: "ach automates routine work you does when you participate AtCoder contests",
		Long:  `ach automates routine work you does when you participate AtCoder contests. `,
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ach/config.yaml)")
	cmd.PersistentFlags().StringVar(&taskCfgFile, "task-config", "", "task config file (default is ./achTaskConfig.yaml")

	registerSubcommands(cmd)

	return cmd
}

func registerSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(version.NewVersionCmd())
	cmd.AddCommand(contest.NewContestCmd())
	cmd.AddCommand(test.NewTestCmd())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	readAppConfig()
	readTaskConfig()
}

func readAppConfig() {
	var configFileName string

	if cfgFile != "" {
		viper.SetConfigName(cfgFile)
		configFileName = cfgFile
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(path.Join(home, ".ach"))
		viper.SetConfigName("config")
		configFileName = "config"
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("failed to read app config %s: %w", configFileName, err))
	}

	if err := viper.UnmarshalExact(&config.GlobalAppConfig); err != nil {
		log.Fatal(fmt.Errorf("failed to parse app config %s: %w", configFileName, err))
	}
}

func readTaskConfig() {
	var configFileName string

	if taskCfgFile != "" {
		viper.SetConfigFile(taskCfgFile)

		configFileName = cfgFile
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("achTaskConfig")
		configFileName = "achTaskConfig"
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Print(fmt.Errorf("failed to read task config %s: %w", configFileName, err))
	}

	if err := viper.UnmarshalExact(&config.GlobalTaskConfig); err != nil {
		log.Fatal(fmt.Errorf("failed to parse app config %s: %w", configFileName, err))
	}
}
