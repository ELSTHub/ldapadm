package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"ldapadm/utils/crypt/md5_crypt"
	"ldapadm/utils/crypt/sha256_crypt"
	"ldapadm/utils/crypt/sha512_crypt"
	"strings"
)

// EncipherLdapPass 加密LDAP密码
func EncipherLdapPass(passwd, encryption string) string {
	var result string
	switch strings.ToLower(encryption) {
	case "md5":
		result = ToMD5Pass(passwd)
	case "md5-crypt":
		result = ToMD5CryptPass(passwd)
	case "smd5":
		result = ToSMD5Pass(passwd)
	case "sha1", "sha":
		result = ToSHA1Pass(passwd)
	case "ssha":
		result = ToSSHAPass(passwd)
	case "sha256", "sha-256", "sha-256-crypt":
		result = ToSHA256CryptPass(passwd)
	case "sha512", "sha-512", "sha-512-crypt":
		result = ToSHA512CryptPass(passwd)
	default:
		result = passwd
	}
	return result
}

// CheckLdapPass 校验LDAP密码
func CheckLdapPass(passwd, ciphertext, encryption string) bool {
	var result bool
	switch strings.ToLower(encryption) {
	case "md5", "{md5}":
		result = CheckMD5Pass(passwd, ciphertext)
	case "md5-crypt", "{crypt}$1":
		result = CheckMD5CryptPass(passwd, ciphertext)
	case "smd5", "{smd5}":
		result = CheckSMD5Pass(passwd, ciphertext)
	case "sha1", "sha", "{sha1}", "{sha}":
		result = CheckSHA1Pass(passwd, ciphertext)
	case "ssha", "{ssha}":
		result = CheckSSHAPass(passwd, ciphertext)
	case "sha256", "sha-256", "sha-256-crypt", "{crypt}$5":
		result = CheckSHA256CryptPass(passwd, ciphertext)
	case "sha512", "sha-512", "sha-512-crypt", "{crypt}$6":
		result = CheckSHA512CryptPass(passwd, ciphertext)
	default:
		result = passwd == ciphertext
	}
	return result
}

// GetEncryption 获取加密方式
func GetEncryption(ciphertext string) string {
	var result string
	if len(ciphertext) <= 6 {
		return result
	}
	if ciphertext[0] != '{' {
		return result
	}
	if ciphertext[4] == '}' {
		result = ciphertext[:5]
	}
	if ciphertext[5] == '}' {
		result = ciphertext[:6]
	}
	result = strings.ReplaceAll(result, "{", "")
	result = strings.ReplaceAll(result, "}", "")
	return result
}

func ToMD5Pass(passwd string) string {
	md5sum := md5.Sum([]byte(passwd))
	base64Encode := base64.StdEncoding.EncodeToString(md5sum[:])
	return "{MD5}" + base64Encode
}

func CheckMD5Pass(passwd, ciphertext string) bool {
	if len(ciphertext) < 5 || ciphertext[:5] != "{MD5}" {
		return false
	}
	md5sum := md5.Sum([]byte(passwd))
	base64Encode := base64.StdEncoding.EncodeToString(md5sum[:])
	return base64Encode == ciphertext[5:]
}

func ToSMD5Pass(passwd string) string {
	salt := make([]byte, 8)
	combined := append([]byte(passwd), salt...)
	md5sum := md5.Sum(combined)
	smd5 := append(md5sum[:], salt...)
	base64Encode := base64.StdEncoding.EncodeToString(smd5)
	return "{SMD5}" + base64Encode
}

func CheckSMD5Pass(passwd, ciphertext string) bool {
	if len(ciphertext) < 6 || ciphertext[:6] != "{SMD5}" {
		return false
	}
	smd5, err := base64.StdEncoding.DecodeString(ciphertext[6:])
	if err != nil {
		return false
	}

	if len(smd5) <= md5.Size {
		return false
	}
	hash := smd5[:md5.Size]
	salt := smd5[md5.Size:]
	combined := append([]byte(passwd), salt...)
	md5sum := md5.Sum(combined)
	return string(md5sum[:]) == string(hash)
}

func ToSHA1Pass(passwd string) string {
	sha1sum := sha1.Sum([]byte(passwd))
	base64Encode := base64.StdEncoding.EncodeToString(sha1sum[:])
	return "{SHA}" + base64Encode
}

func CheckSHA1Pass(passwd, ciphertext string) bool {
	if len(ciphertext) < 5 || ciphertext[:5] != "{SHA}" {
		return false
	}
	sha1sum := sha1.Sum([]byte(passwd))
	base64Encode := base64.StdEncoding.EncodeToString(sha1sum[:])
	return base64Encode == ciphertext[5:]
}

func ToSSHAPass(passwd string) string {
	salt := make([]byte, 12)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}
	sha := sha1.New()
	sha.Write([]byte(passwd))
	sha.Write(salt)
	sshasum := sha.Sum(nil)
	base64Encode := base64.StdEncoding.EncodeToString(sshasum)
	return "{SSHA}" + base64Encode
}

func CheckSSHAPass(passwd, ciphertext string) bool {
	if len(ciphertext) < 6 || ciphertext[:6] != "{SSHA}" {
		return false
	}
	ssha, err := base64.StdEncoding.DecodeString(ciphertext[6:])
	if err != nil {
		return false
	}
	if len(ssha) < sha1.Size {
		return false
	}
	hash := ssha[:sha1.Size]
	salt := ssha[sha1.Size:]

	sha := sha1.New()
	sha.Write([]byte(passwd))
	sha.Write(salt)
	inHash := sha.Sum(nil)

	return string(inHash) == string(hash)
}

func ToSHA512CryptPass(passwd string) string {
	sha512 := sha512_crypt.New()
	generate, err := sha512.Generate([]byte(passwd), nil)
	if err != nil {
		return ""
	}
	return "{CRYPT}" + generate
}

func CheckSHA512CryptPass(passwd, ciphertext string) bool {
	if len(ciphertext) < 7 || ciphertext[:7] != "{CRYPT}" {
		return false
	}
	sha512Ciphertext := ciphertext[7:]
	sha512 := sha512_crypt.New()
	err := sha512.Verify(sha512Ciphertext, []byte(passwd))
	if err != nil {
		return false
	}
	return true
}

func ToSHA256CryptPass(passwd string) string {
	sha256 := sha256_crypt.New()
	generate, err := sha256.Generate([]byte(passwd), nil)
	if err != nil {
		return ""
	}
	return "{CRYPT}" + generate
}

func CheckSHA256CryptPass(passwd, ciphertext string) bool {
	if len(ciphertext) < 7 || ciphertext[:7] != "{CRYPT}" {
		return false
	}
	sha256Ciphertext := ciphertext[7:]
	sha256 := sha256_crypt.New()
	err := sha256.Verify(sha256Ciphertext, []byte(passwd))
	if err != nil {
		return false
	}
	return true
}

func ToMD5CryptPass(passwd string) string {
	md5 := md5_crypt.New()
	generate, err := md5.Generate([]byte(passwd), nil)
	if err != nil {
		return ""
	}
	return "{CRYPT}" + generate
}

func CheckMD5CryptPass(passwd, ciphertext string) bool {
	if len(ciphertext) < 7 || ciphertext[:7] != "{CRYPT}" {
		return false
	}
	md5Ciphertext := ciphertext[7:]
	md5 := md5_crypt.New()
	err := md5.Verify(md5Ciphertext, []byte(passwd))
	if err != nil {
		return false
	}
	return true
}
