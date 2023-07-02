package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey2 = []byte(common.SecretKey)

type Claims2 struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Test2(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InhpYW95aSIsImV4cCI6MTY4ODM2Nzg2Mn0.zHK06J4HOlN3HNj06RWQVD7j-wIs7WcDOe37EpJJvX4"

	claims2 := &Claims2{}

	token, err := jwt.ParseWithClaims(tokenString, claims2, func(token *jwt.Token) (interface{}, error) {
		return jwtKey2, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			t.Error("Invalid JWT signature")
			return
		}
		t.Error("Error parsing JWT:", err)
		return
	}
	if !token.Valid {
		t.Error("Invalid JWT")
		return
	}

	fmt.Println("JWT is valid")
	fmt.Println("Username:", claims2.Username)
	fmt.Println("Expires at:", time.Unix(claims2.ExpiresAt, 0))
}
