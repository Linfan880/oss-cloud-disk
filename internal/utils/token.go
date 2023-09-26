package utils

import (
	"errors"
	"file-system/internal/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(identity, name string, second int) (string, error) {
	uc := middleware.UserClaim{
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(middleware.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*middleware.UserClaim, error) {
	uc := new(middleware.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, err
}
