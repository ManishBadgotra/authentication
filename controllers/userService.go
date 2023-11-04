package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manishbadgotra/authentication/models"
)

var (
	err  error
	user models.EmailUser
)

func CheckStatus(c *gin.Context) {
	log.Println("CheckStatus API")
	c.JSON(http.StatusOK, gin.H{"message": "Connection Established"})
}

func Signup(c *gin.Context) {
	authType := c.Request.Header.Get("AuthenticationType")

	if authType != "" {

		if err = c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
