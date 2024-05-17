package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintMessageWithError(t *testing.T) {

	// Test case 1: Error is nil
	msg := "Hello"
	err := error(nil)
	expected := "Hello, err: <nil>"
	result := PrintMessageWithError(msg, err)
	assert.Equal(t, expected, result, "Result should match expected message")

	// Test case 2: Error is not nil
	errorMsg := "something went wrong"
	err = fmt.Errorf(errorMsg)
	expected = "Hello, err: " + errorMsg
	result = PrintMessageWithError(msg, err)
	assert.Equal(t, expected, result, "Result should match expected message")

}
