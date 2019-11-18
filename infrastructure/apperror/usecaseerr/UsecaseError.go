package usecaseerr

import "errors"

var (
	AuthorizationIsNotFoundOnParam = errors.New("authorization can not be found on param")
	UserIDCanNotBeFoundOnStore     = errors.New("user id can not be found on store")
)
