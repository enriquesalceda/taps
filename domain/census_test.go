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
		cmd := command.CreateCensus{
			CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
			TargetGroup: command.TargetGroup{
				SixToFiftyNineMonthsOld: true,
				SixtyMonthsAndMore:      false,
			},
		}
		clock := clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))

		census, err := domain.BuildCensus(cmd, clock)
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
				ApplicationDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				TargetGroup: domain.TargetGroup{
					SixToFiftyNineMonthsOld: true,
					SixtyMonthsAndMore:      false,
				},
			},
			census,
		)
	})

	t.Run("returns an error when TargetGroup values are the same", func(t *testing.T) {
		cmd := command.CreateCensus{
			CURP: "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|",
			TargetGroup: command.TargetGroup{
				SixToFiftyNineMonthsOld: true,
				SixtyMonthsAndMore:      true,
			},
		}
		clock := clk.NewFrozenClock(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))

		_, err := domain.BuildCensus(cmd, clock)
		require.EqualError(t, err, "target group values cannot be the same")
	})
}
