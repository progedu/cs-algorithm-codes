package main

import "fmt"

//! [marker_gcd]

// ユークリッドの互除法
// 整数 a, b (a >= b > 0) の最大公約数を求める。
func EuclideanAlgorithm(a int, b int) int {
	var r = a % b
	if r == 0 {
		return b
	} else {
		return EuclideanAlgorithm(b, r)
	}
}

//! [marker_gcd]

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(EuclideanAlgorithm(a, b))
}
