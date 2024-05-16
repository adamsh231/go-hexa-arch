package utils

import "fmt"

func PrintMessageWithError(msg string, err error) string{
	return fmt.Sprintf("%s, err: %v", msg, err)
}