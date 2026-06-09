package medio

import (
	"fmt"
	"math/rand"
)

type Lampada struct {
	isLigada   bool
	potencia   float64
	voltagem   float64
	isQueimada bool
}

func NewLampada(isLigada, queimada bool, potencia, voltagem float64) *Lampada {
	return &Lampada{
		isLigada:   isLigada,
		potencia:   potencia,
		voltagem:   voltagem,
		isQueimada: queimada,
	}
}

func (l *Lampada) Ligar() {
	if l.isQueimada {
		fmt.Println("A lâmpada está queimada e não pode ser ligada.")
		return
	}

	if l.isLigada {
		fmt.Println("A lâmpada já está ligada.")
		return
	}

	if rand.Intn(100) < 15 {
		l.isQueimada = true
		l.isLigada = false
		fmt.Println("A lâmpada queimou ao tentar ligar.")
		return
	}

	l.isLigada = true
	fmt.Println("Lâmpada ligada.")
}

func (l *Lampada) Desligar() {
	if l.isQueimada {
		fmt.Println("A lâmpada está queimada.")
		return
	}

	if !l.isLigada {
		fmt.Println("A lâmpada já está desligada.")
		return
	}

	l.isLigada = false
	fmt.Println("Lâmpada desligada.")
}

func (l *Lampada) EditVoltagem(valor float64) {
	if valor <= 0 {
		fmt.Println("Voltagem inválida.")
		return
	}

	l.voltagem = valor
}

func (l *Lampada) EditPotencia(valor float64) {
	if valor <= 0 {
		fmt.Println("Potência inválida.")
		return
	}

	l.potencia = valor
}

func (l *Lampada) Display() {
	if l.isLigada {
		fmt.Println("Estado: ligada")
	} else {
		fmt.Println("Estado: desligada")
	}

	fmt.Println("Queimada:", l.isQueimada)
	fmt.Println("Voltagem:", l.voltagem)
	fmt.Println("Potência:", l.potencia)
}
