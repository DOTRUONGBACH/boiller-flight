package validation

import (
    "strings"
)

func ValidateAvailableSeat(capacity int, economy int, business int) bool{
	if (economy + business) > capacity || (economy + business) != capacity{
		return false
	}
	return true
}

func ValidateDestination(from string, to string) bool{
	return strings.EqualFold(strings.ToLower(from),strings.ToLower(to))
}

