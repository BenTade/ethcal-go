package ethcal

import (
	"testing"
	"time"
)

func TestNewDateTime(t *testing.T) {
	// Test with a known Gregorian date
	gregorian := time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC)
	dt := NewDateTime(gregorian)

	// May 12, 2017 should be Ginbot 4, 2009 in Ethiopian calendar
	expectedYear := 2009
	expectedMonth := 9
	expectedDay := 4

	if dt.GetYear() != expectedYear {
		t.Errorf("NewDateTime year = %d, want %d", dt.GetYear(), expectedYear)
	}
	if dt.GetMonth() != expectedMonth {
		t.Errorf("NewDateTime month = %d, want %d", dt.GetMonth(), expectedMonth)
	}
	if dt.GetDay() != expectedDay {
		t.Errorf("NewDateTime day = %d, want %d", dt.GetDay(), expectedDay)
	}
}

func TestOf(t *testing.T) {
	tests := []struct {
		year       int
		month      int
		day        int
		gregYear   int
		gregMonth  int
		gregDay    int
	}{
		{2000, 1, 1, 2007, 9, 12},  // Ethiopian millennium
		{2009, 9, 4, 2017, 5, 12},
		{1983, 9, 20, 1991, 5, 28}, // Fall of Derg
	}

	for _, tt := range tests {
		dt, err := Of(tt.year, tt.month, tt.day, 0, 0, 0, time.UTC)
		if err != nil {
			t.Errorf("Of(%d, %d, %d) returned error: %v", tt.year, tt.month, tt.day, err)
			continue
		}

		greg := dt.ToGregorian()
		if greg.Year() != tt.gregYear || int(greg.Month()) != tt.gregMonth || greg.Day() != tt.gregDay {
			t.Errorf("Of(%d, %d, %d) = Gregorian(%d, %d, %d), want Gregorian(%d, %d, %d)",
				tt.year, tt.month, tt.day,
				greg.Year(), int(greg.Month()), greg.Day(),
				tt.gregYear, tt.gregMonth, tt.gregDay)
		}
	}
}

func TestFromTimestamp(t *testing.T) {
	// Use a known timestamp: 1494547200 = 2017-05-12 00:00:00 UTC
	timestamp := int64(1494547200)
	dt := FromTimestamp(timestamp, time.UTC)

	expectedYear := 2009
	expectedMonth := 9
	expectedDay := 4

	if dt.GetYear() != expectedYear || dt.GetMonth() != expectedMonth || dt.GetDay() != expectedDay {
		t.Errorf("FromTimestamp(%d) = (%d, %d, %d), want (%d, %d, %d)",
			timestamp, dt.GetYear(), dt.GetMonth(), dt.GetDay(),
			expectedYear, expectedMonth, expectedDay)
	}

	if dt.GetTimestamp() != timestamp {
		t.Errorf("GetTimestamp() = %d, want %d", dt.GetTimestamp(), timestamp)
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2007, true},
		{2011, true},
		{2003, true},
		{2000, false},
		{2001, false},
		{2009, false},
	}

	for _, tt := range tests {
		dt, err := Of(tt.year, 1, 1, 0, 0, 0, time.UTC)
		if err != nil {
			t.Errorf("Of(%d, 1, 1) returned error: %v", tt.year, err)
			continue
		}

		if dt.IsLeapYear() != tt.expected {
			t.Errorf("IsLeapYear() for year %d = %v, want %v", tt.year, dt.IsLeapYear(), tt.expected)
		}
	}
}

func TestGetDaysInMonth(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		expected int
	}{
		{2000, 1, 30},
		{2000, 6, 30},
		{2000, 12, 30},
		{2000, 13, 5},  // Non-leap year Pagume
		{2007, 13, 6},  // Leap year Pagume
	}

	for _, tt := range tests {
		dt, err := Of(tt.year, tt.month, 1, 0, 0, 0, time.UTC)
		if err != nil {
			t.Errorf("Of(%d, %d, 1) returned error: %v", tt.year, tt.month, err)
			continue
		}

		if dt.GetDaysInMonth() != tt.expected {
			t.Errorf("GetDaysInMonth() for %d/%d = %d, want %d",
				tt.year, tt.month, dt.GetDaysInMonth(), tt.expected)
		}
	}
}

