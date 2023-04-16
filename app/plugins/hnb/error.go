package hnb

import "errors"

var (
	ErrInvalidCode            = errors.New("invalid code")
	ErrConversionCodeNotFound = errors.New("conversion code not found")
)
