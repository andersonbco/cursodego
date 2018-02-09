package matematica

/*
Divisor executa uma divisao sem resto
*/
func Divisor(x int, y int) (resultado int) {
	resultado = x / y
	return
}

/*
DivisorComResto executa uma divisao com resto
*/
func DivisorComResto(x int, y int) (resultado int, resto int) {
	resultado = x / y
	resto = x % y
	return
}
