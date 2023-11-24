package repository

import (
	"taps/domain"
)

type CensusStore interface {
	Save(fieldCensus domain.Census) error
	Find(id string, date string) (domain.Census, bool, error)
}
