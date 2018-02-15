package main

import (
	"fmt"
)

func main() {

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Printf("Qual a capacidade deste array?  %d\r\n", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lennon"}

	fmt.Printf("O que há neste array?  %v\r\n", cantores)

	capitais := [...]string{"Brasilia", "Lisboa", "Luanda", "Maputo"}

	for indice, capital := range capitais {
		fmt.Printf("A Capital[%d] é %s\r\n", indice, capital)
	}

	fmt.Printf("Qual a capacidade deste array?  %d\r\n", len(capitais))
}
