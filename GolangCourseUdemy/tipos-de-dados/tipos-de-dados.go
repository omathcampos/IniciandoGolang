package tipos_de_dados

import (
	"errors"
	"fmt"
)

func TiposDados() {
	fmt.Println("Tipos de Dados")

	var numeroInt8 int8 = 48
	fmt.Println(numeroInt8)
	var numeroInt16 int16 = 4888
	fmt.Println(numeroInt16)
	var numeroInt32 int32 = -488888888
	fmt.Println(numeroInt32)
	var numeroInt64 int64 = 4888888888888888888
	fmt.Println(numeroInt64)

	var numeroUint uint = 1000
	fmt.Println(numeroUint)

	//alias
	// Rune = int32
	var numeroAlias rune = 123456
	fmt.Println(numeroAlias)

	//byte = uint8
	var numeroByteAlias byte = 123
	fmt.Println(numeroByteAlias)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234554687.45
	fmt.Println(numeroReal2)

	var numeroReal3 = 1234554687.45
	fmt.Println(numeroReal3)

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	var booleano bool = false
	fmt.Println(booleano)

	var erro error = errors.New("Erro criado")
	fmt.Println(erro)

}
