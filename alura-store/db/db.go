package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaNoBancoDeDados() *sql.DB {
	conexao := "host=localhost user=root password=root dbname=alura_loja port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
