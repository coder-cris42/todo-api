package utils

import (
	"log"
	"todo-api/internal/infrastructure/api"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	pwdBytes := []byte(password)
	hashed_Salted, err := bcrypt.GenerateFromPassword(pwdBytes, 7)
	if err != nil {
		log.Println(err)
		return api.ERROR_CODE_HASH_PASS_FAILED
	}
	return string(hashed_Salted)
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
