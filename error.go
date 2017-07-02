package util

import (
	"fmt"
)

func Err(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
