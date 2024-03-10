package user

import (
	"errors"
	"fmt"
)

var ErrFistNameRequired = errors.New("first name is required")
var ErrLastRequired = errors.New("last name is required")
var ErrThereArentFields = errors.New("there arent fields")

type ErrorNotFound struct {
	ID uint64
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("user id '%d' doesnt exist", e.ID)
}
