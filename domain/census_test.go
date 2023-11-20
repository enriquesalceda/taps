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
	cmd := command.CreateCensus{
		CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
	}
	clock := clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))

	census, err := domain.BuildCensus(cmd, clock)
	require.NoError(t, err)

	require.Equal(t,
		domain.FieldCensus{
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
			ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		census,
	)
}
