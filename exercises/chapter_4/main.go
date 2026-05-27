package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	randomNums := make([]int, 0)
	for i := 1; i < 100; i++ {
		randomNumber := rand.Intn(100)
		fmt.Println(randomNumber)
		switch {
		case randomNumber%2 == 0 && randomNumber%3 == 0:
			fmt.Println("Six")
		case randomNumber%2 == 0:
			fmt.Println("Two")
		case randomNumber%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never mind")
		}
		randomNums = append(randomNums, randomNumber)

	}
	fmt.Println(randomNums)

	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
}
