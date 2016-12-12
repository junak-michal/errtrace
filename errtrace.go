/*
Copyright 2016 Michal Junák

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package errtrace provides a way of adding stack trace information to errors.
package errtrace

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/debug"
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
