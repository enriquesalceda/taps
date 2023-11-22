package command

type CreateCensus struct {
	CURP                                 string
	Address                              Address
	TargetGroup                          TargetGroup
	RiskGroup                            RiskGroup
	OtherRiskGroup                       bool
	SeasonalInfluenzaVaccinationSchedule SeasonalInfluenzaVaccinationSchedule
	BatchNumber                          string
	Rights                               string
}

type Address struct {
	StreetNumber string
	StreetName   string
	SuburbName   string
}

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
}

type RiskGroup struct {
	PregnantWomen                                               bool
	WellnessPerson                                              bool
	AIDS                                                        bool
	Diabetes                                                    bool
	Obesity                                                     bool
	AcuteOrChronicHeartDisease                                  bool
	ChronicLungDiseaseIncludesCOPDAndAsthma                     bool
	Cancer                                                      bool
	ChronicConditionsThatRequireProlongedConsumptionOfSalicylic bool
	RenalInsufficiency                                          bool
	AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS  bool
	EssentialHypertension                                       bool
}

type SeasonalInfluenzaVaccinationSchedule struct {
	FirstDose  bool
	SecondDose bool
	AnnualDose bool
}
