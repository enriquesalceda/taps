package influenzacensus

import (
	"encoding/json"
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

	if t.store.Find(fieldCensus.ID) {
		return events.APIGatewayProxyResponse{
			Body:       "census already exists",
			StatusCode: 409,
		}
	}

	err = t.store.Save(fieldCensus)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server error",
			StatusCode: 500,
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       "success",
		StatusCode: 200,
	}
}
