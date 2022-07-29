package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-vtubeAPI/structs"
	"github.com/dgrijalva/jwt-go"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token");
	log.Println(c)
	
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := c.Value
	claims := &structs.Claims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !jwtToken.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}