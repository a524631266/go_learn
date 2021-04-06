package main

import "fmt"

/**
 * 字典 hash表
 *
 */
func main() {
	//  定义一个字典
	var idNames map[int]string
	// 分配空间
	idNames = make(map[int]string, 10)
	fmt.Println(idNames)
	fmt.Println(idNames == nil)
	idNames[0] = "zhangll"
	idNames[1] = "jiangln"
	fmt.Println("删除前：", idNames)
	fmt.Println("map length: ", len(idNames))

	// 遍历map
	for key, value := range idNames {
		fmt.Println("key: ", key, ", value", value)
	}
	// 是否存在一个key
	// 在map 中， 所有key都是存在的有效的，所以不会出现异常或奔溃，返回nil（零值）
	// bool -》 false ，int -》 0 ， string -》 ""
	num100 := idNames[100]
	fmt.Println("num100:", num100 == "")

	// 常用方法
	idNames1 := make(map[int]string, 10)
	fmt.Println(idNames1)

	// 无法通过返回值是否不为零值来判断，map中的key是否存在
	value, ok := idNames[10]
	if ok {
		fmt.Println("id = 1 存在 value: ", value)
	} else {
		fmt.Println("id = 10不 存在 value: ", value)
	}
	//	 delete
	delete(idNames, 1)
	fmt.Println("删除后：", idNames)
}
