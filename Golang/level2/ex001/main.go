package main

import (
	"fmt"
	_"reflect"
)

type Pessoa struct {
	Nome string;
	Numero string;
	Email string;
}

type pessoas []Pessoa;

func addPessoa(pessoa Pessoa, p *pessoas){
	*p = append(*p, pessoa)
}

func removePessoa(pessoa string, p *pessoas) {
	for _, pessoa_ := range *p {
		if pessoa_.Nome == pessoa {
			fmt.Println()
		}
	}
}

func verPessoas(p *pessoas) {
	for _, pessoa_ := range *p {
		fmt.Println(pessoa_.Nome)
	}
}

func main() {
	var lista pessoas;

	addPessoa(Pessoa{Nome: "João Cleber", Numero: "00000001", Email: "cleber@exemplo.com"}, &lista)
	addPessoa(Pessoa{Nome: "Matheus Henrique", Numero: "00000000", Email: "matheus@exemplo.com"}, &lista)
	addPessoa(Pessoa{Nome: "Cleber Arroz", Numero: "00000001", Email: "cleber@exemplo.com"}, &lista)
	addPessoa(Pessoa{Nome: "Henrique Gustavo", Numero: "00000000", Email: "matheus@exemplo.com"}, &lista)
	addPessoa(Pessoa{Nome: "Pedro Kaio", Numero: "00000001", Email: "cleber@exemplo.com"}, &lista)
	addPessoa(Pessoa{Nome: "Kaio de boca", Numero: "00000000", Email: "matheus@exemplo.com"}, &lista)

	fmt.Println(lista[2:])

	//verPessoas(&lista)
}