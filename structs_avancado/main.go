package main

import (
	"encoding/json"
	"fmt"

	"github.com/andersonbco/cursodego/structs_avancado/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)

	fmt.Println("A casa em JSON: ", string(objJSON))
}
