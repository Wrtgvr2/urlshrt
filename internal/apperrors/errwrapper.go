package apperrors

type AppError struct {
	Err        error
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func WrapError(err error, code int, msg string) *AppError {
	return &AppError{
		Err:        err,
		StatusCode: code,
		Message:    msg,
	}
}
