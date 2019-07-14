package main

import (
	"Go-Playground/basics/simple-linked-list/LinkedList"
	"fmt"
)

func main() {
	list := LinkedList.LinkedList{}
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
