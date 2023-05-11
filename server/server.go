package main

import (
	"authentication/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	UserController := new(controllers.UserController)

	r.GET("/users", UserController.Users)
	r.PATCH("/users", UserController.Update)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
