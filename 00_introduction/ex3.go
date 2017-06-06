package main

import "fmt"

//! [marker_calc]
func calc(n int) int {
	var result = 0
	result += function1(n)
	for i := 0; i < n; i++ {
		if n%2 == 0 {
			result += function2(n)
		} else {
			result += function3(n)
		}
	}
	return result
}

//! [marker_calc]
// O(n)
func function1(n int) int {
	var result = 0
	for i := 0; i < n; i++ {
		result += 1
	}
	return result
}

// O(n^2)
func function2(n int) int {
	var result = 0
	for i, m := int64(0), int64(n); i < m*m; i++ {
		result += 1
	}
	return result
}

// O(n^3)
func function3(n int) int {
	var result = 0
	for i, m := int64(0), int64(n); i < m*m*m; i++ {
		result += 1
	}
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(calc(n))
}
