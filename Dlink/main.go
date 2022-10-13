package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}
type Dlink struct {
	head *Node
	tail *Node
}

func (l *Dlink) append(n int) {
	newNode := &Node{data: n}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
	}
	l.tail = newNode
}

func (l *Dlink) push(n int, pos int) {
	if pos == 0 {
		newNode := &Node{data: n}
		temp := l.head
		l.head.prev = newNode
		newNode.next = temp
		l.head = newNode
	} else if pos > 0 && pos < l.len()-1 {
		newNode := &Node{data: n}
		temp := l.head
		i := 0
		for i < pos-1 {
			temp = temp.next
			i += 1
		}
		newNode.prev = temp
		newNode.next = temp.next
		temp.next.prev = newNode
		temp.next = newNode
	} else {
		newNode := &Node{data: n}
		temp := l.tail
		newNode.prev = temp
		temp.next = newNode
		l.tail = newNode

	}
}

func (l *Dlink) pop(pos int) {
	if pos == 0 {
		temp := l.head
		l.head = temp.next
		l.head.prev = nil
		temp.next = nil
	} else if pos > 0 && pos < l.len()-1 {
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		temp := p.next
		p.next = temp.next
		temp.next.prev = temp
		temp.next = nil
		temp.prev = nil
	} else {
		temp := l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		temp.prev = nil
		temp.next = nil
	}
}
func (l *Dlink) len() int {
	p := l.head
	count := 0
	for p != nil {
		count++
		p = p.next
	}
	return count

}

func (l *Dlink) reverce() {
	var prev *Node
	cuurentNode := l.head
	nextNode := l.head
	for nextNode != nil {
		nextNode = nextNode.next
		cuurentNode.next = prev
		prev = cuurentNode
		cuurentNode = nextNode
	}
	// temp := l.head
	l.head, l.tail = l.tail, l.head
	// l.tail = temp
	p := prev
	fmt.Println(l.head.data)
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println("")
	fmt.Println(l.tail.data)

}
func (l *Dlink) display() {
	p := l.head
	fmt.Println(l.head.data)
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println("")
	fmt.Println(l.tail.data)
}
func main() {
	l := Dlink{}
	l.append(10)
	l.append(20)
	l.append(30)
	l.append(40)
	l.append(50)
	l.append(60)
	l.push(200, 6)
	l.reverce()

}
