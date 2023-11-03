package influenzacensus

type FieldCensus struct {
	ID            string
	FirstLastName string
	LastLastName  string
	FirstName     string
	DOB           string
	State         string
	Gender        string
	Number        int
}

func NewFieldCensus(
	ID string,
	firstLastName string,
	lastLastName string,
	firstName string,
	DOB string,
	state string,
	gender string,
	number int,
) *FieldCensus {
	return &FieldCensus{
		ID:            ID,
		FirstLastName: firstLastName,
		LastLastName:  lastLastName,
		FirstName:     firstName,
		DOB:           DOB,
		State:         state,
		Gender:        gender,
		Number:        number,
	}
}

type InfluenzaCensusTaker struct {
	store CensusStore
}

func NewInfluenzaCensusTaker(store CensusStore) *InfluenzaCensusTaker {
	return &InfluenzaCensusTaker{store: store}
}

func (t *InfluenzaCensusTaker) Take(
	ID string,
	firstLastName string,
	lastLastName string,
	firstName string,
	DOB string,
	state string,
	gender string,
	number int,
) error {
	fieldCensus := &FieldCensus{
		ID:            ID,
		FirstLastName: firstLastName,
		LastLastName:  lastLastName,
		FirstName:     firstName,
		DOB:           DOB,
		State:         state,
		Gender:        gender,
		Number:        number,
	}

	return t.store.Save(fieldCensus)
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
	Save(fieldCensus *FieldCensus) error
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

func (i *InMemoryInfluenzaStore) Save(fieldCensus *FieldCensus) error {
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