func TestGetDayOfYear(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		day      int
		expected int
	}{
		{2000, 1, 1, 1},
		{2000, 1, 30, 30},
		{2000, 2, 1, 31},
		{2000, 13, 5, 365},
		{2007, 13, 6, 366}, // Leap year
	}

	for _, tt := range tests {
		dt, err := Of(tt.year, tt.month, tt.day, 0, 0, 0, time.UTC)
		if err != nil {
			t.Errorf("Of(%d, %d, %d) returned error: %v", tt.year, tt.month, tt.day, err)
			continue
		}

		if dt.GetDayOfYear() != tt.expected {
			t.Errorf("GetDayOfYear() for %d/%d/%d = %d, want %d",
				tt.year, tt.month, tt.day, dt.GetDayOfYear(), tt.expected)
		}
	}
}

func TestAddAndSub(t *testing.T) {
	dt, _ := Of(2000, 1, 1, 12, 0, 0, time.UTC)

	// Test Add
	dt2 := dt.Add(24 * time.Hour)
	if dt2.GetDay() != 2 {
		t.Errorf("Add(24h) day = %d, want 2", dt2.GetDay())
	}

	// Test Sub
	dt3 := dt.Sub(24 * time.Hour)
	if dt3.GetMonth() != 13 || dt3.GetYear() != 1999 {
		t.Errorf("Sub(24h) = %d/%d, want 1999/13", dt3.GetYear(), dt3.GetMonth())
	}

	// Test AddDate
	dt4 := dt.AddDate(0, 0, 30)
	if dt4.GetMonth() != 2 || dt4.GetDay() != 1 {
		t.Errorf("AddDate(0, 0, 30) = day %d, month %d, want day 1, month 2", dt4.GetDay(), dt4.GetMonth())
	}
}

func TestDiff(t *testing.T) {
	dt1, _ := Of(2000, 1, 1, 0, 0, 0, time.UTC)
	dt2, _ := Of(2000, 1, 2, 0, 0, 0, time.UTC)

	diff := dt2.Diff(dt1)
	expected := 24 * time.Hour

	if diff != expected {
		t.Errorf("Diff() = %v, want %v", diff, expected)
	}
}

func TestGetDayOfWeek(t *testing.T) {
	// Test a known date: September 12, 2007 (Gregorian) = Meskerem 1, 2000 (Ethiopian) was a Wednesday
	dt, _ := Of(2000, 1, 1, 0, 0, 0, time.UTC)
	dayOfWeek := dt.GetDayOfWeek()

	// Wednesday is day 3 (1=Monday, ..., 7=Sunday)
	expectedDayOfWeek := 3

	if dayOfWeek != expectedDayOfWeek {
		t.Errorf("GetDayOfWeek() = %d, want %d (Wednesday)", dayOfWeek, expectedDayOfWeek)
	}
}

func TestRoundTripConversion(t *testing.T) {
	tests := []struct {
		year  int
		month int
		day   int
	}{
		{2000, 1, 1},
		{2009, 9, 4},
		{1995, 6, 15},
		{2007, 13, 6}, // Leap year Pagume
	}

	for _, tt := range tests {
		// Create Ethiopian date
		dt1, err := Of(tt.year, tt.month, tt.day, 12, 30, 45, time.UTC)
		if err != nil {
			t.Errorf("Of(%d, %d, %d) returned error: %v", tt.year, tt.month, tt.day, err)
			continue
		}

		// Convert to Gregorian
		greg := dt1.ToGregorian()

		// Convert back to Ethiopian
		dt2 := NewDateTime(greg)

		// Check if we get the same date
		if dt2.GetYear() != tt.year || dt2.GetMonth() != tt.month || dt2.GetDay() != tt.day {
			t.Errorf("Round-trip conversion failed for (%d, %d, %d): got (%d, %d, %d)",
				tt.year, tt.month, tt.day, dt2.GetYear(), dt2.GetMonth(), dt2.GetDay())
		}

		// Check time components
		if dt2.GetHour() != 12 || dt2.GetMinute() != 30 || dt2.GetSecond() != 45 {
			t.Errorf("Time components not preserved: got %02d:%02d:%02d, want 12:30:45",
				dt2.GetHour(), dt2.GetMinute(), dt2.GetSecond())
		}
	}
}
