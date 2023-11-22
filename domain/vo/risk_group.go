package vo

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

func NewRiskGroup(
	pregnantWomen bool,
	wellnessPerson bool,
	AIDS bool,
	diabetes bool,
	obesity bool,
	acuteOrChronicHeartDisease bool,
	chronicLungDiseaseIncludesCOPDAndAsthma bool,
	cancer bool,
	chronicConditionsThatRequireProlongedConsumptionOfSalicylic bool,
	renalInsufficiency bool,
	acquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS bool,
	essentialHypertension bool,
) RiskGroup {
	return RiskGroup{
		PregnantWomen:                           pregnantWomen,
		WellnessPerson:                          wellnessPerson,
		AIDS:                                    AIDS,
		Diabetes:                                diabetes,
		Obesity:                                 obesity,
		AcuteOrChronicHeartDisease:              acuteOrChronicHeartDisease,
		ChronicLungDiseaseIncludesCOPDAndAsthma: chronicLungDiseaseIncludesCOPDAndAsthma,
		Cancer:                                  cancer,
		ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: chronicConditionsThatRequireProlongedConsumptionOfSalicylic,
		RenalInsufficiency: renalInsufficiency,
		AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: acquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS,
		EssentialHypertension: essentialHypertension,
	}
}
