package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestSeasonalInfluenzaVaccinationSchedule(t *testing.T) {
	t.Run("Try New", func(t *testing.T) {
		t.Run("returns a new SeasonalInfluenzaVaccinationSchedule", func(t *testing.T) {
			actual, err := vo.TryNewSeasonalInfluenzaVaccinationSchedule(true, true, true)

			require.NoError(t, err)
			require.Equal(
				t,
				actual,
				vo.SeasonalInfluenzaVaccinationSchedule{FirstDose: true, SecondDose: true, AnnualDose: true},
			)
		})

		t.Run("returns an error when the first dose is false and the second dose is true", func(t *testing.T) {
			_, err := vo.TryNewSeasonalInfluenzaVaccinationSchedule(false, true, false)
			require.EqualError(t, err, "doses should be in order")
		})

		t.Run("returns an error when the first dose is true, the second dose is false and the third dose is true", func(t *testing.T) {
			_, err := vo.TryNewSeasonalInfluenzaVaccinationSchedule(true, false, true)
			require.EqualError(t, err, "doses should be in order")
		})
	})

	t.Run("Must New", func(t *testing.T) {
		t.Run("returns a new SeasonalInfluenzaVaccinationSchedule", func(t *testing.T) {
			actual := vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, true, true)

			require.Equal(
				t,
				actual,
				vo.SeasonalInfluenzaVaccinationSchedule{FirstDose: true, SecondDose: true, AnnualDose: true},
			)
		})

		t.Run("panics when the first dose is false and the second dose is true", func(t *testing.T) {
			t.Run("returns an error when the first dose is true, the second dose is false and the third dose is true", func(t *testing.T) {
				require.PanicsWithError(t,
					"doses should be in order",
					func() {
						vo.MustNewSeasonalInfluenzaVaccinationSchedule(true, false, true)
					})
			})
		})
	})
}
