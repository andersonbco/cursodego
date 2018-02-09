package main

import (
	"fmt"
)

//Imovel struct que armazena os dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("A apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Nome:  "Minha chacara",
		X:     7,
		Y:     8,
		valor: 20000,
	}
	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Minha casa"
	casa.X = 5
	casa.Y = 7
	casa.valor = 45000
	fmt.Printf("A casa é: %+v\r\n", casa)
}
