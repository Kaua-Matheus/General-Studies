package main;
import (
	"fmt"
)

func calculator(operacao string, entrada1 float32, entrada2 float32) (float32) {
	switch operacao {
		case "+":
			return entrada1 + entrada2;
		case "-":
			return entrada1 - entrada2;
		case "*":
			return entrada1 * entrada2;
		case "/":
			if (entrada2 == 0) {
				fmt.Println("Não é possível realizar divisões por zero! ");
				return 0;
			} else {
				return entrada1 / entrada2;
			}
		default:
			fmt.Println("Você escolheu uma operação inválida! ");
			return 0;
		}
	}

func main() {
	var operacao string;
	var entrada1 float32;
	var entrada2 float32;


	fmt.Println("Escolha uma operação [+, -, *, /]: ")
	fmt.Scan(&operacao)

	fmt.Println("Digite o primeiro valor e o segundo valor: ")
	fmt.Scan(&entrada1, &entrada2)

	var resultado = calculator(operacao, entrada1, entrada2);
	fmt.Println("O valor final é ", resultado);

}
