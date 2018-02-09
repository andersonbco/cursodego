package model

import (
	"errors"
)

//Imovel struct que armazena os dados de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// ErrValorDeImovelInvalido erro apresentado quando o valor do imóvel é muito baixo
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

// ErrValorDeImovelMuitoAlto erro apresentado quando o valor do imóvel é muito alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//SetValor seta o valor
func (imovel *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	imovel.valor = valor
	return
}

//GetValor coleta o valor
func (imovel *Imovel) GetValor() int {
	return imovel.valor
}
