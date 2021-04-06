package main

import "fmt"

/**
 * 指针会自动释放
 */
func main() {
	message := `as我`
	fmt.Println(message)
	fmt.Println("len : ", len(message))

	for i := 0; i < len(message); i++ {
		fmt.Printf("index: %d, string: %c \n", i, message[i])
	}
}
