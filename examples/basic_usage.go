package main

import (
	"fmt"
	"time"

	"github.com/BenTade/ethcal-go"
	"github.com/BenTade/ethcal-go/converter"
	"github.com/BenTade/ethcal-go/holiday"
	"github.com/BenTade/ethcal-go/validator"
)

func main() {
	fmt.Println("=== Ethiopian Calendar (ethcal-go) Examples ===")
	fmt.Println()

	// Example 1: Get current Ethiopian date
	fmt.Println("1. Current Ethiopian Date:")
	now := ethcal.Now()
	fmt.Printf("   Ethiopian: %s\n", now.String())
	fmt.Printf("   Gregorian: %s\n", now.ToGregorian().Format("January 2, 2006 15:04:05"))
	fmt.Printf("   Year: %d, Month: %d, Day: %d\n\n", now.GetYear(), now.GetMonth(), now.GetDay())

	// Example 2: Create Ethiopian date (Millennium - Meskerem 1, 2000)
	fmt.Println("2. Ethiopian Millennium (Meskerem 1, 2000):")
	millennium, _ := ethcal.Of(2000, 1, 1, 0, 0, 0, time.UTC)
	fmt.Printf("   Ethiopian: %s\n", millennium.Format("l፣ F d ቀን Y ዓ/ም"))
	fmt.Printf("   Gregorian: %s\n", millennium.ToGregorian().Format("Monday, January 2, 2006"))
	fmt.Println()

	// Example 3: Convert Gregorian to Ethiopian
	fmt.Println("3. Convert Gregorian to Ethiopian:")
	gregorian := time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC)
	ethiopian := ethcal.NewDateTime(gregorian)
	fmt.Printf("   Gregorian: %s\n", gregorian.Format("January 2, 2006"))
	fmt.Printf("   Ethiopian: %s\n", ethiopian.Format("F d, Y"))
	fmt.Printf("   Formatted: %s\n\n", ethiopian.Format("l፣ F d ቀን Y ዓ/ም"))

	// Example 4: Date manipulation
	fmt.Println("4. Date Manipulation:")
	dt, _ := ethcal.Of(2000, 1, 1, 12, 0, 0, time.UTC)
	fmt.Printf("   Original: %s\n", dt.Format("F d, Y H:i:s"))
	tomorrow := dt.Add(24 * time.Hour)
	fmt.Printf("   Add 1 day: %s\n", tomorrow.Format("F d, Y H:i:s"))
	nextMonth := dt.AddDate(0, 1, 0)
	fmt.Printf("   Add 1 month: %s\n", nextMonth.Format("F d, Y H:i:s"))
	fmt.Println()

	// Example 5: Formatting with different styles
	fmt.Println("5. Various Date Formats:")
	dt2, _ := ethcal.Of(2009, 9, 4, 14, 30, 45, time.Local)
	fmt.Printf("   Basic: %s\n", dt2.Format("Y-m-d H:i:s"))
	fmt.Printf("   Ethiopian: %s\n", dt2.Format("l፣ F d ቀን Y ዓ/ም"))
	fmt.Printf("   Geez: %s\n", dt2.Format("l፣ F V ቀን K ዓ/ም"))
	fmt.Printf("   Orthodox: %s\n", dt2.Format("l፣ F d ቀን (x) Y (X) ዓ/ም"))
	fmt.Println()

	// Example 6: Leap year checking
	fmt.Println("6. Leap Year Validation:")
	leapYears := []int{2003, 2007, 2009, 2011}
	for _, year := range leapYears {
		v := validator.NewLeapYearValidator(year)
		status := "not a leap year"
		if v.IsValid() {
			status = "LEAP YEAR"
		}
		fmt.Printf("   Year %d: %s\n", year, status)
	}
	fmt.Println()

	// Example 7: Date validation
	fmt.Println("7. Date Validation:")
	testDates := []struct {
		day   int
		month int
		year  int
	}{
		{4, 9, 2009},
		{31, 1, 2000},
		{6, 13, 2007},
		{6, 13, 2009},
	}
	for _, td := range testDates {
		v := validator.NewDateValidator(td.day, td.month, td.year)
		status := "invalid"
		if v.IsValid() {
			status = "valid"
		}
		fmt.Printf("   %d/%d/%d: %s\n", td.day, td.month, td.year, status)
	}
	fmt.Println()

	// Example 8: JDN conversion
	fmt.Println("8. Julian Day Number Conversion:")
	conv, _ := converter.NewToJdnConverter(4, 9, 2009)
	jdn := conv.GetJdn()
	fmt.Printf("   Ethiopian 4/9/2009 = JDN %d\n", jdn)
	
	backConv, _ := converter.NewFromJdnConverter(jdn)
	fmt.Printf("   JDN %d = Ethiopian %d/%d/%d\n\n",
		jdn, backConv.GetDay(), backConv.GetMonth(), backConv.GetYear())

	// Example 9: Easter calculation
	fmt.Println("9. Ethiopian Orthodox Easter:")
	easter := holiday.NewEaster()
	for _, year := range []int{2006, 2007, 2008, 2009, 2010} {
		easterDate, err := easter.Get(year)
		if err == nil {
			fmt.Printf("   Easter %d: %s\n", year, easterDate.Format("l፣ F d ቀን"))
		}
	}
	fmt.Println()

	// Example 10: Month and day names
	fmt.Println("10. Ethiopian Month Names:")
	for i := 1; i <= 13; i++ {
		dt, _ := ethcal.Of(2009, i, 1, 0, 0, 0, time.UTC)
		fmt.Printf("    Month %2d: %s\n", i, dt.Format("F"))
	}
	fmt.Println()

	fmt.Println("11. Ethiopian Day Names:")
	// Create a week starting from Monday
	startDate, _ := ethcal.Of(2009, 9, 1, 0, 0, 0, time.UTC)
	for i := 0; i < 7; i++ {
		dt := startDate.AddDate(0, 0, i)
		fmt.Printf("    Day %d: %s\n", dt.GetDayOfWeek(), dt.Format("l"))
	}

	fmt.Println("\n=== Examples Complete ===")
}
