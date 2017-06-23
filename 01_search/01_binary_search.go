package main

import "fmt"

//! [marker_binary_search]
// 二分探索
// a: 本棚にある本の巻数が入った配列
// x: 探したい本の巻数
func BinarySearch(a []int, x int) int {
	// 配列が空ならば、ただちに存在しないことを報告する
	if len(a) == 0 {
		return -1
	}

	var low, high = 0, len(a)
	// 探す値が配列の最初の値より小さいか、最後の値より大きいならば、ただちに存在しないことを報告する
	if x < a[low] || a[high-1] < x {
		return -1
	}
	// 探す値が配列の最後の値と等しければ、それを返す
	if a[high-1] == x {
		return high - 1
	}

	// (1)

	for low+1 < high {
		// low と high の中央の位置を middle とする
		var middle = low + (high-low)/2
		if a[middle] <= x {
			// 本は昇順に並んでいることから、low と middle の間に x は存在しないことがわかる
			low = middle
			// (2)
		} else {
			// 同様に、middle と high の間に x は存在しないことがわかる
			high = middle
			// (3)
		}
	}

	if a[low] == x {
		return low
	} else {
		return -1
	}
}

//! [marker_binary_search]

func main() {
	var n int
	fmt.Scan(&n)

	var a []int = make([]int, n)
	for i := 0; i < n; i += 1 {
		fmt.Scan(&a[i])
	}

	var x int
	fmt.Scan(&x)

	var result = BinarySearch(a, x)
	fmt.Println(result)
}
