package cmd

import (
	"fmt"
	"github.com/lyquocnam/tiki_password_manager/app"
	"github.com/spf13/cobra"
	"log"
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
			log.Fatal(err)
		} else {
			fmt.Println("Login successfully")
		}
	},
}

var createUserCommand = &cobra.Command{
	Use:   "create [username] [password]",
	Short: "create new user with username and password",
	Long:  `create new user with username and password`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := app.CreateUser(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Create user successfully")
		}
	},
}

var changePasswordCommand = &cobra.Command{
	Use:   "change_password [current_password] [new_password]",
	Short: "change user's password with current password and new password",
	Long:  `change user's password with current password and new password`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := app.ChangePassword(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("password has been changed !")
		}
	},
}


var vaLidatePasswordCommand = &cobra.Command{
	Use:   "validate_password [password]",
	Short: "validate password",
	Long:  `validate password`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ok, err := app.User.ValidatePassword(args[0])
		if err != nil {
			log.Fatal(err)
		} else if ok {
			fmt.Println("password is ok")
		} else {
			fmt.Println("password is not valid")
		}
	},
}

