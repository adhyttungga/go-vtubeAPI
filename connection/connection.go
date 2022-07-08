package connection

import (
	"github.com/adhyttungga/go-vtubeAPI/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func Connect() {
	dsn := "host=localhost user=postgres password=NA282d6f42 dbname=vtube port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	
	if Err != nil {
		panic("failed to connect database")
	}
	
	DB.AutoMigrate(&structs.User{})
}