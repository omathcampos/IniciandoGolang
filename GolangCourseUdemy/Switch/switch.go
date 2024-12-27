package main

import "fmt"

//func diaDaSemana(numero int) string {
//	switch numero {
//	case 1:
//		return "Domingo"
//	case 2:
//		return "Segunda-feira"
//	case 3:
//		return "Terça-Feira"
//	case 4:
//		return "Quarta-Feira"
//	case 5:
//		return "Quinta-Feira"
//	case 6:
//		return "Sexta-Feira"
//	case 7:
//		return "Sábado"
//	default:
//		return "Número inválido"
//	}
//}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		return "Número inválido"
	}
	return diaDaSemana
}

func dia() {
	dia := diaDaSemana2(2)
	fmt.Println(dia)
}

func main() {
	dia()
}
