package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BabbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var a bool
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				a = true
			}
		}
		if !a {
			break
		}
	}

}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func InsertionSort(arr []int) {
	//O(N)
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func ShellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for j := gap; j < len(arr); j++ {
			for i := j - gap; i >= 0; i = i - gap {
				if arr[i+gap] > arr[i] {
					break
				} else {
					arr[i], arr[i+gap] = arr[i+gap], arr[i]
				}

			}
		}
	}
}

func Partion(arr []int, lb int, ub int) int {
	Pivot := arr[lb]
	start := lb
	end := ub
	for start < end {
		if arr[start] <= Pivot {
			start++
		}
		for arr[end] > Pivot {
			end--
		}
		if start < end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
		}
	}
	temp := arr[lb]
	arr[lb] = arr[end]
	arr[end] = temp
	return end
}

func QueKSort(arr []int, lb int, ub int) {
	if lb < ub {
		Mid := Partion(arr, lb, ub)
		QueKSort(arr, lb, Mid-1)
		QueKSort(arr, Mid+1, ub)
	}
}

func Merge(a []int, b []int) []int {
	L := 0
	R := 0
	arr := []int{}
	for L < len(a) && R < len(b) {
		if a[L] < b[R] {
			arr = append(arr, a[L])
			L++
		} else {
			arr = append(arr, b[R])
			R++
		}
	}
	if L < len(arr) {
		arr = append(arr, a[L:]...)
	}
	if R < len(arr) {
		arr = append(arr, b[R:]...)
	}
	return arr

}

func MergeSort(arr []int) (arr2 []int) {
	if len(arr) < 2 {
		return arr
	}
	a := MergeSort(arr[:len(arr)/2])
	b := MergeSort(arr[len(arr)/2:])
	s := Merge(a, b)
	for i, v := range s {
		arr[i] = v
	}
	return arr

}

func ReadLine(reader *bufio.Reader) string {
	val, _, _ := reader.ReadLine()
	return strings.TrimSpace(string(val))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// val, _, _ := reader.ReadLine()
	// b := strings.TrimSpace(string(val))
	strarr := strings.Split(strings.TrimSpace(ReadLine(reader)), " ")
	var arr []int
	for i := 0; i < len(strarr); i++ {
		val2, _ := strconv.Atoi(strarr[i])
		arr = append(arr, val2)
	}

	//BabbleSort(arr)
	//SelectionSort(arr)
	//InsertionSort(arr)
	//ShellSort(arr)

	//QueKSort(arr, 0, len(arr)-1)
	MergeSort(arr)
	fmt.Println(arr)

}
