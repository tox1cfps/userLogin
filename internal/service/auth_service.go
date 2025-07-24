package service

import (
	"fmt"
	"log"
	"userLogin/config"
	"userLogin/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erro ao gerar senha:", err)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, string(hashedPassword))
	if err != nil {
		log.Println("Erro ao criar usuário:", err)
		return
	}

	log.Println("Usuário criado com sucesso!")
}

func LoginUser() {
	var email, password string

	fmt.Print("\nDigite seu email: ")
	fmt.Scan(&email)

	fmt.Print("\nDigite sua senha: ")
	fmt.Scan(&password)

	var username, hashedPassword string

	err := config.DB.QueryRow("SELECT username, password_hash FROM users WHERE email=$1", email).Scan(&username, &hashedPassword)
	if err != nil {
		log.Println("Email não encontrado ou erro na consulta", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Senha incorreta!")
		return
	}

	fmt.Printf("Login bem-sucedido! Bem vindo de volta, %s.\n", username)

}

func UpdateUser() {
	repository.ListUsers()

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Erro ao atualizar a senha!", err)
		return
	}

	result, err := config.DB.Exec("UPDATE users SET username=$1, email=$2, password_hash=$3 WHERE id=$4", novoUsername, novoEmail, string(hashedPassword), id)
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
	repository.ListUsers()

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
