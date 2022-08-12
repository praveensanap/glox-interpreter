package errors

import "fmt"
import go_error "errors"

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Println(fmt.Sprintf("[line %d] Error %s: %s", line, where, message))
}

func New(text string) error {
	return go_error.New(text)
}
