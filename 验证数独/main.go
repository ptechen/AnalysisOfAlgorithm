package main

import "fmt"

func PrintBoard(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func sudoku() (array [][]int) {
	array = make([][]int, 0, 9)

	array = append(array, []int{5, 3, 4, 6, 7, 8, 9, 1, 2})
	array = append(array, []int{6, 7, 2, 1, 9, 5, 3, 4, 8})
	array = append(array, []int{1, 9, 8, 3, 4, 2, 5, 6, 7})
	array = append(array, []int{8, 5, 9, 7, 6, 1, 4, 2, 3})
	array = append(array, []int{4, 2, 6, 8, 5, 3, 7, 9, 1})
	array = append(array, []int{7, 1, 3, 9, 2, 4, 8, 5, 6})
	array = append(array, []int{9, 6, 1, 5, 3, 7, 2, 8, 4})
	array = append(array, []int{2, 8, 7, 4, 1, 9, 6, 3, 5})
	array = append(array, []int{3, 4, 5, 2, 8, 6, 1, 7, 9})
	PrintBoard(array)
	return array
}

func validationSudoku(array [][]int) {
	length := len(array)
	a := ""
	fortag:
	for i := 0; i < length; i ++ {
		row := make([]int, 9, 9)
		col := make([]int, 9, 9)
		blk := make([]int, 9, 9)
		for j := 0; j < length; j ++ {
			if row[array[i][j] - 1] == 0 {
				row[array[i][j] - 1] = array[i][j]
			} else {
				a = fmt.Sprintf("row: %d col: %d :fail", i, j)
				break fortag
			}

			if col[array[j][i] - 1] == 0 {
				col[array[j][i] - 1] = array[j][i]
			} else {
				a = fmt.Sprintf("row: %d col: %d :fail", i, j)
				break fortag
			}

			// check block
			idx_row := (i / 3) * 3 + j / 3
			idx_col := (i % 3) * 3 + j % 3
			if blk[array[idx_row][idx_col] -1] == 0{
				blk[array[idx_row][idx_col] -1] = array[idx_row][idx_col]
			} else {
				a = fmt.Sprintf("row: %d col: %d :fail", idx_row, idx_col)
				break fortag
			}


		}
	}
	fmt.Println(a)
}

func main() {
	array := sudoku()
	validationSudoku(array)
}
