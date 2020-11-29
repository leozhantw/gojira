package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/leozhantw/gojira/internal/config"
)

var (
	cfg config.Config

	rootCmd = &cobra.Command{
		Use:           "gojira",
		Short:         "Save your time on jira operations",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if os.Args[1] == "configure" {
		return
	}

	var err error
	cfg, err = config.New()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln(fmt.Errorf("Config file not found. Please Execute `gojira configure` first."))
		}
		log.Fatalln(err)
	}
}
