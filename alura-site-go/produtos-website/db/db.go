package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:sua-senha@/alura_site")
	if err != nil {
		panic(err.Error())
	}
	return db
}
