package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestRiskGroup(t *testing.T) {
	t.Run("should create a risk group", func(t *testing.T) {
		require.Equal(t,
			vo.NewRiskGroup(true, true, true, true, true, true, true, true, true, true, true, true),
			vo.RiskGroup{
				PregnantWomen:                           true,
				WellnessPerson:                          true,
				AIDS:                                    true,
				Diabetes:                                true,
				Obesity:                                 true,
				AcuteOrChronicHeartDisease:              true,
				ChronicLungDiseaseIncludesCOPDAndAsthma: true,
				Cancer:                                  true,
				ChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
				RenalInsufficiency: true,
				AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
				EssentialHypertension: true,
			})
	})
}
