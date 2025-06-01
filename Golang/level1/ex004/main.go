package main

import (
	"fmt"
)

func convCelsius(entrada float32) (float32){
	return (entrada - 32) * 5/9;
}

func convFahrenheit(entrada float32) (float32) {
	return (entrada * 9/5) + 32;
}

func main(){
	var opcao string;
	var entrada float32;

	fmt.Println("Olá, selecione qual conversão quer fazer! [0] (F -> C) - [1] (C -> F)");
	fmt.Scanln(&opcao);

	switch (opcao) {
		case "0":
			fmt.Println("Digita a temperatura em Fahrenheit: ");
			fmt.Scanln(&entrada);
			entrada = convCelsius(entrada);
			fmt.Println("Temperatura em Celcius: ", entrada);
			break;

		case "1":
			fmt.Println("Digita a temperatura em Celcius: ");
			fmt.Scanln(&entrada);
			entrada = convFahrenheit(entrada);
			fmt.Println("Temperatura em Fahrenheit: ", entrada);
			break;

		default:
			fmt.Println("Você digitou uma opção inválida..")
			break;
	}
}