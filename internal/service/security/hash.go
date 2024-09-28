package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashString(str string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckHashString(str, strHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(strHash), []byte(str))
	return err == nil
}
