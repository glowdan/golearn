package main

import (
	"time"
	"fmt"
)

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v %s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"can't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
