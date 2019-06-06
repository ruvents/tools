package tools

import (
	"errors"
	"fmt"
)

func NewError(msg string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, args...))
}
