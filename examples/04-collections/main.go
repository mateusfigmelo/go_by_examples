package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Go Collections Examples ===")

	// 1. Arrays
	fmt.Println("1. Arrays:")

	// Array declaration and initialization
	var arr1 [3]int              // Zero-valued array
	arr2 := [3]int{1, 2, 3}      // Array literal
	arr3 := [...]int{4, 5, 6, 7} // Array with inferred size
	arr4 := [5]int{1: 10, 3: 30} // Array with specific index initialization

	fmt.Printf("  Zero-valued array: %v\n", arr1)
	fmt.Printf("  Array literal: %v\n", arr2)
	fmt.Printf("  Array with inferred size: %v (length: %d)\n", arr3, len(arr3))
	fmt.Printf("  Array with specific indices: %v\n", arr4)

	// Array operations
	fmt.Println("\n  Array Operations:")
	arr2[1] = 20 // Modifying element
	fmt.Printf("  Modified array: %v\n", arr2)

	// Multi-dimensional arrays
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("  2D Array: %v\n", matrix)

	// 2. Slices
	fmt.Println("\n2. Slices:")

	// Slice creation
	slice1 := []int{1, 2, 3}    // Slice literal
	slice2 := make([]int, 3)    // Make with length
	slice3 := make([]int, 3, 5) // Make with length and capacity

	fmt.Printf("  Slice literal: %v\n", slice1)
	fmt.Printf("  Made slice: %v\n", slice2)
	fmt.Printf("  Slice with capacity: len=%d cap=%d slice=%v\n",
		len(slice3), cap(slice3), slice3)

	// Slice from array
	arrayForSlice := [5]int{1, 2, 3, 4, 5}
	slice4 := arrayForSlice[1:4] // Slice from index 1 to 3
	fmt.Printf("  Slice from array: %v\n", slice4)

	// Slice operations
	fmt.Println("\n  Slice Operations:")

	// Append
	slice1 = append(slice1, 4) // Append single element
	fmt.Printf("  After append: %v\n", slice1)

	slice1 = append(slice1, 5, 6, 7) // Append multiple elements
	fmt.Printf("  After multiple append: %v\n", slice1)

	// Copy
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("  Copied slice: %v (elements copied: %d)\n", dst, copied)

	// Slicing
	fmt.Println("\n  Slice Slicing:")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("  Original: %v\n", numbers)
	fmt.Printf("  First three: %v\n", numbers[:3])
	fmt.Printf("  Last three: %v\n", numbers[len(numbers)-3:])
	fmt.Printf("  Middle: %v\n", numbers[3:7])

	// 3. Maps
	fmt.Println("\n3. Maps:")

	// Map creation
	map1 := make(map[string]int) // Empty map
	map2 := map[string]int{      // Map literal
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Printf("  Empty map: %v\n", map1)
	fmt.Printf("  Map literal: %v\n", map2)

	// Map operations
	fmt.Println("\n  Map Operations:")

	// Adding and updating
	map1["apple"] = 5
	map1["banana"] = 8
	map1["apple"] = 6 // Updating value
	fmt.Printf("  After adding/updating: %v\n", map1)

	// Accessing and checking existence
	value, exists := map2["two"]
	fmt.Printf("  Value of 'two': %d, exists: %t\n", value, exists)

	value, exists = map2["four"]
	fmt.Printf("  Value of 'four': %d, exists: %t\n", value, exists)

	// Deleting
	delete(map2, "two")
	fmt.Printf("  After delete: %v\n", map2)

	// Length
	fmt.Printf("  Map length: %d\n", len(map2))

	// Iterating over map
	fmt.Println("\n  Map Iteration:")
	for key, value := range map2 {
		fmt.Printf("  Key: %s, Value: %d\n", key, value)
	}

	// 4. Advanced Operations
	fmt.Println("\n4. Advanced Operations:")

	// Sorting slice
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("  Unsorted: %v\n", unsorted)
	sort.Ints(unsorted)
	fmt.Printf("  Sorted: %v\n", unsorted)

	// Sorting slice of strings
	fruits := []string{"banana", "apple", "orange"}
	sort.Strings(fruits)
	fmt.Printf("  Sorted strings: %v\n", fruits)

	// Map keys to slice and sort
	fmt.Println("\n  Sorted Map Keys:")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 20,
	}

	// Get all keys
	var names []string
	for name := range ages {
		names = append(names, name)
	}

	// Sort keys
	sort.Strings(names)

	// Print sorted key-value pairs
	for _, name := range names {
		fmt.Printf("  %s: %d\n", name, ages[name])
	}

	// 5. Practical Examples
	fmt.Println("\n5. Practical Examples:")

	// Using slice as a stack
	fmt.Println("  Stack Operations:")
	stack := []int{}

	// Push
	stack = append(stack, 1) // Push 1
	stack = append(stack, 2) // Push 2
	fmt.Printf("  After push: %v\n", stack)

	// Pop
	x, stack := stack[len(stack)-1], stack[:len(stack)-1]
	fmt.Printf("  Popped value: %d, Remaining stack: %v\n", x, stack)

	// Using map as a frequency counter
	fmt.Println("\n  Frequency Counter:")
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	fmt.Printf("  Word frequencies: %v\n", freq)
}
