package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/arlomcwalter/pylai/util"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies that a TOTP code matches the account.",
	Run: func(cmd *cobra.Command, args []string) {
		nickname, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		code, err := cmd.Flags().GetString("code")
		if err != nil {
			panic(err)
		}

		var account *database.Account

		if isPresent(nickname) && isPresent(code) {
			account, err = database.Get(nickname)
			if err != nil {
				panic(err)
			}
		} else {
			account = accountSelector()
			code = getOtp()
		}

		if util.VerifyTotp(code, account.Secret) {
			fmt.Println("That code is valid.")
		} else {
			fmt.Println("That code is invalid.")
		}
	},
}
