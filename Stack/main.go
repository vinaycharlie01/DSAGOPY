// package main

// import "fmt"

// type stack []interface{}

// func (k *stack) push(s interface{}) {
// 	*k = append(*k, s)
// }

// func (k *stack) pop(n int) (s interface{}, ok bool) {
// 	if k.empty() {
// 		return
// 	}
// 	if n > len(*k) {
// 		fmt.Println("stack out of range")
// 	} else {
// 		for i := 0; i < n; i++ {
// 			last := len(*k) - 1
// 			s = (*k)[last]
// 			*k = (*k)[:last]
// 		}
// 	}
// 	return s, true
// }

// func (k *stack) peek() (s interface{}, ok bool) {
// 	if k.empty() {
// 		return
// 	}
// 	last := len(*k) - 1
// 	s = (*k)[last]
// 	return s, true
// }

// func (k *stack) empty() bool {
// 	return len(*k) == 0
// }

// func (k *stack) diplays() {
// 	i := len(*k) - 1
// 	for i >= 0 {
// 		fmt.Println((*k)[i], " ")
// 		i -= 1
// 	}
// }

// func main() {
// 	// s := stack{}
// 	// s.push(1)
// 	// s.push(2)
// 	// s.push(3)
// 	// s.push(4)
// 	// s.push(5)
// 	// fmt.Println(s)
// 	// s.diplays()
// 	// w, q := s.pop(1)
// 	// fmt.Println(w, " ", q)
// 	// s.diplays()
// 	// fmt.Println("")
// 	// s.diplays()
// 	// fmt.Println(s)
// 	// y, _ := s.peek()
// 	// fmt.Println("peak", y)

// 	a := []int{1}
// 	a = append(a, 10)
// 	fmt.Println(cap(a))
// }

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stackLink struct {
	top *Node
}

func (s *stackLink) push(n int) {
	newNode := &Node{data: n}
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *stackLink) display() {
	p := s.top
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
func (s *stackLink) pop() {
	if s.top == nil {
		fmt.Println("stack is empty")
	} else {
		temp := s.top
		s.top = temp.next
		temp = nil
	}
}

func (s *stackLink) peek() {
	if s.top == nil {
		fmt.Println("This is empty")
	} else {
		fmt.Println(s.top.data)
	}
}

func main() {
	l := stackLink{}
	l.push(10)
	l.push(20)
	l.push(30)
	l.push(40)
	l.pop()
	l.display()

}
