package main;
import (
	"fmt"
)

func imparPar(entrada int) {
	if (entrada % 2 == 0) {
		fmt.Println("É par! ")
	} else {
		fmt.Println("É impar! ")
	}
}

func main() {
	var entrada int;
	fmt.Println("Digite um número inteiro: ");
	fmt.Scanln(&entrada);
	imparPar(entrada);
}