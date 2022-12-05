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
			panic(err)
		}

		var account *database.Account

		if isPresent(nickname) {
			account, err = database.Get(nickname)
			if err != nil {
				panic(err)
			}
		} else {
			account = accountSelector()
		}

		if getConfirm() {
			if err := account.Delete(); err != nil {
				panic(err)
			}
			fmt.Printf("%s Deleted %s.\n", Green("✔"), Bold(account.Nickname))
		} else {
			fmt.Println(Red("Aborted delete operation."))
		}
	},
}
