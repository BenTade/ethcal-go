package validator

// LeapYearValidator validates Ethiopian leap years
type LeapYearValidator struct {
	year int
}

// NewLeapYearValidator creates a new LeapYearValidator
func NewLeapYearValidator(year int) *LeapYearValidator {
	return &LeapYearValidator{year: year}
}

// IsValid checks if the year is valid and is a leap year
func (v *LeapYearValidator) IsValid() bool {
	return isValidInteger(v.year) && v.isLeapYear()
}

// isLeapYear checks if the year is a leap year in the Ethiopian calendar
// In Ethiopian calendar, a year is a leap year if (year + 1) % 4 == 0
func (v *LeapYearValidator) isLeapYear() bool {
	return (v.year+1)%4 == 0
}

// isValidInteger checks if values are valid integers (non-negative)
func isValidInteger(values ...int) bool {
	for _, v := range values {
		if v < 0 {
			return false
		}
	}
	return true
}
