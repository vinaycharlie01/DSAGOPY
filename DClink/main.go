package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type link struct {
	head *Node
	tail *Node
}

func (l *link) creating(v int) {
	newNode := &Node{data: v}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
	}
	l.tail = newNode
}

func (l *link) displayforword() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}

}
func diplay(head *Node) *Node {
	p := head
	for p != nil {
		p = p.next
	}
	return p.next

}
func (l *link) displayreverce() {
	p := l.tail
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.prev
	}
}

func main() {
	// l := link{}
	l := new(link)

	l.creating(10)
	l.creating(20)
	l.creating(30)
	l.creating(40)
	l.creating(50)
	//l.head.display()
	l.displayforword()
}

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type Node struct {
// 	data int
// 	next *Node
// }

// type singleLinklist struct {
// 	head *Node
// 	tail *Node
// }

// func (l *singleLinklist) push(n int) {
// 	newNode := &Node{data: n}
// 	if l.head == nil {
// 		l.head = newNode
// 	} else {
// 		l.tail.next = newNode

// 	}
// 	l.tail = newNode
// }

// func (l *singleLinklist) pushhead(n int) {
// 	newnode := &Node{data: n}
// 	newnode.next = l.head
// 	l.head = newnode

// }
// func (l *singleLinklist) display() {
// 	p := l.head
// 	for p != nil {
// 		fmt.Print(p.data, " ")
// 		p = p.next
// 	}
// }

// func sortList(head *Node) *Node {
// 	//sorting
// 	arr := []int{}
// 	for head != nil {
// 		arr = append(arr, head.data)
// 		head = head.next
// 	}
// 	sort.Ints(arr)

// 	newNode := &Node{}
// 	s := newNode
// 	for _, v := range arr {
// 		s.next = &Node{data: v}
// 		s = s.next
// 	}
// 	return newNode.next

// }
// func main() {
// 	//Enter your code here. Read input from STDIN. Print output to STDOUT
// 	l := singleLinklist{}
// 	var a int
// 	fmt.Print()
// 	fmt.Scanln(&a)
// 	for i := 0; i < a; i++ {
// 		var b int
// 		fmt.Print()
// 		fmt.Scanln(&b)
// 		sortList(&Node{data: b})
// 	}
// 	l.display()
// }
