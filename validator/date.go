package validator

const (
	FirstDay                = 1
	FirstMonth              = FirstDay
	LastDay                 = 30
	LastMonth               = 13
	PagumeLastDay           = 5
	PagumeLeapYearLastDay   = 6
)

// DateValidator validates Ethiopian dates
type DateValidator struct {
	day   int
	month int
	year  int
}

// NewDateValidator creates a new DateValidator
func NewDateValidator(day, month, year int) *DateValidator {
	return &DateValidator{
		day:   day,
		month: month,
		year:  year,
	}
}

// IsValid validates the Ethiopian date
func (v *DateValidator) IsValid() bool {
	validators := []func() bool{
		v.isDateValuesIntegers,
		v.isValidDayRange,
		v.isValidMonthRange,
		v.isValidPagumeDayRange,
		v.isValidLeapDay,
	}

	for _, validator := range validators {
		if !validator() {
			return false
		}
	}
	return true
}

// isValidDayRange checks if day is in valid range
func (v *DateValidator) isValidDayRange() bool {
	return v.day >= FirstDay && v.day <= LastDay
}

// isValidMonthRange checks if month is in valid range
func (v *DateValidator) isValidMonthRange() bool {
	return v.month >= FirstMonth && v.month <= LastMonth
}

// isValidPagumeDayRange checks if day is valid for Pagume (13th month)
func (v *DateValidator) isValidPagumeDayRange() bool {
	if v.month == LastMonth {
		return v.day <= PagumeLeapYearLastDay
	}
	return true
}

// isValidLeapDay checks if the 6th day of Pagume is valid (only in leap years)
func (v *DateValidator) isValidLeapDay() bool {
	if v.month == LastMonth && v.day == PagumeLeapYearLastDay {
		return NewLeapYearValidator(v.year).IsValid()
	}
	return true
}

// isDateValuesIntegers checks if all date values are valid integers
func (v *DateValidator) isDateValuesIntegers() bool {
	return isValidInteger(v.day, v.month, v.year)
}
