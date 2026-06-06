package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	tree := &Tree{}
	tree.Add(5)
	tree.Add(7)
	tree.Add(6)
	tree.Add(19)
	tree.Add(10)
	fmt.Println(tree.Root.Val)
	fmt.Println(tree.Root.Left.Val)
	fmt.Println(tree.Root.Right.Val)
	fmt.Println(tree.Root.Left.Left.Val)
	fmt.Println(tree.Root.Left.Right.Val)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	} else {
		return fmt.Sprintf("%v", n.Val)
	}
}

func (t *Tree) Add(val int) {
	if t.Root == nil {
		t.Root = &Node{val, nil, nil}
		return
	}
	queue := []*Node{t.Root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil {
			current.Left = &Node{val, nil, nil}
			return
		} else {
			queue = append(queue, current.Left)
		}
		if current.Right == nil {
			current.Right = &Node{val, nil, nil}
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}
