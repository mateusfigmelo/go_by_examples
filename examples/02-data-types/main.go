package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Numeric Types
	// Integer types
	var int8Val int8 = 127                   // -128 to 127
	var int16Val int16 = 32767               // -32768 to 32767
	var int32Val int32 = 2147483647          // -2147483648 to 2147483647
	var int64Val int64 = 9223372036854775807 // -9223372036854775808 to 9223372036854775807
	var intVal int = 42                      // Platform dependent (32 or 64 bit)

	// Unsigned integer types
	var uint8Val uint8 = 255                    // 0 to 255
	var uint16Val uint16 = 65535                // 0 to 65535
	var uint32Val uint32 = 4294967295           // 0 to 4294967295
	var uint64Val uint64 = 18446744073709551615 // 0 to 18446744073709551615
	var uintVal uint = 42                       // Platform dependent (32 or 64 bit)

	// Floating point types
	var float32Val float32 = 3.14159       // IEEE-754 32-bit
	var float64Val float64 = 3.14159265359 // IEEE-754 64-bit

	// Complex numbers
	var complex64Val complex64 = 3.0 + 4.0i   // Complex numbers with float32 real and imaginary parts
	var complex128Val complex128 = 3.0 + 4.0i // Complex numbers with float64 real and imaginary parts

	// Boolean type
	var boolVal bool = true // true or false

	// String type
	var stringVal string = "Hello, Go!" // UTF-8 encoded string

	// Type inference using :=
	inferredInt := 42              // Type inferred as int
	inferredFloat := 3.14          // Type inferred as float64
	inferredString := "Auto typed" // Type inferred as string

	// Zero values (uninitialized variables)
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	var zeroComplex complex128

	// Printing all values and their types
	fmt.Println("\n=== Basic Types and Values ===")
	fmt.Printf("int8: %v (Type: %T)\n", int8Val, int8Val)
	fmt.Printf("int16: %v (Type: %T)\n", int16Val, int16Val)
	fmt.Printf("int32: %v (Type: %T)\n", int32Val, int32Val)
	fmt.Printf("int64: %v (Type: %T)\n", int64Val, int64Val)
	fmt.Printf("int: %v (Type: %T)\n", intVal, intVal)

	fmt.Println("\n=== Unsigned Integer Types ===")
	fmt.Printf("uint8: %v (Type: %T)\n", uint8Val, uint8Val)
	fmt.Printf("uint16: %v (Type: %T)\n", uint16Val, uint16Val)
	fmt.Printf("uint32: %v (Type: %T)\n", uint32Val, uint32Val)
	fmt.Printf("uint64: %v (Type: %T)\n", uint64Val, uint64Val)
	fmt.Printf("uint: %v (Type: %T)\n", uintVal, uintVal)

	fmt.Println("\n=== Floating Point Types ===")
	fmt.Printf("float32: %v (Type: %T)\n", float32Val, float32Val)
	fmt.Printf("float64: %v (Type: %T)\n", float64Val, float64Val)

	fmt.Println("\n=== Complex Number Types ===")
	fmt.Printf("complex64: %v (Type: %T)\n", complex64Val, complex64Val)
	fmt.Printf("complex128: %v (Type: %T)\n", complex128Val, complex128Val)
	fmt.Printf("complex magnitude: %v\n", cmplx.Abs(complex128Val))

	fmt.Println("\n=== Other Basic Types ===")
	fmt.Printf("bool: %v (Type: %T)\n", boolVal, boolVal)
	fmt.Printf("string: %v (Type: %T)\n", stringVal, stringVal)

	fmt.Println("\n=== Type Inference Examples ===")
	fmt.Printf("inferred int: %v (Type: %T)\n", inferredInt, inferredInt)
	fmt.Printf("inferred float: %v (Type: %T)\n", inferredFloat, inferredFloat)
	fmt.Printf("inferred string: %v (Type: %T)\n", inferredString, inferredString)

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("zero int: %v (Type: %T)\n", zeroInt, zeroInt)
	fmt.Printf("zero float: %v (Type: %T)\n", zeroFloat, zeroFloat)
	fmt.Printf("zero string: %q (Type: %T)\n", zeroString, zeroString)
	fmt.Printf("zero bool: %v (Type: %T)\n", zeroBool, zeroBool)
	fmt.Printf("zero complex: %v (Type: %T)\n", zeroComplex, zeroComplex)

	// Type conversion examples
	fmt.Println("\n=== Type Conversion Examples ===")
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Printf("int to float64: %v -> %v\n", i, f)
	fmt.Printf("float64 to uint: %v -> %v\n", f, u)
}
