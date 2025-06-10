package ldap

import (
	"fmt"
	"ldapadm/config"
	"testing"
)

func TestGetGroupInfo(t *testing.T) {
	err := config.ConfInit()
	if err != nil {
		fmt.Printf("load config file error:%v", err)
		return
	}

	var groupInfo = new(GroupInfo)
	groupInfo.GroupName = "elst"
	GetGroupInfo(groupInfo)
}
