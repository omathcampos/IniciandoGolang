package maps

import "fmt"

func Maps() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Matheus",
		"sobrenome": "Campos",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":    "Matheus",
			"último nome": "Campos",
		},
		"faculdade": {
			"Curso":  "ADS",
			"Campus": "Unisanta",
		},
	}

	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["faculdade"])
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
