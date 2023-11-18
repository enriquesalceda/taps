package store

import (
	"taps/domain"
)

type CensusStore interface {
	All() map[string]domain.FieldCensus
	Save(fieldCensus domain.FieldCensus) error
	Find(id string) bool
}
