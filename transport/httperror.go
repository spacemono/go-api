package transport

import (
	"errors"
	"github.com/spacemono/go-api/service"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func FromError(err error) APIError {
	var apiErr APIError
	var srvError service.Error
	if errors.As(err, &srvError) {
		apiErr.Message = srvError.SvcErr().Error()
		appErr := srvError.AppErr()
		switch {
		case errors.Is(appErr, service.ErrBadRequest):
			apiErr.Code = http.StatusBadRequest
		case errors.Is(appErr, service.ErrUnauthorized):
			apiErr.Code = http.StatusUnauthorized
		case errors.Is(appErr, service.ErrNotFound):
			apiErr.Code = http.StatusNotFound
		case errors.Is(appErr, service.ErrInternalServerError):
			apiErr.Code = http.StatusInternalServerError
		}
	}
	return apiErr
}
