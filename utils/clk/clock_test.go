package clk_test

import (
	"github.com/stretchr/testify/assert"
	"taps/utils/clk"
	"testing"
)

func TestNow(t *testing.T) {
	t.Run("returns the current time", func(t *testing.T) {
		clock := clk.NewClock()

		now := clock.Now()

		assert.NotZero(t, now)
	})
}
