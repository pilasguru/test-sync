package main

import "fmt"

//func calcula(n1, n2 float32, op string) float32 {}

func main() {
	var n1 float32 = 10
	var n2 float32 = 8

	// Suma
	fmt.Print("La suma es: ")
	fmt.Println(n1 + n2)

	// Resta
	fmt.Print("La resta es: ")
	fmt.Println(n1 - n2)

	// Multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(n1 * n2)

	// División
	fmt.Print("La división es: ")
	fmt.Println(n1 / n2)

	// Resto
	//fmt.Print("El resto es: ")
	//fmt.Println(n1 % n2)

	/* var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Rodolfo"
	fmt.Println("hola mundo desde Go con " + nombre)
	fmt.Println(suma)
	fmt.Println(resta) */
}
