package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adhyttungga/go-vtubeAPI/connection"
	"github.com/adhyttungga/go-vtubeAPI/helper"
	"github.com/adhyttungga/go-vtubeAPI/structs"
	"golang.org/x/crypto/bcrypt"
)

func UserWelcome(w http.ResponseWriter, r *http.Request) {
	key := structs.ContextKey("props")
	props, _ := r.Context().Value(key).(structs.Claims)

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: fmt.Sprintf("Welcome %s!", props.Username), Message:"User login Successful!"})
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var payload, dbuser structs.User

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := connection.DB.Model(&structs.User{}).Where("email =?", payload.Email).Find(&dbuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if dbuser.Id != 0 {
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Message: "Your email already exist, please use other email!"})
		return
	}

	payload.Password = helper.GetHash([]byte(payload.Password))

	if err := connection.DB.Model(&structs.User{}).Create(&payload).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Message: "User registration successful!"})
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

	if dbuser.Id == 0 {
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Message: "Your email do not exist, please try again!"})
		return
	}

	payloadpass := []byte(payload.Password)
	dbpass := []byte(dbuser.Password)

	if err := bcrypt.CompareHashAndPassword(dbpass, payloadpass); err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Message: "Your password invalid, please try again!"})
		return
	}
	
	jwtToken, expireat, err := helper.GenerateJWT(dbuser.Email); 

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtToken,
		Expires: expireat,
	})

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Message: "User login successful!"})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	key := structs.ContextKey("props")
	props, _ := r.Context().Value(key).(structs.Claims)
	
	fmt.Println("Expire at: ", props.ExpiresAt)
	fmt.Println("Duration = ", time.Until(time.Unix(props.ExpiresAt, 0)))
	fmt.Println("30 second = ", 30 * time.Second)

	if time.Until(time.Unix(props.ExpiresAt, 0)) > (12 *time.Hour - 5 * time.Minute) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwtToken, expireat, err := helper.GenerateJWT(props.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtToken,
		Expires: expireat,
	})

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Message: "Refresh token successful!"})
}