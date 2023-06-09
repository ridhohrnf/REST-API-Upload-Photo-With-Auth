package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helpers.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
