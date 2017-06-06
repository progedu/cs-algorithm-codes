package main

import "fmt"

//! [marker_gcd]

// 整数 a, b (a >= b > 0) の最大公約数 (greatest common divisor) を求める。
func GreatestCommonDivisor(a int, b int) int {
	for i := b; i > 1; i -= 1 {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

//! [marker_gcd]

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(GreatestCommonDivisor(a, b))
}
