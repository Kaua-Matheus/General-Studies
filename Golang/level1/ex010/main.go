package main

import (
	"fmt"
)

func fatorial(entrada int) int {
	if entrada == 0 || entrada == 1 {
		return 1
	} else {
		return entrada * fatorial(entrada - 1)
	}
}

func main() {
	var entrada int

	fmt.Println("Digite um número! ")
	fmt.Scan(&entrada)

	entrada = fatorial(entrada)
	fmt.Printf("Esse é o resultado final: %d", entrada)

}
