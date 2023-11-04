package influenzacensus

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
	"taps/modules/store"
)

type InfluenzaCensusTaker struct {
	store store.CensusStore
}

func NewInfluenzaCensusTaker(store store.CensusStore) *InfluenzaCensusTaker {
	return &InfluenzaCensusTaker{store: store}
}

func (t *InfluenzaCensusTaker) Take(fieldCensusParameters events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
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

// -- store
