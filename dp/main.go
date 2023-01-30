package main

import "fmt"

var count int

// top down approch
func Fibrecurtion(num int) int {
	//count++
	if num < 0 {
		fmt.Println("err")
	}
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return Fibrecurtion(num-1) + Fibrecurtion(num-2)
}

func Fib(n int, cache map[int]int64) int64 {
	// If answer already found for n, return it
	count++
	if val, found := cache[n]; found {
		return val
	} else {
		cache[n] = Fib(n-1, cache) + Fib(n-2, cache)
	}
	fmt.Println(cache)
	//cache[n] = Fib(n-1, cache) + Fib(n-2, cache)
	return cache[n]
}

// bottom down approch
func Bott0m_Up_Fib(n int) int {
	ans := make([]int, n+1)

	ans[0], ans[1] = 0, 1
	for i := 2; i <= n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	fmt.Println(ans)
	return ans[n]
}

func longestPalindrome(s string) {
	n := len(s)
	start, end := 0, 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	fmt.Println(dp)
	//fmt.Println(dp)
	//	fmt.Println(0 > 0-0)
	// // "babad" 5
	for len := 0; len < n; len++ {
		for i := 0; i+len < n; i++ {
			fmt.Println(len, i)
			dp[i][i+len] = (s[i] == s[i+len] && (len < 2 || dp[i+1][i+len-1]))
			if dp[i][i+len] && len > end-start {
				start = i
				end = i + len
			}
		}
	}
	fmt.Println(s[start : end+1])
	//return s[start : end+1]
}

func longestPalindrome2(s string) string {
	longestPalindrom := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// filling out the diagonal by 1
	for i := range s {
		dp[i][i] = true
		longestPalindrom = s[i : i+1]
	}
	fmt.Println(dp)
	fmt.Println(longestPalindrom)

	// filling the dp table
	for i := len(s) - 1; i >= 0; i-- {
		// j starts from the i location : to only work on the upper side of the diagonal
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				// if the chars mathces
				// if len slicied sub_string is just one letter if the characters are equal, we can say they are palindomr dp[i][j] =True
				//if the slicied sub_string is longer than 1, then we should check if the inner string is also palindrom (check dp[i+1][j-1] is True)
				if j-i == 1 || dp[i+1][j-1] {
					dp[i][j] = true
					// we also need to keep track of the maximum palindrom sequence
					if len(longestPalindrom) < len(s[i:j+1]) {
						longestPalindrom = s[i : j+1]
					}
				}
			}
		}
	}

	return longestPalindrom
}

func longestPalindrome3(s string) string {
	longestPalindrom := ""
	dp := make([][]bool, len(s))
	st := 0
	en := 0
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		longestPalindrom = s[i : i+1]
	}
	// for i := range s {
	// 	dp[i][i] = true
	// 	//longestPalindrom = s[i : i+1]
	// }
	for start := len(s) - 1; start >= 0; start-- {
		for end := start + 1; end < len(s); end++ {
			//fmt.Println(start, end)
			if s[start] == s[end] {
				if end-start == 1 || dp[start+1][end-1] {
					dp[start][end] = true
					if len(longestPalindrom) < len(s[start:end+1]) {
						longestPalindrom = s[start : end+1]
					}

				}
			}
		}
	}
	fmt.Println(s[st : en+1])

	return longestPalindrom
}

func coinChange(coins []int, amount int) {

	a := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		a[i] = make([]int, amount)
	}
	for i := 0; i < len(coins); i++ {
		a[i][0] = 1
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j < count; j++ {

		}
	}
}

func main() {

	coinChange([]int{2, 3, 5, 10}, 15)
	//a := "babad"
	// longestPalindrome(a)
	// fmt.Println("hh")
	//fmt.Println(longestPalindrome3(a))
	//longestPalindrome(a)
	//FibonacciTopDown(10)
	// a := 3
	//fmt.Println(Fibrecurtion(a))
	// fmt.Println(Fib(a, map[int]int64{0: 0, 1: 1}))
	// fmt.Println(Bott0m_Up_Fib(a))

}
