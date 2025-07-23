package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Conect() {
	connStr := "host=localhost port=5432 user=admin password=senha12345 dbname=usersdb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar com o banco", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao verificar conexão:", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
