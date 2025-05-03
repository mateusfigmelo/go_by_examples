package main

import (
	"fmt"
	"strings"
)

// 1. Basic function with parameters and return value
func add(x, y int) int {
	return x + y
}

// 2. Multiple return values
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

// 3. Named return values
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return uses named return values
}

// 4. Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 5. Function with a function parameter
func applyFunc(x int, fn func(int) int) int {
	return fn(x)
}

// 6. Recursive function - Factorial
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 7. Recursive function with multiple returns - Fibonacci
func fibonacci(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative index")
	}
	if n < 2 {
		return uint64(n), nil
	}
	n1, _ := fibonacci(n - 1)
	n2, _ := fibonacci(n - 2)
	return n1 + n2, nil
}

// 8. Function that returns a function (closure)
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 9. Function with deferred function call
func processWithDefer(data string) string {
	// Deferred functions are executed in LIFO order
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	return strings.ToUpper(data)
}

// 10. Function that modifies its input
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	fmt.Println("=== Go Functions Examples ===")

	// 1. Basic function
	fmt.Println("1. Basic Function:")
	result := add(5, 3)
	fmt.Printf("  add(5, 3) = %d\n", result)

	// 2. Multiple return values
	fmt.Println("\n2. Multiple Return Values:")
	if result, err := divide(10, 2); err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  10 / 2 = %.2f\n", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  10 / 0 = %.2f\n", result)
	}

	// 3. Named return values
	fmt.Println("\n3. Named Return Values:")
	area, perim := rectangle(3.0, 4.0)
	fmt.Printf("  Rectangle (3x4) - Area: %.2f, Perimeter: %.2f\n", area, perim)

	// 4. Variadic function
	fmt.Println("\n4. Variadic Function:")
	fmt.Printf("  sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("  sum(%v...) = %d\n", numbers, sum(numbers...))

	// 5. Function as parameter
	fmt.Println("\n5. Function as Parameter:")
	double := func(x int) int { return x * 2 }
	triple := func(x int) int { return x * 3 }

	fmt.Printf("  applyFunc(5, double) = %d\n", applyFunc(5, double))
	fmt.Printf("  applyFunc(5, triple) = %d\n", applyFunc(5, triple))

	// 6. Recursive function
	fmt.Println("\n6. Recursive Function (Factorial):")
	fmt.Printf("  factorial(5) = %d\n", factorial(5))
	fmt.Printf("  factorial(0) = %d\n", factorial(0))

	// 7. Recursive with error handling
	fmt.Println("\n7. Recursive with Error Handling (Fibonacci):")
	if fib, err := fibonacci(10); err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  fibonacci(10) = %d\n", fib)
	}

	// 8. Closure
	fmt.Println("\n8. Closure:")
	times2 := makeMultiplier(2)
	times3 := makeMultiplier(3)
	fmt.Printf("  times2(5) = %d\n", times2(5))
	fmt.Printf("  times3(5) = %d\n", times3(5))

	// 9. Deferred function calls
	fmt.Println("\n9. Deferred Function Calls:")
	resultStr := processWithDefer("hello")
	fmt.Printf("  Result: %s\n", resultStr)

	// 10. Function modifying slice
	fmt.Println("\n10. Function Modifying Slice:")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("  Before: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("  After:  %v\n", slice)

	// 11. Anonymous function
	fmt.Println("\n11. Anonymous Function:")
	func() {
		fmt.Println("  Executed immediately")
	}()

	// 12. Closure capturing local variables
	fmt.Println("\n12. Closure Capturing Variables:")
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Printf("  First call: %d\n", increment())
	fmt.Printf("  Second call: %d\n", increment())
	fmt.Printf("  Third call: %d\n", increment())
}
