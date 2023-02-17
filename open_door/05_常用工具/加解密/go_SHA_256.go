package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func sha256ThenHex(data string) string {
	sha256Ctx := sha256.New()
	sha256Ctx.Write([]byte(data))
	return hex.EncodeToString(sha256Ctx.Sum(nil))
}

func main() {
	const user_password = "password142344@#$%&*&**("
	fmt.Println(user_password, " sha256 to [sha256]", sha256ThenHex(user_password))

}
