package main

import "fmt"

// go语言通过首字母的大小写 控制权限
// 如果想要在不同包中，调用数据

type Person2 struct {
	name   string
	age    int
	gender string
	score  float32
}

// 类 p := 直接copy
func (p Person2) Eat() {
	fmt.Println(p)
	// 不会修改原先的对象
	p.name = "Jianglina"
}

// 类 p := &对象 ，指针
func (p *Person2) Eat2() {
	fmt.Println(p)
	// 因为是指针引用，所以会修改原先对象
	p.name = "Jianglina"
}

type Son struct {
	Person2 // 这就是go的继承 ，没有变量名
	school  string
}

func main() {
	son := Son{
		Person2: Person2{},
		school:  "添加",
	}
	son.name = "zz"
	son.age = 10
	son.score = 1.2
	// 如下操作
	son.Person2.gender = "A" //
	son.gender = "男"

	fmt.Println(son)
}
