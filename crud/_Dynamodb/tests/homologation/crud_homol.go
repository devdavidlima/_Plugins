package main

import (
	"_Dynamodb/crud"
	"fmt"
	"utils/utils"
)

func main() {
	// Setting local test
	awsConfig := crud.AwsConfig{
		DBEndpoint: "http://localhost:8000",
		DBRegion:   "us-west-1",
	}

	userModel := crud.NewModel(awsConfig, "UserTeste", "ID")

	// Creating item
	user := map[string]interface{}{
		"ID":    "2",
		"Name":  "joao pedro",
		"Email": "jp@gmail.com",
	}
	err := userModel.CreateItem(user)
	utils.CheckErrAbortProgram(err, "Unable to create item in table")

	// Read item
	id := "2"

	user, err = userModel.ReadItem(id)
	utils.CheckErrAbortProgram(err, "Unable to read item in table")

	fmt.Println("Usuário encontrado:", user)

	// Update the item
	/*
		updatedData := map[string]interface{}{
			"Name": "Jane Doe",
			"Age":  35,
		}
		err = userModel.EditItem("item_id_1", updatedData)
		if err != nil {
			fmt.Println("Error updating item:", err)
			return
		}
	*/

	// Deleting item
	id = "2"
	err = userModel.DelItem(id)
	utils.CheckErrAbortProgram(err, "Unable to delete item in table")

	fmt.Println("Usuário deletado:", user)

	// Resto das operações CRUD...
}
