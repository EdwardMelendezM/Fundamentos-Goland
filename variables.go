package main

import "fmt"

func main() {
	// Declaracion de variables
	// Con var y el tipo de dato
	var nombre string = "Juan pipa"
	var edad, memoria int = 22, 88
	edad = 23
	memoria = 55
	nombre = "Amelio"

	fmt.Println(edad)
	fmt.Println(memoria)
	fmt.Println(nombre)

	//Declaracion de variables de varios tipos
	var (
		pi       float64
		booleano bool
		cadena   = "text 01"
		edads    = 25
	)
	pi = 3.141592
	booleano = true
	fmt.Println(pi, booleano, cadena, edads)

	// Definiar variables sin tipo de datos
	variable1 := 24
	variable2 := "Texto 02"

	fmt.Println(variable1, variable2)

	//Constantes
	const s string = "Constante"
	fmt.Println(s)

}
