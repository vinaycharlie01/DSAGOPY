package main

import "fmt"

func SlidingWindow3(arr []int, target int, Windowsize int) {
	Windowsize2 := Windowsize
	var right int
	var left int
	var sum int
	var count int
	//[1, 2, 1, 3, 2]
	for right < len(arr) {
		sum += arr[right]
		if right-left != Windowsize2-1 {
			right++
		} else if right-left == Windowsize2-1 {
			if sum == target {
				count++
			}
			sum -= arr[left]
			left++
			right++
		}

	}
	fmt.Println(count)

}
func main() {

}
