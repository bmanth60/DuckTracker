package errors

import (
	"github.com/pkg/errors"
)

var (
	//ErrPageNotFound requested route is not available
	ErrPageNotFound = errors.New("page not found")

	//ErrDbNotConnected database connection cannot be established
	ErrDbNotConnected = errors.New("database connection is not established")

	//ErrDbFailed database failed to execute
	ErrDbFailed = errors.New("failed to add entry into database")

	//ErrInvalidValue value supplied was invalid
	ErrInvalidValue = errors.New("value is invalid")
)
