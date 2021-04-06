package main

import "fmt"

/**
 * 指针会自动释放
 */
func main() {
	// 定义1-个数组
	nums := [10]int{1, 2, 3, 4, 5}
	for i := range nums {
		fmt.Printf("i: %d, v:  %d\n", i, nums[i])
	}
	fmt.Printf("total : %d", len(nums))
}
