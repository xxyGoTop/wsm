package db

import (
	"log"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/xxyGoTop/wsm/internal/app/config"
	"github.com/xxyGoTop/wsm/internal/app/exception"
)

type UserStatus int32

type Gender int

const (
	// 用户状态
	UserStatusBanned      UserStatus = -100 //账号被禁用
	UserStatusInactivated UserStatus = -1   //账号未激活
	UserStatusInit        UserStatus = 1

	// 用户性别
	GenderUnknown Gender = 0 // 未知性别
	GenderMale    Gender = 1 // 男性
	GenderFemale  Gender = 2 // 女性
)

var (
	userId *snowflake.Node
)

func init() {
	node, err := snowflake.NewNode(config.Common.MachineId)

	if err != nil {
		log.Panicln(err)
	}

	userId = node
}

type User struct {
	Id       string         `gorm:"primary_key;not null;unique;index;type:varchar(32)" json:"id"` //用户ID
	Username string         `gorm:"not null;unique;type:varchar(36);index" json:"username"`       //用户名
	Password string         `gorm:"not null;type:varchar(64);index" json:"password"`              //密码
	NickName *string        `gorm:"null;index;type:varchar(36)" json:"nickname`                   //昵称
	Phone    *string        `gorm:"null;type:varchar(16);index" json:"phone"`                     //电话
	Email    *string        `gorm:"null;type:varchar(36);index" json:"email"`                     //邮箱
	Status   UserStatus     `gorm:"not null" json:"status"`                                       //账号状态
	Role     pq.StringArray `gorm:"not null;type:varchar(36)[]" json:"role"`                      //角色
	Avatar   string         `gorm:"not null;type:varchar(128)" json:"avatar"`                     //头像
	Gender   Gender         `gorm:"not null;default(0)" json:"gender"`                            //性别

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time `sql:"index"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	// 生成uid
	uid := userId.Generate().String()

	if err := scope.SetColumn("id", uid); err != nil {
		return err
	}

	if err := scope.SetColumn("gender", GenderUnknown); err != nil {
		return err
	}

	return nil
}

func (u *User) CheckStatusValid() error {
	if u.Status != UserStatusInit {
		switch true {
		case u.Status == UserStatusInactivated:
			return exception.UserIsInActive
		case u.Status == UserStatusBanned:
			return exception.UserHaveBeenBan
		}
	}
	return nil
}
