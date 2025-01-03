package bancoDeDados

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver de conexão com MySQL
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexaoSQL := "root@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexaoSQL)
	if erro != nil {
		return nil, erro
	}
	// defer db.Close()

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
