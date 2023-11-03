package influenzacensus_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"taps/modules/influenzacensus"
	"testing"
)

func TestInfluenzaCesus(t *testing.T) {
	t.Run("save census", func(t *testing.T) {
		influenzaMemoryStore := influenzacensus.NewInMemoryInfluenzaStore()
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(influenzaMemoryStore)

		require.NoError(t, influenzaCensus.Take(
			&influenzacensus.FieldCensusParameters{
				ID:            "RAHE190116MMCMRSA7",
				LastLastName:  "RAMIREZ",
				FirstLastName: "HERRERA",
				FirstName:     "ESTHER ELIZABETH",
				Gender:        "MUJER",
				DOB:           "16/01/2019",
				State:         "MEXICO",
				Number:        15,
			},
		))

		require.Equal(
			t,
			[]influenzacensus.InfluenzaCensus{
				{
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
			influenzaMemoryStore.All())
	})

	t.Run("fails if the fields ID FirstLastName LastLastName FirstName DOB State Gender Number are not present", func(t *testing.T) {
		influenzaMemoryStore := influenzacensus.NewInMemoryInfluenzaStore()
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(influenzaMemoryStore)
		type testScenario struct {
			name       string
			parameters *influenzacensus.FieldCensusParameters
		}

		testScenarios := []testScenario{
			{
				name: "No ID",
				parameters: &influenzacensus.FieldCensusParameters{
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No LastLastName",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No FirstLastName",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:           "RAHE190116MMCMRSA7",
					LastLastName: "RAMIREZ",
					FirstName:    "ESTHER ELIZABETH",
					Gender:       "MUJER",
					DOB:          "16/01/2019",
					State:        "MEXICO",
					Number:       15,
				},
			},
			{
				name: "No FirstName",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No Gender",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No DOB",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No State",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					Number:        15,
				},
			},
			{
				name: "No Number",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
				},
			},
			{
				name: "No ID\nNo State",
				parameters: &influenzacensus.FieldCensusParameters{
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					Number:        15,
				},
			},
		}

		for _, ts := range testScenarios {
			t.Run(fmt.Sprintf("Test %s", ts.name), func(t *testing.T) {
				require.EqualError(t, influenzaCensus.Take(ts.parameters), ts.name)
			})
		}
	})
}
