package main

import "fmt"

func main() {
	// First exercise
	fmt.Println("Hello world")
	greetings := []string{"Hello", "Hola", "Cześć", "sdasa", "sadasdas"}
	twoFirstVal := greetings[:2]
	fmt.Println(twoFirstVal)
	thirdFourth := greetings[2:4]
	fmt.Println(thirdFourth)
	lastTwo := greetings[3:5]
	fmt.Println(lastTwo)
	// Second exercise
	message := "Hi 👧 and 👨"
	fmt.Println(message)
	bs := []byte(message)
	fmt.Println(bs)
	runeEx := []rune(message)
	fmt.Println(runeEx)
	// Third exercise
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	f := Employee{
		"Imie1", "Nazwisko1", 1,
	}
	s := Employee{firstName: "Imie2", lastName: "Nazwisko2", id: 2}
	var t Employee
	t.firstName = "Imie3"
	t.lastName = "Nazwisko3"
	t.id = 3
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(t)
}
