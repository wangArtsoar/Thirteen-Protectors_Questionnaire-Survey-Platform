package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Test1(t *testing.T) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: "my_username",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		t.Error("Error generating JWT:", err)
		return
	}

	fmt.Println("Generated JWT:", tokenString)
}
