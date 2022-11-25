package main

import "fmt"

func majorityElement(nums []int) int {
	///Boyer-Moore Majority Voting Algorithm
	//1
	candidate := -1
	//1
	votes := 0
	//# Finding majority candidate
	//[1, 1, 1, 1, 2, 3, 4] //we know 1
	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			candidate = nums[i]
			votes = 1
		} else if nums[i] == candidate {
			votes += 1
		} else {
			votes -= 1
		}
	}
	return (candidate)
}

func main() {
	/*
			Example 1:

		Input: nums = [3,2,3]
		Output: 3
		Example 2:

		Input: nums = [2,2,1,1,1,2,2]
		Output: 2
	*/
	a := []int{3, 2, 3}
	fmt.Println(majorityElement(a))

}
