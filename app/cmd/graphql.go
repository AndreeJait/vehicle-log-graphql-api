package cmd

import (
	"github.com/spf13/cobra"
	"vehicle-log-graphql-api/app/graphql"
)

var graphQlCmd = &cobra.Command{
	Use:     "graphql",
	Aliases: []string{"gql"},
	Run: func(cmd *cobra.Command, args []string) {
		gql := graphql.New()
		gql.Start()
	},
}
