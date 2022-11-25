package main

import (
	"fmt"
	"sort"
)

type Node struct {
	data int
	next *Node
}
type Slink struct {
	head *Node
	tail *Node
}

func (l *Slink) append(n int) {
	newNode := &Node{data: n}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func (l *Slink) push(n int, pos int) {
	if pos == 0 {
		newNode := &Node{data: n}
		newNode.next = l.head
		l.head = newNode
	} else if pos > 0 && pos < l.len() {
		newNode := &Node{data: n}
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		newNode.next = p.next
		p.next = newNode
	} else {
		newNode := &Node{data: n}
		l.tail.next = newNode
		l.tail = newNode
	}

}
func (l *Slink) len() int {
	p := l.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count

}

func (l *Slink) pop(pos int) int {
	b := 0
	if pos == 0 {
		temp := l.head
		l.head = l.head.next
		b = temp.data
		temp.next = nil
	} else if pos > 0 && pos < l.len()-1 {
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		temp := p.next
		p.next = p.next.next
		b = temp.data
		temp.next = nil
	} else {
		p := l.head
		for p.next != l.tail {
			p = p.next
		}
		b = p.next.data
		l.tail = p
		p.next = nil
	}
	return b
}

func (l *Slink) Reverse() *Node {

	// var prev *ListNode
	// currNode := head
	// nextNode := head
	// for nextNode != nil {
	// 	nextNode = nextNode.Next
	// 	currNode.Next = prev
	// 	prev = currNode
	// 	currNode = nextNode
	// }
	// return prev
	var prev *Node
	currNode := l.head
	nextNode := l.head
	for nextNode != nil {
		nextNode = nextNode.next
		currNode.next = prev
		prev = currNode
		currNode = nextNode
	}

	// p := prev
	// for p != nil {
	// 	fmt.Print(p.data, " ")
	// 	p = p.next
	// }
	l.head = prev
	l.tail = currNode

	return prev

}
func (l *Slink) sort() {
	var a []int
	p := l.tail
	for p != nil {
		a = append(a, p.data)
		p = p.next
	}
	sort.Ints(a)
	newNode := &Node{}
	s := newNode
	for _, v := range a {
		s.next = &Node{data: v}
		s = s.next
	}
	q := newNode.next
	for q != nil {
		fmt.Print(q.data, " ")
		q = q.next
	}

}
func (l *Slink) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}

}

func disp(n *Node) {
	for n != nil {
		fmt.Print(n.data, " ")
		n = n.next
	}

}

// func removeval(head *Node, val int) {
// 	fake := &Node{data: 0}
// 	fake.next = head
// 	curr := fake
// 	for curr.next != nil {
// 		if curr.next.data == val {
// 			curr.next = curr.next.next
// 		} else {
// 			curr = curr.next
// 		}
// 	}
// 	fmt.Println(fake.next.data) // Return the linked list...
// }

func main() {
	l := Slink{}
	l.append(20)
	l.append(10)
	l.append(2)
	l.append(2)
	l.append(3)
	l.append(4)
	l.append(60)
	// l.push(70, 3)
	l.pop(7)
	l.display()
	l.Reverse()
	l.display()

	// fmt.Println("")
	// // l.Reverse()
	// fmt.Println("")
	// b := l.Reverse()
	// disp(b)

	// disp(b)
	// fmt.Println("")
	// l.sort()
	// l.Reverse()

}
