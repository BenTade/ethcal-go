# ethcal-go

[![Go Report Card](https://goreportcard.com/badge/github.com/BenTade/ethcal-go)](https://goreportcard.com/report/github.com/BenTade/ethcal-go)
[![GoDoc](https://godoc.org/github.com/BenTade/ethcal-go?status.svg)](https://godoc.org/github.com/BenTade/ethcal-go)
![From Ethiopia](https://img.shields.io/badge/From-Ethiopia-brightgreen.svg)

> A Go implementation of the [andegna/calender](https://github.com/andegna/calender) package for Ethiopian calendar conversion and manipulation.

**ethcal-go** is a comprehensive Go library for working with the Ethiopian calendar (ዘመን አቆጣጠር). It provides functionality for converting between Gregorian and Ethiopian calendars, date manipulation, formatting, and holiday calculations.

## Features

- ✅ **Calendar Conversion**: Seamlessly convert between Gregorian and Ethiopian calendars
- ✅ **Date Validation**: Validate Ethiopian dates including leap year handling
- ✅ **Date Manipulation**: Add, subtract, and compare dates
- ✅ **Formatting**: Format dates in Amharic with multiple format options
- ✅ **Holiday Calculations**: Calculate Ethiopian Orthodox holidays (Easter)
- ✅ **Julian Day Number (JDN) Support**: Low-level conversion via JDN
- ✅ **Comprehensive Testing**: Well-tested with extensive test coverage

## Installation

```bash
go get github.com/BenTade/ethcal-go
```

## Quick Start

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/BenTade/ethcal-go"
)

func main() {
	// Get current Ethiopian date
	now := ethcal.Now()
	fmt.Println(now.String())
	
	// Create Ethiopian date
	dt, _ := ethcal.Of(2000, 1, 1, 0, 0, 0, time.Local)
	fmt.Printf("Ethiopian Millennium: %s\n", dt.String())
	
	// Convert to Gregorian
	gregorian := dt.ToGregorian()
	fmt.Printf("Gregorian: %s\n", gregorian.Format("January 2, 2006"))
	
	// Format in Amharic
	formatted := dt.Format("l፣ F d ቀን Y ዓ/ም")
	fmt.Printf("Formatted: %s\n", formatted)
}
```

## Table of Contents

- [Basic Usage](#basic-usage)
- [Creating DateTime](#creating-datetime)
- [Conversion](#conversion)
- [Date Manipulation](#date-manipulation)
- [Formatting](#formatting)
- [Validation](#validation)
- [Low-level Conversion (JDN)](#low-level-conversion-jdn)
- [Holidays](#holidays)
- [API Documentation](#api-documentation)

## Basic Usage

### Creating DateTime

**From current time:**

```go
// Current time in Ethiopian calendar
now := ethcal.Now()

// Current time in a specific location
loc, _ := time.LoadLocation("Africa/Addis_Ababa")
nowInAddis := ethcal.NowInLocation(loc)
```

**From Ethiopian date:**

```go
// Create Ethiopian date: Meskerem 1, 2000 (Ethiopian Millennium)
dt, err := ethcal.Of(2000, 1, 1, 0, 0, 0, time.UTC)
if err != nil {
	log.Fatal(err)
}

// With specific time
dt, _ := ethcal.Of(2009, 9, 4, 14, 30, 0, time.Local)
```

**From Gregorian time:**

```go
gregorian := time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC)
ethiopian := ethcal.NewDateTime(gregorian)
fmt.Printf("Year: %d, Month: %d, Day: %d\n", 
	ethiopian.GetYear(), ethiopian.GetMonth(), ethiopian.GetDay())
// Output: Year: 2009, Month: 9, Day: 4
```

**From Unix timestamp:**

```go
timestamp := time.Now().Unix()
dt := ethcal.FromTimestamp(timestamp, time.Local)
```

## Conversion

### Ethiopian to Gregorian

```go
ethiopian, _ := ethcal.Of(2009, 9, 4, 12, 0, 0, time.UTC)
gregorian := ethiopian.ToGregorian()
fmt.Println(gregorian.Format("Monday, January 2, 2006"))
// Output: Friday, May 12, 2017
```

### Gregorian to Ethiopian

```go
gregorian := time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC)
ethiopian := ethcal.NewDateTime(gregorian)
fmt.Printf("%s %d, %d\n", 
	ethiopian.Format("F"), 
	ethiopian.GetDay(), 
	ethiopian.GetYear())
// Output: ግንቦት 4, 2009
```

## Date Manipulation

### Adding and Subtracting

```go
dt, _ := ethcal.Of(2000, 1, 1, 12, 0, 0, time.UTC)

// Add duration
tomorrow := dt.Add(24 * time.Hour)

// Subtract duration
yesterday := dt.Sub(24 * time.Hour)

// Add years, months, days
future := dt.AddDate(1, 2, 15) // Add 1 year, 2 months, 15 days
```

### Calculating Differences

```go
dt1, _ := ethcal.Of(2000, 1, 1, 0, 0, 0, time.UTC)
dt2, _ := ethcal.Of(2000, 1, 10, 0, 0, 0, time.UTC)

diff := dt2.Diff(dt1)
fmt.Printf("Difference: %v\n", diff)
// Output: Difference: 216h0m0s (9 days)
```

## Formatting

The library supports custom date formatting with both standard and Ethiopian-specific format characters.

### Format Characters

| Character | Description | Example |
|-----------|-------------|---------|
| `Y` | Year (4 digits) | 2009 |
| `m` | Month (2 digits) | 09 |
| `d` | Day (2 digits) | 04 |
| `H` | Hour (24-hour, 2 digits) | 14 |
| `i` | Minute (2 digits) | 30 |
| `s` | Second (2 digits) | 45 |
| `F` | Month name in Amharic | ግንቦት |
| `l` | Day name in Amharic | ዓርብ |
| `x` | Orthodox day name | ዮሐንስ |
| `X` | Orthodox year name | ማርቆስ |
| `K` | Year in Geez numbers | ፳፻፱ |
| `V` | Day in Geez numbers | ፬ |
| `E` | Era (ዓ/ም or ዓ/ዓ) | ዓ/ም |
| `A` | Time of day in Amharic | ጡዋት |
| `T` | Timezone | EAT |

### Format Examples

```go
dt, _ := ethcal.Of(2009, 9, 4, 14, 30, 45, time.Local)

// Ethiopian format
fmt.Println(dt.Format("l፣ F d ቀን Y ዓ/ም"))
// Output: ዓርብ፣ ግንቦት 04 ቀን 2009 ዓ/ም

// Geez numbers
fmt.Println(dt.Format("l፣ F V ቀን K ዓ/ም"))
// Output: ዓርብ፣ ግንቦት ፬ ቀን ፳፻፱ ዓ/ም

// With orthodox names
fmt.Println(dt.Format("l፣ F d ቀን (x) Y (X) ዓ/ም"))
// Output: ዓርብ፣ ግንቦት 04 ቀን (ዮሐንስ) 2009 (ማርቆስ) ዓ/ም

// Custom format
fmt.Println(dt.Format("Y-m-d H:i:s"))
// Output: 2009-09-04 14:30:45
```

### Predefined Constants

```go
// Use predefined format constants
fmt.Println(dt.Format(ethcal.DateEthiopian))
fmt.Println(dt.Format(ethcal.DateEthiopianOrthodox))
fmt.Println(dt.Format(ethcal.DateGeez))
fmt.Println(dt.Format(ethcal.DateGeezOrthodox))
```

## Validation

### Validate Ethiopian Dates

```go
import "github.com/BenTade/ethcal-go/validator"

// Validate a date
v := validator.NewDateValidator(4, 9, 2009)
if v.IsValid() {
	fmt.Println("Valid Ethiopian date")
}

// Invalid date
v = validator.NewDateValidator(31, 1, 2000)
if !v.IsValid() {
	fmt.Println("Invalid day: Ethiopian months have max 30 days")
}

// 13th month (Pagume) validation
v = validator.NewDateValidator(6, 13, 2007) // Leap year
if v.IsValid() {
	fmt.Println("Valid: 6th day of Pagume in leap year")
}

v = validator.NewDateValidator(6, 13, 2009) // Non-leap year
if !v.IsValid() {
	fmt.Println("Invalid: 6th day only valid in leap years")
}
```

### Leap Year Validation

```go
import "github.com/BenTade/ethcal-go/validator"

v := validator.NewLeapYearValidator(2007)
if v.IsValid() {
	fmt.Println("2007 is a leap year")
}
// In Ethiopian calendar: (year + 1) % 4 == 0
```

## Low-level Conversion (JDN)

For advanced use cases, you can work directly with Julian Day Numbers (JDN):

### Ethiopian to JDN

```go
import "github.com/BenTade/ethcal-go/converter"

conv, err := converter.NewToJdnConverter(4, 9, 2009)
if err != nil {
	log.Fatal(err)
}
jdn := conv.GetJdn()
fmt.Printf("JDN: %d\n", jdn)
```

### JDN to Ethiopian

```go
import "github.com/BenTade/ethcal-go/converter"

conv, err := converter.NewFromJdnConverter(2457886)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Ethiopian: %d/%d/%d\n", 
	conv.GetDay(), conv.GetMonth(), conv.GetYear())
```

### Practical Example: Convert to Other Calendars

Since most calendar systems can convert to/from JDN, you can use this to convert Ethiopian dates to other calendar systems:

```go
// Ethiopian to Jewish calendar (using Go's extended libraries)
ethiopian, _ := converter.NewToJdnConverter(4, 9, 2009)
jdn := ethiopian.GetJdn()
// Now use jdn with other calendar conversion libraries
```

## Holidays

### Ethiopian Orthodox Easter

```go
import "github.com/BenTade/ethcal-go/holiday"

easter := holiday.NewEaster()

// Get Easter date for Ethiopian year 2009
dt, err := easter.Get(2009)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Easter 2009: %s\n", dt.Format("l፣ F d ቀን Y ዓ/ም"))

// Get Gregorian date of Easter
gregorian, err := easter.GetGregorian(2009)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Easter (Gregorian): %s\n", gregorian.Format("Monday, January 2, 2006"))
```

## API Documentation

### DateTime Methods

```go
// Getters
GetYear() int           // Ethiopian year
GetMonth() int          // Ethiopian month (1-13)
GetDay() int            // Ethiopian day (1-30)
GetHour() int           // Hour (0-23)
GetMinute() int         // Minute (0-59)
GetSecond() int         // Second (0-59)
GetNanosecond() int     // Nanosecond
GetDayOfWeek() int      // Day of week (1=Monday, 7=Sunday)
GetDayOfYear() int      // Day of year (1-366)
GetDaysInMonth() int    // Days in current month
GetTimestamp() int64    // Unix timestamp
IsLeapYear() bool       // Check if leap year

// Conversion
ToGregorian() time.Time  // Convert to Gregorian time.Time

// Manipulation
Add(time.Duration) *DateTime              // Add duration
Sub(time.Duration) *DateTime              // Subtract duration
AddDate(years, months, days int) *DateTime // Add years, months, days
Diff(*DateTime) time.Duration             // Calculate difference

// Formatting
Format(string) string          // Format with custom format
FormatRFC3339() string        // RFC3339 format (Gregorian)
FormatISO8601() string        // ISO8601 format (Gregorian)
String() string               // String representation
```

## Month Names

The Ethiopian calendar has 13 months:

1. መስከረም (Meskerem) - 30 days
2. ጥቅምት (Tikimt) - 30 days
3. ኅዳር (Hidar) - 30 days
4. ታኅሣሥ (Tahsas) - 30 days
5. ጥር (Tir) - 30 days
6. የካቲት (Yekatit) - 30 days
7. መጋቢት (Megabit) - 30 days
8. ሚያዝያ (Miyazya) - 30 days
9. ግንቦት (Ginbot) - 30 days
10. ሰኔ (Sene) - 30 days
11. ሐምሌ (Hamle) - 30 days
12. ነሐሴ (Nehase) - 30 days
13. ጳጉሜን (Pagume) - 5 or 6 days (6 in leap years)

## Leap Years

In the Ethiopian calendar, a year is a leap year if `(year + 1) % 4 == 0`.

Examples:
- 2003, 2007, 2011 are leap years
- 2000, 2001, 2002, 2009, 2010 are not leap years

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Acknowledgments

- Based on the [andegna/calender](https://github.com/andegna/calender) PHP package
- Ethiopian calendar conversion algorithms from [ethiopic.org](http://ethiopic.org/)
- Test cases from [ICU4J](https://github.com/unicode-org/icu/tree/main/icu4j)

## Related Projects

- [andegna/calender](https://github.com/andegna/calender) - Original PHP implementation

---

Made with ❤️ from Ethiopia 🇪🇹