package ach

import (
	"fmt"
	"log"
	"os/user"
	"path"
	"path/filepath"
	"strings"

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

	user, err := user.Current()
	if err != nil {
		log.Fatal(fmt.Errorf("NewAchCmd: %w", err))
	}

	defaultConfigFile := path.Join(user.HomeDir, ".ach", "config.yaml")

	cmd.PersistentFlags().StringVar(&cfgFile, "config", defaultConfigFile, "config file (default is $HOME/.ach/config.yaml)")
	cmd.PersistentFlags().StringVar(&taskCfgFile, "task-config", "./achTaskConfig.yaml", "task config file (default is ./achTaskConfig.yaml")

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
	v := viper.New()

	absCfgFile, err := filepath.Abs(cfgFile)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to convert config to its absolute path: %w", err))
	}

	v.SetConfigName(strings.TrimSuffix(absCfgFile, path.Ext(absCfgFile)))
	v.AddConfigPath("/")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("failed to read app config %s: %w", absCfgFile, err))
	}

	if err := v.UnmarshalExact(&config.GlobalAppConfig); err != nil {
		log.Fatal(fmt.Errorf("failed to parse app config %s: %w", absCfgFile, err))
	}

	config.GlobalAppConfig.ConfigDir = filepath.Dir(absCfgFile)
}

func readTaskConfig() {
	v := viper.New()

	absTaskCfgFile, err := filepath.Abs(taskCfgFile)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to convert task config to its absolute path: %w", err))
	}

	v.SetConfigName(strings.TrimSuffix(absTaskCfgFile, path.Ext(absTaskCfgFile)))
	v.AddConfigPath("/")

	if err := v.ReadInConfig(); err != nil {
		return
	}

	if err := v.UnmarshalExact(&config.GlobalTaskConfig); err != nil {
		log.Fatal(fmt.Errorf("failed to parse app config %s: %w", absTaskCfgFile, err))
	}
}
