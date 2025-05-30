package main;
import (
	"os"
	"strconv"
	"fmt"
	"path/filepath"
)

// Cria novos arquivos .go
func main() {

	var nomeArquivo string = "main.go";

	for i := 4; i < 10; i++ {
		pasta := ("../ex00" + strconv.Itoa(i))
		err := os.MkdirAll(pasta, os.ModePerm);
		if err != nil {
			panic(err);
		}

		arquivoFinal := filepath.Join(pasta, nomeArquivo)
		arquivo, err := os.Create(arquivoFinal)
		if err != nil {
			panic(err)
		}

		defer arquivo.Close()
	}

	fmt.Println("Arquivos criados com sucesso! ")
}
