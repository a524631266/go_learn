package main

import "fmt"

/**
 * 函数方式1
 */
func test2(a int, b int, c string) (int, bool) {

	return a + b, c == ""
}

/**
 * 函数方式2
 * 缩写参数， 以及添加默认 返回值
 */
func test3(a, b int, c string) (res int, bl bool) {
	res = a + b
	bl = c == ""
	return
}

/**
 * fun 函数
 */
func main() {
	i, b := test2(1, 2, "")
	fmt.Println("i: ", i, ", b:", b)
	res, bl := test3(3, 4, "")
	fmt.Println("res: ", res, ", bl:", bl)
}
