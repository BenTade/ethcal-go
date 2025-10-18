package converter

import (
	"testing"
)

// Test data from the original PHP package
// https://github.com/svn2github/icu4j/blob/master/main/tests/core/src/com/ibm/icu/dev/test/calendar/EthiopicTest.java

func TestToJdnConverter(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
		jdn   int
	}{
		// Test cases from ICU4J
		{20, 2, 1855, 2401443},
		{29, 10, 1857, 2402423},
		{22, 5, 1858, 2402631},
		{10, 8, 1858, 2402709},
		{28, 4, 1859, 2402972},
		{5, 5, 1860, 2403345},

		// More test cases
		{1, 1, 0, 1723856},
		{1, 1, 1, 1724221},
		{1, 1, 2, 1724586},
		{1, 1, 3, 1724951},
		{1, 1, 4, 1725317},
		{5, 13, 0, 1724220},
		{5, 13, 1, 1724585},
		{5, 13, 2, 1724950},
		{5, 13, 3, 1725315},
		{6, 13, 3, 1725316},
		{5, 13, 4, 1725681},

		// Additional test cases
		{6, 2, 1575, 2299159},
		{7, 2, 1575, 2299160},
		{8, 2, 1575, 2299161},
		{9, 2, 1575, 2299162},

		{23, 4, 1892, 2415021},
		{23, 4, 1997, 2453372},
		{5, 13, 2000, 2454720},

		{22, 4, 1893, 2415385},
		{22, 4, 1985, 2448988},
		{22, 4, 1989, 2450449},
		{22, 4, 1993, 2451910},
		{22, 4, 1997, 2453371},

		{14, 4, 2993, 2817152},
		{7, 4, 3993, 3182395},
		{22, 3, 5993, 3912880},
	}

	for _, tt := range tests {
		conv, err := NewToJdnConverter(tt.day, tt.month, tt.year)
		if err != nil {
			t.Errorf("NewToJdnConverter(%d, %d, %d) returned error: %v", tt.day, tt.month, tt.year, err)
			continue
		}

		if conv.GetJdn() != tt.jdn {
			t.Errorf("ToJdnConverter(%d, %d, %d) = %d, want %d", tt.day, tt.month, tt.year, conv.GetJdn(), tt.jdn)
		}
	}
}

func TestFromJdnConverter(t *testing.T) {
	tests := []struct {
		jdn   int
		day   int
		month int
		year  int
	}{
		// Test cases from ICU4J
		{2401443, 20, 2, 1855},
		{2402423, 29, 10, 1857},
		{2402631, 22, 5, 1858},
		{2402709, 10, 8, 1858},
		{2402972, 28, 4, 1859},
		{2403345, 5, 5, 1860},

		// More test cases
		{1723856, 1, 1, 0},
		{1724221, 1, 1, 1},
		{1724586, 1, 1, 2},
		{1724951, 1, 1, 3},
		{1725317, 1, 1, 4},
		{1724220, 5, 13, 0},
		{1724585, 5, 13, 1},
		{1724950, 5, 13, 2},
		{1725315, 5, 13, 3},
		{1725316, 6, 13, 3},
		{1725681, 5, 13, 4},

		// Additional test cases
		{2299159, 6, 2, 1575},
		{2299160, 7, 2, 1575},
		{2299161, 8, 2, 1575},
		{2299162, 9, 2, 1575},

		{2415021, 23, 4, 1892},
		{2453372, 23, 4, 1997},
		{2454720, 5, 13, 2000},

		{2415385, 22, 4, 1893},
		{2448988, 22, 4, 1985},
		{2450449, 22, 4, 1989},
		{2451910, 22, 4, 1993},
		{2453371, 22, 4, 1997},

		{2817152, 14, 4, 2993},
		{3182395, 7, 4, 3993},
		{3912880, 22, 3, 5993},
	}

	for _, tt := range tests {
		conv, err := NewFromJdnConverter(tt.jdn)
		if err != nil {
			t.Errorf("NewFromJdnConverter(%d) returned error: %v", tt.jdn, err)
			continue
		}

		if conv.GetDay() != tt.day || conv.GetMonth() != tt.month || conv.GetYear() != tt.year {
			t.Errorf("FromJdnConverter(%d) = (%d, %d, %d), want (%d, %d, %d)",
				tt.jdn, conv.GetDay(), conv.GetMonth(), conv.GetYear(), tt.day, tt.month, tt.year)
		}
	}
}

func TestRoundTripConversion(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
	}{
		{1, 1, 2000},
		{15, 6, 1995},
		{30, 12, 2010},
		{5, 13, 2007},
		{1, 1, 1},
		{22, 4, 1997},
	}

	for _, tt := range tests {
		// Convert to JDN
		toConv, err := NewToJdnConverter(tt.day, tt.month, tt.year)
		if err != nil {
			t.Errorf("NewToJdnConverter(%d, %d, %d) returned error: %v", tt.day, tt.month, tt.year, err)
			continue
		}
		jdn := toConv.GetJdn()

		// Convert back to Ethiopian
		fromConv, err := NewFromJdnConverter(jdn)
		if err != nil {
			t.Errorf("NewFromJdnConverter(%d) returned error: %v", jdn, err)
			continue
		}

		if fromConv.GetDay() != tt.day || fromConv.GetMonth() != tt.month || fromConv.GetYear() != tt.year {
			t.Errorf("Round-trip conversion failed for (%d, %d, %d): got (%d, %d, %d)",
				tt.day, tt.month, tt.year, fromConv.GetDay(), fromConv.GetMonth(), fromConv.GetYear())
		}
	}
}

func TestInvalidDate(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
	}{
		{0, 1, 2000},
		{31, 1, 2000},
		{1, 0, 2000},
		{1, 14, 2000},
		{6, 13, 2000}, // 6th of Pagume is only valid in leap years
		{7, 13, 2000},
	}

	for _, tt := range tests {
		_, err := NewToJdnConverter(tt.day, tt.month, tt.year)
		if err == nil {
			t.Errorf("NewToJdnConverter(%d, %d, %d) should return error for invalid date", tt.day, tt.month, tt.year)
		}
	}
}

func TestInvalidJDN(t *testing.T) {
	tests := []int{-1, -100}

	for _, jdn := range tests {
		_, err := NewFromJdnConverter(jdn)
		if err == nil {
			t.Errorf("NewFromJdnConverter(%d) should return error for invalid JDN", jdn)
		}
	}
}
