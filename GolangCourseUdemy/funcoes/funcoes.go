package funcoes

import "fmt"

func Somar(n1, n2 int) int {
	return n1 + n2
}

func Importar() {
	soma := Somar(10, 30)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Print(txt)
	}
	f("texto da função 1\n")

	var E = func(txt string) string {
		return txt
	}
	resultado := E("\ntexto da função 2")
	fmt.Println(resultado)
}
