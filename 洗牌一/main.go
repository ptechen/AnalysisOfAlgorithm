package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffleCorrect(cards []int)  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(cards); i ++ {
		index := rand.Intn(len(cards) - i) + i
		cards[i], cards[index] = cards[index], cards[i]
	}
}

func testShuffle()  {
	result := make([][]int, 0)

	for n := 0; n < 10; n ++ {
		array := make([]int, 0)
		for m :=0; m < 10; m++ {
			array = append(array, 0)
		}
		result = append(result, array)
	}

	for i := 0; i <100000; i ++ {
		A := make([]int, 0)
		for m := 0; m < 10; m++ {
			A = append(A, m)
		}
		shuffleCorrect(A)
		for j := 0; j < len(A); j ++ {
			result[A[j]][j] =  result[A[j]][j] + 1
		}
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func main()  {
	testShuffle()
	//array := []int{1, 2 ,3 ,4 ,5}
	//shuffleCorrect(array)
}
