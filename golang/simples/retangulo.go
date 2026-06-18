package simples

import "fmt"

type Retangulo struct {
	largura float64
	altura  float64
}

func NewRetangulo(largura float64, altura float64) *Retangulo {
	if largura <= 0 || altura <= 0 {
		fmt.Println("Largura e altura devem ser maiores que zero.")
		return nil
	}

	return &Retangulo{
		largura: largura,
		altura:  altura,
	}
}

func (r *Retangulo) Area() float64 {
	return r.largura * r.altura
}

func (r *Retangulo) Perimetro() float64 {
	return 2 * (r.largura + r.altura)
}

func (r *Retangulo) Imprimir() {
	fmt.Println("Altura:", r.altura)
	fmt.Println("Largura:", r.largura)
	fmt.Println("Área:", r.Area())
	fmt.Println("Perímetro:", r.Perimetro())
}
