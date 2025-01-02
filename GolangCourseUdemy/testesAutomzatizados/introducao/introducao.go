package main

import (
	"GolangCourseUdemy/testesAutomzatizados/enderecos"
	"fmt"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndereco)
}
