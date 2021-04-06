package main

import (
	"DAY01/import1/add"
	s "DAY01/import1/sub"
	"fmt"
)

func main() {
	sum := add.Addd(1, 5)
	fmt.Println("sum:", sum)
	sub := s.Sub(10, 8)
	fmt.Println("sub: ", sub)

	fmt.Println("utils: ", s.Utils(9, 2))
}
