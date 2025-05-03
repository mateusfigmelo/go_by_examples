package main

import (
	"fmt"
	"math"
)

// Direction represents a type-based enum
type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

// String method for Direction enum
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// Shape interface defines behavior for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Point struct represents a 2D point
type Point struct {
	X, Y float64
}

// Move method for Point
func (p *Point) Move(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

// String method for Point
func (p Point) String() string {
	return fmt.Sprintf("Point(%.2f, %.2f)", p.X, p.Y)
}

// Circle struct embeds a Point
type Circle struct {
	Center Point
	Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// String method for Circle
func (c Circle) String() string {
	return fmt.Sprintf("Circle(center: %v, radius: %.2f)", c.Center, c.Radius)
}

// Rectangle struct with two points
type Rectangle struct {
	TopLeft, BottomRight Point
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	width := math.Abs(r.BottomRight.X - r.TopLeft.X)
	height := math.Abs(r.TopLeft.Y - r.BottomRight.Y)
	return width * height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	width := math.Abs(r.BottomRight.X - r.TopLeft.X)
	height := math.Abs(r.TopLeft.Y - r.BottomRight.Y)
	return 2 * (width + height)
}

// String method for Rectangle
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(topLeft: %v, bottomRight: %v)", r.TopLeft, r.BottomRight)
}

// Person struct demonstrates nested structs and tags
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}

// Birthday method for Person
func (p *Person) Birthday() {
	p.Age++
}

// String method for Person
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old) from %s", p.Name, p.Age, p.Address.City)
}

// Logger interface demonstrates interface composition
type Logger interface {
	Log(message string)
}

// Writer interface demonstrates interface composition
type Writer interface {
	Write(message string)
}

// LogWriter interface composes Logger and Writer
type LogWriter interface {
	Logger
	Writer
}

// ConsoleLogger implements LogWriter
type ConsoleLogger struct {
	prefix string
}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] LOG: %s\n", c.prefix, message)
}

func (c ConsoleLogger) Write(message string) {
	fmt.Printf("[%s] WRITE: %s\n", c.prefix, message)
}

// Status represents a type-based enum with custom behavior
type Status int

const (
	StatusPending Status = iota
	StatusActive
	StatusInactive
)

// String method for Status
func (s Status) String() string {
	return [...]string{"Pending", "Active", "Inactive"}[s]
}

// IsActive method for Status
func (s Status) IsActive() bool {
	return s == StatusActive
}

func main() {
	fmt.Println("=== Go Structs, Methods, and Interfaces Example ===")

	// 1. Working with Points
	fmt.Println("\n1. Points:")
	p1 := Point{X: 1, Y: 2}
	fmt.Printf("    Original point: %v\n", p1)
	p1.Move(2, 3)
	fmt.Printf("    After move: %v\n", p1)

	// 2. Shapes Interface
	fmt.Println("\n2. Shapes Interface:")
	shapes := []Shape{
		Circle{Center: Point{X: 0, Y: 0}, Radius: 5},
		Rectangle{
			TopLeft:     Point{X: 0, Y: 5},
			BottomRight: Point{X: 5, Y: 0},
		},
	}

	for i, shape := range shapes {
		fmt.Printf("    Shape %d:\n", i+1)
		fmt.Printf("        Type: %T\n", shape)
		fmt.Printf("        String: %v\n", shape)
		fmt.Printf("        Area: %.2f\n", shape.Area())
		fmt.Printf("        Perimeter: %.2f\n", shape.Perimeter())
	}

	// 3. Person Struct with Nested Fields
	fmt.Println("\n3. Person Struct:")
	person := Person{
		Name: "Alice",
		Age:  25,
	}
	person.Address.Street = "123 Main St"
	person.Address.City = "New York"
	person.Address.Country = "USA"

	fmt.Printf("    Initial: %v\n", person)
	person.Birthday()
	fmt.Printf("    After birthday: %v\n", person)

	// 4. Direction Enum
	fmt.Println("\n4. Direction Enum:")
	directions := []Direction{North, East, South, West}
	for _, d := range directions {
		fmt.Printf("    %d: %v\n", d, d)
	}

	// 5. Interface Composition
	fmt.Println("\n5. Interface Composition:")
	logger := ConsoleLogger{prefix: "APP"}
	logger.Log("This is a log message")
	logger.Write("This is a write message")

	// 6. Status Enum with Behavior
	fmt.Println("\n6. Status Enum:")
	statuses := []Status{StatusPending, StatusActive, StatusInactive}
	for _, s := range statuses {
		fmt.Printf("    %v (Active: %v)\n", s, s.IsActive())
	}

	// 7. Interface Type Assertions
	fmt.Println("\n7. Interface Type Assertions:")
	var shape Shape = Circle{Center: Point{X: 0, Y: 0}, Radius: 5}

	if circle, ok := shape.(Circle); ok {
		fmt.Printf("    Circle radius: %.2f\n", circle.Radius)
	}

	// Type switch example
	switch v := shape.(type) {
	case Circle:
		fmt.Printf("    Circle with radius %.2f\n", v.Radius)
	case Rectangle:
		fmt.Printf("    Rectangle with width %.2f\n",
			math.Abs(v.BottomRight.X-v.TopLeft.X))
	default:
		fmt.Printf("    Unknown shape type\n")
	}
}
