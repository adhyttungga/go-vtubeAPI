package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adhyttungga/go-vtubeAPI/connection"
	"github.com/adhyttungga/go-vtubeAPI/helper"
	"github.com/adhyttungga/go-vtubeAPI/structs"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload structs.User

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload.Password = helper.GetHash([]byte(payload.Password))

	if err := connection.DB.Model(&structs.User{}).Create(&payload).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Message: "User Registration Successful!"})
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var payload, dbuser structs.User


	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := connection.DB.Model(&structs.User{}).Where("email =?", payload.Email).Find(&dbuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(dbuser)

	if dbuser.Id == 0 {
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Message: "Your Email Do Not Exist, Please Try Again!"})
		return
	}

	payloadpass := []byte(payload.Password)
	dbpass := []byte(dbuser.Password)

	if err := bcrypt.CompareHashAndPassword(dbpass, payloadpass); err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Message: "Your Password is Invalid, Please Try Again!"})
		return
	}

	jwtToken, err := helper.GenerateJWT(); 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: []byte(`{"token":"`+jwtToken+`"}`), Message: "User Login Successful!"})
}