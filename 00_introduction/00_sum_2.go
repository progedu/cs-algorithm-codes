package main

import "fmt"

//! [marker_sum]

// 1 から n までの整数の合計を等差数列の和の公式で求める。
func Sum(n int) int {
	return n * (n + 1) / 2
}

//! [marker_sum]

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Sum(n))
}
