package main

import (
	"crud/servidor"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE - Operações básicas em uma interação com banco de dados

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) //poderia utilizar "POST" somente em string dentro do () .Methods
	fmt.Println("Executando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
