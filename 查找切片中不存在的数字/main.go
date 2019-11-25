package main

import "fmt"

func find(array []int)  {
	for _, v := range array {
		if v >= 0 {
			array[v - 1] = -v
		} else {
			array[-v -1] = v
		}

	}
	fmt.Println(array)
}

func main() {
	a := []int{1, 2, 3, 5, 7, 2, 7, 8, 9}
	find(a)
}
