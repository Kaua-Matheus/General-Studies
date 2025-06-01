package main;
import (
	"fmt"
)

func fatorial(numero int) (int) {
	if numero == 0 {
		return 1;
	} else {
		valor := numero;

		for i := 1; i != numero; i++ {
			valor = valor * i;
		}

		return valor;
	}
}

func main() {
	var entrada int;

	fmt.Println("Digite um número: ")
	fmt.Scan(&entrada)

	fmt.Println(fatorial(entrada))
}