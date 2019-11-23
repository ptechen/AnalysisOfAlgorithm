package main

import (
	"fmt"
)
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
func main()  {
	eventNumber := 100
	nums ,err := FindPrimeNumbers(eventNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
}
