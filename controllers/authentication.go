package controllers

import (
	"strings"

	"github.com/dgrijalva/rest-api/pkg/models"
	"github.com/gin-gonic/gin"
)

// Auth Function
func Auth(c *gin.Context) {
	var claims models.JwtClaims
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	token = strings.Split(token, "Bearer")[1]
	err := models.VerifyToken(token, &claims)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	c.Set("claims", claims)
}
