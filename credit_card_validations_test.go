package sumUp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValidCardNumber(t *testing.T) {
	t.Run("should return error for empty card number", func(t *testing.T) {
		isValid, err := ValidCardNumber("")

		require.False(t, isValid)

		require.Error(t, err)
		require.ErrorContains(t, err, "card number  contains non digits")
	})

	t.Run("should return error for card number with letters in them", func(t *testing.T) {
		isValid, err := ValidCardNumber("23352352f345")

		require.False(t, isValid)

		require.Error(t, err)
		require.ErrorContains(t, err, "card number 23352352f345 contains non digits")
	})

	t.Run("should return false and no error for card number that is not divisible by 10", func(t *testing.T) {
		isValid, err := ValidCardNumber("23352352345")

		require.False(t, isValid)

		require.NoError(t, err)
	})

	t.Run("should return true and no error for card number that is completely valid", func(t *testing.T) {
		isValid, err := ValidCardNumber("5237251624778133")

		require.True(t, isValid)

		require.NoError(t, err)
	})
}
