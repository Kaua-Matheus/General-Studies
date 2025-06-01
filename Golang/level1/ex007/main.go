package main;
import (
	"fmt"
)

func multi(entrada int) {
	for i := 0; i <= 10; i ++ {
		fmt.Println(entrada, " X ", i, " = ", entrada * i);
	}
}

func main() {
	var entrada int;
	fmt.Println("Digite um número: ");
	fmt.Scan(&entrada);

	multi(entrada);
}