package influenzacensus

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
	"taps/domain/command"
	"taps/modules/influenzacensus/store"
	"taps/utils/clk"
)

type Taker struct {
	store store.CensusStore
	clock clk.Clk
}

func NewInfluenzaCensusTaker(store store.CensusStore, clock clk.Clk) *Taker {
	return &Taker{store: store, clock: clock}
}

func (t *Taker) Handle(fieldCensusParameters events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	censusInput := command.CreateCensus{}
	json.Unmarshal([]byte(fieldCensusParameters.Body), &censusInput)

	fieldCensus, err := domain.BuildCensus(censusInput, t.clock)
	if err != nil {
		fmt.Println(err)
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
