package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports an existing account's secret.",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		if all {
			accounts, err := database.GetAll()
			if err != nil {
				Quit("Error getting accounts from database.")
			}

			for _, account := range accounts {
				printAccount(account)
			}
		} else {
			nickname, err := cmd.Flags().GetString("name")
			if err != nil {
				Quit("Error reading flag.")
			}

			var account *database.Account

			if isPresent(nickname) {
				account, err = database.Get(nickname)
				if err != nil {
					Quit("Error getting account from database.")
				}
			} else {
				account = accountSelector()
			}

			printAccount(account)
		}
	},
}

func printAccount(account *database.Account) {
	fmt.Printf("%s: %s\n", Bold(account.Nickname), account.Secret)
}

func addAllFlag(c *cobra.Command) {
	c.Flags().BoolP("all", "a", false, "Exports all accounts.")
}
