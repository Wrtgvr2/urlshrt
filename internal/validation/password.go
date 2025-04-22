package validation

import (
	"fmt"
	"regexp"

	"github.com/wrtgvr2/errsuit"
)

var PasswordMinLength = 8
var PasswordMaxLength = 30

var passwordRegexp = regexp.MustCompile(`^[a-zA-Z0-9@!*&$%\\/"'-_]+$`)

func ValidatePassword(s string) *errsuit.AppError {
	if len(s) < PasswordMinLength {
		return errsuit.NewBadRequest(fmt.Sprintf("password too long. Max: %d character", PasswordMinLength), nil, false)
	}
	if len(s) > PasswordMaxLength {
		return errsuit.NewBadRequest(fmt.Sprintf("password too long. Max: %d character", PasswordMaxLength), nil, false)
	}
	if !passwordRegexp.Match([]byte(s)) {
		return errsuit.NewBadRequest("password contain unallowed characters", nil, false)
	}

	return nil
}
