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
	router.HandleFunc("/cadastrar-usuario", servidor.CriarUsuario).Methods(http.MethodPost) //poderia utilizar "POST" somente em string dentro do () .Methods
	router.HandleFunc("/buscar-usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/buscar-usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/alterar-usuario/{id}", servidor.AtualizoUsuario).Methods(http.MethodPut)
	router.HandleFunc("/deletar-usuario/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Executando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
