package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Clink struct {
	tail *Node
}

func (l *Clink) append(n int) {
	newNode := &Node{data: n}
	if l.tail == nil {
		l.tail = newNode
		l.tail.next = newNode
	} else {
		newNode.next = l.tail.next
		l.tail.next = newNode
		l.tail = newNode
	}

}

func (l *Clink) push(n int, pos int) {
	if pos == 0 {
		newNode := &Node{data: n}
		newNode.next = l.tail.next
		l.tail.next = newNode
	} else if pos >= l.len() {
		newNode := &Node{data: n}
		newNode.next = l.tail.next
		l.tail.next = newNode
		l.tail = newNode
	} else if pos > 0 && pos < l.len() {
		newNode := &Node{data: n}
		p := l.tail
		i := 0
		for i < pos {
			p = p.next
			i += 1
		}
		newNode.next = p.next
		p.next = newNode
	}
}

func (l *Clink) pop(pos int) {
	if pos == 0 {
		temp := l.tail.next.next
		l.tail.next = temp
		temp = nil
		// fmt.Println("")
		// fmt.Println(l.tail.next.next.data)
	} else if pos >= l.len()-1 {
		p := l.tail
		for p.next.next != l.tail {
			p = p.next
		}
		temp := l.tail.next
		p.next.next = temp
		l.tail = p.next
		temp = nil
	} else if pos > 0 && pos < l.len() {
		p := l.tail
		i := 1
		for i < pos {
			p = p.next
			i += 1
		}
		temp := p.next.next
		p.next.next = temp.next
		temp.next = nil
	}
}

func (l *Clink) len() int {
	temp := l.tail
	count := 1
	for l.tail != temp.next {
		count += 1
		temp = temp.next
	}
	return count
}

func (l *Clink) reverse() {
	var prev, currNode, nextNode *Node
	currNode = l.tail.next
	nextNode = currNode.next
	for currNode != l.tail {
		prev = currNode
		currNode = nextNode
		nextNode = currNode.next
		currNode.next = prev
	}
	nextNode.next = l.tail
	l.tail = nextNode
	p := l.tail
	temp := l.tail
	for l.tail != temp.next {
		temp = temp.next
		fmt.Print(temp.data, " ")
	}
	fmt.Print(p.data)
}

func (l *Clink) disp() {
	p := l.tail
	temp := l.tail
	for l.tail != temp.next {
		temp = temp.next
		fmt.Print(temp.data, " ")
	}
	fmt.Print(p.data)

}

func main() {
	l := Clink{}
	l.append(10)
	l.append(20)
	l.append(30)
	l.append(40)
	l.append(50)
	l.push(100, 2)
	l.push(300, 6)
	//l.pop(4)
	//l.reverse()
	l.disp()
	b := l.len()
	fmt.Println("")
	fmt.Println(b)

}

// package main

// import (
// 	"container/ring"
// 	"fmt"
// )

// func main() {
// 	numbers := []int{10, 20, 30, 40, 50, 60, 70}
// 	r := ring.New(len(numbers))
// 	// initialize ring
// 	for i := 0; i < r.Len(); i++ {
// 		r.Value = numbers[i]
// 		r = r.Next()
// 	}

// 	//Do calls function f on each element of the ring, in forward order.

// 	r.Do(func(x interface{}) { fmt.Print(x, " ") })
// 	fmt.Println()
// 	// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// 	// in the ring and returns that ring element. r must not be empty.
// 	//r = r.Move(5)

// 	// print elements in reverse order
// 	for i := 0; i < r.Len(); i++ {
// 		r = r.Prev()
// 		fmt.Print(r.Value, " ")
// 	}
// 	fmt.Println()

// 	// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// 	// in the ring and returns that ring element. r must not be empty.
// 	r = r.Move(3)
// 	r.Do(func(x interface{}) {
// 		fmt.Print(x, " ")
// 	})
// 	fmt.Println()

// 	//r.Move(3)
// 	r.Unlink(3)
// 	r.Do(func(p any) {
// 		fmt.Println(p.(int))
// 	})

// }
