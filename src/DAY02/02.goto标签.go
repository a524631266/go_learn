package main

import "fmt"

func main() {
	//	标签 label1
	//	goto label1
	//	break label1
	//	continue label1

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//	label1 中 的状态重新来过，i会从0开始重新计算
				//goto LABEL1
				//label1 会记住上次的状态， i从 0-4 开始
				//continue LABEL1
				break LABEL1 // 直接 over ，跳出 外循环
			}
			fmt.Println("i:", i, ", j:", j)

		}
	}
	fmt.Println("over!")
}
