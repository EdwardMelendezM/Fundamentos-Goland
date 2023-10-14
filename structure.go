package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es ", p.nombre)
}

func main() {
	p := Persona{"Alex", 28, "alex@gmail.com"}
	p.diHola()
}
