package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ldapadm/ldap"
	"os"
	"os/user"
)

var passwdOpts = new(ldap.UserInfo)

var passwdCmd = &cobra.Command{
	Use:   "passwd",
	Short: "Manage passwd",
}

var modifyPasswdCmd = &cobra.Command{
	Use:   "modify",
	Short: "Non-root users can only modify their own password.",
	Run: func(cmd *cobra.Command, args []string) {
		current, err := user.Current()
		if err != nil {
			return
		}
		if current.Username != passwdOpts.Username && os.Getuid() != 0 {
			fmt.Println("only modify user.")
			return
		}
		ldap.UpdateUser(passwdOpts)
	},
}

func init() {
	modifyPasswdCmd.Flags().StringVar(&passwdOpts.Username, "username", "", "Username")
	modifyPasswdCmd.Flags().StringVar(&passwdOpts.Password, "password", "", "Password")
	passwdCmd.AddCommand(addUserCmd)
}
