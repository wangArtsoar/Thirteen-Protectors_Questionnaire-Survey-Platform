package util

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Claims JWT声明
type Claims struct {
	Username string `json:"username"`
	jwt.ClaimStrings
}

// CreateNewToken return a new token
func CreateNewToken(name string, roleName string, isLoggedOut bool) string {
	// 过期时间
	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"name":        name,
		"isLoggedOut": isLoggedOut,
		"extractAt":   expirationTime,
		"role":        roleName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constant.SecretKey))
	if err != nil {
		return errors.New("jwt token create fail").Error()
	}
	return tokenString
}

// ExtractJwtToken 解析jwt
func ExtractJwtToken(jwtToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.SecretKey), nil
	})
}

// GetMapClaims 获取MapClaims
func GetMapClaims(jwtToken string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := ExtractJwtToken(jwtToken)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, nil, errors.New("invalid JWT signature")
		}
		return nil, nil, errors.New("Error parsing JWT: " + err.Error())
	}
	if !token.Valid {
		return nil, nil, errors.New("invalid JWT")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("invalid JWT")
	}
	return token, claims, nil
}
