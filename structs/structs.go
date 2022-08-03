package structs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id          int      	`json:"-"`
	Name        string    `json:"name" gorm:"size:100"`
	Email       string    `json:"email" gorm:"size:100 unique"`
	Password    string    `json:"password" gorm:"size:200"`
	CreatedDate time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedDate time.Time `json:"-" gorm:"type:timestamp"`
}

type Result struct {
	Code 		int 				`json:"code"`
	Data 		interface{} `json:"data"`
	Message string			`json:"message"`
}		

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ContextKey string