package medio

import "fmt"

type Livro struct {
	titulo          string
	autor           string
	publicacao      int
	paginas         int
	disponibilidade bool
}

func NewLivro(titulo, autor string, publicacao, paginas int, disponibilidade bool) *Livro {
	return &Livro{
		titulo:          titulo,
		autor:           autor,
		publicacao:      publicacao,
		paginas:         paginas,
		disponibilidade: disponibilidade,
	}
}

func (l *Livro) Emprestar() {
	if !l.disponibilidade {
		return
	}

	l.disponibilidade = false
}

func (l *Livro) Devolver() {
	if l.disponibilidade {
		return
	}

	l.disponibilidade = true
}

func (l *Livro) Disponivel() bool {
	return l.disponibilidade
}

func (l *Livro) ImprimirInfos() {
	fmt.Println("Título:", l.titulo)
	fmt.Println("Autor:", l.autor)
	fmt.Println("Ano de publicação:", l.publicacao)
	fmt.Println("Páginas:", l.paginas)

	if l.disponibilidade {
		fmt.Println("Disponibilidade: disponível")
	} else {
		fmt.Println("Disponibilidade: emprestado")
	}
}
