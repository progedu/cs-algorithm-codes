package main

import (
	"fmt"
	"os"
	"time"
)

// a: 本棚にある本の巻数が入った配列
// x: 探したい本の巻数
func linear_search(a []int, x int) int {
	for i := 0; i < len(a); i += 1 {
		if a[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Scan(&n)

	var a []int = make([]int, n, 1000000)
	for i := 0; i < n; i += 1 {
		fmt.Scan(&a[i])
	}

	var x int
	fmt.Scan(&x)

	// 計測開始
	var start_time = time.Now()
	var result = linear_search(a, x)
	// 計測終了
	var end_time = time.Now()

	var duration = end_time.Sub(start_time)
	fmt.Println(result)
	fmt.Fprintln(os.Stderr, duration
}
