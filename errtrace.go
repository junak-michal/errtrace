// Package errtrace provides a way of adding stack trace information to errors.
package errtrace

import (
	"runtime/debug"
	"bytes"
	"errors"
	"fmt"
)

// Creates a new error that contains given message and a stack trace.
func New(message string) error {
	var contentBuff bytes.Buffer
	contentBuff.WriteString(message)
	contentBuff.WriteByte('\n')
	contentBuff.Write(debug.Stack())
	return errors.New(contentBuff.String())
}

// Formats given string with given arguments and creates new error with stack trace information.
func Newf(format string, args ...interface{}) error {
	return New(fmt.Sprintf(format, args...))
}
