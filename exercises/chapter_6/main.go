package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("Hello world")
	test1 := []string{"sadsda", "sadasdasd", "dsa", "adwa", "adww"}
	test2 := []string{"sadsda", "sadasdasd", "dsa", "adwa", "adww"}
	fmt.Println(test1)
	fmt.Println("przed")
	UpdateSlice(test1, "www")
	fmt.Println(test1)
	fmt.Println("Przed")
	GrowSlice(test2, "ddd")
	fmt.Println(test2)
	fmt.Println("PO2")
	count := 0
	start := time.Now()
	var arr []Person
	for {
		count++
		p := Person{
			"dsadsa",
			"sadsda",
			12,
		}
		arr = append(arr, p)
		if count >= 10000000 {
			break
		}
	}
	duration := time.Since(start)
	fmt.Printf("Loop time %v", duration)
}

func MakePerson(firstName string, lastName string, age int) Person {
	p := Person{
		firstName,
		lastName,
		age,
	}
	return p
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	p := Person{
		firstName,
		lastName,
		age,
	}
	return &p
}

func UpdateSlice(arr []string, upd string) {
	length := len(arr)
	arr[length-1] = upd
	fmt.Println(arr)
}

func GrowSlice(arr []string, add string) {
	arr = append(arr, add)
	fmt.Println(arr)
}
