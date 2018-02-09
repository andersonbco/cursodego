package main

import (
	"fmt"
)

var (
	//Endereco é uma variável pública
	Endereco string
	//Telefone é uma variável pública
	Telefone            string
	quantidade, estoque int
	comprou             bool
	valor               float64
	palavras            rune
)

func main() {
	teste := "variavel teste"
	fmt.Printf("Endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("Variável teste: %s\r\n", teste)
}
