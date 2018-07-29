package main

import "fmt"

type Notebook struct {
	tmk   string
	size  int
	metal bool
}

func main() {

	var MiMac = Notebook{"Macbook", 17, true}
	fmt.Println(MiMac.tmk)
	/*
		var suma int = 8 + 9
		var resta int = 6 - 4
		var nombre string = "Rodolfo"
		fmt.Println("hola mundo desde Go con " + nombre)
		fmt.Println(suma)
		fmt.Println(resta)
	*/
}
