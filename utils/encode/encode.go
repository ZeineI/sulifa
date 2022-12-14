package encode

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(password string) string {
	return fmt.Sprintf("%v", sha256.Sum256([]byte(password)))
}

func ComparePasswordHash(pass1, pass2 string) bool {
	passFromClient := GenerateHash(pass2)
	fmt.Println(passFromClient, pass1)
	return passFromClient == pass1
}
