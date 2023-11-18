package domain

import (
	"taps/domain/command"
	"taps/domain/vo"
)

type FieldCensus struct {
	ID   string
	CURP vo.Curp
}

func CreateFieldCensus(censusInput command.CreateCensus) (FieldCensus, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return FieldCensus{}, err
	}

	fieldCensus := FieldCensus{
		ID:   curp.ID,
		CURP: curp,
	}

	return fieldCensus, nil
}
