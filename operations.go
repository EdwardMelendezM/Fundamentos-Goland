package main

import "fmt"

func main() {

	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	suma := a + b
	mul := a * b
	div := a / b
	rest := a % b
	res := a - b
	fmt.Println("La suma es :", suma)
	fmt.Println("La mult es :", mul)
	fmt.Println("La divi es :", div)
	fmt.Println("El resto es :", rest)
	fmt.Println("La resta es :", res)

}
