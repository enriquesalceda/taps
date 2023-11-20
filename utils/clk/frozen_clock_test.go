package clk_test

import (
	"github.com/stretchr/testify/require"
	"taps/utils/clk"
	"testing"
	"time"
)

func TestFrozenClockNow(t *testing.T) {
	t.Run("returns the frozen time", func(t *testing.T) {
		fc := clk.NewFrozenClock(
			time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		)

		now := fc.Now()

		require.Equal(t,
			now,
			time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		)
	})
}
