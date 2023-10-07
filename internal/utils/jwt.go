package utils

import (
	"errors"
	"file-system/internal/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(identity, name string, expire int) (string, error) {
	uc := middleware.UserClaim{
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expire)).Unix(),
		},
	}
	// 使用指定的签名方法（HS256）创建带有声明的新令牌。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 使用秘密密钥签名令牌并获取编码后的字符串令牌。
	tokenString, err := token.SignedString([]byte(middleware.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
