package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "api server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

var (
	cfg *viper.Viper
)

func initConfig() {
	cfg = viper.New()
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath("./config/")
	cfg.SetConfigName("conf")

	err := cfg.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalln("failed to bind flags", err)
	}

	err = cfg.ReadInConfig()
	if err != nil {
		log.Fatalln("failed to read config after merge", cfg.ConfigFileUsed(), err)
	}
}
