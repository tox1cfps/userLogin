package ui

import (
	"fmt"
	"userLogin/internal/repository"
	"userLogin/internal/service"
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
			service.CreateUser()
		case 2:
			service.LoginUser()
		case 3:
			repository.ListUsers()
		case 4:
			service.UpdateUser()
		case 5:
			service.DeleteUser()
		case 0:
			fmt.Println("Encerrando programa...")
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}
