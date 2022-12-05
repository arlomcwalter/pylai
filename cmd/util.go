package cmd

import (
	"fmt"
	"github.com/arlomcwalter/pylai/database"
	"github.com/arlomcwalter/pylai/util"
	"github.com/manifoldco/promptui"
	"strings"
)

var (
	Bold  = promptui.Styler(promptui.FGBold)
	Green = promptui.Styler(promptui.FGGreen)
	Red   = promptui.Styler(promptui.FGRed)
	Faint = promptui.Styler(promptui.FGFaint)
)

func isPresent(arg string) bool {
	return strings.TrimSpace(arg) != ""
}

func getConfirm() bool {
	prompt := promptui.Prompt{
		Label:     "Are you sure",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return false
	}

	return true
}

func accountSelector() *database.Account {
	accounts, err := database.GetAll()
	if err != nil {
		panic(err)
	}

	var names []string
	for _, account := range accounts {
		names = append(names, account.Nickname)
	}

	prompt := promptui.Select{
		Label: "Account",
		Items: names,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	account, err := database.Get(result)
	if err != nil {
		panic(err)
	}

	return account
}

func getNewName() string {
	prompt := promptui.Prompt{
		Label: "Nickname",
		Validate: func(input string) error {
			input = strings.TrimSpace(input)

			if !util.IsValidNickname(input) {
				return fmt.Errorf("invalid nickname format or length")
			}

			existingAcc, err := database.Get(input)
			if err == nil && existingAcc != nil {
				return fmt.Errorf("account with nickname '%s' already exists", input)
			}

			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}

func getSecret() string {
	prompt := promptui.Prompt{
		Label: "TOTP Secret",
		Validate: func(input string) error {
			if util.IsValidSecret(input) {
				return nil
			} else {
				return fmt.Errorf("invalid TOTP format")
			}
		},
	}

	result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}

func getOtp() string {
	prompt := promptui.Prompt{
		Label: "OTP Code",
		Validate: func(input string) error {
			if util.IsValidOtp(input) {
				return nil
			} else {
				return fmt.Errorf("invalid TOTP format")
			}
		},
	}

	result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return result
}
