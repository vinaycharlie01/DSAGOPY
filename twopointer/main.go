package main

import "fmt"

func merge(arr1 []int, m int, arr2 []int, n int) []int {
	L := 0
	R := 0
	var a []int
	for L < m && R < n {
		if arr1[L] < arr2[R] {
			a = append(a, arr1[L])
			L++
		} else {
			a = append(a, arr2[R])
			R++
		}
	}
	if L < m {
		a = append(a, arr1[L:m]...)
	} else {
		a = append(a, arr2[R:n]...)
	}
	copy(arr1, a)
	return a
	// for i := 0; i < len(a); i++ {
	// 	arr1[i] = a[i]
	// }
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 3, 4}
	fmt.Println(merge(arr1, len(arr1), arr2, len(arr2)))

}
