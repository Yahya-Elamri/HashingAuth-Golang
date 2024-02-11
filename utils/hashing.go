package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hashing(str string) (string, error) {
	HashedString, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", fmt.Errorf("faild to hash the text : %s", err)
	}
	return string(HashedString), nil
}

func CompairHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {
		return false
	}
	return true
}
