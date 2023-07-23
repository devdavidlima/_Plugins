package crud

import (
	"fmt"
	"utils/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Generic model to represent a database data
type Model struct {
	TableName  string
	PrimaryKey string
	svc        *dynamodb.DynamoDB
}

// CreateItem: insert a new item respecting the PK
func (m *Model) CreateItem(data interface{}) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(m.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", data.(map[string]interface{})[m.PrimaryKey]))},
		},
	}

	_, err := m.svc.PutItem(input)
	utils.CheckErr(err, "")

	return nil
}

// ReadItem: Get item by PK
func (m *Model) ReadItem(id interface{}) (map[string]interface{}, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(m.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			m.PrimaryKey: {S: aws.String(fmt.Sprintf("%v", id))},
		},
	}

	result, err := m.svc.GetItem(input)
	utils.CheckErr(err, "Unable to get item")

	item := make(map[string]interface{})
	for k, v := range result.Item {
		item[k] = v.String()
	}

	return item, nil
}

// Resto das funções do CRUD...

// NewModel: Create db models using a struct. It`s like a "constructor of my interface Model"
func NewModel(tableName, primaryKey string) *Model {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-west-1"),
	})
	utils.CheckErrAbortProgram(err, "Unable to create a db model")

	return &Model{
		TableName:  tableName,
		PrimaryKey: primaryKey,
		svc:        dynamodb.New(sess),
	}
}
