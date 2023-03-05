package user

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	ErrorFaildToFatchRecord = "Failed to fetch record"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email)
			}
		}
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)

	if err != nil{
		return nil, errors.New(ErrorFaildToFatchRecord)
	}

}

func FetchUsers() {

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() error {

}
