package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/const"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

var jwtKey2 = []byte(_const.SecretKey)

func Test2(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHRyYWN0QXQiOiIyMDIzLTA3LTA2VDE2OjE4OjUxLjMzMzAxNTcrMDg6MDAiLCJpc0xvZ2dlZE91dCI6ZmFsc2UsIm5hbWUiOiJ4aWFveWlAaWNsb3VkLmNvbSJ9.-7EOh2QEGg6bY3YyVzVM47dx6cLoWdAi5UUlLQthB78"

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if ok && token.Valid {
		fmt.Println("JWT is valid")
		name := claims["name"].(string)
		isLoggedOut := claims["isLoggedOut"].(bool)
		fmt.Println("name:", name)
		fmt.Println("is loggedOut:", isLoggedOut)
		extractAt := claims["extractAt"].(string)
		extractTime, err := time.Parse(time.RFC3339Nano, extractAt)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("extract at:", extractTime)
		return
	}
	t.Error("Invalid JWT")
}
