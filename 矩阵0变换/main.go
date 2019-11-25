package main

import (
	"fmt"
)

// 给一个 m * n 的矩阵，如果有一个元素为0，则把该元素对应的行与列的所有元素全部变成0。

func PrintBoard(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func transform(array [][]int)  {
	m := make([]int, 0, 2)
	n := make([]int, 0, 2)
	for i := 0; i < len(array); i ++ {
		for j := 0; j < len(array[i]); j ++ {
			if array[i][j] == 0 {
				m = append(m, i)
				n = append(n, j)
			}
		}
	}

	for i := 0; i < len(m); i ++ {
		array[m[i]] = make([]int, len(array[0]))
	}
	for i := 0; i < len(n); i ++ {
		for j := 0; j < len(array); j ++ {
			array[j][n[i]] = 0
		}
	}
}

func main()  {
	M := 10
	N := 9
	board := make([][]int, 0, M)
	for m := 0; m < M; m++ {
		array := make([]int, 0, N)
		for n := 0; n < N; n++ {
			array = append(array, 1)
		}
		board = append(board, array)
	}
	board[0][5] = 0
	board[7][8] = 0
	transform(board)
	PrintBoard(board)


}
