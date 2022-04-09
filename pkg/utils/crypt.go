package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	
	return string(passwordBytes), err
}

func CheckPassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

    return err == nil
}