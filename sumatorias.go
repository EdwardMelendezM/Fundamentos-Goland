package main

import "fmt"

func main() {

	var numero int
	var sumatoria int

	fmt.Scanln(&numero)

	if numero == 0 {
		fmt.Println("Es un cero")
	} else {
		for i := 0; i <= numero; i++ {
			sumatoria += i
			if i%2 == 0 {
				continue
			}
			fmt.Println("Aqui entro", i)
		}
		fmt.Println("La suma es :", sumatoria)

	}

}
