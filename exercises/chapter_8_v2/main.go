package main

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	Val  T
	Next *Node[T]
}

func main() {
	fmt.Println("Hello world")
	node := &Node[int]{}
	node.Add(5)
	node.Add(10)
	node.Add(15)
	tenIndex := node.Index(10)
	fmt.Printf("Index: %d\n", tenIndex)
	fifteenIndex := node.Index(15)
	fmt.Printf("Index: %d\n", fifteenIndex)
	zeroIndex := node.Index(0)
	fmt.Printf("INdex: %d\n", zeroIndex)
	node.Insert(7, 1)
	node.Insert(12, 3)
	sevenIndex := node.Index(7)
	twelveIndex := node.Index(12)
	tenIndex = node.Index(10)
	fifteenIndex = node.Index(15)
	fmt.Printf("7 Index %d\n", sevenIndex)
	fmt.Printf("12 Index %d\n", twelveIndex)
	fmt.Printf("0 INdex: %d\n", zeroIndex)
	fmt.Printf("10 Index: %d\n", tenIndex)
	fmt.Printf("15 Index: %d\n", fifteenIndex)
}

func (n *Node[T]) Add(val T) {
	nextNode := &Node[T]{val, nil}
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = nextNode
}

func (n *Node[T]) Index(val T) int {
	current := n
	counter := 0
	for current != nil {
		if current.Val == val {
			return counter - 1
		}
		current = current.Next
		counter++
	}
	return -1
}

func (n *Node[T]) Insert(val T, index int) {
	current := n
	counter := 0
	for current != nil {
		if counter == index {
			insertNode := &Node[T]{val, current.Next}
			current.Next = insertNode
			return
		}
		current = current.Next
		counter++
	}
	insertNode := &Node[T]{val, nil}
	current.Next = insertNode
}
