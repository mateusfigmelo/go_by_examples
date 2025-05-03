package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Point struct for pointer examples
type Point struct {
	X, Y int
}

// Function that modifies a point through pointer
func movePoint(p *Point, dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Function that demonstrates pointer arithmetic (which is limited in Go)
func swapInts(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println("=== Go Types and Pointers Examples ===")

	// 1. Range over Array
	fmt.Println("1. Range over Array:")
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("  Original array:", arr)

	fmt.Println("  Using range:")
	for i, v := range arr {
		fmt.Printf("    Index: %d, Value: %d\n", i, v)
	}

	// 2. Range over Slice
	fmt.Println("\n2. Range over Slice:")
	slice := []string{"apple", "banana", "cherry"}
	for i, v := range slice {
		fmt.Printf("    Index: %d, Value: %s\n", i, v)
	}

	// 3. Range over Map
	fmt.Println("\n3. Range over Map:")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Printf("    Key: %s, Value: %d\n", k, v)
	}

	// 4. Basic Pointer Operations
	fmt.Println("\n4. Basic Pointer Operations:")
	x := 42
	ptr := &x // Get pointer to x
	fmt.Printf("    Value of x: %d\n", x)
	fmt.Printf("    Address of x: %p\n", ptr)
	fmt.Printf("    Value through pointer: %d\n", *ptr)

	*ptr = 100 // Modify x through pointer
	fmt.Printf("    Modified x through pointer: %d\n", x)

	// 5. Struct Pointers
	fmt.Println("\n5. Struct Pointers:")
	point := Point{X: 10, Y: 20}
	fmt.Printf("    Original point: %+v\n", point)

	movePoint(&point, 5, 5)
	fmt.Printf("    After move: %+v\n", point)

	// 6. Pointer Swap Example
	fmt.Println("\n6. Pointer Swap:")
	a, b := 1, 2
	fmt.Printf("    Before swap: a=%d, b=%d\n", a, b)
	swapInts(&a, &b)
	fmt.Printf("    After swap: a=%d, b=%d\n", a, b)

	// 7. String Operations
	fmt.Println("\n7. String Operations:")
	str := "Hello, 世界" // String with ASCII and Unicode characters
	fmt.Printf("    String: %s\n", str)
	fmt.Printf("    Length (bytes): %d\n", len(str))
	fmt.Printf("    Length (runes): %d\n", utf8.RuneCountInString(str))

	// 8. Range over String
	fmt.Println("\n8. Range over String:")
	for i, r := range str {
		fmt.Printf("    Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	// 9. Rune Operations
	fmt.Println("\n9. Rune Operations:")
	runes := []rune(str)
	fmt.Printf("    Runes: %v\n", runes)
	for i, r := range runes {
		fmt.Printf("    Rune at %d: %c, Unicode: %U\n", i, r, r)
	}

	// 10. String Manipulation with Runes
	fmt.Println("\n10. String Manipulation with Runes:")
	text := "Hello, 世界! 123"
	var letters, numbers, spaces, others int

	for _, r := range text {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsNumber(r):
			numbers++
		case unicode.IsSpace(r):
			spaces++
		default:
			others++
		}
	}

	fmt.Printf("    Text: %s\n", text)
	fmt.Printf("    Letters: %d\n", letters)
	fmt.Printf("    Numbers: %d\n", numbers)
	fmt.Printf("    Spaces: %d\n", spaces)
	fmt.Printf("    Others: %d\n", others)

	// 11. Nil Pointer Handling
	fmt.Println("\n11. Nil Pointer Handling:")
	var nilPtr *int
	fmt.Printf("    Nil pointer value: %v\n", nilPtr)

	if nilPtr == nil {
		fmt.Println("    Pointer is nil")
	}

	// Safe way to handle potentially nil pointer
	if nilPtr != nil {
		fmt.Printf("    Value: %d\n", *nilPtr)
	}

	// 12. Array Pointers
	fmt.Println("\n12. Array Pointers:")
	arr2 := [3]int{1, 2, 3}
	arrPtr := &arr2

	fmt.Printf("    Original array: %v\n", arr2)
	arrPtr[0] = 100 // Syntactic sugar for (*arrPtr)[0]
	fmt.Printf("    Modified through pointer: %v\n", arr2)

	// 13. Unicode Categories
	fmt.Println("\n13. Unicode Categories:")
	text = "Hello, 世界! 12３"
	for _, r := range text {
		fmt.Printf("    Character: %c\n", r)
		fmt.Printf("      Is letter: %t\n", unicode.IsLetter(r))
		fmt.Printf("      Is number: %t\n", unicode.IsNumber(r))
		fmt.Printf("      Is symbol: %t\n", unicode.IsSymbol(r))
		fmt.Printf("      Is punctuation: %t\n", unicode.IsPunct(r))
		fmt.Printf("      Is space: %t\n", unicode.IsSpace(r))
	}
}
