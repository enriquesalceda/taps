package influenzacensus_test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"taps/modules/influenzacensus"
	"testing"
)

func TestInfluenzaCensus(t *testing.T) {
	t.Run("save census", func(t *testing.T) {
		influenzaMemoryStore := influenzacensus.NewInMemoryInfluenzaStore()
		influenzaCensus := influenzacensus.NewInfluenzaCensusTaker(influenzaMemoryStore)

		response := influenzaCensus.Take(
			events.APIGatewayProxyRequest{Body: PreparePayload(t, CensusPayload{
				ID:            "RAHE190116MMCMRSA7",
				LastLastName:  "RAMIREZ",
				FirstLastName: "HERRERA",
				FirstName:     "ESTHER ELIZABETH",
				Gender:        "MUJER",
				DOB:           "16/01/2019",
				State:         "MEXICO",
				Number:        15,
			})},
		)

		require.Equal(t, response.StatusCode, 200)
		require.Equal(t, response.Body, "success")
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
			parameters string
		}

		testScenarios := []testScenario{
			{
				name: "No ID",
				parameters: PreparePayload(t, CensusPayload{
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				}),
			},
			{
				name: "No LastLastName",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				}),
			},
			{
				name: "No FirstLastName",
				parameters: PreparePayload(t, CensusPayload{
					ID:           "RAHE190116MMCMRSA7",
					LastLastName: "RAMIREZ",
					FirstName:    "ESTHER ELIZABETH",
					Gender:       "MUJER",
					DOB:          "16/01/2019",
					State:        "MEXICO",
					Number:       15,
				}),
			},
			{
				name: "No FirstName",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				}),
			},
			{
				name: "No Gender",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					DOB:           "16/01/2019",
					State:         "MEXICO",
					Number:        15,
				}),
			},
			{
				name: "No DOB",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					State:         "MEXICO",
					Number:        15,
				}),
			},
			{
				name: "No State",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					Number:        15,
				}),
			},
			{
				name: "No Number",
				parameters: PreparePayload(t, CensusPayload{
					ID:            "RAHE190116MMCMRSA7",
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					State:         "MEXICO",
				}),
			},
			{
				name: "No ID\nNo State",
				parameters: PreparePayload(t, CensusPayload{
					LastLastName:  "RAMIREZ",
					FirstLastName: "HERRERA",
					FirstName:     "ESTHER ELIZABETH",
					Gender:        "MUJER",
					DOB:           "16/01/2019",
					Number:        15,
				}),
			},
		}

		for _, ts := range testScenarios {
			t.Run(fmt.Sprintf("Test %s", ts.name), func(t *testing.T) {
				response := influenzaCensus.Take(events.APIGatewayProxyRequest{Body: ts.parameters})
				require.Equal(t, 400, response.StatusCode)
				require.Equal(t, ts.name, response.Body)
			})
		}
	})
}

type CensusPayload struct {
	ID            string
	LastLastName  string
	FirstLastName string
	FirstName     string
	Gender        string
	DOB           string
	State         string
	Number        int
}

func PreparePayload(t *testing.T, censusPayload CensusPayload) string {
	body, err := json.Marshal(censusPayload)
	require.NoError(t, err)
	return string(body)
}
