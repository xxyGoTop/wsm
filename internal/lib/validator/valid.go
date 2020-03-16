package validator

import (
	"github.com/asaskevich/govalidator"
	"github.com/xxyGoTop/wsm/internal/app/exception"
	"github.com/xxyGoTop/wsm/internal/lib/util"
	"regexp"
)

var (
	usernameReg = regexp.MustCompile("^[\\w\\-]$]")
)

func IsValidUsername(username string) bool  {
	return usernameReg.MatchString(username)
}

func ValidateUsername(username string) error {
	if !IsValidUsername(username) {
		return exception.InvalidFormat
	}
	return nil
}

func IsEmail(email string) bool  {
	return govalidator.IsEmail(email)
}

func IsPhone(phone stting) bool  {
	return util.IsPhone(phone)
}
