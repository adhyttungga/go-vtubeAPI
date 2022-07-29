package helper

import (
	"log"
	"time"

	"github.com/adhyttungga/go-vtubeAPI/structs"
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

func GenerateJWT(username string, expirationTime time.Time) (string, error) {
	var SECRET_KEY = []byte("my_secret_key")
	
	claims := &structs.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)

	if err != nil {
		log.Println("Error in JWT Token Generator")
		return "", err
	}

	return tokenString, nil
}