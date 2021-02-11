package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

// 将字符串md5加密
func Str2md5(str string) string {
	//方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// base64加密
func Base64Encode(str string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(str))
	return sEnc
}

// base64解密
func Base64Decode(str string) string {
	sDec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(sDec)
}
