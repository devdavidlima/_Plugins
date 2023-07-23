package main

import "_Dynamodb/crud"

func main() {
	userModel := crud.NewModel(User{}, "ID")

	user := User{ID: 1, Name: "Carlos", Email: "carlinhos@gmail.com"}
	userModel.CreateItem(user)
	userModel.ReadItem(1)
	// userModel.UpdateItem(user)
	// userModel.DeleteItem(1)
}

type User struct {
	ID    int
	Name  string
	Email string
}
