package utils

// import (
// 	"fmt"
// 	"github.com/golang-jwt/jwt/v5"
// 	"time"
// )

// // Add a new global variable for the secret key
// var secretKey = []byte("your-secret-key")

// // Function to create JWT tokens with claims
// func createToken(username string) (string, error) {
// 	// Create a new JWT token with claims
// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": username,   // Subject (user identifier)
// 		"iss": "todo-app", // Issuer
// 		"aud": getRole(username),
// 		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
// 		"iat": time.Now().Unix(),                // Issued at
// 	})

// 	// Print information about the created token
// 	fmt.Printf("Token claims added: %+v\n", claims)
// 	return tokenString, nil
// }
