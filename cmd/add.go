package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new account.",
	Run: func(cmd *cobra.Command, args []string) {
		nickname, err := cmd.Flags().GetString("name")
		if err != nil {
			Quit("Error reading flag.")
		}

		secret, err := cmd.Flags().GetString("secret")
		if err != nil {
			Quit("Error reading flag.")
		}

		if !isPresent(nickname) && !isPresent(secret) {
			nickname = getNewName()
			secret = getSecret()
		}

		account := &database.Account{
			Nickname: nickname,
			Secret:   secret,
		}

		if err := account.Save(); err != nil {
			Quit("Error saving account.")
		}

		fmt.Printf("%s Added %s.\n", Green("✔"), Bold(account.Nickname))
	},
}
