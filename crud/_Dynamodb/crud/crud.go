package crud

import (
	"fmt"
	"reflect"
)

type Model struct {
	PartitionKey string
}

func (m *Model) CreateItem(data interface{}) error {
	// Simula a inserção no banco de dados
	fmt.Println("Inserindo novo registro:", data)
	return nil
}

func (m *Model) ReadItem(id interface{}) error {
	// Simula a leitura do banco de dados
	fmt.Println("Recuperando registro com chave primária:", id)
	return nil
}

func (m *Model) UpdateItem(data interface{}) error {
	// Simula a atualização no banco de dados
	fmt.Println("Atualizando registro:", data)
	return nil
}

func (m *Model) Delete(id interface{}) error {
	// Simula a exclusão no banco de dados
	fmt.Println("Excluindo registro com chave primária:", id)
	return nil
}

func NewModel(modelStruct interface{}, primaryKey string) *Model {
	// Obtém o tipo reflect.Type da struct
	modelType := reflect.TypeOf(modelStruct)

	// Verifica se a struct tem o campo especificado como chave primária
	if _, found := modelType.FieldByName(primaryKey); !found {
		panic("A struct não tem o campo especificado como chave primária.")
	}

	return &Model{
		PartitionKey: primaryKey,
	}
}
