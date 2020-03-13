package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xxyGoTop/wsm/internal/app/config"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"net/http"
)

func GracefulExit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Common.Exiting {
			err := exception.SystemMaintenance

			c.JSON(http.StatusOK, schema.Response{
				Message: err.Error(),
				Data:    nil,
				Status:  err.Code(),
			})
			c.Abort()
		}
	}
}
