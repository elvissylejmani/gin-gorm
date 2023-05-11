package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Users(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (u UserController) Update(c *gin.Context) {
	c.JSON(200, gin.H{"Updated": "True"})
}
