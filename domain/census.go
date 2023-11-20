package domain

import (
	"errors"
	"taps/domain/command"
	"taps/domain/vo"
	"taps/utils/clk"
	"time"
)

type Census struct {
	ID              string
	CURP            vo.Curp
	ApplicationDate time.Time
	TargetGroup     TargetGroup
}

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
}

func BuildCensus(censusInput command.CreateCensus, clock clk.Clk) (Census, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return Census{}, err
	}

	if censusInput.TargetGroup.SixToFiftyNineMonthsOld == censusInput.TargetGroup.SixtyMonthsAndMore {
		return Census{}, errors.New("target group values cannot be the same")
	}

	fieldCensus := Census{
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
