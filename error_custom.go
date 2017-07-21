package main

import (
	"fmt"
)

var (
	ErrRabbimqConnNotOpen = &CustomError{1001, "Not Open Conn"}
	ErrMysqlConnNotOpen   = &CustomError{1001, "Not Open Conn"}
)

type CustomError struct {
	Code int
	Msg  string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Code: %v Msg: %s", c.Code, c.Msg)
}

func run() (string, error) {
	return "", ErrMysqlConnNotOpen
}

func main() {
	_, err := run()
	if err == nil {
		fmt.Println("None")
	} else if err == ErrRabbimqConnNotOpen {
		fmt.Println(30)
		fmt.Println(err)
	} else if err == ErrMysqlConnNotOpen {
		fmt.Println(33)
		fmt.Println(err)
	}

}
