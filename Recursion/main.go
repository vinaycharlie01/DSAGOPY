package main

import "fmt"

//direct recursive function
func Vinay(a int) int {
	var s int
	if a == 0 {
		return a
	}
	s = a + Vinay(a-1)
	return s

}

//indirect recursive function
func fun1(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fun2(n-1)
	}
}
func fun2(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fun1(n-1)
	}
}

// tail recursive function
func tail(n int) {
	if n < 1 {
		return
	} else {
		fmt.Printf("%d  ", n)
		tail(n / 2)
	}

}

// /non tail recursive function
// /-----> non tail recursive function if there is no expre,ver,cond  after calling function
func non_tail(n int) {
	if n < 1 {
		return
	} else {
		fmt.Printf("%d  ", n)
		non_tail(n / 2)
		fmt.Printf("%d", n)
	}
}

// non tail recursive function
func hello2(a int) int {
	if a < 1 {
		return 0
	}
	return 1 + hello2(a/2)
}

func main() {

	//fmt.Println(fun2(2))
	fmt.Println(fun1(5))

	//fmt.Println(Vinay(5))
	// fmt.Println(0 < 1)
}
