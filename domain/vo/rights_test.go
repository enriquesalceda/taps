package vo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestRights(t *testing.T) {
	t.Run("try parse rights", func(t *testing.T) {
		issste, err := vo.TryNewRights("ISSSTE")
		require.NoError(t, err)
		assert.Equal(t, vo.Rights.ISSSTE, issste)

		imss, err := vo.TryNewRights("IMSS")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.IMSS, imss)

		sedena, err := vo.TryNewRights("SEDENA")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SEDENA, sedena)

		semar, err := vo.TryNewRights("SEMAR")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SEMAR, semar)

		pemex, err := vo.TryNewRights("PEMEX")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.PEMEX, pemex)

		ssa, err := vo.TryNewRights("SSA")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SSA, ssa)

		none, err := vo.TryNewRights("NONE")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.NONE, none)

		_, err = vo.TryNewRights("INVALID-RIGHT")
		require.EqualError(t, err, "invalid rights")
	})

	t.Run("must parse rights", func(t *testing.T) {
		issste := vo.MustNewRights("ISSSTE")
		assert.Equal(t, vo.Rights.ISSSTE, issste)

		imss := vo.MustNewRights("IMSS")
		assert.Equal(t, vo.Rights.IMSS, imss)

		sedena := vo.MustNewRights("SEDENA")
		assert.Equal(t, vo.Rights.SEDENA, sedena)

		semar := vo.MustNewRights("SEMAR")
		assert.Equal(t, vo.Rights.SEMAR, semar)

		pemex := vo.MustNewRights("PEMEX")
		assert.Equal(t, vo.Rights.PEMEX, pemex)

		ssa := vo.MustNewRights("SSA")
		assert.Equal(t, vo.Rights.SSA, ssa)

		none := vo.MustNewRights("NONE")
		assert.Equal(t, vo.Rights.NONE, none)

		require.Panics(t, func() {
			vo.MustNewRights("INVALID-RIGHT")
		})
	})
}
