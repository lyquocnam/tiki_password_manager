package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var newUserCommand = &cobra.Command{
	Use:   "new username password",
	Short: "create new user with username and password",
	Long:  `create new user with username and password`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}
