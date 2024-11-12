package operadores

import "fmt"

func OperadoresAritmeticos() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//Atribuição
	var varivael1 string = "texto"
	variavel2 := "texto2"
	fmt.Println(varivael1, variavel2)

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	//Operadores Lógicos
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // AND
	fmt.Println(verdadeiro || falso) // OR
	fmt.Println(!verdadeiro)         // Negação

	//Operadores Unários
	fmt.Println("--------------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 2
	fmt.Println(numero)

	numero /= 2
	fmt.Println(numero)

	numero %= 2
	fmt.Println(numero)

	if numero == 10 {
		fmt.Println("Falso")
	} else {
		fmt.Println("Verdadeiro")
	}
}
