package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("O número é ", i)
	}

	teste := true
	valor := 0

	for teste {
		valor++

		if valor%5 == 0 {
			teste = false
		}

		fmt.Println("O valor é ", valor)
	}

	for {
		valor--

		fmt.Println("O valor é ", valor)

		if valor == 0 {
			break
		}
	}

	texto := "Eu adoro escrever programas usando Go"

	for indice, letra := range texto {
		fmt.Printf("O texto é [%d] = %q\r\n", indice, letra)
	}
}
