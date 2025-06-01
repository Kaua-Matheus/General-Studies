package main
import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var palavra string;

	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&palavra)
	palavra = strings.ToLower(palavra)

	palavraSeparada := strings.Split(palavra, "");
	palavraAuxiliar := []string{};

	for i := len(palavraSeparada) - 1; i >= 0; i-=1 {
		palavraAuxiliar = append(palavraAuxiliar, palavraSeparada[i])
	}

	fmt.Println(palavraAuxiliar)
	fmt.Println(palavraSeparada)


	for j := 0; j != len(palavraSeparada); j++ {
		if (palavraAuxiliar[0] == palavraSeparada[0]) {
			continue;
		} else {
			fmt.Println("Não são iguais! ")
			os.Exit(0);
		}
	}
	fmt.Println("São iguais! ")
}