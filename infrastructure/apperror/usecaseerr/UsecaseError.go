package apperror

import "errors"

var (
	UsecaseErr UsecaseError
)


type UsecaseError struct {
	AuthorizationIsNotFoundOnParam error
	UserIDCanNotBeFoundOnStore error
}

func (e *UsecaseError) Init() {
	e.AuthorizationIsNotFoundOnParam = errors.New("authorization can not be found on param")
	e.UserIDCanNotBeFoundOnStore = errors.New("user id can not be found on store")
}