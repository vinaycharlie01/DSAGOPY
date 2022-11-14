package main

import "fmt"

var N = 5

type queue [5]int

var front = -1
var rear = -1

func (q *queue) append(x int) {
	if rear == N-1 {
		fmt.Println("Overflow")
	} else if front == -1 && rear == -1 {
		front = 0
		rear = 0
		(*q)[front] = x
	} else {
		rear++
		(*q)[rear] = x
	}
}

func (q *queue) remove() {
	if front == -1 && rear == -1 {
		fmt.Println("list is empty")
	} else if front == rear {
		front = -1
		rear = -1
	} else {
		fmt.Println((*q)[front])
		front++
	}

}

func (q *queue) display() {
	if front == -1 && rear == -1 {
		fmt.Println("queue is empty")
	} else {
		p := front
		for p != rear+1 {
			fmt.Println((*q)[p])
			p += 1
		}
	}
}

func main() {
	l := queue{}
	l.append(10)
	l.append(20)
	l.append(30)
	l.append(40)
	l.append(50)
	l.display()
	l.remove()
}
