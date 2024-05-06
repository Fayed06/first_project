// app/handlers/auth_handler.go
package handlers

import (
	"net/http"

	"github.com/fayed06/first_project/app/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var requestBody services.RegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := services.Register(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func Login(c *gin.Context) {
	var requestBody services.LoginRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := services.Login(requestBody)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func GetCurrentUserData(c *gin.Context) {
	user, err := services.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
