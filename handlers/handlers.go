package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/adhyttungga/go-vtubeAPI/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			c, err := r.Cookie("token")

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
			tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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

			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			key := structs.ContextKey("props")
			ctx := context.WithValue(r.Context(), key,  *claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func HandleRequest() {
	log.Println("Server Up and Running on Port 11000")

	myRouter := mux.NewRouter().StrictSlash(true)
	
	myRouter.HandleFunc("/", HomePage)

	myRouter.Handle("/user/register", http.HandlerFunc(UserRegister)).Methods("OPTIONS", "POST")
	myRouter.Handle("/user/login", http.HandlerFunc(UserLogin)).Methods("OPTIONS", "POST")
	myRouter.Handle("/user/welcome", MiddlewareAuth(http.HandlerFunc(UserWelcome)))
	myRouter.Handle("/user/refresh-token", MiddlewareAuth(http.HandlerFunc(RefreshToken))).Methods("OPTIONS", "GET")

	handler := cors.AllowAll().Handler(myRouter)

	log.Fatal(http.ListenAndServe(":11000", handler))
}