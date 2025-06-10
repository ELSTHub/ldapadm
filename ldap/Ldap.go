package ldap

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
	"os"
	"regexp"
)

func InitLDAP() (*ldap.Conn, error) {
	host := viper.GetString("ldap_server_conf.host")
	port := viper.GetInt("ldap_server_conf.port")
	ldapUser := viper.GetString("ldap_server_conf.login_dn")
	ldapPass := viper.GetString("ldap_server_conf.password")
	ldapConn, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	err = ldapConn.Bind(ldapUser, ldapPass)
	if err != nil {
		return nil, err
	}
	return ldapConn, nil
}

func CreateLdapInfo(attributeMap map[string][]string, dn string) error {
	ldapConn, err := InitLDAP()
	if err != nil {
		fmt.Printf("ldap init err: %v\n", err)
		os.Exit(0)
	}
	if val, ok := attributeMap["objectClass"]; ok {
		for _, v := range val {
			if v == "posixAccount" {
				dn = fmt.Sprintf("uid=%s,%s", attributeMap["uid"][0], dn)
			}
			if v == "organizationalUnit" {
				dn = fmt.Sprintf("ou=%s,%s", attributeMap["ou"][0], dn)
			}
			if v == "posixGroup" {
				dn = fmt.Sprintf("cn=%s,%s", attributeMap["cn"][0], dn)
			}
		}
	}
	addRequest := ldap.NewAddRequest(dn, []ldap.Control{})
	for key, val := range attributeMap {
		addRequest.Attribute(key, val)
	}
	err = ldapConn.Add(addRequest)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLdapInfo(dn string) error {
	ldapConn, err := InitLDAP()
	if err != nil {
		fmt.Printf("ldap init err: %v\n", err)
		os.Exit(0)
	}
	// baseDN格式验证
	baseDNCompile, err := regexp.Compile(`^(\S+=\S+,)*(dc=\S*,)+dc=\S*$`)
	if err != nil {
		return err
	}
	baseStatus := baseDNCompile.MatchString(dn)
	if !baseStatus {
		return errors.New("ABNORMAL DATA FORMAT")
	}

	delRequest := ldap.NewDelRequest(dn, []ldap.Control{})

	err = ldapConn.Del(delRequest)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLdapInfo(attributeMap map[string][]string, dn string) error {
	ldapConn, err := InitLDAP()
	if err != nil {
		fmt.Printf("ldap init err: %v\n", err)
		os.Exit(0)
	}
	replace := ldap.NewModifyRequest(dn, []ldap.Control{})
	for key, val := range attributeMap {
		replace.Replace(key, val)
	}

	err = ldapConn.Modify(replace)
	if err != nil {
		return err
	}
	return nil
}

func SearchLDAPInfo(baseDN, filter string) ([]*ldap.Entry, error) {
	ldapConn, err := InitLDAP()
	if err != nil {
		fmt.Printf("ldap init err: %v\n", err)
		os.Exit(0)
	}
	// baseDN格式验证
	baseDNCompile, err := regexp.Compile(`^(\S+=\S+,)*(dc=\S*,)+dc=\S*$`)
	if err != nil {
		return nil, err
	}
	baseStatus := baseDNCompile.MatchString(baseDN)
	if !baseStatus {
		return nil, errors.New("ABNORMAL DATA FORMAT")
	}

	if filter == "" {
		filter = "(objectClass=*)"
	}

	ldapSql := &ldap.SearchRequest{
		BaseDN:       baseDN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       filter,
		Attributes:   []string{},
		Controls:     nil,
	}
	result, err := ldapConn.Search(ldapSql)
	if err != nil {
		return nil, err
	}
	return result.Entries, nil
}
