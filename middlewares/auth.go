package middlewares

import (
	"example.com/rest-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized.  Auth token required"})
		return
	}

	userId, err := utils.VerifyToken(token)
	fmt.Println(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
