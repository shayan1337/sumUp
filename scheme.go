package sumUp

import (
	"strconv"
	"strings"
)

type bound struct {
	lower int
	upper int
}

type Scheme struct {
	name                       string
	startingDigitsRange        []bound
	numberOfDigitsAllowedRange []bound
}

func NewScheme(name string, ranges []bound, numberOfDigitsAllowed []bound) *Scheme {
	return &Scheme{
		name:                       name,
		startingDigitsRange:        ranges,
		numberOfDigitsAllowedRange: numberOfDigitsAllowed,
	}
}

func (s *Scheme) Supports(cardNumber string) bool {
	return s.startingDigitsWithinBounds(cardNumber) && s.numberOfDigitsAllowedWithinBounds(cardNumber)
}

func (s *Scheme) startingDigitsWithinBounds(cardNumber string) bool {
	for _, bound := range s.startingDigitsRange {
		for start := bound.lower; start <= bound.upper; start++ {
			prefix := strconv.Itoa(start)
			if strings.HasPrefix(cardNumber, prefix) {
				return true
			}
		}
	}
	return false
}

func (s *Scheme) numberOfDigitsAllowedWithinBounds(cardNumber string) bool {
	for _, bound := range s.numberOfDigitsAllowedRange {
		if bound.lower <= len(cardNumber) && len(cardNumber) <= bound.upper {
			return true
		}
	}
	return false
}
