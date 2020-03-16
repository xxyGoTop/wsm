package token

import "github.com/dgrijalva/jwt-go"

const (
	Prefix = "Bearer"
	AuthField = "Authorization"
)

type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

type ClaimsInternal struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func JoinPrefixToken(token string) string  {
	return Prefix + " " + token
}