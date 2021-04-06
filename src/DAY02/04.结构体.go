package main

import "fmt"

// c 语言，我们可以使用 typedef int MyInt，重新定义一个类型
// 使用结构体，来模拟 类
type MyInt int

// 定义了一个结构体
type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 10, 10
	fmt.Println("i:", i, ",j:", j)

	liLy := Student{
		name:   "LiLy",
		age:    0,
		gender: "female",
		score:  80,
	}
	fmt.Println("liLy:", liLy.name, liLy.age, liLy.gender, liLy.score)

	// 在go语言中， 没有对指针的->操作，没有
	s1 := &liLy

	liLy.age = 90
	fmt.Println("s1:", s1.name, s1.age, s1.gender, s1.score)
	ZhangLL := Student{
		"zhangll",
		30,
		"male",
		100,
	}
	fmt.Println(ZhangLL)
}
