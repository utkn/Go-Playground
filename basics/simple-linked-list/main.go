package main

import (
	"Go-Playground/basics/simple-linked-list/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)

	fmt.Println(list)

	list.Remove(0)
	fmt.Println(list)

	list.Remove(2)
	fmt.Println(list)
}
