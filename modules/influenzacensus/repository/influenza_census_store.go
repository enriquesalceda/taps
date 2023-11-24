package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"taps/domain"
	"taps/utils/clk"
)

type DynamoInfluenzaCensusRepository struct {
	dynamoDB *dynamodb.DynamoDB
	table    string
}

type InfluezaCensusItem struct {
	CurpID    string
	TimeStamp string
	Census    domain.Census
}

func NewDynamoInfluenzaCensusRepository(
	dynamoDB *dynamodb.DynamoDB,
	table string,
) *DynamoInfluenzaCensusRepository {
	return &DynamoInfluenzaCensusRepository{dynamoDB: dynamoDB, table: table}
}

func (d *DynamoInfluenzaCensusRepository) Save(fieldCensus domain.Census, clock clk.Clock) error {
	census, err := dynamodbattribute.MarshalMap(
		InfluezaCensusItem{
			CurpID:    fieldCensus.ID,
			TimeStamp: fmt.Sprint(clock.Now().Unix()),
			Census:    fieldCensus,
		},
	)

	if err != nil {
		return err
	}

	_, err = d.dynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: &d.table,
		Item:      census,
	})

	return nil
}

func (d DynamoInfluenzaCensusRepository) All() map[string]domain.Census {
	//TODO implement me
	panic("implement me")
}

func (d DynamoInfluenzaCensusRepository) Find(id string) bool {
	//TODO implement me
	panic("implement me")
}
