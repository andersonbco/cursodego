package main

import (
	"fmt"
)

func main() {

	//fmt.Println(nums, len(nums), cap(nums))
	capitais := []string{"Lisboa"}
	//fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	//fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Toquio"
	cidades[4] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))
	capitais[1] = "Brasilia"
	fmt.Println(capitais, len(capitais), cap(capitais))

	for indice, cidades := range cidades {
		fmt.Printf("Cidade[%d] - %s\r\n", indice, cidades)
	}

	//Primeiro item começa com o índice 0
	//Segundo item começa com o índice 1
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia, len(capitaisAsia), cap(capitaisAsia))

	//Caso o primeiro item não seja informado, será considerado desde o início do slice
	temp1 := cidades[:2]
	fmt.Println(temp1)

	//Caso o segundo item não seja informado, será considerado até o final do slice
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)

	//Para remover um item de um slice, é necessário criar um slice temporário sem o item a ser excluído
	//e depois copiar o conteúdo dele para o slice original
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	fmt.Println(temp)
	fmt.Println(cidades[indiceDoItemARetirar+1:])
	//"..." são usados para informar ao append vários valores
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	fmt.Println(temp)
	copy(cidades, temp)
	fmt.Println(cidades)
}
