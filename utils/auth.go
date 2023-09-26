package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)

// var (
// 	key []byte
// 	t   *jwt.Token
// 	s   string
//   )
  
func GenerateJWT(email string, secretKey string) (string, error) {
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (e.g., 24 hours)

	// Sign the token with a secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(tokenString string, secretKey string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Failed to extract claims")
	}

	return claims, nil
}
  