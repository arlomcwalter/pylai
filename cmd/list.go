package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts.",
	Run: func(cmd *cobra.Command, args []string) {
		accounts, err := database.GetAll()
		if err != nil {
			Quit("Error getting accounts from database.")
		}

		count := len(accounts)
		if count < 1 {
			fmt.Println("Couldn't find any accounts in the database.")
			return
		}

		fmt.Printf("%s %s\n", Bold("Accounts"), Faint(fmt.Sprintf("(%d)", len(accounts))))

		for _, account := range accounts {
			fmt.Printf("%s %s\n", Bold("▸"), account.Nickname)
		}
	},
}
