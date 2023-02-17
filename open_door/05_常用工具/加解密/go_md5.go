package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	const user_password = "password142344@#$%&*&**("
	fmt.Println(user_password, " md5 to [md5]", Md5ThenHex(user_password))
}

func Md5ThenHex(data string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(data))                     //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}
