package converter

import (
	"errors"
)

// FromJdnConverter converts Julian Day Number (JDN) to Ethiopian date
type FromJdnConverter struct {
	day   int
	month int
	year  int
	jdn   int
}

// NewFromJdnConverter creates a new FromJdnConverter
func NewFromJdnConverter(jdn int) (*FromJdnConverter, error) {
	c := &FromJdnConverter{}
	if err := c.Set(jdn); err != nil {
		return nil, err
	}
	return c, nil
}

// Set sets the JDN for processing
func (c *FromJdnConverter) Set(jdn int) error {
	if !isValidInteger(jdn) {
		return errors.New("invalid JDN")
	}

	c.jdn = jdn
	day, month, year := c.process(jdn)
	c.day = day
	c.month = month
	c.year = year

	return nil
}

// process converts JDN to Ethiopian date
func (c *FromJdnConverter) process(jdn int) (day, month, year int) {
	r := (jdn - 1723856) % 1461
	n := (r % 365) + 365*(r/1460)

	year = 4*((jdn-1723856)/1461) + (r / 365) - (r / 1460)
	month = (n / 30) + 1
	day = (n % 30) + 1

	return day, month, year
}

// isValidInteger checks if the value is a valid integer (not negative in this context)
func isValidInteger(values ...int) bool {
	for _, v := range values {
		if v < 0 {
			return false
		}
	}
	return true
}

// GetJdn returns the Julian Day Number
func (c *FromJdnConverter) GetJdn() int {
	return c.jdn
}

// GetDay returns the day
func (c *FromJdnConverter) GetDay() int {
	return c.day
}

// GetMonth returns the month
func (c *FromJdnConverter) GetMonth() int {
	return c.month
}

// GetYear returns the year
func (c *FromJdnConverter) GetYear() int {
	return c.year
}
