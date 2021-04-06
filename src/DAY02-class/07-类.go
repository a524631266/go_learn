package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
	score  float32
}

// 类 p := 直接copy
func (p Person) Eat() {
	fmt.Println(p)
	// 不会修改原先的对象
	p.name = "Jianglina"
}

// 类 p := &对象 ，指针
func (p *Person) Eat2() {
	fmt.Println(p)
	// 因为是指针引用，所以会修改原先对象
	p.name = "Jianglina"
}
func main() {
	lily := Person{
		name:   "ZHANGLL",
		age:    30,
		gender: "female",
		score:  50,
	}
	lily.Eat()
	fmt.Println(lily)

	lily.Eat2()
	fmt.Println(lily)

}
