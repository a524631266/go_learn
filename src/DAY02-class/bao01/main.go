package main

import (
	"DAY02-class/bao01/src"
	"fmt"
)

func main() {
	person := src.Person{}
	person.Name = "abcd"
	fmt.Println(person)
	src.Hello()
}
