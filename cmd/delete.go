package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete an account.",
	Run: func(cmd *cobra.Command, args []string) {
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

		if getConfirm() {
			if err := account.Delete(); err != nil {
				Quit("Error deleting account.")
			}
			fmt.Printf("%s Deleted %s.\n", Green("✔"), Bold(account.Nickname))
		} else {
			fmt.Println(Red("Aborted delete operation."))
		}
	},
}
