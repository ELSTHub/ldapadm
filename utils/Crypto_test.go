package utils

import (
	"fmt"
	"testing"
)

func TestSSHA(t *testing.T) {
	var pass = "11111111"

	sshaPass := ToSSHAPass(pass)
	status := CheckSSHAPass(pass, sshaPass)
	fmt.Println(status)
}

func TestSHA256Crypt(t *testing.T) {
	var pass = "11111111"

	sshaPass := ToSHA256CryptPass(pass)
	status := CheckSHA256CryptPass(pass, sshaPass)
	fmt.Println(status)
}
