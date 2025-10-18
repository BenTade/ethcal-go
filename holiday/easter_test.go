package holiday

import (
	"testing"
)

func TestEasterGet(t *testing.T) {
	tests := []struct {
		year         int
		wantMonth    int
		minDay       int
		maxDay       int
	}{
		{2006, 8, 1, 30},  // Easter should be in Miyazya (month 8)
		{2007, 8, 1, 30},
		{2008, 8, 1, 30},
		{2009, 8, 1, 30},
		{2010, 7, 1, 30},  // Could also be in Megabit (month 7)
		{2011, 8, 1, 30},
		{2012, 8, 1, 30},
	}

	easter := NewEaster()

	for _, tt := range tests {
		dt, err := easter.Get(tt.year)
		if err != nil {
			t.Errorf("Easter.Get(%d) returned error: %v", tt.year, err)
			continue
		}

		// Check that it's on a Sunday
		if dt.GetDayOfWeek() != 7 {
			t.Errorf("Easter.Get(%d) day of week = %d, want 7 (Sunday)", tt.year, dt.GetDayOfWeek())
		}

		// Check that the month is either Megabit or Miyazya
		month := dt.GetMonth()
		if month != 7 && month != 8 {
			t.Errorf("Easter.Get(%d) month = %d, want 7 (Megabit) or 8 (Miyazya)", tt.year, month)
		}

		// Check day is in valid range
		day := dt.GetDay()
		if day < 1 || day > 30 {
			t.Errorf("Easter.Get(%d) day = %d, want 1-30", tt.year, day)
		}
	}
}

func TestEasterGetGregorian(t *testing.T) {
	tests := []int{2006, 2007, 2008, 2009, 2010, 2011, 2012}

	easter := NewEaster()

	for _, year := range tests {
		greg, err := easter.GetGregorian(year)
		if err != nil {
			t.Errorf("Easter.GetGregorian(%d) returned error: %v", year, err)
			continue
		}

		// Check that it's on a Sunday
		if greg.Weekday() != 0 { // 0 is Sunday in Go's time package
			t.Errorf("Easter.GetGregorian(%d) weekday = %v, want Sunday", year, greg.Weekday())
		}
	}
}

func TestEasterConsistency(t *testing.T) {
	// Test that Get and GetGregorian return consistent dates
	tests := []int{2000, 2005, 2010, 2015}

	easter := NewEaster()

	for _, year := range tests {
		dt, err := easter.Get(year)
		if err != nil {
			t.Errorf("Easter.Get(%d) returned error: %v", year, err)
			continue
		}

		greg1 := dt.ToGregorian()

		greg2, err := easter.GetGregorian(year)
		if err != nil {
			t.Errorf("Easter.GetGregorian(%d) returned error: %v", year, err)
			continue
		}

		// Check that both methods return the same Gregorian date
		if greg1.Year() != greg2.Year() || greg1.Month() != greg2.Month() || greg1.Day() != greg2.Day() {
			t.Errorf("Easter methods inconsistent for year %d: Get gives %v, GetGregorian gives %v",
				year, greg1, greg2)
		}
	}
}
