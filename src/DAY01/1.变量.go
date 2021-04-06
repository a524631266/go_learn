package main

import "fmt"

func main() {
	//test2.sayHello()
	// 变量定义： var
	//	 常量定义： const

	// 定义 之后再赋值
	var name string
	name = "duck"

	var age int
	age = 30
	fmt.Println("name :", name)
	fmt.Printf("name:   aa%s, age: %d", name, age)

	// 定义赋值
	var gender string = "男"
	//	定义直接赋值 并推到
	address := "北京"
	fmt.Printf("gender: %s, addres: %s", gender, address)

	fmt.Println("#######")
	a := 10
	b := "abc"
	test(a, b)
}

func test(a int, b string) {
	fmt.Printf("a : %d, b:%s", a, b)
}
