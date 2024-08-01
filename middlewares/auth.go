package middlewares

import (
	"net/http"

	"github.com/VisarutJDev/social-media-api/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Example authentication without user/pass
// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token != "Bearer your-secret-token" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		claims := &controllers.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while parsing token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
