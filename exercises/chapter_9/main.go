package main

import (
	"errors"
	"fmt"
)

var ErrDataNotFound = errors.New("data not found")

func main() {
	fmt.Println("Hello World")
	_, err := biggerThanSeven(8)
	if err != nil {
		if errors.Is(err, ErrDataNotFound) {
			fmt.Println("Data Not Found")
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
