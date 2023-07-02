package common

import (
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
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return errors.New("jwt token create fail").Error()
	}
	return tokenString
}

//// ExtractJwt 解析jwt
//func ExtractJwt(jwtToken string) (*Claims, error) {
//	claims := &Claims{}
//	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
//		return []byte(SecretKey), nil
//	})
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			return nil, errors.New("invalid JWT signature")
//		}
//		return nil, errors.New("Error parsing JWT: " + err.Error())
//	}
//	if !token.Valid {
//		return nil, errors.New("invalid JWT")
//	}
//	return claims, nil
//}
