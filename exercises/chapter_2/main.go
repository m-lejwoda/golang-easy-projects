package main

import "fmt"

func main() {
	integerValue := 20
	fmt.Println(integerValue)
	fmt.Printf("%T\n", integerValue)
	floatValue := float64(integerValue)
	fmt.Println(floatValue)
	fmt.Printf("%T\n", floatValue)
	const value = 30
	intVal := int(value)
	floatVal := float64(value)
	fmt.Println(intVal)
	fmt.Println(floatVal)
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
	b++
	b++
	b++
	smallI++
	bigI++
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
