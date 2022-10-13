package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	data int
	next *Node
}

type Link struct {
	head *Node
	tail *Node
}

func (l *Link) append(n int) {
	newNode := &Node{data: n}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func main() {
	var a = []any{"vinay", 10, true}
	fmt.Println(reflect.TypeOf(a))
}
