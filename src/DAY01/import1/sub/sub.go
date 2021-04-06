package sub

import "fmt"

func Sub(a, b int) (sub int) {
	fmt.Println("add")
	sub = a - b
	return
}

func sub(a, b int) (sub int) {
	fmt.Println("add")
	sub = a - b
	return
}
