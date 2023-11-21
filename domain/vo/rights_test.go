package vo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"taps/domain/vo"
	"testing"
)

func TestRights(t *testing.T) {
	t.Run("try parse rights", func(t *testing.T) {
		issste, err := vo.TryParseRights("ISSSTE")
		require.NoError(t, err)
		assert.Equal(t, vo.Rights.ISSSTE, issste)

		imss, err := vo.TryParseRights("IMSS")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.IMSS, imss)

		sedena, err := vo.TryParseRights("SEDENA")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SEDENA, sedena)

		semar, err := vo.TryParseRights("SEMAR")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SEMAR, semar)

		pemex, err := vo.TryParseRights("PEMEX")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.PEMEX, pemex)

		ssa, err := vo.TryParseRights("SSA")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.SSA, ssa)

		none, err := vo.TryParseRights("NONE")
		require.NoError(t, err)
		require.Equal(t, vo.Rights.NONE, none)

		_, err = vo.TryParseRights("INVALID-RIGHT")
		require.EqualError(t, err, "invalid rights")
	})

	t.Run("must parse rights", func(t *testing.T) {
		issste := vo.MustParseRights("ISSSTE")
		assert.Equal(t, vo.Rights.ISSSTE, issste)

		imss := vo.MustParseRights("IMSS")
		assert.Equal(t, vo.Rights.IMSS, imss)

		sedena := vo.MustParseRights("SEDENA")
		assert.Equal(t, vo.Rights.SEDENA, sedena)

		semar := vo.MustParseRights("SEMAR")
		assert.Equal(t, vo.Rights.SEMAR, semar)

		pemex := vo.MustParseRights("PEMEX")
		assert.Equal(t, vo.Rights.PEMEX, pemex)

		ssa := vo.MustParseRights("SSA")
		assert.Equal(t, vo.Rights.SSA, ssa)

		none := vo.MustParseRights("NONE")
		assert.Equal(t, vo.Rights.NONE, none)

		require.Panics(t, func() {
			vo.MustParseRights("INVALID-RIGHT")
		})
	})
}
