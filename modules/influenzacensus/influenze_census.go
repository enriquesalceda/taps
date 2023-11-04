package influenzacensus

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"taps/domain"
)

type InfluenzaCensusTaker struct {
	store CensusStore
}

type invalidCensusResponse struct {
	errors string
}

func NewInfluenzaCensusTaker(store CensusStore) *InfluenzaCensusTaker {
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
type InMemoryInfluenzaStore struct {
	all map[string]InfluenzaCensus
}

type InfluenzaCensus struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

type CensusStore interface {
	All() []InfluenzaCensus
	Save(fieldCensus domain.FieldCensus) error
}

func NewInMemoryInfluenzaStore() *InMemoryInfluenzaStore {
	return &InMemoryInfluenzaStore{
		all: map[string]InfluenzaCensus{},
	}
}

func (i *InMemoryInfluenzaStore) All() []InfluenzaCensus {
	var allCensus []InfluenzaCensus
	for _, census := range i.all {
		allCensus = append(allCensus, census)
	}
	return allCensus
}

func (i *InMemoryInfluenzaStore) Save(fieldCensus domain.FieldCensus) error {
	influenzaCensus := InfluenzaCensus{
		ID:            fieldCensus.ID,
		FirstLastName: fieldCensus.FirstLastName,
		LastLastName:  fieldCensus.LastLastName,
		FirstName:     fieldCensus.FirstName,
		DOB:           fieldCensus.DOB,
		State:         fieldCensus.State,
		Gender:        fieldCensus.Gender,
		Number:        fieldCensus.Number,
	}

	i.all[fieldCensus.ID] = influenzaCensus
	return nil
}
