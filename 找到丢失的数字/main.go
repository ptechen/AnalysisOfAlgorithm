package main

import "fmt"

// 找到丢失的数字
// 有 n-1 个数字， 这些数字的范围是 [1, n] 且这 n-1 个 数字中没有重复的数字。
// 有上述条件可知：你丢失一个数字。请编写一段高效的找到该丢失数字的代码。
// 方法：排序 二分法
// 方法：排序 线性查找方式
// 方法：求和 [1, n]范围的和 减去 [1, n] 中 n-1 个数字
// 方法：计数排序
// 方法：XOR 异或
func main()  {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 19, 20,}
	for k := 1; k < len(array); k ++ {
		v := array[k - 1]
		r := k ^ v
		if r != 0 {
			fmt.Println(k)
			break
		}
	}

	a := 98764 << 5
	fmt.Println(a)
	b := 98764 * 32
	fmt.Println(b)
	if a == b {
		fmt.Println("a = b")
	}
}
