package structs

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func Structs() {
	fmt.Println("structs")

	var u usuario
	u.nome = "Matheus"
	u.idade = 19
	u.endereco = endereco{"Rua Proost", 18}
	fmt.Println(u)

	usuario2 := usuario{"Matheus Campos", 19, endereco{"Rua Proost", 18}}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Matheus Dias Campos", idade: 21}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Math"}
	fmt.Println(usuario4)

	enderecoExemplo := endereco{"Rua Proost", 18}
	fmt.Println(usuario3, enderecoExemplo)
}
