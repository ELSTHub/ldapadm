package ldap

import (
	"fmt"
	"ldapadm/config"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	err := config.ConfInit()
	if err != nil {
		fmt.Printf("load config file error:%v", err)
		return
	}

	var userInfo = new(UserInfo)
	userInfo.Username = "elst"
	GetUserInfo(userInfo)
}
