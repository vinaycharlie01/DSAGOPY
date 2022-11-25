package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	data  int
	right *BinaryTreeNode
}

func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PreOrder(root.left)
	fmt.Printf("%d ", root.data)
	PreOrder(root.right)
}

// func inorderTraversal(root *BinaryTreeNode) []int {
// 	var a []int

// 	return a
// }

// func postorderTraversal(root *BinaryTreeNode) (a []int) {

// }

func postorder(root *BinaryTreeNode, arr *[]int) {
	if root == nil {
		return
	}
	postorder(root.left, arr)
	postorder(root.right, arr)
	*arr = append(*arr, root.data)
}

func PostOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%d ", root.data)
}

func height(root *BinaryTreeNode) int {
	if root == nil {
		return -1
	} else {
		//return Math.max(height(root.left)+1, height(root.right)+1)
		b := math.Max(float64(height(root.left)+1), float64(height(root.right)+1))
		return int(b)
	}
}

func rightSideView(root *BinaryTreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}
	bfs := []*BinaryTreeNode{}
	bfs = append(bfs, root)

	for len(bfs) > 0 {
		size := len(bfs)
		last := size - 1
		for i := 0; i < size; i++ {
			var current *BinaryTreeNode
			current, bfs = bfs[0], bfs[1:]
			if i == last {
				result = append(result, current.data)
			}
			if current.left != nil {
				bfs = append(bfs, current.left)
			}
			if current.right != nil {
				bfs = append(bfs, current.right)
			}
		}
	}

	return result
}
func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

// def topView(root):
//     d = {}
//     def traverse(root,key,level):
//         if root:
//             if key not in d:
//                 d[key] = [root.info, level]
//             elif d[key][1]>level:
//                 d[key] = [root.info, level]
//             traverse(root.left,key -1,level+1)
//             traverse(root.right,key+1,level+1)
//     traverse(root,0,0)
//     for key in sorted(d):
//         print(d[key][0],end = ' ')
var d = map[int][]int{}

func traverse(root *BinaryTreeNode, key int, level int) {
	if root != nil {
		_, ok := d[key]
		if !ok {
			d[key] = []int{root.data, level}

		} else if d[key][1] > level {
			d[key] = []int{root.data, level}

		}
		traverse(root.left, key-1, level+1)
		traverse(root.right, key+1, level+1)
	}
}
func topView(root *BinaryTreeNode) {
	traverse(root, 0, 0)
	keys := make([]int, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Print(d[k][0], " ")
	}
}
func NewBinaryTree() *BinaryTreeNode {
	var root *BinaryTreeNode
	reader := bufio.NewReader(os.Stdin)
	val, _, _ := reader.ReadLine()
	input, _ := strconv.ParseInt(string(val), 10, 64)
	val2, _, _ := reader.ReadLine()
	ss := strings.Split(string(val2), " ")
	for i := 0; i < int(input); i++ {
		ss1, _ := strconv.Atoi(ss[i])
		root = insert(root, ss1)
	}
	return root
}

func levelOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	res := [][]int{}
	treeNodeList := []*BinaryTreeNode{root}
	bfs(treeNodeList, &res)
	for _, v := range res {
		for _, data := range v {
			fmt.Print(data, " ")
		}
	}
}

func bfs(node []*BinaryTreeNode, res *[][]int) {
	currentLevel := []int{}
	nextLevel := []*BinaryTreeNode{}
	for _, e := range node {
		if e != nil {
			currentLevel = append(currentLevel, e.data)
			nextLevel = append(nextLevel, e.left)
			nextLevel = append(nextLevel, e.right)
		}
	}
	if len(currentLevel) > 0 {
		*res = append(*res, currentLevel)
		bfs(nextLevel, res)
	}
}

func main() {
	t1 := NewBinaryTree()
	topView(t1)
	//PreOrder(t1)
	//InOrder(t1)
	//PostOrder(t1)
	//fmt.Println(inorderTraversal(t1))
	//inorderTraversal(t1)
	//fmt.Println(height(t1))

}
