package simples

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func NewCirculo(raio float64) *Circulo {
	return &Circulo{raio: raio}
}

func (c *Circulo) Area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c *Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (c *Circulo) ImprimirInfo() {
	fmt.Println("Raio:", c.raio)
	fmt.Println("Área:", c.Area())
	fmt.Println("Perímetro:", c.Perimetro())
}
