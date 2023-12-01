package sumUp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScheme_Supports(t *testing.T) {
	maestroScheme := NewScheme("Maestro",
		[]bound{
			{
				lower: 50,
				upper: 50,
			},
			{
				lower: 56,
				upper: 58,
			},
			{
				lower: 6,
				upper: 6,
			},
		},
		[]bound{
			{
				lower: 12,
				upper: 19,
			},
		},
	)

	t.Run("Should not support card scheme when range is out of bounds", func(t *testing.T) {
		supports := maestroScheme.Supports("5159649826438453")
		require.False(t, supports)
	})

	t.Run("Should not support card scheme when length of card is lesser than the lower bound", func(t *testing.T) {
		supports := maestroScheme.Supports("615964923")
		require.False(t, supports)
	})

	t.Run("Should not support card scheme when length of card is higher than the upper bound", func(t *testing.T) {
		supports := maestroScheme.Supports("61596492334324232342412")
		require.False(t, supports)
	})

	t.Run("Should support card scheme when length and range are within bounds", func(t *testing.T) {
		supports := maestroScheme.Supports("6759649826438453")
		require.True(t, supports)
	})
}
