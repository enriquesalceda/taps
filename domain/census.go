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
	RiskGroup                            RiskGroup
	ApplicationDate                      time.Time
}

type RiskGroup struct {
	PregnantWomen                                                                                        bool
	WellnessPerson                                                                                       bool
	AIDS                                                                                                 bool
	Diabetes                                                                                             bool
	Obesity                                                                                              bool
	AcuteOrChronicHeartDisease                                                                           bool
	ChronicLungDiseaseIncludesCOPDAndAsthma                                                              bool
	Cancer                                                                                               bool
	CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic bool
	RenalInsufficiency                                                                                   bool
	AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS                                           bool
	EssentialHypertension                                                                                bool
}

func BuildCensus(censusInput command.CreateCensus, clock clk.Clock) (Census, error) {
	curp, err := vo.TryParseCURP(censusInput.CURP)
	if err != nil {
		return Census{}, err
	}

	seasonalInfluenzaVaccinationSchedule, err := vo.TryNewSeasonalInfluenzaVaccinationSchedule(censusInput.SeasonalInfluenzaVaccinationSchedule.FirstDose, censusInput.SeasonalInfluenzaVaccinationSchedule.SecondDose, censusInput.SeasonalInfluenzaVaccinationSchedule.AnnualDose)
	if err != nil {
		return Census{}, err
	}

	targetGroup, err := vo.TryNewTargetGroup(censusInput.TargetGroup.SixToFiftyNineMonthsOld, censusInput.TargetGroup.SixtyMonthsAndMore)
	if err != nil {
		return Census{}, err
	}

	address, err := vo.TryNewAddress(censusInput.Address.StreetNumber, censusInput.Address.StreetName, censusInput.Address.SuburbName)
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
		RiskGroup: RiskGroup{
			PregnantWomen:                           censusInput.RiskGroup.PregnantWomen,
			WellnessPerson:                          censusInput.RiskGroup.WellnessPerson,
			AIDS:                                    censusInput.RiskGroup.AIDS,
			Diabetes:                                censusInput.RiskGroup.Diabetes,
			Obesity:                                 censusInput.RiskGroup.Obesity,
			AcuteOrChronicHeartDisease:              censusInput.RiskGroup.AcuteOrChronicHeartDisease,
			ChronicLungDiseaseIncludesCOPDAndAsthma: censusInput.RiskGroup.ChronicLungDiseaseIncludesCOPDAndAsthma,
			Cancer:                                  censusInput.RiskGroup.Cancer,
			CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic: censusInput.RiskGroup.CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic,
			RenalInsufficiency: censusInput.RiskGroup.RenalInsufficiency,
			AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: censusInput.RiskGroup.AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS,
			EssentialHypertension: censusInput.RiskGroup.EssentialHypertension,
		},
		SeasonalInfluenzaVaccinationSchedule: seasonalInfluenzaVaccinationSchedule,
		Rights:                               right,
	}

	return fieldCensus, nil
}
