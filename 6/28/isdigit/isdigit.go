package main

import (
	"errors"
	"fmt"
	"os"
)


// Errors that could occur.
var (
	ErrInvalidDigit = errors.New("invalid digit")
)

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func newDigit(digit int8) error {
	if (!validDigit(digit)){
		return ErrInvalidDigit
	}
	return nil
}

func main() {
	err := newDigit(18)
	if err != nil {
		switch err {
		case ErrInvalidDigit:
			fmt.Println("Wrong digit range!")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
