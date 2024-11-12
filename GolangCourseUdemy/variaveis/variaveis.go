package variaveis

import "fmt"

func Variaveis() {
	//String Explicita
	var variavel1 string = "var teste 1"
	//String Implicita - Inferencia de tipo
	variavel2 := "var teste 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração explicita de mais de uma varíavel
	var (
		variavel3 string = "var teste 3"
		variavel4 string = "var teste 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var teste 5", "var teste 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
