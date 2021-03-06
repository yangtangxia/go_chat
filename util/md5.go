package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//小写md5 加密
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//大写md5 加密
func MD5Encode(data string) string  {
	return strings.ToUpper(Md5Encode(data))
}

//小写加密
func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

//验证
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd + salt) == passwd
}





