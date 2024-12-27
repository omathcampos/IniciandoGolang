package ifElse

import "fmt"

func EstruturaDeControle() {
	numero := 15

	if numero > 10 {
		fmt.Println(numero, " é maior que 10")
	} else {
		fmt.Println(numero, "é menor que 10")
	}

	if numero > 100 {
		fmt.Println(numero, " é maior que 100")
	}

	//IF INIT

	if outroNumero := numero; outroNumero > 10 {
		fmt.Println(outroNumero, " Teste 2 é maior que 10")
	} else {
		fmt.Println(numero, " é menor que 10")
	}
}
