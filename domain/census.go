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
	RiskGroup       RiskGroup
}

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
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
	}

	return fieldCensus, nil
}
