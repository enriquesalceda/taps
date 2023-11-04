package influenzacensus

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
	"taps/modules/store"
)

type Taker struct {
	store store.CensusStore
}

func NewInfluenzaCensusTaker(store store.CensusStore) *Taker {
	return &Taker{store: store}
}

func (t *Taker) Take(fieldCensusParameters events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	fieldCensus := domain.FieldCensus{}
	json.Unmarshal([]byte(fieldCensusParameters.Body), &fieldCensus)

	err := fieldCensus.Validate()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}
	}

	err = t.store.Save(fieldCensus)
	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{Body: "success", StatusCode: 200}
}
