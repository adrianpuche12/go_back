package user

import (
	"errors"
	"fmt"
)

var ErrFistNameRequired = errors.New("first name is required")
var ErrlastRequired = errors.New("last name is required")
var ErrEmailRequired = errors.New("email name is required")

type ErrorNotFound struct {
	ID uint64
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("user id %d doesnt exist", e.ID)
}
