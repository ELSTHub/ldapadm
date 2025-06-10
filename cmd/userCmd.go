package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ldapadm/ldap"
	"os"
)

var userOpts = new(ldap.UserInfo)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
}

var addUserCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can add user.")
			return
		}
		ldap.AddUser(userOpts)
	},
}

var delUserCmd = &cobra.Command{
	Use:   "del",
	Short: "Del a user",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can del user.")
			return
		}
		ldap.DelUser(userOpts)
	},
}

var modifyUserCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify a user",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can modify user.")
			return
		}
		ldap.UpdateUser(userOpts)
	},
}

var showUserInfo = &cobra.Command{
	Use:   "show",
	Short: "Show a user info",
	Run: func(cmd *cobra.Command, args []string) {
		ldap.GetUserInfo(userOpts)
	},
}

func init() {
	if os.Getuid() == 0 {
		addUserCmd.Flags().StringVarP(&userOpts.Username, "username", "U", "", "Username")
		addUserCmd.Flags().StringVarP(&userOpts.Password, "password", "p", "", "Password")
		addUserCmd.Flags().StringVarP(&userOpts.Group, "group", "G", "", "Group")
		addUserCmd.Flags().IntVarP(&userOpts.UID, "uid", "u", -1, "UID")
		addUserCmd.Flags().IntVarP(&userOpts.GID, "gid", "g", -1, "GID")
		addUserCmd.Flags().StringVarP(&userOpts.HomeDir, "home_dir", "d", "", "Home Dir")
		addUserCmd.Flags().StringVarP(&userOpts.Shell, "shell", "s", "", "Shell")
		addUserCmd.Flags().StringVarP(&userOpts.ExpireAt, "expire", "e", "", "Expire Date Example: 2006-01-02T15:04:05")
		userCmd.AddCommand(addUserCmd)
		delUserCmd.Flags().StringVarP(&userOpts.Username, "username", "U", "", "Username")
		userCmd.AddCommand(delUserCmd)
		modifyUserCmd.Flags().StringVarP(&userOpts.Username, "username", "U", "", "Username")
		modifyUserCmd.Flags().StringVarP(&userOpts.Password, "password", "p", "", "Password")
		modifyUserCmd.Flags().StringVarP(&userOpts.Group, "group", "G", "", "Group")
		modifyUserCmd.Flags().IntVarP(&userOpts.UID, "uid", "u", -1, "UID")
		modifyUserCmd.Flags().IntVarP(&userOpts.GID, "gid", "g", -1, "GID")
		modifyUserCmd.Flags().StringVarP(&userOpts.HomeDir, "home_dir", "d", "", "Home Dir")
		modifyUserCmd.Flags().StringVarP(&userOpts.Shell, "shell", "s", "", "Shell")
		addUserCmd.Flags().StringVarP(&userOpts.ExpireAt, "expire", "e", "", "Expire Date Example: 2006-01-02T15:04:05")
		userCmd.AddCommand(modifyUserCmd)
		showUserInfo.Flags().StringVarP(&userOpts.Username, "username", "U", "", "Username")
		userCmd.AddCommand(showUserInfo)
	} else {
		showUserInfo.Flags().StringVarP(&userOpts.Username, "username", "U", "", "Username")
		userCmd.AddCommand(showUserInfo)
	}
}
