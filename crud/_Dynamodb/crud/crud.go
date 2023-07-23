package crud

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Model representa o modelo genérico
type Model struct {
	TableName  string // Nome da tabela no DynamoDB
	PrimaryKey string // Nome do campo usado como chave primária
	svc        *dynamodb.DynamoDB
}

// CreateItem insere um novo registro no banco de dados
func (m *Model) CreateItem(data interface{}) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(m.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", data.(map[string]interface{})[m.PrimaryKey]))},
		},
	}

	_, err := m.svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

// ReadItem recupera um registro do banco de dados com base na chave primária
func (m *Model) ReadItem(id interface{}) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(m.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", id))},
		},
	}

	result, err := m.svc.GetItem(input)
	if err != nil {
		return nil, err
	}

	item := make(map[string]interface{})
	for k, v := range result.Item {
		item[k] = v.String()
	}

	return item, nil
}

// Resto das funções do CRUD...

// Crie um novo modelo genérico com base em uma struct fornecida
func NewModel(tableName, primaryKey string) *Model {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-west-1"), // Escolha a região correta, mesmo que o banco esteja em execução local
	})
	if err != nil {
		panic(err)
	}

	return &Model{
		TableName:  tableName,
		PrimaryKey: primaryKey,
		svc:        dynamodb.New(sess),
	}
}
