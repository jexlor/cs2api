package utils

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var JwtKey []byte

// func GenerateJWT(userID string) (string, error) {
// 	claims := &jwt.MapClaims{
// 		"user_id": userID,
// 		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(JwtKey)
// }
