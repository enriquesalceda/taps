package store

import "taps/domain"

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
