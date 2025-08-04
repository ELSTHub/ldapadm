package ldap

import (
	"fmt"
	"github.com/spf13/viper"
	"ldapadm/utils"
	"strconv"
	"time"
)

type UserInfo struct {
	Username        string
	Password        string
	UID             int
	GID             int
	Group           string
	HomeDir         string
	Shell           string
	ExpireAt        string
	AutoCreateGroup bool
}

func AddUser(info *UserInfo) {
	if info.UID <= 0 {
		uid, err := utils.GetIdleUID()
		if err != nil {
			fmt.Printf("GetIdleUID err: %v\n", err)
			return
		}
		info.UID = uid
	}

	if info.GID < 0 {
		gid, err := utils.GetIdleGID()
		if err != nil {
			fmt.Printf("GetIdleGID err: %v\n", err)
			return
		}
		info.GID = gid
	}

	if info.AutoCreateGroup {
		var groupInfo = new(GroupInfo)
		if info.Group == "" {
			info.Group = info.Username
		}
		groupInfo.GroupName = info.Group
		groupInfo.GID = info.GID
		groupInfo.UserList = make([]string, 0)
		groupInfo.UserList = append(groupInfo.UserList, info.Username)
		AddGroup(groupInfo)
	}

	var attributeMap = make(map[string][]string, 0)
	attributeMap["uid"] = []string{info.Username}
	attributeMap["uidNumber"] = []string{strconv.Itoa(info.UID)}
	attributeMap["sn"] = []string{info.Username}
	attributeMap["gecos"] = []string{info.Username}
	attributeMap["cn"] = []string{info.Group}
	attributeMap["gidNumber"] = []string{strconv.Itoa(info.GID)}
	if info.HomeDir == "" {
		attributeMap["homeDirectory"] = []string{fmt.Sprintf("%s/%s", viper.GetString("ldap_server_conf.default_home_path"), info.Username)}
	} else {
		attributeMap["homeDirectory"] = []string{info.HomeDir}
	}
	if info.Password != "" {
		attributeMap["userPassword"] = []string{info.Password}
	}

	if info.Shell == "" {
		attributeMap["loginShell"] = []string{viper.GetString("ldap_server_conf.default_bash")}
	} else {
		attributeMap["loginShell"] = []string{info.Shell}
	}
	attributeMap["objectClass"] = []string{"posixAccount", "top", "inetOrgPerson", "shadowAccount"}
	if info.ExpireAt != "" {
		expireAt, err := time.Parse("2006-01-02T15:04:05", info.ExpireAt)
		if err != nil {
			fmt.Printf("Expire date format err: %v\n", err)
			return
		}
		expire := expireAt.Sub(time.UnixMilli(0)).Hours() / 24
		attributeMap["shadowExpire"] = []string{strconv.Itoa(int(expire))}
	}
	err := CreateLdapInfo(attributeMap, viper.GetString("ldap_server_conf.user_dn"))
	if err != nil {
		fmt.Printf("创建LDAP账号异常：%v", err)
		return
	}
}

func DelUser(info *UserInfo) {
	err := DeleteLdapInfo(fmt.Sprintf("uid=%s,%s", info.Username, viper.GetString("ldap_server_conf.user_dn")))
	if err != nil {
		fmt.Printf("删除LDAP账号异常：%v", err)
		return
	}
}

func UpdateUser(info *UserInfo) {
	var attributeMap = make(map[string][]string, 0)
	attributeMap["uidNumber"] = []string{strconv.Itoa(info.UID)}
	attributeMap["sn"] = []string{info.Username}
	attributeMap["gecos"] = []string{info.Username}
	attributeMap["cn"] = []string{info.Group}
	if info.UID > 0 {
		attributeMap["uidNumber"] = []string{strconv.Itoa(info.UID)}
	}
	if info.GID >= 0 {
		attributeMap["gidNumber"] = []string{strconv.Itoa(info.GID)}
	}
	if info.HomeDir != "" {
		attributeMap["homeDirectory"] = []string{info.HomeDir}
	}
	if info.Password != "" {
		attributeMap["userPassword"] = []string{info.Password}
	}

	if info.Shell != "" {
		attributeMap["loginShell"] = []string{info.Shell}
	}
	//attributeMap["objectClass"] = []string{"posixAccount", "top", "inetOrgPerson", "shadowAccount"}
	if info.ExpireAt != "" {
		expireAt, err := time.Parse("2006-01-02T15:04:05", info.ExpireAt)
		if err != nil {
			fmt.Printf("Expire date format err: %v\n", err)
			return
		}
		expire := expireAt.Sub(time.UnixMilli(0)).Hours() / 24
		attributeMap["shadowExpire"] = []string{strconv.Itoa(int(expire))}
	}
	err := UpdateLdapInfo(attributeMap, fmt.Sprintf("uid=%s,%s", info.Username, viper.GetString("ldap_server_conf.user_dn")))
	if err != nil {
		fmt.Printf("创建LDAP账号异常：%v", err)
		return
	}
}

func ModifyPassword(info *UserInfo) {
	var attributeMap = make(map[string][]string, 0)
	if info.Password != "" {
		attributeMap["userPassword"] = []string{info.Password}
	}
	err := UpdateLdapInfo(attributeMap, fmt.Sprintf("uid=%s,%s", info.Username, viper.GetString("ldap_server_conf.user_dn")))
	if err != nil {
		fmt.Printf("创建LDAP账号异常：%v", err)
		return
	}
}

func GetUserInfo(info *UserInfo) {
	var filter = "(objectClass=posixAccount)"
	entries, err := SearchLDAPInfo(fmt.Sprintf("uid=%s,%s", info.Username, viper.GetString("ldap_server_conf.user_dn")), filter)
	if err != nil {
		fmt.Printf("获取LDAP账号异常：%v", err)
		return
	}
	for _, entry := range entries {
		for _, attribute := range entry.Attributes {
			if attribute.Name == "userPassword" {
				continue
			}
			fmt.Printf("%s: %v\n", attribute.Name, attribute.Values)
		}
	}
}
