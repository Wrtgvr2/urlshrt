package httpmaps

import (
	"net/http"

	"github.com/wrtgvr/urlshrt/internal/apperrors/defs"
)

var AuthErrorsMap = map[error]int{
	defs.ErrInvalidPassword: http.StatusBadRequest,
}
