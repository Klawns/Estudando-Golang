package simples

import (
	"fmt"
)

type ContaBancaria struct {
	nomeConta   string
	nomeTitular string
	saldo       float64
}

func NewContaBancaria(saldo float64, nomeConta, nomeTitular string) *ContaBancaria {
	return &ContaBancaria{
		nomeConta:   nomeConta,
		nomeTitular: nomeTitular,
		saldo:       saldo,
	}
}

func (c *ContaBancaria) Depositar(valor float64) {
	if valor <= 0 {
		return
	}
	c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) {
	if valor <= 0 {
		return
	}

	if valor <= c.saldo {
		c.saldo -= valor
	} else {
		return
	}

}

func (c *ContaBancaria) ImprimirSaldo() {
	fmt.Println("Imprimir Saldo Atual:", c.saldo)
}

func (c *ContaBancaria) ImprimirInfos() {
	fmt.Println("Imprimir Infos:", c.nomeConta)
	fmt.Println("Imprimir Titular:", c.nomeTitular)
	fmt.Println("Imprimir Saldo:", c.saldo)
}
