package main

import (
	"encoding/json"
	"fmt"

	"github.com/andersonbco/cursodego/gobuild/model"
)

/*
GOOS=windows GOARCH=386 go build -o meuapp.exe
GOOS=linux GOARCH=arm go build -o meuappraspberry -v github.com/andersonbco/cursodego/gobuild
*/
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
