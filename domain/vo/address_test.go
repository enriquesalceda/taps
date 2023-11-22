package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestAddress(t *testing.T) {
	t.Run("try parse address", func(t *testing.T) {
		address, err := vo.TryNewAddress(
			"1",
			"Calle Benito Juarez",
			"El Centro",
		)
		require.NoError(t, err)

		require.Equal(t,
			vo.Address{
				StreetNumber: "1",
				StreetName:   "Calle Benito Juarez",
				SuburbName:   "El Centro",
			},
			address,
		)
	})

	t.Run("try parse address with empty street number", func(t *testing.T) {
		_, err := vo.TryNewAddress(
			"",
			"Calle Benito Juarez",
			"El Centro",
		)

		require.EqualError(t, err, "address should have: street number")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryNewAddress(
			"1",
			"",
			"El Centro",
		)

		require.EqualError(t, err, "address should have: street name")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryNewAddress(
			"1",
			"Calle Benito Juarez",
			"",
		)

		require.EqualError(t, err, "address should have: suburb name")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryNewAddress(
			"",
			"",
			"",
		)

		require.EqualError(t, err, "address should have: street number, street name, suburb name")
	})

	t.Run("must parse address", func(t *testing.T) {
		t.Run("successfully", func(t *testing.T) {
			address := vo.MustNewAddress(
				"1",
				"Calle Benito Juarez",
				"El Centro",
			)

			require.Equal(t,
				vo.Address{
					StreetNumber: "1",
					StreetName:   "Calle Benito Juarez",
					SuburbName:   "El Centro",
				},
				address,
			)
		})

		t.Run("panics when address is invalid", func(t *testing.T) {
			require.Panics(t, func() {
				vo.MustNewAddress(
					"",
					"",
					"",
				)
			})
		})
	})
}
