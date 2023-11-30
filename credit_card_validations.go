package sumUp

import (
	"fmt"
	"unicode"
)

func ValidCardNumber(cardNumber string) (bool, error) {
	err := validate(cardNumber)
	if err != nil {
		return false, err
	}

	isSecondDigit := false
	sumOfAllDigits := 0

	for idx := len(cardNumber) - 1; idx >= 0; idx-- {
		digit := int(cardNumber[idx] - '0')

		if isSecondDigit {
			double := digit * 2
			sumOfAllDigits += (double / 10) + (double % 10)
		} else {
			sumOfAllDigits += digit
		}

		isSecondDigit = !isSecondDigit
	}
	return sumOfAllDigits%10 == 0, nil
}

func validate(cardNumber string) error {
	if cardNumber == "" {
		return fmt.Errorf("card number %s contains non digits", cardNumber)
	}

	if !isDigitsOnly(cardNumber) {
		return fmt.Errorf("card number %s contains non digits", cardNumber)
	}

	return nil
}

func isDigitsOnly(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
