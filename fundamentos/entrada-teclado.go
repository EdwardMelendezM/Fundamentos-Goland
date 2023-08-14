package main

import "fmt"

func main() {
	fmt.Println("Entrada por el teclado")
	var (
		nombre string
		edad   int
		pi     float64
	)

	//Entrada por teclado
	fmt.Print("Ingrese su nombre")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su edad")
	fmt.Scanln(&edad)

	fmt.Print("Ingrese pi")
	fmt.Scanln(&pi)

	//Salida
	fmt.Printf("Nombre: %s , Edad: %d \nValor de pi: %f", nombre, edad, pi)

}
