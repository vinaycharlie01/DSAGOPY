package main

import (
	"fmt"
	"math/big"
)

func LienearSearch(arr []int, key int) int {
	i := 0             //0(1)
	for i < len(arr) { //O(N)
		if arr[i] == key { //O(N)
			return i
		}
		i += 1
	}
	return -1
}

func BinarySearch(arr []int, key int) int {
	//O(log(N))
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == key {
			return mid
		} else if key > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func removeDuplicates(arr []int) {
	var a bool = true
	for a {
		a = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] == arr[i+1] {
				a = true
				arr = append(arr[i:], arr[i+1:]...)
				break
			}
		}
	}
	fmt.Println(arr)

}

func Prime(a int) bool {
	if big.NewInt(int64(a)).ProbablyPrime(20) {
		return true
	} else {
		return false
	}
}

func main() {

	//arr := []int{1, 2, 3, 4}

	// fmt.Println(LienearSearch(arr, 4))
	//fmt.Println(BinarySearch(arr, 4))

	//fmt.Println(Prime(5))
	a := []int{0, 0, 0, 0}
	removeDuplicates(a)
	// s := "_"
	// b, _ := strconv.Atoi(s)
	// fmt.Println(b)

}
