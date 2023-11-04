package vo_test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestCurp(t *testing.T) {
	t.Run("Parse CURP from string", func(t *testing.T) {
		rawCURPData := "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"

		curp, err := vo.ParseCURP(rawCURPData)

		require.NoError(t, err)
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

	t.Run("Parse errors when the 10 expected items are not present", func(t *testing.T) {
		rawCURPData := "RAHE190116MMCMRSA7||RAMIREZ|HERRERA"

		_, err := vo.ParseCURP(rawCURPData)

		require.Error(t, err)
		require.EqualError(t, err, "curp should have 10 items, it has 4")
	})

	t.Run("Parse errors when the expected number is a special character", func(t *testing.T) {
		rawCURPData := "RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|?|"

		_, err := vo.ParseCURP(rawCURPData)

		require.Error(t, err)
		require.EqualError(t, err, "strconv.Atoi: parsing \"?\": invalid syntax")
	})

	t.Run("Must parse a valid CURP string", func(t *testing.T) {
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
			vo.MustParseCURP("RAHE190116MMCMRSA7||RAMIREZ|HERRERA|ESTHER ELIZABETH|MUJER|16/01/2019|MEXICO|15|"),
		)
	})

	t.Run("Must parse panics if invalid string", func(t *testing.T) {
		require.PanicsWithError(t,
			errors.New("curp should have 10 items, it has 4").Error(),
			func() { vo.MustParseCURP("RAHE190116MMCMRSA7||RAMIREZ|HERRERA") },
		)
	})
}
