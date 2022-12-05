package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BTree struct {
	root *Node
}

func (b *BTree) Create(data int) {
	NewNode := &Node{data: data}
	if b.root == nil {
		b.root = NewNode
		return
	}
	b.root.Insert(data)

}

func (b *Node) Insert(data int) {
	NewNode := &Node{data: data}
	if data < b.data {
		if b.left == nil {
			b.left = NewNode
			return
		} else {
			b.left.Insert(data)
			return
		}

	} else if data > b.data {
		if b.right == nil {
			b.right = NewNode
			return
		} else {
			b.right.Insert(data)
			return
		}
	}
}

func (b *BTree) PreOrder(Br *Node) {
	if Br == nil {
		return
	}
	fmt.Print(Br.data, " ")
	b.PreOrder(Br.left)
	b.PreOrder(Br.right)
}

func Input() *BTree {
	in := new(BTree)
	reader := bufio.NewReader(os.Stdin)
	val, _, _ := reader.ReadLine()
	input, _ := strconv.ParseInt(string(val), 10, 64)
	val2, _, _ := reader.ReadLine()
	ss := strings.Split(string(val2), " ")
	for i := 0; i < int(input); i++ {
		ss1, _ := strconv.Atoi(ss[i])
		in.Create(ss1)
	}
	return in
}

func main() {
	b := Input()
	b.PreOrder(b.root)

}
