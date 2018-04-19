package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccsv",
	Short: "Export contributor user data to a csv",
	Run: func(cmd *cobra.Command, args []string) {
		client := github.NewClient(nil)
		contributors, _, err := client.Repositories.ListContributors(context.Background(), args[0], args[1], nil)
		var userIds []int64

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, contributor := range contributors {
			userIds = append(userIds, *contributor.ID)
		}

		var users []github.User
		for _, id := range userIds {
			user, _, err := client.Users.GetByID(context.Background(), id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			users = append(users, *user)
		}

		gocsv.Marshal(&users, os.Stdout)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
