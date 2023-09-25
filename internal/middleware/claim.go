package middleware

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}
