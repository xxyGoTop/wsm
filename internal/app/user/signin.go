package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"github.com/xxyGoTop/wsm/internal/app/config"
	"github.com/xxyGoTop/wsm/internal/app/db"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"github.com/xxyGoTop/wsm/internal/lib/controller"
	"github.com/xxyGoTop/wsm/internal/lib/helper"
	"github.com/xxyGoTop/wsm/internal/lib/password"
	"github.com/xxyGoTop/wsm/internal/lib/token"
	"github.com/xxyGoTop/wsm/internal/lib/validator"
	"time"
)

type SignInParams struct {
	Account string `json:"account" valid:"required~请输入登录账号"`
	Password string `json:"password" valid:"required~请输入密码"`
}

func SigninWithUsername(c *controller.Context) (res schema.Response)  {
	var (
		err error
		input SignInParams
		data = &schema.ProfileWithToken{}
		tx *gorm.DB
	)

	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = exception.Unknown
			}
		}

		if tx != nil {
			if err != nil {
				_ = tx.Rollback().Error
			} else {
				err = tx.Commit().Error
			}
		}

		helper.Response(&res, data, nil , err)
	}()

	if err = c.ShouldBindJSON(&input); err != nil {
		return
	}

	userInfo := db.User{}

	if validator.IsPhone(input.Account) {
		userInfo.Phone = &input.Account
	} else if validator.IsEmail(input.Account) {
		userInfo.Email = &input.Account
	} else {
		userInfo.Username = input.Account
	}

	tx = db.Db.Begin()

	if err = tx.Where(&userInfo).Last(&userInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = exception.InvalidAccountOrPassword
		}
		return
	}

	if password.Verify(input.Password, userInfo.Password) == false {
		err = exception.InvalidAccountOrPassword
		return
	}

	if err = userInfo.CheckStatusValid(); err != nil {
		return
	}

	if err = mapstructure.Decode(userInfo, &data.ProfilePure); err != nil {
		return
	}

	data.CreatedAt = userInfo.CreatedAt.Format(time.RFC3339Nano)
	data.UpdateAt = userInfo.UpdatedAt.Format(time.RFC3339Nano)

	if t, er := token.Generate(config.Http.Secret, userInfo.Id); er != nil {
		err = er
		return
	} else {
		data.Token = t
	}

	return
}
