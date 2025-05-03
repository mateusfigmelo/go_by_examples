package main

import (
	"fmt"
	"time"
)

// basicTimeOperations demonstrates basic time operations
func basicTimeOperations() {
	fmt.Println("    Basic Time Operations:")

	// Current time
	now := time.Now()
	fmt.Printf("        Current time: %v\n", now)
	fmt.Printf("        Unix epoch: %d\n", now.Unix())
	fmt.Printf("        Unix nano: %d\n", now.UnixNano())

	// Time components
	fmt.Printf("        Year: %d\n", now.Year())
	fmt.Printf("        Month: %s\n", now.Month())
	fmt.Printf("        Day: %d\n", now.Day())
	fmt.Printf("        Hour: %d\n", now.Hour())
	fmt.Printf("        Minute: %d\n", now.Minute())
	fmt.Printf("        Second: %d\n", now.Second())
	fmt.Printf("        Nanosecond: %d\n", now.Nanosecond())
	fmt.Printf("        Weekday: %s\n", now.Weekday())
	fmt.Printf("        Day of year: %d\n", now.YearDay())

	// Time zones
	location, _ := time.LoadLocation("America/New_York")
	nyTime := now.In(location)
	fmt.Printf("        New York time: %v\n", nyTime)
}

// timeArithmetic demonstrates time arithmetic operations
func timeArithmetic() {
	fmt.Println("    Time Arithmetic:")

	now := time.Now()

	// Adding durations
	future := now.Add(24 * time.Hour)
	fmt.Printf("        Tomorrow: %v\n", future)

	// Subtracting times
	duration := future.Sub(now)
	fmt.Printf("        Duration: %v\n", duration)

	// Adding components
	nextMonth := now.AddDate(0, 1, 0)
	fmt.Printf("        Next month: %v\n", nextMonth)

	// Time comparison
	fmt.Printf("        Is future after now? %v\n", future.After(now))
	fmt.Printf("        Is now before future? %v\n", now.Before(future))
	fmt.Printf("        Are they equal? %v\n", now.Equal(future))

	// Rounding times
	rounded := now.Round(time.Hour)
	truncated := now.Truncate(time.Hour)
	fmt.Printf("        Rounded to hour: %v\n", rounded)
	fmt.Printf("        Truncated to hour: %v\n", truncated)
}

// epochOperations demonstrates working with Unix epochs
func epochOperations() {
	fmt.Println("    Epoch Operations:")

	// Current epoch
	now := time.Now()
	epoch := now.Unix()
	nanoEpoch := now.UnixNano()

	fmt.Printf("        Current epoch: %d\n", epoch)
	fmt.Printf("        Current nano epoch: %d\n", nanoEpoch)

	// Creating time from epoch
	epochTime := time.Unix(epoch, 0)
	nanoTime := time.Unix(0, nanoEpoch)

	fmt.Printf("        Time from epoch: %v\n", epochTime)
	fmt.Printf("        Time from nano epoch: %v\n", nanoTime)

	// Working with specific epochs
	y2k := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("        Y2K epoch: %d\n", y2k.Unix())

	// Future epoch calculations
	future := time.Date(2038, time.January, 19, 3, 14, 7, 0, time.UTC)
	fmt.Printf("        32-bit epoch limit: %d\n", future.Unix())
}

// timeFormatting demonstrates time formatting and parsing
func timeFormatting() {
	fmt.Println("    Time Formatting and Parsing:")

	now := time.Now()

	// Standard formats
	fmt.Printf("        ANSIC: %v\n", now.Format(time.ANSIC))
	fmt.Printf("        UnixDate: %v\n", now.Format(time.UnixDate))
	fmt.Printf("        RFC822: %v\n", now.Format(time.RFC822))
	fmt.Printf("        RFC3339: %v\n", now.Format(time.RFC3339))
	fmt.Printf("        Kitchen: %v\n", now.Format(time.Kitchen))

	// Custom formats
	fmt.Printf("        Custom (DD-MM-YYYY): %v\n", now.Format("02-01-2006"))
	fmt.Printf("        Custom (MM/DD/YY): %v\n", now.Format("01/02/06"))
	fmt.Printf("        Custom (HH:mm:ss): %v\n", now.Format("15:04:05"))

	// Parsing times
	timeStr := "2024-03-15 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("        Error parsing time: %v\n", err)
	} else {
		fmt.Printf("        Parsed time: %v\n", parsedTime)
	}

	// Parsing with location
	loc, _ := time.LoadLocation("Europe/London")
	parsedWithLoc, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("        Error parsing time with location: %v\n", err)
	} else {
		fmt.Printf("        Parsed time (London): %v\n", parsedWithLoc)
	}
}

// timeZoneHandling demonstrates working with time zones
func timeZoneHandling() {
	fmt.Println("    Time Zone Handling:")

	now := time.Now()

	// Load different time zones
	locations := []string{
		"America/New_York",
		"Europe/London",
		"Asia/Tokyo",
		"Australia/Sydney",
		"UTC",
	}

	for _, loc := range locations {
		location, err := time.LoadLocation(loc)
		if err != nil {
			fmt.Printf("        Error loading location %s: %v\n", loc, err)
			continue
		}

		localTime := now.In(location)
		fmt.Printf("        Time in %s: %v\n", loc, localTime)

		// Zone information
		name, offset := localTime.Zone()
		fmt.Printf("        Zone: %s, Offset: %d hours\n", name, offset/3600)
	}

	// DST information
	ny, _ := time.LoadLocation("America/New_York")
	summer := time.Date(2024, time.July, 1, 0, 0, 0, 0, ny)
	winter := time.Date(2024, time.January, 1, 0, 0, 0, 0, ny)

	fmt.Printf("\n        DST in New York:\n")
	fmt.Printf("        Summer: %v, Zone: %s\n", summer, summer.Format("MST"))
	fmt.Printf("        Winter: %v, Zone: %s\n", winter, winter.Format("MST"))
}

func main() {
	fmt.Println("=== Go Time Handling Examples ===")

	fmt.Println("\n1. Basic Time Operations:")
	basicTimeOperations()

	fmt.Println("\n2. Time Arithmetic:")
	timeArithmetic()

	fmt.Println("\n3. Epoch Operations:")
	epochOperations()

	fmt.Println("\n4. Time Formatting and Parsing:")
	timeFormatting()

	fmt.Println("\n5. Time Zone Handling:")
	timeZoneHandling()
}
