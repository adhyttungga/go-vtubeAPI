package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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
			json.NewEncoder(w).Encode(make(map[string]interface{}))
		}

		next.ServeHTTP(w, r)
	})
}

func HandleRequest() {
	log.Println("Server Up and Running on Port 11000")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.Handle("/user/register", http.HandlerFunc(RegisterUser)).Methods("OPTIONS", "POST")
	myRouter.Handle("/user/login", http.HandlerFunc(UserLogin)).Methods("OPTIONS", "POST")

	handler := cors.AllowAll().Handler(myRouter)

	log.Fatal(http.ListenAndServe(":11000", handler))
}