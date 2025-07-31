package service

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)

type Error struct {
	appErr error
	svcErr error
}

func (e Error) AppErr() error {
	return e.appErr
}

func (e Error) SvcErr() error {
	return e.svcErr
}

func NewError(appErr, svcErr error) error {
	return Error{svcErr: svcErr, appErr: appErr}
}

func (e Error) Error() string {
	return errors.Join(e.appErr, e.svcErr).Error()
}
