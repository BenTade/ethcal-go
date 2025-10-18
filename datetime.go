package ethcal

import (
	"time"

	"github.com/BenTade/ethcal-go/converter"
	"github.com/BenTade/ethcal-go/validator"
)

// DateTime represents an Ethiopian calendar date with time
type DateTime struct {
	gregorianTime time.Time
	year          int
	month         int
	day           int
	leapYear      bool
	dayOfYear     int
	daysInMonth   int
}

// NewDateTime creates a new DateTime from a Gregorian time.Time
func NewDateTime(t time.Time) *DateTime {
	dt := &DateTime{
		gregorianTime: t,
	}
	dt.updateComputedFields()
	return dt
}

// Now creates a DateTime representing the current time
func Now() *DateTime {
	return NewDateTime(time.Now())
}

// NowInLocation creates a DateTime representing the current time in a specific location
func NowInLocation(loc *time.Location) *DateTime {
	return NewDateTime(time.Now().In(loc))
}

// Of creates a DateTime from Ethiopian date components
func Of(year, month, day, hour, minute, second int, loc *time.Location) (*DateTime, error) {
	// Convert Ethiopian date to JDN
	conv, err := converter.NewToJdnConverter(day, month, year)
	if err != nil {
		return nil, err
	}

	jdn := conv.GetJdn()

	// Convert JDN to Gregorian
	gregorian := jdnToGregorian(jdn)

	// Create time with the specified time components
	if loc == nil {
		loc = time.Local
	}
	t := time.Date(gregorian.year, time.Month(gregorian.month), gregorian.day, hour, minute, second, 0, loc)

	return NewDateTime(t), nil
}

// FromTimestamp creates a DateTime from a Unix timestamp
func FromTimestamp(timestamp int64, loc *time.Location) *DateTime {
	if loc == nil {
		loc = time.Local
	}
	return NewDateTime(time.Unix(timestamp, 0).In(loc))
}

// FromConverter creates a DateTime from a converter
func FromConverter(conv *converter.FromJdnConverter, loc *time.Location) (*DateTime, error) {
	return Of(conv.GetYear(), conv.GetMonth(), conv.GetDay(), 0, 0, 0, loc)
}

// updateComputedFields updates the Ethiopian date fields from the Gregorian time
func (dt *DateTime) updateComputedFields() {
	// Convert Gregorian to JDN
	jdn := gregorianToJdn(dt.gregorianTime.Year(), int(dt.gregorianTime.Month()), dt.gregorianTime.Day())

	// Convert JDN to Ethiopian
	conv, _ := converter.NewFromJdnConverter(jdn)
	dt.year = conv.GetYear()
	dt.month = conv.GetMonth()
	dt.day = conv.GetDay()

	// Calculate other fields
	dt.leapYear = validator.NewLeapYearValidator(dt.year).IsValid()
	dt.dayOfYear = (dt.month-1)*30 + dt.day
	if dt.month == 13 {
		dt.daysInMonth = 5
		if dt.leapYear {
			dt.daysInMonth = 6
		}
	} else {
		dt.daysInMonth = 30
	}
}

// GetYear returns the Ethiopian year
func (dt *DateTime) GetYear() int {
	return dt.year
}

// GetMonth returns the Ethiopian month
func (dt *DateTime) GetMonth() int {
	return dt.month
}

// GetDay returns the Ethiopian day
func (dt *DateTime) GetDay() int {
	return dt.day
}

// IsLeapYear returns true if the year is a leap year
func (dt *DateTime) IsLeapYear() bool {
	return dt.leapYear
}

// GetDayOfYear returns the day of the year
func (dt *DateTime) GetDayOfYear() int {
	return dt.dayOfYear
}

// GetDaysInMonth returns the number of days in the month
func (dt *DateTime) GetDaysInMonth() int {
	return dt.daysInMonth
}

// GetHour returns the hour
func (dt *DateTime) GetHour() int {
	return dt.gregorianTime.Hour()
}

// GetMinute returns the minute
func (dt *DateTime) GetMinute() int {
	return dt.gregorianTime.Minute()
}

// GetSecond returns the second
func (dt *DateTime) GetSecond() int {
	return dt.gregorianTime.Second()
}

// GetNanosecond returns the nanosecond
func (dt *DateTime) GetNanosecond() int {
	return dt.gregorianTime.Nanosecond()
}

// GetDayOfWeek returns the day of the week (1 for Monday, 7 for Sunday)
func (dt *DateTime) GetDayOfWeek() int {
	weekday := int(dt.gregorianTime.Weekday())
	if weekday == 0 {
		return 7 // Sunday
	}
	return weekday
}

// GetTimestamp returns the Unix timestamp
func (dt *DateTime) GetTimestamp() int64 {
	return dt.gregorianTime.Unix()
}

// ToGregorian returns a copy of the underlying Gregorian time
func (dt *DateTime) ToGregorian() time.Time {
	return dt.gregorianTime
}

// Add adds a duration to the DateTime
func (dt *DateTime) Add(d time.Duration) *DateTime {
	newTime := dt.gregorianTime.Add(d)
	return NewDateTime(newTime)
}

// Sub subtracts a duration from the DateTime
func (dt *DateTime) Sub(d time.Duration) *DateTime {
	newTime := dt.gregorianTime.Add(-d)
	return NewDateTime(newTime)
}

// AddDate adds years, months, and days to the DateTime
func (dt *DateTime) AddDate(years, months, days int) *DateTime {
	newTime := dt.gregorianTime.AddDate(years, months, days)
	return NewDateTime(newTime)
}

// Diff returns the duration between two DateTimes
func (dt *DateTime) Diff(other *DateTime) time.Duration {
	return dt.gregorianTime.Sub(other.gregorianTime)
}

// Helper types for conversion
type gregorianDate struct {
	year  int
	month int
	day   int
}

// jdnToGregorian converts JDN to Gregorian date
func jdnToGregorian(jdn int) gregorianDate {
	a := jdn + 32044
	b := (4*a + 3) / 146097
	c := a - (146097*b)/4

	d := (4*c + 3) / 1461
	e := c - (1461*d)/4
	m := (5*e + 2) / 153

	day := e - (153*m+2)/5 + 1
	month := m + 3 - 12*(m/10)
	year := 100*b + d - 4800 + m/10

	return gregorianDate{year: year, month: month, day: day}
}

// gregorianToJdn converts Gregorian date to JDN
func gregorianToJdn(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3

	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}
