package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "application",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use -h to show available commands")
	},
}

func Run() {
	rootCmd.AddCommand(graphQlCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
