package LinkedList

import "fmt"

type Error struct {
	errString string
}

func (err *Error) Error() string {
	return fmt.Sprintf("[LinkedList] %v", err.errString)
}
