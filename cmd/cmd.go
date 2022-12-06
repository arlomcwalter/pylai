package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

var (
	rootCmd = &cobra.Command{
		Use:   "pylai",
		Short: "A simple cli totp manager and authenticator.",
	}
	vFlag bool
)

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&vFlag, "version", "v", false, "Print the version number")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if vFlag {
			fmt.Println(version)
			return
		} else {
			err := cmd.Help()
			if err != nil {
				Quit("Error displaying help message.")
			}
		}
	}

	rootCmd.AddCommand(addCmd)
	addNameFlag(addCmd)
	addSecretFlag(addCmd)

	rootCmd.AddCommand(deleteCmd)
	addNameFlag(deleteCmd)

	rootCmd.AddCommand(exportCmd)
	addNameFlag(exportCmd)
	addAllFlag(exportCmd)

	rootCmd.AddCommand(listCmd)

	rootCmd.AddCommand(genCmd)
	addNameFlag(genCmd)

	rootCmd.AddCommand(verifyCmd)
	addNameFlag(verifyCmd)
	addOtpFlag(verifyCmd)

	if err := rootCmd.Execute(); err != nil {
		Quit("Error executing command.")
	}
}

func addNameFlag(c *cobra.Command) {
	c.Flags().StringP("name", "n", "", "The nickname for the account.")
}

func addSecretFlag(c *cobra.Command) {
	c.Flags().StringP("secret", "s", "", "The secret for the account.")
}

func addOtpFlag(c *cobra.Command) {
	c.Flags().StringP("code", "c", "", "The one-time code to verify.")
}
