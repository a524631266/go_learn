package main

import "fmt"

// go语言中没有枚举的类型
// 使用 const + iota 模拟
func main() {
	//var number int
	//var name string
	//var flag bool

	// 变量组的形式定义，比上方优雅一些
	var (
		number int
		name   string
		flag   bool
	)
	number = 10
	name = "zhangll"
	flag = false
	fmt.Println("number: ", number, "; name :", name, "; flag :", flag)

	// 枚举 类型，类似于 变量组定义方式
	// iota 从0开始
	const (
		MONDAY    = iota // 0
		TUESDAY   = iota // 1
		WEDNESDAY = iota
		THURSDAY  = iota
		FRIDAY    = iota
		SATURDAY  = iota
		SUNDAY    = iota
		M, N      = iota, iota // const 为编译器赋值，所有不需要m，n := 进行自动 推导 ， 默认在同一行，为 7，7
	)
	// iota 的作用域 为一个圆括号中
	const (
		MONDAY2    = iota // 0
		TUESDAY2   = iota // 1
		WEDNESDAY2 = iota
		THURSDAY2  = iota
		FRIDAY2
		SATURDAY2
		SUNDAY2
	)
	fmt.Println("MONDAY: ", MONDAY)
	fmt.Println("FRIDAY2: ", FRIDAY2)
	fmt.Println("MONDAY2: ", MONDAY2)

	// 可以通过iota 自动加 1 来从 1开始计数
	const (
		JANU = iota + 1
		FER
		MAR
	)
	fmt.Println("FER", FER)
}
