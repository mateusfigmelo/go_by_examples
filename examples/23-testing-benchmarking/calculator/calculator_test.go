package calculator

import (
	"fmt"
	"testing"
)

// TestCalculator runs unit tests for the Calculator type
func TestCalculator(t *testing.T) {
	calc := &Calculator{}

	// Test Add
	t.Run("Add", func(t *testing.T) {
		result := calc.Add(2, 3)
		if result != 5 {
			t.Errorf("Add(2, 3) = %f; want 5", result)
		}
	})

	// Test Subtract
	t.Run("Subtract", func(t *testing.T) {
		result := calc.Subtract(5, 3)
		if result != 2 {
			t.Errorf("Subtract(5, 3) = %f; want 2", result)
		}
	})

	// Test Multiply
	t.Run("Multiply", func(t *testing.T) {
		result := calc.Multiply(4, 3)
		if result != 12 {
			t.Errorf("Multiply(4, 3) = %f; want 12", result)
		}
	})

	// Test Divide
	t.Run("Divide", func(t *testing.T) {
		result, err := calc.Divide(6, 2)
		if err != nil {
			t.Errorf("Divide(6, 2) returned unexpected error: %v", err)
		}
		if result != 3 {
			t.Errorf("Divide(6, 2) = %f; want 3", result)
		}
	})

	// Test Divide by Zero
	t.Run("DivideByZero", func(t *testing.T) {
		_, err := calc.Divide(6, 0)
		if err == nil {
			t.Error("Divide(6, 0) should return an error")
		}
	})

	// Test Memory Operations
	t.Run("Memory", func(t *testing.T) {
		calc.Store(42)
		if calc.Recall() != 42 {
			t.Errorf("Memory recall = %f; want 42", calc.Recall())
		}
		calc.Clear()
		if calc.Recall() != 0 {
			t.Errorf("Memory after clear = %f; want 0", calc.Recall())
		}
	})
}

// TestStringCalculator runs unit tests for the StringCalculator type
func TestStringCalculator(t *testing.T) {
	sc := &StringCalculator{}

	// Table test for Reverse
	reverseTests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
	}

	for _, tt := range reverseTests {
		t.Run("Reverse_"+tt.input, func(t *testing.T) {
			result := sc.Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}

	// Table test for IsPalindrome
	palindromeTests := []struct {
		input    string
		expected bool
	}{
		{"radar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
		{"", true},
		{"a", true},
	}

	for _, tt := range palindromeTests {
		t.Run("IsPalindrome_"+tt.input, func(t *testing.T) {
			result := sc.IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestSliceCalculator runs unit tests for the SliceCalculator type
func TestSliceCalculator(t *testing.T) {
	sc := &SliceCalculator{}

	// Test Sum
	t.Run("Sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := sc.Sum(numbers)
		if result != 15 {
			t.Errorf("Sum(%v) = %d; want 15", numbers, result)
		}
	})

	// Test Average
	t.Run("Average", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result, err := sc.Average(numbers)
		if err != nil {
			t.Errorf("Average(%v) returned unexpected error: %v", numbers, err)
		}
		if result != 3 {
			t.Errorf("Average(%v) = %f; want 3", numbers, result)
		}
	})

	// Test Empty Slice
	t.Run("EmptySlice", func(t *testing.T) {
		numbers := []int{}
		_, err := sc.Average(numbers)
		if err == nil {
			t.Error("Average(empty slice) should return an error")
		}
	})
}

// Benchmarks for Calculator operations
func BenchmarkCalculator(b *testing.B) {
	calc := &Calculator{}

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Add(2, 3)
		}
	})

	b.Run("Power", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Power(2, 10)
		}
	})
}

// Benchmarks for StringCalculator operations
func BenchmarkStringCalculator(b *testing.B) {
	sc := &StringCalculator{}
	longString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"

	b.Run("Reverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sc.Reverse(longString)
		}
	})

	b.Run("IsPalindrome", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sc.IsPalindrome(longString)
		}
	})

	b.Run("Anagrams", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sc.Anagrams("listen", "silent")
		}
	})
}

// Benchmarks for SliceCalculator operations
func BenchmarkSliceCalculator(b *testing.B) {
	sc := &SliceCalculator{}
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}

	b.Run("Sum", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sc.Sum(numbers)
		}
	})

	b.Run("Sort", func(b *testing.B) {
		b.ResetTimer() // Reset the timer to exclude setup time
		for i := 0; i < b.N; i++ {
			sc.Sort(numbers)
		}
	})
}

// Example functions for documentation
func ExampleCalculator_Add() {
	calc := &Calculator{}
	result := calc.Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func ExampleStringCalculator_IsPalindrome() {
	sc := &StringCalculator{}
	result := sc.IsPalindrome("A man a plan a canal Panama")
	fmt.Println(result)
	// Output: true
}
