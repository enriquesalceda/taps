package influenzacensus_test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"taps/domain"
	"taps/domain/vo"
	"taps/modules/influenzacensus"
	"taps/modules/influenzacensus/store"
	"taps/utils/clk"
	"testing"
	"time"
)

func TestInfluenzaCensus(t *testing.T) {
	t.Run("save census", func(t *testing.T) {
		influenzaMemoryStore := store.NewInMemoryInfluenzaStore(map[string]domain.Census{})
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(
			influenzaMemoryStore,
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		response := influenzaCensus.Handle(
			events.APIGatewayProxyRequest{
				Body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
		)

		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, response.Body, "success")
		require.Equal(
			t,
			map[string]domain.Census{
				"RAHE190116MMCMRSA7": {
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
					Address:         vo.MustParseAddress("123", "Main Street", "Greenwood"),
					ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					TargetGroup: domain.TargetGroup{
						SixToFiftyNineMonthsOld: true,
						SixtyMonthsAndMore:      false,
					},
					SeasonalInfluenzaVaccinationSchedule: domain.SeasonalInfluenzaVaccinationSchedule{
						FirstDose:  true,
						SecondDose: false,
						AnnualDose: false,
					},
				},
			},
			influenzaMemoryStore.All())
	})

	t.Run("save multiple census", func(t *testing.T) {
		influenzaMemoryStore := store.NewInMemoryInfluenzaStore(map[string]domain.Census{})
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(
			influenzaMemoryStore,
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		response := influenzaCensus.Handle(
			events.APIGatewayProxyRequest{
				Body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
		)

		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, response.Body, "success")

		response = influenzaCensus.Handle(
			events.APIGatewayProxyRequest{
				Body: Body(t, "AABBCC112233||PEREZ|PEREZ|PEPE|HOMBRE|16/01/2019|MEXICO|7|"),
			},
		)

		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, response.Body, "success")
		require.Equal(
			t,
			map[string]domain.Census{
				"AABBCC112233": {
					ID: "AABBCC112233",
					CURP: vo.Curp{
						ID:            "AABBCC112233",
						LastLastName:  "PEREZ",
						FirstLastName: "PEREZ",
						FirstName:     "PEPE",
						Gender:        "HOMBRE",
						DOB:           "16/01/2019",
						State:         "MEXICO",
						Number:        7,
					},
					Address:         vo.MustParseAddress("123", "Main Street", "Greenwood"),
					ApplicationDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
					TargetGroup: domain.TargetGroup{
						SixToFiftyNineMonthsOld: true,
						SixtyMonthsAndMore:      false,
					},
					SeasonalInfluenzaVaccinationSchedule: domain.SeasonalInfluenzaVaccinationSchedule{
						FirstDose:  true,
						SecondDose: false,
						AnnualDose: false,
					},
				},
				"RAHE190116MMCMRSA7": {
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
					Address:         vo.MustParseAddress("123", "Main Street", "Greenwood"),
					ApplicationDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
					TargetGroup: domain.TargetGroup{
						SixToFiftyNineMonthsOld: true,
						SixtyMonthsAndMore:      false,
					},
					SeasonalInfluenzaVaccinationSchedule: domain.SeasonalInfluenzaVaccinationSchedule{
						FirstDose:  true,
						SecondDose: false,
						AnnualDose: false,
					},
				},
			},
			influenzaMemoryStore.All())
	})

	t.Run("fails if the fields ID FirstLastName LastLastName FirstName DOB State Gender Number are not present", func(t *testing.T) {
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(
			store.NewInMemoryInfluenzaStore(map[string]domain.Census{}),
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)
		type testScenario struct {
			name string
			body string
		}

		testScenarios := []testScenario{
			{
				name: "curp is not including: ID",
				body: Body(t, "||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
			{
				name: "curp is not including: LastLastName",
				body: Body(t, "RAHE190116MMCMRSA7|||HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
			{
				name: "curp is not including: FirstLastName",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ||ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
			{
				name: "curp is not including: FirstName",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA||MUJER|16/01/2019|MEXICO|15|"),
			},
			{
				name: "curp is not including: Gender",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH||16/01/2019|MEXICO|15|"),
			},
			{
				name: "curp is not including: DOB",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER||MEXICO|15|"),
			},
			{
				name: "curp is not including: State",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019||15|"),
			},
			{
				name: "curp should have 10 items, it has 9",
				body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|"),
			},
			{
				name: "curp is not including: ID, State",
				body: Body(t, "||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019||15|"),
			},
		}

		for _, ts := range testScenarios {
			t.Run(fmt.Sprintf("Test %s", ts.name), func(t *testing.T) {
				response := influenzaCensus.Handle(events.APIGatewayProxyRequest{Body: ts.body})
				require.Equal(t, response.StatusCode, 400)
				require.Equal(t, ts.name, response.Body)
			})
		}
	})

	t.Run("raises a 409 conflict when the CURP code already exists", func(t *testing.T) {
		influenzaMemoryStore := store.NewInMemoryInfluenzaStore(map[string]domain.Census{
			"RAHE190116MMCMRSA7": {
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
			},
		})
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(
			influenzaMemoryStore,
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		response := influenzaCensus.Handle(
			events.APIGatewayProxyRequest{
				Body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
		)

		require.Equal(t, 409, response.StatusCode)
		require.Equal(t, response.Body, "census already exists")
		require.Equal(
			t,
			map[string]domain.Census{
				"RAHE190116MMCMRSA7": {
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
				},
			},
			influenzaMemoryStore.All())
	})

	t.Run("raises a 500 when the store returns an error", func(t *testing.T) {
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(
			store.NewBrokenInfluenzaStore(),
			clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		)

		response := influenzaCensus.Handle(
			events.APIGatewayProxyRequest{
				Body: Body(t, "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
			},
		)

		require.Equal(t, 500, response.StatusCode)
		require.Equal(t, response.Body, "internal server error")
	})
}

func Body(t *testing.T, curp string) string {
	sample := map[string]any{
		"CURP": curp,
		"Address": map[string]any{
			"StreetNumber": "123",
			"StreetName":   "Main Street",
			"SuburbName":   "Greenwood",
		},
		"TargetGroup": map[string]any{
			"SixToFiftyNineMonthsOld": true,
			"SixtyMonthsAndMore":      false,
		},
		"RiskGroup": map[string]any{
			"PregnantWomen":              false,
			"WellnessPerson":             false,
			"AIDS":                       false,
			"Diabetes":                   false,
			"Obesity":                    false,
			"AcuteOrChronicHeartDisease": false,
			"ChronicLungDiseaseIncludesCOPDAndAsthma": false,
			"Cancer": false,
			"CongenitalHeartOrPulmonaryDiseasesOrOtherChronicConditionsThatRequireProlongedConsumptionOfSalicylic": false,
			"RenalInsufficiency": false,
			"AcquiredImmunosuppressionDueToDiseaseOrTreatmentExceptAIDS": false,
			"EssentialHypertension": false,
		},
		"OtherRiskGroup": false,
		"SeasonalInfluenzaVaccinationSchedule": map[string]any{
			"FirstDose":  true,
			"SecondDose": false,
			"AnnualDose": false,
		},
		"BatchNumber": "Batch12345",
		"Rights":      "Basic Healthcare Rights",
	}

	body, err := json.Marshal(sample)
	require.NoError(t, err)
	return string(body)
}
