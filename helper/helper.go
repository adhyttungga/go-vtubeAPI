package helper

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	} 

	return string(hash)
}

func GenerateJWT() (string, error) {
	var SECRET_KEY = []byte("my_secret_key")

	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		log.Println("Error in JWT Token Generator")
		return "", err
	}

	return tokenString, nil
}