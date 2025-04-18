package httpmaps

import (
	"net/http"

	"github.com/wrtgvr/urlshrt/internal/apperrors/defs"
)

var DbUsersErrorsMap = map[error]int{
	defs.ErrUserNotFound: http.StatusNotFound,
}
