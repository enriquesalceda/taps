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
					ID:            "",
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
					LastLastName:  "",
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
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				},
			},
			{
				name: "No FirstName",
				parameters: &influenzacensus.FieldCensusParameters{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "",
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
					Gender:        "",
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
					DOB:           "",
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
					State:         "",
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
					Number:        0,
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
