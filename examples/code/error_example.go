package main

import (
	"errors"
	"fmt"
)

type TestError struct {
	msg string
	asd int
}

func (e *TestError) Error() string {
	return fmt.Sprintf("TestError: %s | %d", e.msg, e.asd)
}

func getError() error {
	return &TestError{"err msg", 123}
}

func main() {
	err := getError()
	if errors.As(err, &TestError{}) {
		fmt.Println("err is TestError")
	}
}
