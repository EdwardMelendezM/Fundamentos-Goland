package main

import "fmt"

func main() {

	saludar()
	var num1 int
	var num2 int

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	fmt.Println("La suma es :", sumar(num1, num2))

}

func saludar() {
	fmt.Println("Hola a todo gaa")
}
func sumar(a int, b int) int {

	return a + b
}
