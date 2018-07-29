package main

import "fmt"

func operacion(n1 float32, n2 float32, op string) float32 {
	var res float32
	if op == "+" {
		res = n1 + n2
	}
	if op == "-" {
		res = n1 - n2
	}
	if op == "*" {
		res = n1 * n2
	}
	if op == "/" {
		res = n1 / n2
	}
	return res
}

func calculadora(num1 float32, num2 float32) {
	// Suma
	fmt.Print("La suma es: ")
	fmt.Println(operacion(num1, num2, "+"))

	// Resta
	fmt.Print("La resta es: ")
	fmt.Println(operacion(num1, num2, "-"))

	// Multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(operacion(num1, num2, "*"))

	// División
	fmt.Print("La división es: ")
	fmt.Println(operacion(num1, num2, "/"))

	// Resto
	fmt.Print("El resto es: ")
	fmt.Println(operacion(num1, num2, "%"))
}

func main() {
	fmt.Println("Calculo 1")
	var num1 float32 = 20
	var num2 float32 = 15
	calculadora(num1, num2)

	fmt.Println("------")

	fmt.Println("Calculo 2")
	var num3 float32 = 88
	var num4 float32 = 45
	calculadora(num3, num4)
}
