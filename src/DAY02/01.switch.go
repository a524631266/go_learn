package main

import (
	"fmt"
	"os"
)

// c语言 有： argc， **argv
func main() {
	// go语言 os.args
	cmds := os.Args
	for key, value := range cmds {
		fmt.Println("key:", key, ", value:", value, "\n")
	}
	if len(cmds) < 2 {
		fmt.Println("请正确输入参数")
		return
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		// 默认 会加break
		//break
		// go语言 中如果想向下穿透，可以使用关键字 fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("result")
	}
}
