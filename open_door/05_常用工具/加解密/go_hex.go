package main

import (
	"encoding/hex"
	"fmt"
)

func hexEncode(data string) string {
	return hex.EncodeToString([]byte(data))
}

func hexDecode(data string) string {
	hexD, hexE := hex.DecodeString(data)
	if nil != hexE {
		panic("hex decode fail")
	}
	return string(hexD)
}

func main() {
	const user_password = "password142344@#$%&*&**("

	fmt.Println("dohex " + user_password + "=" + hexEncode(user_password))
	fmt.Println("unhex " + user_password + "=" + hexDecode(hexEncode(user_password)))

}
