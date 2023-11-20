package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestAddress(t *testing.T) {
	t.Run("try parse address", func(t *testing.T) {
		address, err := vo.TryParseAddress(
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
		_, err := vo.TryParseAddress(
			"",
			"Calle Benito Juarez",
			"El Centro",
		)

		require.EqualError(t, err, "address should have: street number")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryParseAddress(
			"1",
			"",
			"El Centro",
		)

		require.EqualError(t, err, "address should have: street name")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryParseAddress(
			"1",
			"Calle Benito Juarez",
			"",
		)

		require.EqualError(t, err, "address should have: suburb name")
	})

	t.Run("try parse address with empty street name", func(t *testing.T) {
		_, err := vo.TryParseAddress(
			"",
			"",
			"",
		)

		require.EqualError(t, err, "address should have: street number, street name, suburb name")
	})
}
