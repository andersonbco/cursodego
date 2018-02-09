package main

import (
	"fmt"

	"github.com/andersonbco/cursodego/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Anderson"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
}
