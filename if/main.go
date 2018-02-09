package main

import (
	"fmt"
)

func main() {

	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele está adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente mora na cidade Teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situaçâo do cliente", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {

	if meses > 0 {
		descricao = "O cliente está devendo"
		status = true
		return
	}
	descricao = "O cliente está com o carnê em dia"
	return
}
