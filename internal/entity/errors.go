package entity

import "errors"

var (
	ErrUnauthorizedPassword = errors.New("unauthorized password")
	ErrIDIsRequired         = errors.New("id is required")
	ErrIDIsInvalid          = errors.New("ID is invalid")
	ErrInvalidName          = errors.New("invalid name")
	ErrPriceIsRequired      = errors.New("price is required")
	ErrInvalidPrice         = errors.New("invalid price")
)
