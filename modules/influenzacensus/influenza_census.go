package influenzacensus

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
	"taps/domain/command"
	"taps/modules/influenzacensus/repository"
	"taps/utils/clk"
)

type Taker struct {
	store repository.CensusStore
	clock clk.Clock
}

func NewInfluenzaCensusTaker(store repository.CensusStore, clock clk.Clock) *Taker {
	return &Taker{store: store, clock: clock}
}

func (t *Taker) Handle(fieldCensusParameters events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	censusInput := command.CreateCensus{}
	json.Unmarshal([]byte(fieldCensusParameters.Body), &censusInput)

	fieldCensus, err := domain.BuildCensus(censusInput, t.clock)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}
	}

	census, found, err := t.store.Find(fieldCensus.ID, t.clock.Now().Format("2006-01-02"))

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server error at finding",
			StatusCode: 500,
		}
	}

	if found {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("census with id %s already exists", census.ID),
			StatusCode: 409,
		}
	}

	err = t.store.Save(fieldCensus)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server error at saving",
			StatusCode: 500,
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       "success",
		StatusCode: 200,
	}
}
