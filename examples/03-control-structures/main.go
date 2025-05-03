package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Go Control Structures Examples ===")

	// 1. Basic for loop
	fmt.Println("1. Basic for loop (C-style):")
	for i := 0; i < 3; i++ {
		fmt.Printf("  Count: %d\n", i)
	}

	// 2. While-style for loop
	fmt.Println("\n2. While-style for loop:")
	count := 0
	for count < 3 {
		fmt.Printf("  Count: %d\n", count)
		count++
	}

	// 3. Infinite loop with break
	fmt.Println("\n3. Infinite loop with break:")
	count = 0
	for {
		fmt.Printf("  Count: %d\n", count)
		count++
		if count >= 3 {
			break
		}
	}

	// 4. For loop with continue
	fmt.Println("\n4. For loop with continue (print odd numbers):")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("  Odd number: %d\n", i)
	}

	// 5. For-range over slice
	fmt.Println("\n5. For-range over slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, value := range fruits {
		fmt.Printf("  Index: %d, Value: %s\n", index, value)
	}

	// 6. For-range over map
	fmt.Println("\n6. For-range over map:")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	for key, value := range colors {
		fmt.Printf("  Color: %s, Hex: %s\n", key, value)
	}

	// 7. Basic if statement
	fmt.Println("\n7. Basic if statement:")
	number := 7
	if number%2 == 0 {
		fmt.Println("  Number is even")
	} else {
		fmt.Println("  Number is odd")
	}

	// 8. If with initialization
	fmt.Println("\n8. If with initialization:")
	if value := 100; value > 50 {
		fmt.Printf("  Value %d is greater than 50\n", value)
	}

	// 9. If-else if chain
	fmt.Println("\n9. If-else if chain:")
	score := 85
	if score >= 90 {
		fmt.Println("  Grade: A")
	} else if score >= 80 {
		fmt.Println("  Grade: B")
	} else if score >= 70 {
		fmt.Println("  Grade: C")
	} else {
		fmt.Println("  Grade: F")
	}

	// 10. Basic switch
	fmt.Println("\n10. Basic switch:")
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("  Start of work week")
	case "Wednesday":
		fmt.Println("  Middle of work week")
	case "Friday":
		fmt.Println("  End of work week")
	default:
		fmt.Println("  Other day")
	}

	// 11. Switch with multiple cases
	fmt.Println("\n11. Switch with multiple cases:")
	char := 'B'
	switch char {
	case 'a', 'A':
		fmt.Println("  It's an A")
	case 'b', 'B':
		fmt.Println("  It's a B")
	default:
		fmt.Println("  It's something else")
	}

	// 12. Switch with expression
	fmt.Println("\n12. Switch with expression:")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("  Good morning!")
	case hour < 17:
		fmt.Println("  Good afternoon!")
	default:
		fmt.Println("  Good evening!")
	}

	// 13. Switch with fallthrough
	fmt.Println("\n13. Switch with fallthrough:")
	num := 1
	switch num {
	case 1:
		fmt.Println("  One")
		fallthrough
	case 2:
		fmt.Println("  Two")
		fallthrough
	case 3:
		fmt.Println("  Three")
	}

	// 14. Nested control structures
	fmt.Println("\n14. Nested control structures:")
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			fmt.Printf("  First iteration: %d\n", i)
		case 1:
			if i%2 == 0 {
				fmt.Println("  Even number")
			} else {
				fmt.Println("  Odd number")
			}
		default:
			fmt.Printf("  Last iteration: %d\n", i)
		}
	}

	// 15. Label with break
	fmt.Println("\n15. Label with break:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop
			}
			fmt.Printf("  i: %d, j: %d\n", i, j)
		}
	}
}
