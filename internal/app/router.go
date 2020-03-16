package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxyGoTop/wsm/internal/app/config"
	"github.com/xxyGoTop/wsm/internal/app/middleware"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"github.com/xxyGoTop/wsm/internal/app/user"
	"github.com/xxyGoTop/wsm/internal/lib/controller"
	"github.com/xxyGoTop/wsm/internal/lib/dotenv"
	"net/http"
	"path"
)

var RootRouter *gin.Engine

func init() {
	if config.Common.Mode == config.ModeProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.Use(middleware.GracefulExit())

	router.Use(middleware.CORS())

	router.Static("/public", path.Join(dotenv.RootDir, "public"))

	if config.Common.Mode != config.ModeProduction {
		router.Use(gin.Logger())
	}

	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, schema.Response{
			Message: fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
			Data:    nil,
			Status:  schema.StatusFail,
		})
	})

	{
		v1 := router.Group("/v1")
		v1.Use(middleware.Common)

		v1.GET("", controller.Router(func(c *controller.Context) (res schema.Response) {
			res.Data = "pong"
			res.Status = schema.StatusSuccess
			return
		}))

		// 认证类接口
		{
			authRouter := v1.Group("/auth")
			authRouter.POST("/signup", controller.Router(user.SignUpWithUsername))
			authRouter.POST("/signin", controller.Router(user.SigninWithUsername))
		}
	}

	RootRouter = router
}
