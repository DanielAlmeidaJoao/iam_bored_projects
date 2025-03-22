package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Generate JWT
func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Securely Extract and Store User in Context
func getLoggedUsername(r *http.Request) (string, error) {

	// 🔹 Get Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Missing token")
	}

	// 🔹 Extract Bearer token
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		return "", fmt.Errorf("Invalid token format")
	}

	// 🔹 Parse JWT Token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	return claims.Username, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// Simulate user authentication (Use a database in real-world apps)
	if username != "alice" || password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// 🔹 Return JWT in Authorization header
	w.Header().Set("Authorization", "Bearer "+token)

	// Also return token in the response body for debugging (optional)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, token)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getLoggedUsername(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome, %s! You are an admin.", user)
}

var userRoles = map[string]string{
	"alice": "admin",
	"bob":   "user",
}

func main() {
	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/admin", adminHandler)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
