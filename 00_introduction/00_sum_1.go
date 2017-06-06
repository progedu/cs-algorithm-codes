package main

import "fmt"

//! [marker_sum]

// （Go での実装）
// 1 から n までの整数の合計を求める。
func Sum(n int) int {
	var result = 0
	for i = 1; i <= n; i += 1 {
		result += i
	}
	return result
}

//! [marker_sum]

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Sum(n))
}
