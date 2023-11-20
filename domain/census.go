package domain

import (
	"taps/domain/command"
	"taps/domain/vo"
	"taps/utils/clk"
	"time"
)

type FieldCensus struct {
	ID              string
	CURP            vo.Curp
	ApplicationDate time.Time
}

func BuildCensus(censusInput command.CreateCensus, clock clk.Clk) (FieldCensus, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return FieldCensus{}, err
	}

	fieldCensus := FieldCensus{
		ID:              curp.ID,
		CURP:            curp,
		ApplicationDate: clock.Now(),
	}

	return fieldCensus, nil
}
