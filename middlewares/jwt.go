package middlewares

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtKey = []byte("your-secret-key") // Replace with your secret key, preferably from an environment variable.

// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
// 			c.Abort()
// 			return
// 		}
// 		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 			c.Abort()
// 			return
// 		}

// 		// Parse and validate the JWT token
// 		claims := &jwt.MapClaims{}
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user_id", claims["user_id"])
// 		c.Next()
// 	}
// }
