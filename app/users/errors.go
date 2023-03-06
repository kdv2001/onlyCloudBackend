package users

import (
	"net/http"

	"onlyCloudBackend/internal/appErrors"
)

var (
	ErrorWrongPassword = appErrors.AppError{Msg: "wrong password", Code: http.StatusBadRequest}
	ErrAuthEmailUsed   = appErrors.AppError{Msg: "email already registered", Code: http.StatusConflict}
)
