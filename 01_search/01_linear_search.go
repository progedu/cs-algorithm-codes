package main

import "fmt"

//! [marker_linear_search]
// 線形探索
// a: 本棚にある本の巻数が入った配列
// x: 探したい本の巻数
func LinearSearch(a []int, x int) int {
	var n = len(a)
	// 配列の中身についてループを回す
	for i := 0; i < n; i += 1 {
		// i 番目の要素が探したいものであれば、i を返す
		if a[i] == x {
			return i
		}
		// (*)
	}
	// ループが終わっても見つからなかった場合、-1 を返す
	return -1
}

//! [marker_linear_search]

func main() {
	// 本の冊数の読み込み
	var n int
	fmt.Scan(&n)

	// 本のリストの読み込み
	var a []int = make([]int, n)
	for i := 0; i < n; i += 1 {
		fmt.Scan(&a[i])
	}

	// 探したい本の読み込み
	var x int
	fmt.Scan(&x)

	// 線形探索の実行
	var result = LinearSearch(a, x)
	// 結果の表示
	fmt.Println(result)
}
