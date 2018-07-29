package main

import (
	"fmt"
)

func pantalon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}

func main() {
	pantalon("largo", "jean", "dos bolsillos")
}
