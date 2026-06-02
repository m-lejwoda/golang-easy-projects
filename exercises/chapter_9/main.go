package main

import (
	"errors"
	"fmt"
)

var ErrDataNotFound = errors.New("data not found")

type EmptyFieldError struct {
	FieldName string
}

var ErrInvalidID = errors.New("invalid ID")

type Employee struct {
	FirstName string
	LastName  string
	Id        int
}

func (fe EmptyFieldError) Error() string {
	return fe.FieldName
}

func main() {
	fmt.Println("Hello World")
	_, err := biggerThanSeven(8)
	var fieldErr EmptyFieldError

	if err != nil {
		if errors.Is(err, ErrDataNotFound) {
			fmt.Println("Data Not Found")
		}
	}
	empl := Employee{
		"sadasd",
		"",
		0,
	}
	err = validateEmployee(empl)
	if err != nil {
		fmt.Println(err)
		if errors.As(err, &fieldErr) {
			fmt.Println("Check if as working")
		}
	}
}

func biggerThanSeven(num int) (int, error) {
	if num > 7 {
		return 0, ErrDataNotFound
	} else {
		return num, nil
	}
}

func validateEmployee(e Employee) error {
	if len(e.FirstName) == 0 {
		return EmptyFieldError{FieldName: "FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldError{FieldName: "LastName"}
	}
	if e.Id == 0 {
		return ErrInvalidID
	}
	return nil
}
