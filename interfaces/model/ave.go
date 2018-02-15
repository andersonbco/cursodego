package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo pata
type Pata interface {
	Quack() string
}

//Ave representa uma ave
type Ave struct {
	Nome string
}

//Cacareja representa o som emitido por uma Galinha
func (a Ave) Cacareja() string {
	return "Cocoricó!"
}

//Quack representa o som emitido por uma Pata
func (a Ave) Quack() string {
	return "Quack! Quack!"
}
