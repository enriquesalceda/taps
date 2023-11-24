package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"taps/domain"
	"taps/utils/clk"
)

type DynamoInfluenzaCensusRepository struct {
	dynamoDB *dynamodb.DynamoDB
	table    string
	clock    clk.Clock
}

type InfluenzaCensusItem struct {
	CurpID string
	Date   string
	Census domain.Census
}

func NewDynamoInfluenzaCensusRepository(
	dynamoDB *dynamodb.DynamoDB,
	table string,
	clock clk.Clock,
) *DynamoInfluenzaCensusRepository {
	return &DynamoInfluenzaCensusRepository{dynamoDB: dynamoDB, table: table, clock: clock}
}

func (d *DynamoInfluenzaCensusRepository) Save(fieldCensus domain.Census) error {
	census, err := dynamodbattribute.MarshalMap(
		InfluenzaCensusItem{
			CurpID: fieldCensus.ID,
			Date:   d.clock.Now().Format("2006-01-02"),
			Census: fieldCensus,
		},
	)

	if err != nil {
		return err
	}

	_, err = d.dynamoDB.PutItem(&dynamodb.PutItemInput{
		TableName: &d.table,
		Item:      census,
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoInfluenzaCensusRepository) Find(id string, date string) (domain.Census, bool, error) {
	item, err := d.dynamoDB.GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String(d.table),
			Key: map[string]*dynamodb.AttributeValue{
				"CurpID": {S: aws.String(id)},
				"Date":   {S: aws.String(date)},
			},
		},
	)

	census := domain.Census{}

	if err != nil {
		return census, false, err
	}

	if item.Item == nil {
		return census, false, nil
	}

	var influenzaCensusItem InfluenzaCensusItem
	err = dynamodbattribute.UnmarshalMap(item.Item, &influenzaCensusItem)
	if err != nil {
		return census, false, err
	}

	return influenzaCensusItem.Census, true, nil
}
