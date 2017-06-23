package main

import (
	"fmt"
	"os"
	"time"
)

//! [marker_binary_search]
func binary_search(a []int, x int) int {
	if len(a) == 0 {
		return -1
	}

	var low, high = 0, len(a)
	if x < a[low] || a[high-1] < x {
		return -1
	}
	if a[low] == x {
		return low
	}

	for low+1 < high {
		var middle = low + (high-low)/2
		if a[middle] >= x {
			high = middle
		} else {
			low = middle
		}
	}

	if a[high] == x {
		return high
	} else {
		return -1
	}
}

//! [marker_binary_search]

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
	var result = binary_search(a, x)
	// 計測終了
	var end_time = time.Now()
	var duration = end_time.Sub(start_time)

	fmt.Println(result)
	fmt.Fprintln(os.Stderr, duration)
}
