package main

import (
	"fmt"
)

func main() {
	s := "фываd"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
