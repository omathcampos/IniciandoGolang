package servidor

import (
	"crud/bancoDeDados"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Utilizar io.ReadAll pois ioutil.ReadAll foi descontinuado na versão 1.16
// Insere usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
	}

	db, erro := bancoDeDados.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT EVITA SQL INJECTION
	statement, erro := db.Prepare("INSERT INTO usuarios(nome, email) VALUES(?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obtener o inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated) //Pode-se utilizar o número do status code caso prefira
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %v", idInserido)))
}
