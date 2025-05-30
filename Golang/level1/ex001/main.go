package main;
import "fmt";

var nome string = "Kauã";
var idade int = 20;

func saudacao(nome string, idade int) {
	fmt.Println("Olá " + nome + "!")	
	if idade <= 20 {
		fmt.Println("Tá novinho ainda..")
	} else if idade > 20 && idade < 40 {
		fmt.Println("Já está ficando velho em..")
	} else {
		fmt.Println("Tá velho já..")
	}
}

func main() {
	saudacao(nome, idade);

	fmt.Println("Fim da execução..")
}
