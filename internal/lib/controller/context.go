package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/app/middleware"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"net/http"
)

// request 上下文控制
type Context struct {
	ctx		  *gin.Context
	Uid 	  string `json:"uid"` //用户uid
	UserAgent string `json:"user_agent"` //用户代理
	IP 		  string `json:"ip"` //ip地址
}

type controllerFunc func(c *Context) schema.Response

// 校验 json body
func (c *Context) ShouldBindJSON(inputPointer interface{}) error  {
	if err := c.ctx.ShouldBindJSON(inputPointer); err != nil {
		return exception.InvalidParams.New(err.Error())
	}

	if isValid, err := govalidator.ValidateStruct(inputPointer); err != nil {
		return exception.InvalidParams.New(err.Error())
	} else if !isValid {
		return exception.InvalidParams
	}
	return nil
}

// 校验 url query
func (c *Context) ShouldBindQuery(inputPointer interface{}) error  {
	if err := c.ctx.ShouldBindQuery(inputPointer); err != nil {
		return exception.InvalidParams.New(err.Error())
	}

	if isValid, err := govalidator.ValidateStruct(inputPointer); err != nil {
		return exception.InvalidParams.New(err.Error())
	} else if !isValid {
		return exception.InvalidParams
	}
	return nil
}

func (c *Context) SetHeader(key, value string) {
	c.ctx.Header(key, value)
}

func (c *Context) GetHeader(key string) string  {
	return c.ctx.GetHeader(key)
}

func (c *Context) GetParam(key string) string {
	return c.ctx.Param(key)
}

func (c *Context) GetQuery(key string) string {
	return c.ctx.Query(key)
}

func (c *Context) response(data interface{}) {
	c.ctx.JSON(http.StatusOK, data)
}

func NewContext(c *gin.Context) Context {
	return Context{
		ctx:       c,
		Uid:       c.GetString(middleware.ContextUidField),
		UserAgent: c.GetHeader("user-agent"),
		IP:        c.ClientIP(),
	}
}

func Router(ctrl controllerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := NewContext(c)
		ctx.response(ctrl(&ctx))
	}
}