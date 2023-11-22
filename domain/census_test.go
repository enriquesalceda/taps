package domain_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain"
	"taps/domain/command"
	"taps/domain/vo"
	"taps/utils/clk"
	"testing"
	"time"
)

func TestBuildCensus(t *testing.T) {
	t.Run("builds a census", func(t *testing.T) {
		census, err := domain.BuildCensus(
			command.CreateCensus{
				CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
				Address: command.Address{
					StreetNumber: "18b",
					StreetName:   "chapulin",
					SuburbName:   "arcoiris",
				},
				TargetGroup: command.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      false,
				},
				RiskGroup: command.RiskGroup{
					PregnantWomen:                           true,
					WellnessPerson:                          true,
					AIDS:                                    true,
					Diabetes:                                true,
					Obesity:                                 true,
					AcuteOrChronicHeartDisease:              true,
					ChronicLungDiseaseIncludesCOPDAndAsthma: true,
					Cancer:                                  true,
					CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
					RenalInsufficiency: true,
					AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
					EssentialHypertension: true,
				},
				SeasonalInfluenzaVaccinationSchedule: command.SeasonalInfluenzaVaccinationSchedule{
					FirstDose:  true,
					SecondDose: false,
					AnnualDose: false,
				},
				Rights: "ISSSTE",
			},
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		require.NoError(t, err)
		require.Equal(t,
			domain.Census{
				ID: "RAHE190116MMCMRSA7",
				CURP: vo.Curp{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
				Address:         AddressFixture(t, "18b", "chapulin", "arcoiris"),
				ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				TargetGroup:     vo.MustParseTargetGroup(true, false),
				RiskGroup: domain.RiskGroup{
					PregnantWomen:                           true,
					WellnessPerson:                          true,
					AIDS:                                    true,
					Diabetes:                                true,
					Obesity:                                 true,
					AcuteOrChronicHeartDisease:              true,
					ChronicLungDiseaseIncludesCOPDAndAsthma: true,
					Cancer:                                  true,
					CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic: true,
					RenalInsufficiency: true,
					AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS: true,
					EssentialHypertension: true,
				},
				SeasonalInfluenzaVaccinationSchedule: vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, false),
				Rights:                               vo.Rights.ISSSTE,
			},
			census,
		)
	})

	t.Run("returns an error when TargetGroup values are the same", func(t *testing.T) {
		_, err := domain.BuildCensus(
			command.CreateCensus{
				CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
				TargetGroup: command.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      true,
				},
			},
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		require.EqualError(t, err, "target group values cannot be the same")
	})

	t.Run("returns an error when TargetGroup values are the same", func(t *testing.T) {
		_, err := domain.BuildCensus(
			command.CreateCensus{
				CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
				TargetGroup: command.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      false,
				},
				SeasonalInfluenzaVaccinationSchedule: command.SeasonalInfluenzaVaccinationSchedule{
					FirstDose:  false,
					SecondDose: false,
					AnnualDose: true,
				},
			},
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		require.EqualError(t, err, "doses should be in order")
	})

	t.Run("returns an error when address is not valid", func(t *testing.T) {
		_, err := domain.BuildCensus(
			command.CreateCensus{
				CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
				Address: command.Address{
					StreetNumber: "18b",
					StreetName:   "",
					SuburbName:   "",
				},
				TargetGroup: command.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      false,
				},
			},
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		require.EqualError(t, err, "address should have: street name, suburb name")
	})

	t.Run("returns an error when right is not valid", func(t *testing.T) {
		_, err := domain.BuildCensus(
			command.CreateCensus{
				CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
				Address: command.Address{
					StreetNumber: "18b",
					StreetName:   "Emiliano Zapata",
					SuburbName:   "El Hormiguero",
				},
				TargetGroup: command.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      false,
				},
				Rights: "INVALID-RIGHT",
			},
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		require.EqualError(t, err, "invalid rights")
	})
}

func AddressFixture(t *testing.T, streetNumber, streetName, suburbName string) vo.Address {
	address, err := vo.TryParseAddress(streetNumber, streetName, suburbName)
	require.NoError(t, err)
	return address
}
