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

	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacaraaaa", 280000}
	fmt.Printf("Casa é: %p - %+v\r\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Casa é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 5
	imovel.Y = imovel.Y + 6
}
