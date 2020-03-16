package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xxyGoTop/wsm/internal/lib/util"
	"time"
)

func Generate(secret, userId string) (tokenString string, err error)  {
	c := ClaimsInternal{
		Uid:            util.Base64Encode(userId),
		StandardClaims: jwt.StandardClaims{
			Audience:userId,
			Id:userId,
			ExpiresAt:time.Now().Add(time.Hour * time.Duration(6)).Unix(),
			Issuer:"user",
			IssuedAt:time.Now().Unix(),
			NotBefore:time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)

	tokenString, err = token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	tokenString = JoinPrefixToken(tokenString)
	return
}
