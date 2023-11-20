package store

import (
	"taps/domain"
)

type CensusStore interface {
	All() map[string]domain.Census
	Save(fieldCensus domain.Census) error
	Find(id string) bool
}
