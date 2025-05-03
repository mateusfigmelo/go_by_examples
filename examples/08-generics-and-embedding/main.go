package main

import (
	"fmt"
	"strings"
)

// Generic Stack implementation
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Generic Queue implementation
type Queue[T any] struct {
	items []T
}

// Enqueue adds an item to the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Generic constraints example
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Generic function to find minimum value
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic function to find maximum value
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic Map function
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

// Generic Filter function
func Filter[T any](items []T, f func(T) bool) []T {
	var result []T
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Embedding example: Vehicle hierarchy
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v Vehicle) Description() string {
	return fmt.Sprintf("%d %s %s", v.Year, v.Brand, v.Model)
}

// Car embeds Vehicle
type Car struct {
	Vehicle
	NumDoors   int
	IsElectric bool
}

func (c Car) String() string {
	carType := "Gasoline"
	if c.IsElectric {
		carType = "Electric"
	}
	return fmt.Sprintf("%s - %s car with %d doors",
		c.Description(), carType, c.NumDoors)
}

// Truck embeds Vehicle
type Truck struct {
	Vehicle
	CargoCapacity float64
}

func (t Truck) String() string {
	return fmt.Sprintf("%s - Cargo capacity: %.2f tons",
		t.Description(), t.CargoCapacity)
}

// Database example with embedding
type Logger struct {
	prefix string
}

func (l Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.prefix, message)
}

// Database embeds Logger
type Database struct {
	Logger
	connection string
}

func (d Database) Connect() {
	d.Log(fmt.Sprintf("Connecting to %s", d.connection))
}

func (d Database) Query(query string) {
	d.Log(fmt.Sprintf("Executing query: %s", query))
}

// Iterator interface for generic iteration
type Iterator[T any] interface {
	Next() bool
	Value() T
	Reset()
}

// SliceIterator implements Iterator for slices
type SliceIterator[T any] struct {
	slice []T
	pos   int
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: slice,
		pos:   -1,
	}
}

func (it *SliceIterator[T]) Next() bool {
	it.pos++
	return it.pos < len(it.slice)
}

func (it *SliceIterator[T]) Value() T {
	return it.slice[it.pos]
}

func (it *SliceIterator[T]) Reset() {
	it.pos = -1
}

// RangeIterator generates a sequence of numbers
type RangeIterator struct {
	start, end, current int
}

func NewRangeIterator(start, end int) *RangeIterator {
	return &RangeIterator{
		start:   start,
		end:     end,
		current: start - 1,
	}
}

func (it *RangeIterator) Next() bool {
	it.current++
	return it.current < it.end
}

func (it *RangeIterator) Value() int {
	return it.current
}

func (it *RangeIterator) Reset() {
	it.current = it.start - 1
}

// FilterIterator creates a filtered view of another iterator
type FilterIterator[T any] struct {
	source Iterator[T]
	pred   func(T) bool
	value  T
}

func NewFilterIterator[T any](source Iterator[T], pred func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		source: source,
		pred:   pred,
	}
}

func (it *FilterIterator[T]) Next() bool {
	for it.source.Next() {
		value := it.source.Value()
		if it.pred(value) {
			it.value = value
			return true
		}
	}
	return false
}

func (it *FilterIterator[T]) Value() T {
	return it.value
}

func (it *FilterIterator[T]) Reset() {
	it.source.Reset()
}

// MapIterator transforms values from another iterator
type MapIterator[T, U any] struct {
	source Iterator[T]
	mapper func(T) U
}

func NewMapIterator[T, U any](source Iterator[T], mapper func(T) U) *MapIterator[T, U] {
	return &MapIterator[T, U]{
		source: source,
		mapper: mapper,
	}
}

func (it *MapIterator[T, U]) Next() bool {
	return it.source.Next()
}

func (it *MapIterator[T, U]) Value() U {
	return it.mapper(it.source.Value())
}

func (it *MapIterator[T, U]) Reset() {
	it.source.Reset()
}

// Generic function to collect iterator values into a slice
func Collect[T any](it Iterator[T]) []T {
	var result []T
	for it.Next() {
		result = append(result, it.Value())
	}
	it.Reset()
	return result
}

