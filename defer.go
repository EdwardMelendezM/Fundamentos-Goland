package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)
}
