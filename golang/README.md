# Exercícios Diagnósticos de Programação Orientada a Objetos em Go

Este conjunto de exercícios tem como objetivo praticar conceitos de programação utilizando a linguagem **Go (Golang)**.

Embora Go não possua classes no mesmo formato de linguagens como Java ou C#, os exercícios devem ser resolvidos utilizando:

* `structs`
* métodos associados a `structs`
* construtores quando necessário
* encapsulamento por meio de atributos e métodos
* composição entre structs
* slices ou arrays
* manipulação de datas
* geração de números aleatórios
* entrada e saída no console

---

## Exercícios Simples

### 1. Struct `Circulo`

**Objetivo:**
Desenvolver uma struct `Circulo` com um atributo de raio.

**Tarefas:**

* Crie uma struct chamada `Circulo`.
* Adicione um atributo para armazenar o raio.
* Implemente um método para calcular a área do círculo.
* Implemente um método para calcular o perímetro do círculo.
* Crie um método para exibir as informações do círculo, incluindo:

    * raio
    * área
    * perímetro

---

### 2. Struct `ContaBancaria`

**Objetivo:**
Implementar uma struct `ContaBancaria` com informações básicas de uma conta.

**Atributos sugeridos:**

* número da conta
* nome do titular
* saldo

**Tarefas:**

* Crie uma struct chamada `ContaBancaria`.
* Implemente um método para realizar depósitos.
* Implemente um método para realizar saques.
* Implemente um método para exibir o saldo atual da conta.
* Garanta que o saque só seja realizado caso exista saldo suficiente.

---

### 3. Struct `Pessoa`

**Objetivo:**
Criar uma struct `Pessoa` contendo dados pessoais básicos.

**Atributos sugeridos:**

* nome
* idade
* gênero

**Tarefas:**

* Crie uma struct chamada `Pessoa`.
* Implemente um método para exibir as informações da pessoa.
* Implemente um método para verificar se a pessoa é maior de idade.
* Considere maior de idade a pessoa com idade maior ou igual a 18 anos.

---

### 4. Struct `Retangulo`

**Objetivo:**
Desenvolver uma struct `Retangulo` com largura e altura.

**Atributos sugeridos:**

* largura
* altura

**Tarefas:**

* Crie uma struct chamada `Retangulo`.
* Implemente um método para calcular a área do retângulo.
* Implemente um método para calcular o perímetro do retângulo.
* Crie um método para exibir as dimensões do retângulo.

---

## Exercícios Médios

### 5. Struct `Livro`

**Objetivo:**
Criar uma struct `Livro` para representar um livro de uma biblioteca.

**Atributos sugeridos:**

* título
* autor
* ano de publicação
* número de páginas
* disponibilidade

**Tarefas:**

* Crie uma struct chamada `Livro`.
* Implemente um método para emprestar o livro.
* Implemente um método para devolver o livro.
* Implemente um método para verificar se o livro está disponível para empréstimo.
* Exiba uma mensagem adequada quando o livro já estiver emprestado.

---

### 6. Struct `Lampada`

**Objetivo:**
Criar uma struct `Lampada` que possa ser ligada e desligada.

**Atributos sugeridos:**

* estado da lâmpada
* potência
* voltagem

**Tarefas:**

* Crie uma struct chamada `Lampada`.
* Implemente métodos para ligar e desligar a lâmpada.
* Implemente um método para verificar se a lâmpada está ligada ou desligada.
* Permita ler e alterar os valores de potência e voltagem.
* Crie um programa no `main.go` para testar a lâmpada.
* No teste, exiba no console:

    * se a lâmpada está ligada ou desligada
    * potência
    * voltagem
* Ligue a lâmpada e exiba novamente as informações.

---

### 7. Lâmpada com chance de queimar

**Objetivo:**
Modificar a struct `Lampada` para incluir a possibilidade de queimar ao ser ligada.

**Tarefas:**

* Adicione um atributo para indicar se a lâmpada está queimada.
* Ao tentar ligar a lâmpada, implemente uma chance de 15% dela queimar.
* Utilize a biblioteca padrão do Go para gerar números aleatórios.
* Caso a lâmpada esteja queimada, ela não poderá ser ligada.
* Exiba uma mensagem informando quando a lâmpada queimar.

---

## Exercícios Difíceis

### 8. Struct `Conta`

**Objetivo:**
Criar uma struct `Conta` para representar uma conta bancária mais completa.

**Atributos sugeridos:**

* nome do cliente
* número da conta
* saldo

**Tarefas:**

* Os valores da conta devem ser informados no momento da criação.
* Crie uma função construtora para inicializar uma conta.
* Implemente um método `Depositar`.
* Implemente um método `Sacar`.
* Implemente um método `ObterSaldo`.
* Implemente um método `ObterNumero`.
* Implemente um método `ObterNomeCliente`.

---

### 9. Programa de teste para contas bancárias

**Objetivo:**
Testar a struct `Conta` criada no exercício anterior.

**Tarefas:**

* Crie três contas bancárias diferentes.
* Realize várias operações de depósito e saque em cada conta.
* Ao final, exiba um relatório no console contendo:

    * número da conta
    * nome do titular
    * saldo atual

---

### 10. Struct `Extrato`

**Objetivo:**
Criar uma struct para representar movimentações bancárias.

**Atributos sugeridos:**

* data da movimentação
* valor movimentado

**Tarefas:**

* Crie uma struct chamada `Extrato`.
* A data deve representar o momento em que a movimentação ocorreu.
* O valor movimentado pode ser:

    * positivo, para depósitos
    * negativo, para saques

---

### 11. Conta com extrato bancário

**Objetivo:**
Incrementar a struct `Conta` para armazenar o histórico de movimentações.

**Tarefas:**

* Adicione à struct `Conta` um vetor ou slice de `Extrato`.
* Cada vez que ocorrer um depósito, registre uma movimentação no extrato.
* Cada vez que ocorrer um saque, registre uma movimentação no extrato.
* O extrato deve armazenar a data da movimentação e o valor movimentado.
* Crie um programa de teste com três contas.
* Realize várias operações de saque e depósito.
* Ao final, exiba um relatório contendo:

    * número da conta
    * titular
    * saldo atual
    * extrato completo de movimentações

---

### 12. Struct `Cartao`

**Objetivo:**
Criar uma struct para representar um cartão bancário.

**Atributos sugeridos:**

* número do cartão
* validade do cartão

**Tarefas:**

* Crie uma struct chamada `Cartao`.
* Adicione à struct `Conta` uma informação de cartão.
* Associe cada conta a um cartão.

---

### 13. Saque com cartão

**Objetivo:**
Permitir a realização de saques usando os dados do cartão.

**Tarefas:**

* Crie três contas diferentes, cada uma com seu respectivo cartão.
* Realize operações de depósito e saque.
* Implemente uma operação de saque com cartão.
* Para sacar com cartão, o programa deverá solicitar:

    * número do cartão
    * validade do cartão
    * valor do saque
* O programa deve identificar a conta à qual o cartão pertence.
* Caso o cartão seja válido, realize o saque.
* Caso o cartão não seja encontrado ou os dados estejam incorretos, exiba uma mensagem de erro.
* Ao final, emita um relatório no console contendo:

    * número da conta
    * titular
    * saldo atual
    * extrato de cada conta