func main() {
	fmt.Println("=== Go Generics and Embedding Example ===")

	// 1. Generic Stack
	fmt.Println("\n1. Generic Stack:")
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Print("    Stack contents: ")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	// 2. Generic Queue
	fmt.Println("\n2. Generic Queue:")
	stringQueue := &Queue[string]{}
	stringQueue.Enqueue("First")
	stringQueue.Enqueue("Second")
	stringQueue.Enqueue("Third")

	fmt.Print("    Queue contents: ")
	for {
		if val, ok := stringQueue.Dequeue(); ok {
			fmt.Printf("%s ", val)
		} else {
			break
		}
	}
	fmt.Println()

	// 3. Generic Min/Max
	fmt.Println("\n3. Generic Min/Max:")
	fmt.Printf("    Min(5, 3): %d\n", Min(5, 3))
	fmt.Printf("    Max(5.5, 3.3): %.1f\n", Max(5.5, 3.3))

	// 4. Generic Map/Filter
	fmt.Println("\n4. Generic Map/Filter:")
	numbers := []int{1, 2, 3, 4, 5}

	// Map: Double each number
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("    Doubled numbers: %v\n", doubled)

	// Filter: Keep even numbers
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("    Even numbers: %v\n", evens)

	// Map: Convert numbers to strings
	strNumbers := Map(numbers, func(x int) string {
		return fmt.Sprintf("Num-%d", x)
	})
	fmt.Printf("    String numbers: %v\n", strNumbers)

	// 5. Vehicle Embedding
	fmt.Println("\n5. Vehicle Embedding:")
	car := Car{
		Vehicle: Vehicle{
			Brand: "Tesla",
			Model: "Model 3",
			Year:  2023,
		},
		NumDoors:   4,
		IsElectric: true,
	}
	fmt.Printf("    Car: %v\n", car)

	truck := Truck{
		Vehicle: Vehicle{
			Brand: "Volvo",
			Model: "FH16",
			Year:  2022,
		},
		CargoCapacity: 32.5,
	}
	fmt.Printf("    Truck: %v\n", truck)

	// 6. Database with Logger Embedding
	fmt.Println("\n6. Database with Logger:")
	db := Database{
		Logger:     Logger{prefix: "DB"},
		connection: "postgresql://localhost:5432",
	}
	db.Connect()
	db.Query("SELECT * FROM users")

	// 7. Generic Data Processing
	fmt.Println("\n7. Generic Data Processing:")
	words := []string{"hello", "world", "golang", "generics"}

	// Filter words longer than 5 characters
	longWords := Filter(words, func(s string) bool {
		return len(s) > 5
	})
	fmt.Printf("    Long words: %v\n", longWords)

	// Map words to uppercase
	upperWords := Map(words, strings.ToUpper)
	fmt.Printf("    Uppercase words: %v\n", upperWords)

	// 8. Iterator Examples
	fmt.Println("\n8. Iterator Examples:")

	// Slice Iterator
	fmt.Println("    a) Slice Iterator:")
	numbers = []int{1, 2, 3, 4, 5}
	numIter := NewSliceIterator(numbers)
	fmt.Print("        Basic iteration: ")
	for numIter.Next() {
		fmt.Printf("%d ", numIter.Value())
	}
	fmt.Println()

	// Range Iterator
	fmt.Println("    b) Range Iterator:")
	rangeIter := NewRangeIterator(1, 5)
	fmt.Print("        Range from 1 to 5: ")
	for rangeIter.Next() {
		fmt.Printf("%d ", rangeIter.Value())
	}
	fmt.Println()

	// Filter Iterator
	fmt.Println("    c) Filter Iterator:")
	numIter.Reset()
	evenIter := NewFilterIterator(numIter, func(n int) bool {
		return n%2 == 0
	})
	fmt.Print("        Even numbers: ")
	for evenIter.Next() {
		fmt.Printf("%d ", evenIter.Value())
	}
	fmt.Println()

	// Map Iterator
	fmt.Println("    d) Map Iterator:")
	numIter.Reset()
	squareIter := NewMapIterator(numIter, func(n int) int {
		return n * n
	})
	fmt.Print("        Squared numbers: ")
	for squareIter.Next() {
		fmt.Printf("%d ", squareIter.Value())
	}
	fmt.Println()

	// Chaining Iterators
	fmt.Println("    e) Chaining Iterators:")
	rangeIter.Reset()
	// Filter even numbers and square them
	chainedIter := NewMapIterator(
		NewFilterIterator(rangeIter, func(n int) bool { return n%2 == 0 }),
		func(n int) int { return n * n },
	)
	fmt.Print("        Even numbers squared: ")
	fmt.Printf("%v\n", Collect(chainedIter))

	// String Iterator Example
	fmt.Println("    f) String Iterator:")
	words = []string{"hello", "world", "golang"}
	wordIter := NewSliceIterator(words)
	upperIter := NewMapIterator(wordIter, strings.ToUpper)
	fmt.Print("        Uppercase words: ")
	fmt.Printf("%v\n", Collect(upperIter))
}
