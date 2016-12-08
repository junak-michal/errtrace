package errtrace

import (
	"strings"
	"testing"
)

var testInputsForNew = []string{
	"Short error text",
	"",
	"A little bit \n longer text \t\t\t containing escape \n sequences.",
}

var testInputsForNewf = []struct {
	inputFmt       string
	fmtArgs        []interface{}
	expectedOutput string
}{
	{"Formatted %s message %d!", []interface{}{"error", 5}, "Formatted error message 5!"},
	{"", []interface{}{}, ""},
	{"%s%s %d%s", []interface{}{"A", "B", 1, "C"}, "AB 1C"},
}

func TestNew(t *testing.T) {
	for _, errMsg := range testInputsForNew {
		err := New(errMsg)
		if !strings.Contains(err.Error(), errMsg) {
			t.Errorf("Error does not contain original error message.\nOriginal message: %s\nError: %v", errMsg, err)
		}
		if !containsStackTrace(err) {
			t.Error("Error does not contain stack trace.\nError:", err)
		}
	}
}

func TestNewf(t *testing.T) {
	for _, testArgs := range testInputsForNewf {
		err := Newf(testArgs.inputFmt, testArgs.fmtArgs...)
		if !strings.Contains(err.Error(), testArgs.expectedOutput) {
			t.Errorf("Error does not contain original error message.\nOriginal message: %s\nError: %v", testArgs.expectedOutput, err)
		}
		if !containsStackTrace(err) {
			t.Error("Error does not contain stack trace.\nError:", err)
		}
	}
}

func containsStackTrace(err error) bool {
	// This is a simple yet naive way to check that stack trace is in the string.
	// Stack trace always contain information about calling the method runtime/debug.Stack
	return strings.Contains(err.Error(), "runtime/debug.Stack")
}
