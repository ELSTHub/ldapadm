package utils

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func GetIdleUID() (int, error) {
	lock, err := UidAcquireLock()
	if err != nil {
		fmt.Println("GetIdleUID AcquireLock Error:", err)
		os.Exit(0)
	}
	defer lock.Unlock()
	var id = 0
	numByte, err := os.ReadFile(viper.GetString("ldap_adm.uid"))
	if err != nil {
		id = viper.GetInt("ldap_adm.min_uid")
	}
	file, err := os.OpenFile(viper.GetString("ldap_adm.uid"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	numByte = bytes.Replace(numByte, []byte("\n"), []byte(""), -1)
	numByte = bytes.Replace(numByte, []byte("\r"), []byte(""), -1)
	id, err = strconv.Atoi(string(numByte))
	if err != nil {
		id = viper.GetInt("ldap_adm.min_uid")
	}
	defer file.WriteString(strconv.Itoa(id + 1))
	return id, nil
}

func GetIdleGID() (int, error) {
	lock, err := GidAcquireLock()
	if err != nil {
		fmt.Println("GetIdleGID AcquireLock Error:", err)
		os.Exit(0)
	}
	defer lock.Unlock()
	var id = 0
	numByte, err := os.ReadFile(viper.GetString("ldap_adm.gid"))
	if err != nil {
		id = viper.GetInt("ldap_adm.min_gid")
	}
	file, err := os.OpenFile(viper.GetString("ldap_adm.gid"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	numByte = bytes.Replace(numByte, []byte("\n"), []byte(""), -1)
	numByte = bytes.Replace(numByte, []byte("\r"), []byte(""), -1)
	id, err = strconv.Atoi(string(numByte))
	if err != nil {
		id = viper.GetInt("ldap_adm.min_gid")
	}
	defer file.WriteString(strconv.Itoa(id + 1))
	return id, nil
}
