package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE - Operações básicas em uma interação com banco de dados

	router := mux.NewRouter()

	fmt.Println("Executando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
