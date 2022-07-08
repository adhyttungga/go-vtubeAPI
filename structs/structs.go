package structs

import "time"

type User struct {
	Id          int      	`json:"id"`
	Name        string    `json:"name" gorm:"size:100"`
	Email       string    `json:"email" gorm:"size:100 unique"`
	Password    []byte    `json:"-" gorm:"size:200"`
	CreatedDate time.Time `json:"-" gorm:"type:timestamp;default:current_timestamp`
	UpdatedDate time.Time `json:"-" gorm:"type:timestamp null`
}

type Result struct {
	Code 		int 				`json:"code"`
	Data 		interface{} `json:"data"`
	Message string			`json:"message"`
}		