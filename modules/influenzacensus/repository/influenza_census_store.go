package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (d *DynamoInfluenzaCensusRepository) FindByDate(dt string) ([]domain.Census, error) {
	var items []InfluenzaCensusItem

	condition := expression.Key("Date").Equal(expression.Value(dt))
	expr, err := expression.NewBuilder().WithKeyCondition(condition).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return []domain.Census{}, err
	}

	queryParams := &dynamodb.QueryInput{
		TableName:                 aws.String(d.table),
		IndexName:                 aws.String("DateIndex"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	queryOutput, err := d.dynamoDB.Query(queryParams)
	if err != nil {
		return []domain.Census{}, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, &items)
	if err != nil {
		return []domain.Census{}, err
	}

	var censuses []domain.Census
	for _, item := range items {
		censuses = append(censuses, item.Census)
	}

	return censuses, nil
}
