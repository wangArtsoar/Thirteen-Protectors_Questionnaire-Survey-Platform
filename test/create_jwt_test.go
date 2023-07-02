package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

func Test1(t *testing.T) {
	expirationTime := time.Now().Add(5 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  "xiaoyi",
		"extractAt": expirationTime,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		t.Error("Error generating JWT:", err)
		return
	}
	fmt.Println("Generated JWT:", tokenString)
}
