package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xxyGoTop/wsm/internal/app/config"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"github.com/xxyGoTop/wsm/internal/lib/token"
	"net/http"
)

var (
	ContextUidField = "uid"
)

func Authenticate(isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			tokenString string
			status = schema.StatusFail
		)

		defer func() {
			if err != nil {
				c.JSON(http.StatusOK, schema.Response{
					Message: err.Error(),
					Data:    nil,
					Status:  status,
				})
			}
		}()

		if s, isExist := c.GetQuery(token.AuthField); isExist == true {
			tokenString = s
		} else {
			tokenString = c.GetHeader(token.AuthField)
			if len(tokenString) == 0 {
				if s, err := c.Cookie(token.AuthField); err != nil {
					err = exception.InvalidToken
					status = exception.InvalidToken.Code()
					return
				} else {
					tokenString = s
				}
			}
		}

		if claims, er := token.Parse(config.Http.Secret, tokenString); err != nil {
			err = er
			status = exception.InvalidToken.Code()
			return
		} else {
			c.Set(ContextUidField, claims.Uid)
		}
	}
}