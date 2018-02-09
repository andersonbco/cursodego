package model

//Imovel struct que armazena os dados de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor seta o valor
func (imovel *Imovel) SetValor(valor int) {
	imovel.valor = valor
}

//GetValor coleta o valor
func (imovel *Imovel) GetValor() int {
	return imovel.valor
}
