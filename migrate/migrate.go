package main

import (
	"inventario/03-ormgo/configs"
	"inventario/03-ormgo/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}