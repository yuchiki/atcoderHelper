package ach

import (
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
	if cfgFile != "" {
		viper.SetConfigName(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		viper.AddConfigPath(path.Join(home, ".ach"))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.UnmarshalExact(&config.GlobalAppConfig); err != nil {
		log.Fatal(err)
	}
}

func readTaskConfig() {
	if taskCfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("achTaskConfig")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Print(err)
	}

	if err := viper.UnmarshalExact(&config.GlobalTaskConfig); err != nil {
		log.Fatal(err)
	}
}
