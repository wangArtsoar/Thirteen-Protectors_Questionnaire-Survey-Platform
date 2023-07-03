package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/bean"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

var jwtKey2 = []byte(bean.SecretKey)

func Test2(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InhpYW95aSIsImV4cCI6MTY4ODM2Nzg2Mn0.zHK06J4HOlN3HNj06RWQVD7j-wIs7WcDOe37EpJJvX4"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Error("Invalid JWT")
		return
	}
	fmt.Println("JWT is valid")
	subject, err := claims.GetSubject()
	if err != nil {
		t.Error("Get JWT subject fail" + err.Error())
		return
	}
	fmt.Println("Username:", subject)
	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		t.Error("Get JWT expirationTime fail" + err.Error())
		return
	}
	fmt.Println("Expires at:", expirationTime)
}
