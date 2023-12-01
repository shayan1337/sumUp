package sumUp

import (
	"fmt"
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

func TestSchemeSupported(t *testing.T) {
	t.Run("should support the scheme when length and range is within bounds", func(t *testing.T) {
		schemes := GetSchemes()
		schemeName, err := SchemeSupported(schemes, "3530111333300000")

		require.NoError(t, err)
		require.Equal(t, "JCB", schemeName)
	})

	t.Run("should not support the scheme and return error when length and range is out of bounds", func(t *testing.T) {
		schemes := GetSchemes()
		cardNumber := "123353011100000"
		_, err := SchemeSupported(schemes, cardNumber)

		require.Error(t, err)
		require.ErrorContains(t, err, fmt.Sprintf("card number %s does not support any scheme", cardNumber))
	})
}
