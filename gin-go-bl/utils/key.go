package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

const hash = 10

func GetMd5(str string) (sMd5 string) {
	data := []byte(str)
	has := md5.Sum(data)
	sMd5 = fmt.Sprintf("%x", has)
	//sMd5 = strings.ToUpper(md5str)
	return
}

func ScryptPassword(password string) (out string) {
	salt := make([]byte, 8)
	salt = []byte{1, 2, 3, 4, 5, 6, 7, 8}

	hapw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, hash)
	if err != nil {
		fmt.Println(err.Error())
	}
	out = base64.StdEncoding.EncodeToString(hapw)
	return
}
