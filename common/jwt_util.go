package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims JWT声明
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateNewToken return a new token
func CreateNewToken(name string) string {
	// 过期时间
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return errors.New("jwt token create fail").Error()
	}
	return tokenString
}

// ExtractJwt 解析jwt
func ExtractJwt(jwtToken string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
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
	return claims, nil
}
