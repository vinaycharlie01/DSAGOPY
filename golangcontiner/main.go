package main

import (
	"container/ring"
	"fmt"
	"sort"
)

func display(r *ring.Ring) {
	r.Do(func(a any) { fmt.Print(a, " ") })
}
func reverse(r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		r = r.Prev()
		fmt.Print(r.Value, " ")
	}
}

func move(r *ring.Ring) {
	r = r.Move(3)
	r.Do(func(x interface{}) { fmt.Print(x, " ") })
	fmt.Println()
}

func linktwo(r *ring.Ring, s *ring.Ring) {
	rs := r.Link(s)
	rs.Do(func(p any) {
		fmt.Println(p.(int))
	})
}

func main() {
	a := []int{10, 5, 30, 3, 50, 1, 70}
	sort.Ints(a)
	r := ring.New(len(a))
	for i := 0; i < len(a); i++ {
		r.Value = a[i]
		r = r.Next()

	}
	s := ring.New(len(a))
	for i := 0; i < len(a); i++ {
		s.Value = a[i]
		s = r.Next()

	}
	display(r)
	fmt.Println()
	reverse(r)
	fmt.Println()
	move(r)
	fmt.Println("")
	linktwo(r, s)
	// fmt.Println("")
	// r.Next().Do(func(a any) { fmt.Print(a, " ") })
}
