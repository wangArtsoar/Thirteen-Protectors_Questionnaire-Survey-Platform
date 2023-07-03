package common

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/bean"
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
func CreateNewToken(name string) string {
	// 过期时间
	_ = time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, nil)
	tokenString, err := token.SignedString([]byte(bean.SecretKey))
	if err != nil {
		return errors.New("jwt token create fail").Error()
	}
	return tokenString
}

// ExtractJwt 解析jwt
func ExtractJwt(jwtToken string) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(bean.SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid JWT signature")
		}
		return nil, errors.New("Error parsing JWT: " + err.Error())
	}
	if !token.Valid {
		return nil, errors.New("invalid JWT")
	}
	claims := token.Claims
	return claims, nil
}
