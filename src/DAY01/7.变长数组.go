package main

import "fmt"

/**
 * 指针会自动释放
 */
func main() {

	// a) 定义-个可变长 数组
	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		fmt.Printf("i: %d, v:  %d\n", i, nums[i])
	}
	fmt.Printf("total : %d", len(nums))
	// b) 追加一个元素
	nums = append(nums, 100)
	for key, value := range nums {
		fmt.Printf("i: %d, v:  %d\n", key, value)
	}
	fmt.Printf("length : %d", len(nums))
	fmt.Printf("cap : %d", cap(nums))

	// c) 切片，类似于 python中的
	slice2 := nums[0:3]
	fmt.Println("slice2: ", slice2)
	//  c.1) 全切片
	slice3 := nums[:]
	fmt.Println("slice3：", slice3)

	// d) 浅copy 问题  ， 修改切片的数据，会直接影响 原始数据
	nums2 := nums[:3]
	nums2[0] = 100
	fmt.Println(nums2) // [100 2 3]
	fmt.Println(nums)  // [100 2 3 4 5 100]

	// e) make 数组 ,make=>     len = 20 capacity = 20
	strings := make([]string, 20)
	fmt.Println("str2:len:", len(strings), ",cap:", cap(strings))
	// strings[18] = "ad"
	fmt.Println(strings)

	// f) copy 深度copy
	depTarget := make([]int, len(nums))
	//copy(depTarget, nums)  // 方式1
	copy(depTarget, nums[:]) // 方式2
	depTarget[0] = 101
	// 此时 source [100（k） 2 3 4 5 100]
	fmt.Println(nums)
	// depTarget [101（k） 2 3 4 5 100]
	fmt.Println(depTarget)
}
