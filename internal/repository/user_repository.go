package repository

import (
	"fmt"
	"log"
	"userLogin/config"
)

func ListUsers() {
	rows, err := config.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Println("Erro ao ler usuários:", err)
		return
	}

	defer rows.Close()

	fmt.Println("#### Lista de usuários ####")

	for rows.Next() {
		var id int
		var username string
		var email string

		err := rows.Scan(&id, &username, &email)
		if err != nil {
			log.Println("Erro ao listar usuários:", err)
			continue
		}
		fmt.Printf("\nID: %d | Username: %s | Email: %s |", id, username, email)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Erro ao iterar pelos usuários:", err)
	}
}
