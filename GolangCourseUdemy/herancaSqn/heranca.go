package herancaSqn

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func Heranca() {
	p1 := pessoa{"math", "Campos", 19, 186}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "Unisanta"}
	fmt.Println(e1)

}
