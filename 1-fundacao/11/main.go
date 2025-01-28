package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // composicao
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
func main() {

	alano := Cliente{
		Nome:  "Alano",
		Idade: 34,
		Ativo: true,
	}

	//minhaEmpresa := Empresa{}

	Desativacao(alano)

	alano.Cidade = "Dois Vizinhos"
	alano.Desativar()
	fmt.Println(alano.Ativo)
}
