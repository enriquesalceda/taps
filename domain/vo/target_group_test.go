package vo_test

import (
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestTargetGroup(t *testing.T) {
	t.Run("builds a target group", func(t *testing.T) {
		targetGroup, err := vo.TryNewTargetGroup(true, false)
		require.NoError(t, err)
		require.Equal(t,
			vo.TargetGroup{
				SixToFiftyNineMonthsOld: true,
				SixtyMonthsAndMore:      false,
			},
			targetGroup,
		)
	})

	t.Run("validates only one target group can be true", func(t *testing.T) {
		_, err := vo.TryNewTargetGroup(true, true)
		require.EqualError(t, err, "target group values cannot be the same")

		_, err = vo.TryNewTargetGroup(false, false)
		require.EqualError(t, err, "target group values cannot be the same")
	})

	t.Run("must create a target group", func(t *testing.T) {
		targetGroup := vo.MustNewTargetGroup(true, false)
		require.Equal(t,
			vo.TargetGroup{
				SixToFiftyNineMonthsOld: true,
				SixtyMonthsAndMore:      false,
			},
			targetGroup,
		)
	})
}
