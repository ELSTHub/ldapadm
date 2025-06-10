package ldap

import (
	"fmt"
	"github.com/spf13/viper"
	"ldapadm/utils"
	"strconv"
)

type GroupInfo struct {
	GroupName string
	GID       int
	UserList  []string
}

func AddGroup(info *GroupInfo) {
	if info.GID <= 0 {
		gid, err := utils.GetIdleGID()
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		info.GID = gid
	}

	var attributeMap = make(map[string][]string, 0)
	attributeMap["cn"] = []string{info.GroupName}
	attributeMap["gidNumber"] = []string{strconv.Itoa(info.GID)}
	if info.UserList != nil && len(info.UserList) > 0 {
		attributeMap["memberUid"] = info.UserList
	}
	attributeMap["objectClass"] = []string{"posixGroup", "top"}
	err := CreateLdapInfo(attributeMap, viper.GetString("ldap_server_conf.group_dn"))
	if err != nil {
		fmt.Printf("创建LDAP用户组异常：%v", err)
		return
	}
}

func DelGroup(info *GroupInfo) {
	err := DeleteLdapInfo(fmt.Sprintf("cn=%s,%s", info.GroupName, viper.GetString("ldap_server_conf.group_dn")))
	if err != nil {
		fmt.Printf("删除LDAP用户组异常：%v", err)
		return
	}
}

func UpdateGroup(info *GroupInfo) {
	var attributeMap = make(map[string][]string, 0)
	attributeMap["cn"] = []string{info.GroupName}
	if info.UserList != nil && len(info.UserList) > 0 {
		attributeMap["memberUid"] = info.UserList
	}
	attributeMap["objectClass"] = []string{"posixGroup", "top"}
	err := UpdateLdapInfo(attributeMap, fmt.Sprintf("cn=%s,%s", info.GroupName, viper.GetString("ldap_server_conf.group_dn")))
	if err != nil {
		fmt.Printf("创建LDAP用户组异常：%v", err)
		return
	}
}

func GetGroupInfo(info *GroupInfo) {
	var filter = "(objectClass=posixGroup)"
	entries, err := SearchLDAPInfo(fmt.Sprintf("cn=%s,%s", info.GroupName, viper.GetString("ldap_server_conf.group_dn")), filter)
	if err != nil {
		fmt.Printf("获取LDAP用户组异常：%v", err)
		return
	}
	for _, entry := range entries {
		for _, attribute := range entry.Attributes {
			fmt.Printf("%s: %v\n", attribute.Name, attribute.Values)
		}
	}
}
