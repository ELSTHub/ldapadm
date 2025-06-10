package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ldapadm/ldap"
	"os"
)

var groupOpts = new(ldap.GroupInfo)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage groups",
}

var addGroupCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new group",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can add group.")
			return
		}
		ldap.AddGroup(groupOpts)
	},
}

var delGroupCmd = &cobra.Command{
	Use:   "del",
	Short: "Del a group",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can del group.")
			return
		}
		ldap.DelGroup(groupOpts)
	},
}

var modifyGroupCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify a group",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Only root can modify group.")
			return
		}
		ldap.UpdateGroup(groupOpts)
	},
}

var showGroupInfo = &cobra.Command{
	Use:   "show",
	Short: "Show a user info",
	Run: func(cmd *cobra.Command, args []string) {
		ldap.GetGroupInfo(groupOpts)
	},
}

func init() {
	if os.Getuid() == 0 {
		addGroupCmd.Flags().StringVarP(&groupOpts.GroupName, "group_name", "G", "", "Group Name")
		addGroupCmd.Flags().IntVarP(&groupOpts.GID, "gid", "g", -1, "GID")
		addGroupCmd.Flags().StringArrayVarP(&groupOpts.UserList, "users", "U", []string{}, "User List")
		groupCmd.AddCommand(addGroupCmd)
		delGroupCmd.Flags().StringVarP(&groupOpts.GroupName, "group_name", "G", "", "Group Name")
		groupCmd.AddCommand(delGroupCmd)
		modifyGroupCmd.Flags().StringVarP(&groupOpts.GroupName, "group_name", "G", "", "Group Name")
		modifyGroupCmd.Flags().IntVar(&groupOpts.GID, "gid", -1, "GID")
		modifyGroupCmd.Flags().StringArrayVarP(&groupOpts.UserList, "users", "U", []string{}, "User List")
		groupCmd.AddCommand(modifyGroupCmd)
		showGroupInfo.Flags().StringVarP(&groupOpts.GroupName, "group_name", "G", "", "Group Name")
		groupCmd.AddCommand(showGroupInfo)
	} else {
		showGroupInfo.Flags().StringVarP(&groupOpts.GroupName, "group_name", "G", "", "Group Name")
		groupCmd.AddCommand(showGroupInfo)
	}
}
