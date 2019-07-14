package main

type ID int
type ConnectionMap map[int]chan string

var lastID ID = 0

type Node struct {
	id          ID
	receiver    <-chan string
	connections ConnectionMap
}

func newNode() *Node {
	lastID++
	return &Node{id: lastID - 1,
		receiver:    make(chan string),
		connections: make(ConnectionMap)}
}
