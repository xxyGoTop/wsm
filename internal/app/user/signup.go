package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"github.com/xxyGoTop/wsm/internal/app/db"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/app/schema"
	"github.com/xxyGoTop/wsm/internal/lib/controller"
	"github.com/xxyGoTop/wsm/internal/lib/helper"
	"github.com/xxyGoTop/wsm/internal/lib/password"
	"github.com/xxyGoTop/wsm/internal/lib/validator"
	"time"
)

type SignUpWithUsernameParams struct {
	Username string `json:"username" valid:"required~请输入用户名"`
	Password string `json:"password" valid:"required~请输入密码"`
}

func SignUpWithUsername(c *controller.Context) (res schema.Response)  {
	var (
		input SignUpWithUsernameParams
		err error
		data schema.Profile
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

	if err = validator.ValidateUsername(input.Username); err != nil {
		return
	}

	tx = db.Db.Begin()

	userNum := 0

	if err = tx.Model(db.User{}).Where("username = ?", input.Username).Count(&userNum).Error; err != nil {
		return
	}

	if userNum != 0 {
		err = exception.UserExist
		return
	}

	passwordHash, err := password.Generate(input.Password)

	if err != nil {
		return
	}

	userInfo := db.User{
		Username:  input.Username,
		Password:  passwordHash,
		NickName:  &input.Username,
		Phone:     nil,
		Email:     nil,
		Status:    db.UserStatusInit,
		Role:      pq.StringArray{},
		Gender:    db.GenderUnknown,
	}

	if err = tx.Create(&userInfo).Error; err != nil {
		return
	}

	if err = mapstructure.Decode(userInfo, &data.ProfilePure); err != nil {
		return
	}

	data.CreatedAt = userInfo.CreatedAt.Format(time.RFC3339Nano)
	data.UpdateAt = userInfo.UpdatedAt.Format(time.RFC3339Nano)

	return
}