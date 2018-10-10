package errors

import (
	"github.com/pkg/errors"
)

var (
	//ErrPageNotFound requested route is not available
	ErrPageNotFound = errors.New("page not found")

	//ErrDbNotConnected database connection cannot be established
	ErrDbNotConnected = errors.New("database connection is not established")

	ErrInvalidValue = errors.New("value is invalid")
)
