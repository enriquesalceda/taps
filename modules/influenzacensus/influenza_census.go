package influenzacensus

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
	"taps/domain/command"
	"taps/modules/influenzacensus/store"
)

type Taker struct {
	store store.CensusStore
}

func NewInfluenzaCensusTaker(store store.CensusStore) *Taker {
	return &Taker{store: store}
}

func (t *Taker) Handle(fieldCensusParameters events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	censusInput := command.CreateCensus{}
	json.Unmarshal([]byte(fieldCensusParameters.Body), &censusInput)

	fieldCensus, err := domain.CreateFieldCensus(censusInput)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}
	}

	found := t.store.Find(fieldCensus.ID)
	if found {
		return events.APIGatewayProxyResponse{Body: "census already exists", StatusCode: 409}
	}

	err = t.store.Save(fieldCensus)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "internal server error", StatusCode: 500}
	}

	return events.APIGatewayProxyResponse{Body: "success", StatusCode: 200}
}
