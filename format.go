package ethcal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Format formats the DateTime according to the given format string
// Supports standard Go time format plus Ethiopian-specific formats:
// x - Orthodox day name
// X - Orthodox year name
// E - Era (ዓ/ም or ዓ/ዓ)
// K - Year in Geez numbers
// V - Day in Geez numbers
// F - Month name in Amharic
// l - Day name in Amharic
func (dt *DateTime) Format(format string) string {
	// First, handle Ethiopian-specific format characters
	result := format

	// Replace Ethiopian format characters with placeholders
	replacements := map[string]string{
		"x": dt.getOrthodoxDayName(),
		"X": dt.getOrthodoxYearName(),
		"E": EraAM,
		"K": dt.getYearInGeez(),
		"V": dt.getDayInGeez(),
		"F": dt.getMonthName(),
		"l": dt.getDayName(),
		"Y": strconv.Itoa(dt.year),
		"m": fmt.Sprintf("%02d", dt.month),
		"d": fmt.Sprintf("%02d", dt.day),
		"H": fmt.Sprintf("%02d", dt.GetHour()),
		"i": fmt.Sprintf("%02d", dt.GetMinute()),
		"s": fmt.Sprintf("%02d", dt.GetSecond()),
		"A": dt.getTimeOfDay(),
		"T": dt.gregorianTime.Format("MST"),
	}

	// We need to be careful about the order of replacements
	// to avoid replacing parts of already-replaced strings
	for key := range replacements {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, key, placeholder)
	}

	for key, value := range replacements {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}

// getMonthName returns the Amharic name of the month
func (dt *DateTime) getMonthName() string {
	if dt.month >= 1 && dt.month <= len(MonthNames)-1 {
		return MonthNames[dt.month]
	}
	return ""
}

// getDayName returns the Amharic name of the day
func (dt *DateTime) getDayName() string {
	dayOfWeek := dt.GetDayOfWeek()
	if dayOfWeek >= 1 && dayOfWeek <= len(DayNames)-1 {
		return DayNames[dayOfWeek]
	}
	return ""
}

// getOrthodoxDayName returns the Orthodox saint day name
func (dt *DateTime) getOrthodoxDayName() string {
	if dt.day >= 1 && dt.day <= len(OrthodoxDayNames)-1 {
		return OrthodoxDayNames[dt.day]
	}
	return ""
}

// getOrthodoxYearName returns the Orthodox year name (Evangelist)
func (dt *DateTime) getOrthodoxYearName() string {
	index := (dt.year - 1) % 4
	if index >= 0 && index < len(OrthodoxYearNames) {
		return OrthodoxYearNames[index]
	}
	return ""
}

// getYearInGeez returns the year in Geez numbers
func (dt *DateTime) getYearInGeez() string {
	return numberToGeez(dt.year)
}

// getDayInGeez returns the day in Geez numbers
func (dt *DateTime) getDayInGeez() string {
	if dt.day >= 1 && dt.day < len(GeezNumbers) {
		return GeezNumbers[dt.day]
	}
	return numberToGeez(dt.day)
}

// getTimeOfDay returns the Amharic time of day
func (dt *DateTime) getTimeOfDay() string {
	hour := dt.GetHour()
	switch {
	case hour == 0:
		return TimeOfDayNames["midnight"]
	case hour < 6:
		return TimeOfDayNames["night"]
	case hour < 12:
		return TimeOfDayNames["morning"]
	case hour == 12:
		return TimeOfDayNames["noon"]
	case hour < 17:
		return TimeOfDayNames["afternoon"]
	case hour < 20:
		return TimeOfDayNames["evening"]
	default:
		return TimeOfDayNames["night"]
	}
}

// numberToGeez converts a number to Geez representation
func numberToGeez(n int) string {
	if n >= 1 && n < len(GeezNumbers) {
		return GeezNumbers[n]
	}

	// For larger numbers, build the representation
	geez := []rune{
		' ', '፩', '፪', '፫', '፬', '፭', '፮', '፯', '፰', '፱',
	}
	tens := []rune{
		' ', '፲', '፳', '፴', '፵', '፶', '፷', '፸', '፹', '፺',
	}

	result := ""
	
	// Handle thousands
	if n >= 1000 {
		thousands := n / 1000
		if thousands == 1 {
			result += "፲፻"
		} else if thousands < 10 {
			result += string(geez[thousands]) + "፲፻"
		} else {
			// Recursively handle larger thousands
			result += numberToGeez(thousands) + "፻"
		}
		n %= 1000
	}

	// Handle hundreds
	if n >= 100 {
		hundredsDigit := n / 100
		if hundredsDigit == 1 {
			result += "፻"
		} else if hundredsDigit < 10 {
			result += string(geez[hundredsDigit]) + "፻"
		}
		n %= 100
	}

	// Handle tens
	if n >= 10 {
		tensDigit := n / 10
		if tensDigit < len(tens) {
			result += string(tens[tensDigit])
		}
		n %= 10
	}

	// Handle ones
	if n > 0 && n < len(geez) {
		result += string(geez[n])
	}

	return result
}

// FormatRFC3339 formats the DateTime in RFC3339 format (Gregorian)
func (dt *DateTime) FormatRFC3339() string {
	return dt.gregorianTime.Format(time.RFC3339)
}

// FormatISO8601 formats the DateTime in ISO8601 format (Gregorian)
func (dt *DateTime) FormatISO8601() string {
	return dt.gregorianTime.Format("2006-01-02T15:04:05Z07:00")
}

// String returns a string representation of the DateTime
func (dt *DateTime) String() string {
	return fmt.Sprintf("%s %d, %d %02d:%02d:%02d",
		dt.getMonthName(),
		dt.day,
		dt.year,
		dt.GetHour(),
		dt.GetMinute(),
		dt.GetSecond())
}
