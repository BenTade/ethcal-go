package validator

import (
	"testing"
)

func TestLeapYearValidator(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2007, true},  // (2007 + 1) % 4 == 0
		{2011, true},  // (2011 + 1) % 4 == 0
		{2003, true},  // (2003 + 1) % 4 == 0
		{1999, true},  // (1999 + 1) % 4 == 0
		{2000, false}, // (2000 + 1) % 4 != 0
		{2001, false}, // (2001 + 1) % 4 != 0
		{2002, false}, // (2002 + 1) % 4 != 0
		{2009, false}, // (2009 + 1) % 4 != 0
		{2010, false}, // (2010 + 1) % 4 != 0
		{3, true},     // (3 + 1) % 4 == 0
		{7, true},     // (7 + 1) % 4 == 0
		{0, false},    // Special case: year 0
		{-1, false},   // Negative year (invalid)
	}

	for _, tt := range tests {
		v := NewLeapYearValidator(tt.year)
		result := v.IsValid()
		if result != tt.expected {
			t.Errorf("LeapYearValidator(%d).IsValid() = %v, want %v", tt.year, result, tt.expected)
		}
	}
}

func TestDateValidator(t *testing.T) {
	tests := []struct {
		day      int
		month    int
		year     int
		expected bool
	}{
		// Valid dates
		{1, 1, 2000, true},
		{15, 6, 2000, true},
		{30, 12, 2000, true},
		{5, 13, 2000, true},
		{6, 13, 2007, true},  // Valid in leap year
		{1, 1, 1, true},
		{30, 1, 2000, true},

		// Invalid day range
		{0, 1, 2000, false},
		{31, 1, 2000, false},
		{32, 1, 2000, false},
		{-1, 1, 2000, false},

		// Invalid month range
		{1, 0, 2000, false},
		{1, 14, 2000, false},
		{1, 15, 2000, false},
		{1, -1, 2000, false},

		// Invalid Pagume days
		{6, 13, 2000, false}, // 6th of Pagume not valid in non-leap year
		{7, 13, 2000, false},
		{8, 13, 2000, false},
		{6, 13, 2001, false},
		{6, 13, 2002, false},
		{6, 13, 2004, false},
		{6, 13, 2006, false},

		// Valid Pagume days in leap years
		{6, 13, 2003, true},
		{6, 13, 2007, true},
		{6, 13, 2011, true},
		{5, 13, 1999, true},

		// Negative values
		{-1, 1, 2000, false},
		{1, -1, 2000, false},
		{1, 1, -1, false},
	}

	for _, tt := range tests {
		v := NewDateValidator(tt.day, tt.month, tt.year)
		result := v.IsValid()
		if result != tt.expected {
			t.Errorf("DateValidator(%d, %d, %d).IsValid() = %v, want %v",
				tt.day, tt.month, tt.year, result, tt.expected)
		}
	}
}

func TestDateValidatorMethods(t *testing.T) {
	// Test isValidDayRange
	v := NewDateValidator(15, 6, 2000)
	if !v.isValidDayRange() {
		t.Error("isValidDayRange() should return true for day 15")
	}

	v = NewDateValidator(31, 6, 2000)
	if v.isValidDayRange() {
		t.Error("isValidDayRange() should return false for day 31")
	}

	// Test isValidMonthRange
	v = NewDateValidator(15, 6, 2000)
	if !v.isValidMonthRange() {
		t.Error("isValidMonthRange() should return true for month 6")
	}

	v = NewDateValidator(15, 14, 2000)
	if v.isValidMonthRange() {
		t.Error("isValidMonthRange() should return false for month 14")
	}

	// Test isValidPagumeDayRange
	v = NewDateValidator(5, 13, 2000)
	if !v.isValidPagumeDayRange() {
		t.Error("isValidPagumeDayRange() should return true for day 5 of Pagume")
	}

	v = NewDateValidator(7, 13, 2000)
	if v.isValidPagumeDayRange() {
		t.Error("isValidPagumeDayRange() should return false for day 7 of Pagume")
	}

	// Test isValidLeapDay
	v = NewDateValidator(6, 13, 2007)
	if !v.isValidLeapDay() {
		t.Error("isValidLeapDay() should return true for 6th of Pagume in leap year 2007")
	}

	v = NewDateValidator(6, 13, 2000)
	if v.isValidLeapDay() {
		t.Error("isValidLeapDay() should return false for 6th of Pagume in non-leap year 2000")
	}
}
