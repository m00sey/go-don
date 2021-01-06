package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/m00sey/go-don/pkg/api"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Run:   runStart,
}

func runStart(cmd *cobra.Command, args []string) {
	srv, err := api.New(cfg)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}

	err = srv.Launch()
	if err != nil {
		log.Fatalln("launch fail", err)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
}
