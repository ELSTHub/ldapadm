package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ldapadm/Version"
	"ldapadm/config"
	"os"
)

var mainCmd = &cobra.Command{
	Use:     "ldapadm",
	Short:   "LDAP administrator CLI tool",
	Long:    "LDAP administrator CLI tool",
	Version: Version.VERSION,
}

func Execute() {
	if err := mainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	err := config.ConfInit()
	if err != nil {
		fmt.Printf("load config file error:%v", err)
		return
	}

	mainCmd.AddCommand(userCmd)
	mainCmd.AddCommand(groupCmd)
	mainCmd.AddCommand(passwdCmd)
}
