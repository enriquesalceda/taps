package vo

import "errors"

type SeasonalInfluenzaVaccinationSchedule struct {
	FirstDose  bool
	SecondDose bool
	AnnualDose bool
}

func MustNewSeasonalInfluenzaVaccinationSchedule(firstDose, secondDose, annualDose bool) SeasonalInfluenzaVaccinationSchedule {
	schedule, err := TryNewSeasonalInfluenzaVaccinationSchedule(firstDose, secondDose, annualDose)
	if err != nil {
		panic(err)
	}

	return schedule
}

func TryNewSeasonalInfluenzaVaccinationSchedule(firstDose, secondDose, annualDose bool) (SeasonalInfluenzaVaccinationSchedule, error) {
	if dosesAreInOrder(firstDose, secondDose, annualDose) {
		return SeasonalInfluenzaVaccinationSchedule{FirstDose: firstDose, SecondDose: secondDose, AnnualDose: annualDose}, nil
	}

	return SeasonalInfluenzaVaccinationSchedule{}, errors.New("doses should be in order")
}

func dosesAreInOrder(firstDose, secondDose, annualDose bool) bool {
	return !((annualDose && (!firstDose || !secondDose)) || (secondDose && !firstDose))
}
