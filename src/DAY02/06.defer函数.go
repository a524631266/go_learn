package main

import (
	"fmt"
	"os"
)

/**
 * 1.  延迟，关键字，用于匿名函数 和 函数调用
 */
func main() {
	fileName := "/Users/didi/go/src/DAY02/01.switch.go"
	openFile(fileName)
}

func openFile(fileName string) {
	// 一般默认返回err为nil 则为正常，否则为异常
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error open file : ", fileName, " \n err:", err)
		return
	}
	defer func() {
		fmt.Println("close the file")
		_ = file.Close()
	}()

	defer fmt.Println("222222222")
	defer fmt.Println("111111111")
	defer fmt.Println("000000000")

	buf := make([]byte, 100)
	// 可以忽略
	read, _ := file.Read(buf)

	fmt.Println(" read len : ", read)

	fmt.Println("read message: ", string(buf)) // 类型转换
}
