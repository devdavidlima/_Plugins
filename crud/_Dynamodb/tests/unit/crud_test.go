package crud_test

import (
	"_Dynamodb/crud"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	// Implementação do mock para PutItem
	return &dynamodb.PutItemOutput{}, nil
}

func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	// Implementação do mock para GetItem
	item := map[string]*dynamodb.AttributeValue{
		"ID":    {S: aws.String("1")},
		"Name":  {S: aws.String("John")},
		"Email": {S: aws.String("john@example.com")},
	}
	return &dynamodb.GetItemOutput{Item: item}, nil
}

func TestCreateItem(t *testing.T) {
	// Criando o modelo com o mock do DynamoDB
	model := &crud.Model{
		TableName:  "User",
		PrimaryKey: "ID",
		Svc:        &mockDynamoDBClient{},
	}

	// Chamando a função CreateItem
	err := model.CreateItem(map[string]interface{}{"ID": "1", "Name": "John", "Email": "john@example.com"})
	if err != nil {
		t.Errorf("Erro ao criar item: %v", err)
	}
}

func TestReadItem(t *testing.T) {
	// Criando o modelo com o mock do DynamoDB
	model := &crud.Model{
		TableName:  "User",
		PrimaryKey: "ID",
		svc:        &mockDynamoDBClient{},
	}

	// Chamando a função ReadItem
	item, err := model.ReadItem("1")
	if err != nil {
		t.Errorf("Erro ao ler item: %v", err)
	}

	expectedItem := map[string]interface{}{"ID": "1", "Name": "John", "Email": "john@example.com"}
	for key, value := range expectedItem {
		if item[key] != value {
			t.Errorf("Valor incorreto para a chave %s. Esperado: %v, Obtido: %v", key, value, item[key])
		}
	}
}

func TestNewModel(t *testing.T) {
	model := crud.NewModel("TableName_here", "PrimaryKey_here")

	if model.TableName != "TableName_here" {
		t.Errorf("Nome da tabela incorreto. Esperado: TableName_here, Obtido: %s", model.TableName)
	}

	if model.PrimaryKey != "PrimaryKey_here" {
		t.Errorf("Chave primária incorreta. Esperado: PrimaryKey_here, Obtido: %s", model.PrimaryKey)
	}

	// Verifica se svc não é nulo
	if model.Svc == nil {
		t.Error("O serviço DynamoDB não deve ser nulo")
	}
}
