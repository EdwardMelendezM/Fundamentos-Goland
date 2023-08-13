package main

import "fmt"

func main() {

	// No tienen un tamaño estatico
	var slicen1 []string

	slicen1 = append(slicen1, "dato01", "dato02", "3")

	fmt.Println(slicen1)

}
