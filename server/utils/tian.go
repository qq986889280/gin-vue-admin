package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"hash"
)

func Json_encode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

func Json_decode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

func Hash(src string, CryptTool string, isDoubleCrypt bool) string {
	var hash hash.Hash
	switch CryptTool {
	case "md5":
		hash = md5.New()
	case "sha1":
		hash = sha1.New()
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	}
	hash.Write([]byte(src))
	//一次加密后的结果
	cryptStr := fmt.Sprintf("%x", hash.Sum(nil))
	if isDoubleCrypt {
		//做二次加密
		hash.Reset()
		hash.Write([]byte(cryptStr))
		cryptStr = fmt.Sprintf("%x", hash.Sum(nil))
	}
	return cryptStr
}
