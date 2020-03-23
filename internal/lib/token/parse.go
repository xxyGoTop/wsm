package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/lib/util"
	"strings"
)

func Parse(secret, tokenString string) (claims Claims, err error) {
	var (
		token *jwt.Token
	)

	if strings.HasPrefix(tokenString, Prefix+" ") == false {
		err = exception.InvalidAuth
		return
	}

	tokenString = strings.Replace(tokenString, Prefix+" ", "", 1)

	if tokenString == "" {
		err = exception.InvalidToken
		return
	}

	c := ClaimsInternal{}

	if token, err = jwt.ParseWithClaims(tokenString, &c, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	}); err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			err = exception.TokenExpired
		}
		err = exception.InvalidToken
		return
	}

	if token != nil && token.Valid {
		var (
			uid string
		)

		if uid, err = util.Base64Decode(c.Uid); err != nil {
			return
		}

		claims.Uid = uid
		claims.Audience = c.Audience
		claims.Id = c.Id
		claims.NotBefore = c.NotBefore
		claims.ExpiresAt = c.ExpiresAt
		claims.Issuer = c.Issuer
		claims.IssuedAt = c.IssuedAt
		claims.Subject = c.Subject

		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors & jwt.ValidationErrorMalformed != 0 {
			err = exception.InvalidToken
			return
		} else if ve.Errors & (jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			err = exception.TokenExpired
			return
		} else {
			err = exception.InvalidToken
			return
		}
	} else {
		err = exception.InvalidToken
		return
	}
}
