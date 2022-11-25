package main

import "fmt"

func HashMap(nums []int, target int) (int, bool, int) {
	chek := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		chek[i] = nums[i]
		val, ok := chek[i]
		if val == target {
			return val, ok, i
		}
	}
	return target, false, -1
}

func main() {
	a := []int{1, 2, 3, 4}
	val, ok, index := HashMap(a, 3)
	fmt.Println("targetNumber:", val, "Result:", ok, "Index Numeber:", index)

}
