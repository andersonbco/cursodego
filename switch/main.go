package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3

	switch numero {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	fmt.Println("Você é da família do Unix?")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Não!!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Você pode dormir até mais tarde!")
	default:
		fmt.Println("Hoje é dia de trabalho!!")
	}

	numero = 20

	fmt.Println("Este número cabe em um dígito?")

	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10, numero < 100:
		fmt.Println("Serve dois dígitos?")
	case numero > 100:
		fmt.Println("Não dá, o número é muito grande")
	}
}
