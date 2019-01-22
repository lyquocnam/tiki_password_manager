package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tiki/app"
)

var loginCommand = &cobra.Command{
	Use:   "login username password",
	Short: "Login with username and password",
	Long:  `Login with username and password`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Print: " + strings.Join(args, " "))
		err := app.Login(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	},
}
