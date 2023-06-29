package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaNoBancoDeDados() *sql.DB {
	conexao := "user=vcosta dbname=alura_loja password=postgreSQL152022 host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
