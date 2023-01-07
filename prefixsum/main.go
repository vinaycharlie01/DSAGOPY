package main

import "fmt"

func PrifixSum(arr []int) []int {
	prefix := make([]int, len(arr)+1)

	for i := 1; i <= len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	return prefix
}

func PrifixSum2(nums []int) []int {
	prefix := make([]int, len(nums))
	for i, num := range nums {
		prefix[i] = num
		if i > 0 {
			prefix[i] += prefix[i-1]
		}
	}
	return prefix
}

func main() {
	a := []int{10, 20, 30, 40, 50, 50}

	fmt.Println(PrifixSum(a))
	fmt.Println(PrifixSum2(a))
}
