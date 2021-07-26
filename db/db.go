package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//conectando com banco de dados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=sergio_loja password=Sergim25. host=localhost sslmode=disable" // as informações que tem q ser passadas
	db, err := sql.Open("postgres", conexao)                                                        //abrindo uma conexao com o banco de dados

	//verificando se tem algum erro
	if err != nil {
		panic(err.Error())
	}
	return db
}
