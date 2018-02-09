package matematica

/*
Calculo executa qualquer tipo de calculo baseado na funcao enviada
*/
func Calculo(function func(int, int) int, x int, y int) int {
	return function(x, y)
}
