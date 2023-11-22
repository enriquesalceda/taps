package domain

import (
	"taps/domain/command"
	"taps/domain/vo"
	"taps/utils/clk"
	"time"
)

type Census struct {
	ID                                   string
	CURP                                 vo.Curp
	Address                              vo.Address
	TargetGroup                          vo.TargetGroup
	SeasonalInfluenzaVaccinationSchedule vo.SeasonalInfluenzaVaccinationSchedule
	Rights                               vo.Right
	RiskGroup                            vo.RiskGroup
	ApplicationDate                      time.Time
}

func BuildCensus(censusInput command.CreateCensus, clock clk.Clock) (Census, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return Census{}, err
	}

	seasonalInfluenzaVaccinationSchedule, err := vo.TryNewSeasonalInfluenzaVaccinationSchedule(
		censusInput.SeasonalInfluenzaVaccinationSchedule.FirstDose,
		censusInput.SeasonalInfluenzaVaccinationSchedule.SecondDose,
		censusInput.SeasonalInfluenzaVaccinationSchedule.AnnualDose,
	)
	if err != nil {
		return Census{}, err
	}

	targetGroup, err := vo.TryNewTargetGroup(
		censusInput.TargetGroup.SixToFiftyNineMonthsOld,
		censusInput.TargetGroup.SixtyMonthsAndMore,
	)
	if err != nil {
		return Census{}, err
	}

	address, err := vo.TryNewAddress(
		censusInput.Address.StreetNumber,
		censusInput.Address.StreetName,
		censusInput.Address.SuburbName,
	)
	if err != nil {
		return Census{}, err
	}

	right, err := vo.TryNewRights(censusInput.Rights)
	if err != nil {
		return Census{}, err
	}

	fieldCensus := Census{
		ID:              curp.ID,
		CURP:            curp,
		ApplicationDate: clock.Now(),
		TargetGroup:     targetGroup,
		Address:         address,
		RiskGroup: vo.NewRiskGroup(
			censusInput.RiskGroup.PregnantWomen,
			censusInput.RiskGroup.WellnessPerson,
			censusInput.RiskGroup.AIDS,
			censusInput.RiskGroup.Diabetes,
			censusInput.RiskGroup.Obesity,
			censusInput.RiskGroup.AcuteOrChronicHeartDisease,
			censusInput.RiskGroup.ChronicLungDiseaseIncludesCOPDAndAsthma,
			censusInput.RiskGroup.Cancer,
			censusInput.RiskGroup.ChronicConditionsThatRequireProlongedConsumptionOfSalicylic,
			censusInput.RiskGroup.RenalInsufficiency,
			censusInput.RiskGroup.AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS,
			censusInput.RiskGroup.EssentialHypertension,
		),
		SeasonalInfluenzaVaccinationSchedule: seasonalInfluenzaVaccinationSchedule,
		Rights:                               right,
	}

	return fieldCensus, nil
}
