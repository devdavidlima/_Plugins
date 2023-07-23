package main

import (
	"_Dynamodb/crud"
	"fmt"
	"utils/utils"
)

func main() {
	userModel := crud.NewModel("UserTeste", "ID")

	// Criar um novo usuário
	user := map[string]interface{}{
		"ID":    "2",
		"Name":  "joao pedro",
		"Email": "jp@gmail.com",
	}
	err := userModel.CreateItem(user)
	utils.CheckErrAbortProgram(err, "Unable to create item in table")

	// Ler um usuário pelo ID
	id := "2"

	user, err = userModel.ReadItem(id)
	utils.CheckErrAbortProgram(err, "Unable to read item in table")

	fmt.Println("Usuário encontrado:", user)

	// Resto das operações CRUD...
}
