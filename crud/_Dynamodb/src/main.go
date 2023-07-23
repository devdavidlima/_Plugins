package main

import (
	"_Dynamodb/crud"
	"fmt"
)

func main() {
	userModel := crud.NewModel("UserTeste", "ID")

	// Criar um novo usuário
	user := map[string]interface{}{
		"ID":    "1",
		"Name":  "Carlos",
		"Email": "carlinhos@gmail.com",
	}
	err := userModel.CreateItem(user)
	if err != nil {
		panic(err)
	}

	// Ler um usuário pelo ID
	id := "1"
	user, err = userModel.ReadItem(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Usuário encontrado:", user)

	// Resto das operações CRUD...
}
