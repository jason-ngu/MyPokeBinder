package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateJWTToken(providerKey string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": providerKey,                      // Subject (user identifier)
		"iss": "MyPokeBinder",                   // Issuer
		"aud": "user",                           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(`Unexpected signing method: %v`, token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			http.Error(w, "Invalid or expired JWT token", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 	fmt.Println(claims["foo"], claims["nbf"])
		// } else {
		// 	fmt.Println(err)
		// }
		next.ServeHTTP(w, r)
	})
}
