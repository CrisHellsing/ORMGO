package configs

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	var err error 

	dsn := "root.@tcp(localhost:3319)/inventario"
	DB, err = grom.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}