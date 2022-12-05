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
			panic(err)
		}

		secret, err := cmd.Flags().GetString("secret")
		if err != nil {
			panic(err)
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
			panic(err)
		}

		fmt.Printf("%s Added %s.\n", Green("✔"), Bold(account.Nickname))
	},
}
