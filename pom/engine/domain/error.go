package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrBadRequest          = errors.New("Bad Request")
	ErrRecordAlreadyExists = errors.New("The record already exists")
)
