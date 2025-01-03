package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConexaoSQL := "root@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexaoSQL)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão funcionando")

	linhas, erro := db.Query("Select * from usuarios")
	if erro != nil {
		log.Fatal(erro)

		fmt.Println(linhas)
	}

}
