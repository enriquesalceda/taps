package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestCurp(t *testing.T) {
	t.Run("Parse CURP from string", func(t *testing.T) {
		rawCURPData := "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"

		curp := vo.ParseCURP(rawCURPData)

		require.Equal(t,
			vo.Curp{
				ID:            "RAHE190116MMCMRSA7",
				LastLastName:  "RAMIREZ",
				FirstLastName: "HERRERA",
				FirstName:     "ESTHER ELIZABETH",
				Gender:        "MUJER",
				DOB:           "16/01/2019",
				State:         "MEXICO",
				Number:        15,
			},
			curp,
		)
	})
}
