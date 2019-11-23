package main

import "fmt"

// 1742年， 哥德巴赫提出著名的哥德巴赫猜想。
// 即： 任一大于2的偶数都可以写成两个质数之和。
// 比如：16 = 6 + 13.


// 查找质数
func FindPrimeNumbers(evenNumber int) (primeNumbers []int, err error ){
	if evenNumber <= 2 || evenNumber % 2 != 0{
		return nil, fmt.Errorf("请输入大于2的偶数")
	}
	key := make([]int, 0, 8)
	for i := 2; i < 10; i ++ {
		key = append(key, i)
	}
	array := make([]int, 0, evenNumber)
	for i := 2; i <= evenNumber; i ++ {
		array = append(array, i)
	}
	for _, v := range key{
		for k, n := range array {
			if n == 0 {
				continue
			}
			if n % v == 0 && n != v{
				array[k] = 0
			}
		}
	}
	primeNumbers = make([]int, 0, 20)
	for _, v := range array {
		if v != 0 {
			primeNumbers = append(primeNumbers, v)
		}
	}
	return
}


// 证明猜想
func proveGuess(evenNumber int, primeNumbers []int)  {
	allLength := len(primeNumbers)
	m := 0
	n := allLength - 1
	for m < n {
		before := primeNumbers[m]
		after := primeNumbers[n]
		sum := before + after
		if sum > evenNumber {
			n --
		} else if sum < evenNumber {
			m ++
		} else {
			res := fmt.Sprintf("before: %d + after: %d = %d", before, after, evenNumber)
			fmt.Println(res)
			break
		}
	}
}

func main()  {
	eventNumber := 1000002
	nums ,err := FindPrimeNumbers(eventNumber)
	if err != nil {
		panic(err)
	}
	proveGuess(eventNumber, nums)
}
