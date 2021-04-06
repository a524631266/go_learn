package initt

import "fmt"

func init() {
	fmt.Println(" init 02 method")
}

func init02() {
	// 不能直接调用， 没有init方法，这个方法类似于构造函数，包的构造函数
	//init()
}
