package main

import "fmt"

/**
 * 指针会自动释放
 */
func main() {

	// 第一种方法
	name := "Zhang LL"
	ptr := &name
	// Zhang LL
	fmt.Println("指针具体数值", *ptr)
	// 0xc000010240
	fmt.Println("指针首地址", ptr)

	// 第二种方法
	name2 := new(string)
	*name2 = "ads"
	fmt.Println("name2:", *name2)

	// 可以返回局部 指针, 在c语言是不允许的，但是go语言由于gc的缘故，运行局部变量的返回
	city := showLocalCity()
	fmt.Println(*city)

	if city == nil {
		// 在c语言中空指针为null， c++ nullptr， go 语言是 nil
		fmt.Println("空指针")
	} else {
		fmt.Println("非空")
	}
}

func showLocalCity() *string {
	city := "杭州"
	ptr := &city
	return ptr
}
