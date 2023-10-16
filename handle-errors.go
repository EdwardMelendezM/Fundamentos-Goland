package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "65496"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("Numero", num)
}
