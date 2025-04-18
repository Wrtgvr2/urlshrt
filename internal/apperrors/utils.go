package apperrors

import (
	"errors"
	"net/http"
)

func GetErrorStatusCode(err error) int {
	for _, m := range ErrorCodesMaps {
		for k, v := range m {
			if errors.Is(err, k) {
				return v
			}
		}
	}

	return http.StatusInternalServerError
}
