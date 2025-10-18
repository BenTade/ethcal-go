package holiday

import (
	"github.com/BenTade/ethcal-go"
	"time"
)

// Easter calculates Ethiopian Orthodox Easter dates
type Easter struct{}

// NewEaster creates a new Easter calculator
func NewEaster() *Easter {
	return &Easter{}
}

// Get calculates the Easter date for a given Ethiopian year
// Based on the Computus calculation for Ethiopian Orthodox Easter
func (e *Easter) Get(year int) (*ethcal.DateTime, error) {
	// Ethiopian Easter calculation
	// The formula is based on the Alexandrian computus
	
	// Calculate the Golden Number
	goldenNumber := (year % 19) + 1
	
	// Calculate the Epact
	epact := (11 * goldenNumber) % 30
	
	// Calculate the full moon (14th day of the lunar month)
	fullMoon := 21 + epact
	if fullMoon > 50 {
		fullMoon -= 30
	}
	
	// Easter is the first Sunday after the full moon
	// that occurs on or after the vernal equinox (March 21 in Gregorian, Megabit 13 in Ethiopian)
	
	// For Ethiopian calendar, Easter falls in Miyazya (8th month)
	// The calculation gives us the day in Miyazya
	
	// Simplified calculation for Ethiopian Easter
	// It typically falls between Miyazya 1 and Miyazya 23
	month := 8 // Miyazya
	day := ((19 * (year % 19)) + 15) % 30
	
	// Adjust the calculation
	if day == 0 {
		day = 30
		month = 7 // Megabit
	}
	
	// Fine-tune based on day of week
	// This is a simplified version; the actual calculation is more complex
	dt, err := ethcal.Of(year, month, day, 0, 0, 0, time.Local)
	if err != nil {
		return nil, err
	}
	
	// Ensure it's a Sunday
	dayOfWeek := dt.GetDayOfWeek()
	if dayOfWeek != 7 { // 7 is Sunday
		daysToSunday := (7 - dayOfWeek) % 7
		dt = dt.AddDate(0, 0, daysToSunday)
	}
	
	return dt, nil
}

// GetGregorian returns the Gregorian date of Easter for a given Ethiopian year
func (e *Easter) GetGregorian(year int) (time.Time, error) {
	dt, err := e.Get(year)
	if err != nil {
		return time.Time{}, err
	}
	return dt.ToGregorian(), nil
}
