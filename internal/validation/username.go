package validation

import (
	"fmt"
	"regexp"

	"github.com/wrtgvr2/errsuit"
)

var UsernameMinLength = 5
var UsernameMaxLength = 30

var usernameRegexp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func ValidateUsername(s string) *errsuit.AppError {
	if len(s) < UsernameMinLength {
		return errsuit.NewBadRequest(fmt.Sprintf("username too short. Max: %d character", UsernameMinLength), nil, false)
	}
	if len(s) > UsernameMaxLength {
		return errsuit.NewBadRequest(fmt.Sprintf("username too long. Max: %d character", UsernameMaxLength), nil, false)
	}
	if !usernameRegexp.Match([]byte(s)) {
		return errsuit.NewBadRequest("usrname contain unallowed characters", nil, false)
	}

	return nil
}
