package calculator

import (
	"errors"
	"math"
	"sort"
	"strings"
)

// Calculator represents a basic calculator with memory
type Calculator struct {
	memory float64
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Power returns a raised to the power of b
func (c *Calculator) Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// Store saves a number in memory
func (c *Calculator) Store(value float64) {
	c.memory = value
}

// Recall returns the number from memory
func (c *Calculator) Recall() float64 {
	return c.memory
}

// Clear resets the memory
func (c *Calculator) Clear() {
	c.memory = 0
}

// StringCalculator performs operations on strings
type StringCalculator struct{}

// Reverse returns the reversed string
func (sc *StringCalculator) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func (sc *StringCalculator) IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == sc.Reverse(s)
}

// Anagrams checks if two strings are anagrams
func (sc *StringCalculator) Anagrams(s1, s2 string) bool {
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	if len(s1) != len(s2) {
		return false
	}

	chars1 := strings.Split(s1, "")
	chars2 := strings.Split(s2, "")

	sort.Strings(chars1)
	sort.Strings(chars2)

	return strings.Join(chars1, "") == strings.Join(chars2, "")
}

// SliceCalculator performs operations on slices
type SliceCalculator struct{}

// Sum returns the sum of all numbers in a slice
func (sc *SliceCalculator) Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Average returns the average of all numbers in a slice
func (sc *SliceCalculator) Average(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}
	sum := sc.Sum(numbers)
	return float64(sum) / float64(len(numbers)), nil
}

// Max returns the maximum number in a slice
func (sc *SliceCalculator) Max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max, nil
}

// Sort returns a sorted copy of the slice
func (sc *SliceCalculator) Sort(numbers []int) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)
	return sorted
}
