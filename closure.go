package main

import (
	"fmt"
)

func gorras(pedido int) (string, int, string) {

	precio := func() int {
		return pedido * 7
	}
	return "El precio del pedido es:", precio(), "pesos"
}

func main() {
	fmt.Println(gorras(44))
	fmt.Println(gorras(23))

}
