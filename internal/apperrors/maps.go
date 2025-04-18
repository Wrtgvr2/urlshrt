package apperrors

import "github.com/wrtgvr/urlshrt/internal/apperrors/httpmaps"

var ErrorCodesMaps = []map[error]int{
	httpmaps.DbUsersErrorsMap,
	httpmaps.AuthErrorsMap,
}
