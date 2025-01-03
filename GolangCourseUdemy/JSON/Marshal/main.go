package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Shelbi", "Golden", 3}
	fmt.Println(c)

	cahorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(string(cahorroEmJson))
	fmt.Println(bytes.NewBuffer(cahorroEmJson))

	c2 := map[string]string{
		"nome": "Jackson",
		"raca": "Pastor Alemão",
	}

	fmt.Println(c2)

	cahorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(string(cahorro2EmJson))
	fmt.Println(bytes.NewBuffer(cahorro2EmJson))
	fmt.Println(cahorro2EmJson)
}
