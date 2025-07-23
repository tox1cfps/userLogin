package ui

import (
	"fmt"
	"log"
	"userLogin/config"
)

func Menu() {
	for {
		var option int

		fmt.Println("\n--- Menu do usuário ---")
		fmt.Print("\n1. Criar usuário")
		fmt.Print("\n2. Login")
		fmt.Print("\n3. Mostrar usuários cadastrados")
		fmt.Print("\n4. Atualizar usuário")
		fmt.Print("\n5. Deletar usuário")
		fmt.Print("\n0. Sair")
		fmt.Print("\nSelecione uma opção: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			CreateUser()
		case 2:
			// LoginUser()
		case 3:
			ListUsers()
		case 4:
			UpdateUser()
		case 5:
			DeleteUser()
		case 0:
			fmt.Println("Encerrando programa...")
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}

func CreateUser() {
	var username string
	var email string
	var password string

	fmt.Print("\nDigite o nome do seu usuário: ")
	fmt.Scan(&username)

	fmt.Print("\nDigite o seu email: ")
	fmt.Scan(&email)

	fmt.Print("\nDigite a sua senha: ")
	fmt.Scan(&password)

	_, err := config.DB.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return
	}

	log.Println("Usuário criado com sucesso!")
}

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

func UpdateUser() {
	ListUsers()

	var id int
	fmt.Print("\nSelecione o ID do usuário que deseja atualizar: ")
	fmt.Scan(&id)

	var novoUsername string
	var novoEmail string
	var novaSenha string

	fmt.Print("\nDigite o novo nome do usuário: ")
	fmt.Scan(&novoUsername)

	fmt.Print("\nDigite o novo email do usuário: ")
	fmt.Scan(&novoEmail)

	fmt.Print("\nDigite a nova senha do usuário: ")
	fmt.Scan(&novaSenha)

	result, err := config.DB.Exec("UPDATE users SET username=$1, email=$2, password_hash=$3 WHERE id=$4", novoUsername, novoEmail, novaSenha, id)
	if err != nil {
		log.Println("Erro ao atualizar usuário:", err)
		return
	}
	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		log.Println("Usuário não encontrado")
		return
	}

	log.Println("Usuário atualizado com sucesso!")
}

func DeleteUser() {
	ListUsers()

	var id int
	fmt.Print("\nDigite o ID do usuário que deseja deletar: ")
	fmt.Scan(&id)

	result, err := config.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Erro ao deletar usuário:", err)
		return
	}

	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		log.Println("Usuário não encontrado")
		return
	}

	log.Println("Usuário deletado com sucesso!")
}
