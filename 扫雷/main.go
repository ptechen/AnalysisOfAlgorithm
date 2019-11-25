package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintBoard(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func Algo(M, N, k, j int, array [][]int) {
	num := 0
	for y := k - 1; y <= k+1; y++ {
		if y < 0 || y > M-1 {
			continue
		}
		if array[y][j] == 9 {
			num++
		}
		x := j - 1
		if x >= 0 {
			if array[y][x] == 9 {
				num++
			}
		}
		x = j + 1
		if x < N {
			if array[y][x] == 9 {
				num++
			}
		}

	}
	array[k][j] = num
}

func Create(M, N, p int) {
	board := make([][]int, 0, M)
	for m := 0; m < M; m++ {
		array := make([]int, 0, N)
		for n := 0; n < N; n++ {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(M*N) > p {
				array = append(array, 9)
			} else {
				array = append(array, 0)
			}
		}
		board = append(board, array)
	}

	for k := 0; k < M; k++ {
		for j := 0; j < N; j++ {
			num := board[k][j]
			if num == 0 {
				Algo(M, N, k, j, board)
			}
		}

	}
	PrintBoard(board)
	//for
}

func main() {
	Create(10, 10, 80)
}
