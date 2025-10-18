package converter

import (
	"errors"
	"github.com/BenTade/ethcal-go/validator"
)

// ToJdnConverter converts Ethiopian dates to Julian Day Number (JDN)
type ToJdnConverter struct {
	day   int
	month int
	year  int
	jdn   int
}

// NewToJdnConverter creates a new ToJdnConverter and validates the date
func NewToJdnConverter(day, month, year int) (*ToJdnConverter, error) {
	c := &ToJdnConverter{}
	if err := c.Set(day, month, year); err != nil {
		return nil, err
	}
	return c, nil
}

// Set sets the date for processing
func (c *ToJdnConverter) Set(day, month, year int) error {
	v := validator.NewDateValidator(day, month, year)
	if !v.IsValid() {
		return errors.New("invalid Ethiopian date")
	}

	c.day = day
	c.month = month
	c.year = year
	c.jdn = c.process(day, month, year)

	return nil
}

// process calculates the JDN from Ethiopian date
func (c *ToJdnConverter) process(day, month, year int) int {
	return (1723856 + 365) +
		365*(year-1) +
		(year / 4) +
		30*month +
		day - 31
}

// GetJdn returns the Julian Day Number
func (c *ToJdnConverter) GetJdn() int {
	return c.jdn
}

// GetDay returns the day
func (c *ToJdnConverter) GetDay() int {
	return c.day
}

// GetMonth returns the month
func (c *ToJdnConverter) GetMonth() int {
	return c.month
}

// GetYear returns the year
func (c *ToJdnConverter) GetYear() int {
	return c.year
}
