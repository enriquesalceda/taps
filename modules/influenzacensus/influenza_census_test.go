package influenzacensus_test

import (
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
}
