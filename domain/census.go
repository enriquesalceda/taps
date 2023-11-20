package domain

import (
	"errors"
	"taps/domain/command"
	"taps/domain/vo"
	"taps/utils/clk"
	"time"
)

type FieldCensus struct {
	ID              string
	CURP            vo.Curp
	ApplicationDate time.Time
	TargetGroup     TargetGroup
}

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
}

func BuildCensus(censusInput command.CreateCensus, clock clk.Clk) (FieldCensus, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return FieldCensus{}, err
	}

	if censusInput.TargetGroup.SixToFiftyNineMonthsOld == censusInput.TargetGroup.SixtyMonthsAndMore {
		return FieldCensus{}, errors.New("target group values cannot be the same")
	}

	fieldCensus := FieldCensus{
		ID:              curp.ID,
		CURP:            curp,
		ApplicationDate: clock.Now(),
		TargetGroup: TargetGroup{
			SixToFiftyNineMonthsOld: censusInput.TargetGroup.SixToFiftyNineMonthsOld,
			SixtyMonthsAndMore:      censusInput.TargetGroup.SixtyMonthsAndMore,
		},
	}

	return fieldCensus, nil
}
