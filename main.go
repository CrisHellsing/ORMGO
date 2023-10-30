package main

import (
	"net/http"
	"inventario/03-ormgo/routes"
	"inventario/03-ormgo/configs"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}