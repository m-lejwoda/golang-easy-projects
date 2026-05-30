package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	res, err := divide(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %b \n", res)
	}
	result, _ := fileLen("text.txt")
	fmt.Printf("Number of bytes in file: %d \n", result)
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func divide(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("division by zero")
	}
	return numerator / denominator, nil
}

func fileLen(filePath string) (int, error) {
	f, _ := os.Open(filePath)

	defer f.Close()
	data := make([]byte, 100)
	result := 0
	for {
		count, _ := f.Read(data)
		if count > 0 {
			fmt.Println(count)
			result += count
		} else {
			break
		}
	}
	return result, nil
}

func prefixer(inp string) func(string) string {
	return func(inpSecond string) string {
		res := fmt.Sprintf("%s %s", inp, inpSecond)
		return res
	}
}
