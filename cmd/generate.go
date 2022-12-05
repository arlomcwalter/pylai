package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/arlomcwalter/pylai/util"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a TOTP code.",
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

		code, err := util.GenerateTotp(account.Secret)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s Your code is: %s\n", Green("✔"), Bold(code))
	},
}
