package simples

import "fmt"

type Pessoa struct {
	nome   string
	idade  int
	genero string
}

func NewPessoa(nome, genero string, idade int) *Pessoa {
	return &Pessoa{
		nome:   nome,
		idade:  idade,
		genero: genero,
	}
}

func (p *Pessoa) ImprimirInfos() {
	fmt.Println("Informações da pessoa:")
	fmt.Println("Nome:", p.nome)
	fmt.Println("Idade:", p.idade)
	fmt.Println("Gênero:", p.genero)
}

func (p *Pessoa) MaiorIdade() bool {
	return p.idade >= 18
}
