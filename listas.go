package main

import "fmt"

func main() {

	var array1 [2]string
	array1[0] = "Dato1"
	array1[1] = "Dato2"

	array2 := [3]int{1, 2, 3}
	fmt.Println(array2)

	fmt.Println(array1)

	for i := 0; i < 3; i++ {
		fmt.Println(array1[i])
	}
}
