package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

// basicDefer demonstrates basic defer statement usage
func basicDefer() {
	fmt.Println("    Basic Defer Example:")
	defer fmt.Println("        4. This runs last")
	defer fmt.Println("        3. This runs second to last")
	fmt.Println("        1. This runs first")
	fmt.Println("        2. This runs second")
}

// deferWithFiles demonstrates using defer for file handling
func deferWithFiles() {
	fmt.Println("    Defer with Files Example:")

	// Create a temporary file
	f, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("        Error creating file:", err)
		return
	}
	defer func() {
		fmt.Println("        Closing file...")
		f.Close()
		// Clean up the temporary file
		os.Remove("temp.txt")
	}()

	// Write to file
	fmt.Println("        Writing to file...")
	fmt.Fprintln(f, "Hello, World!")
}

// deferWithArguments demonstrates how defer captures argument values
func deferWithArguments() {
	fmt.Println("    Defer with Arguments Example:")

	i := 0
	defer fmt.Printf("        Deferred i: %d\n", i) // Will print 0
	i++
	fmt.Printf("        Current i: %d\n", i) // Will print 1
}

// basicPanic demonstrates a simple panic
func basicPanic() {
	fmt.Println("    Basic Panic Example:")
	defer fmt.Println("        This will still run before the panic")

	fmt.Println("        About to panic...")
	panic("something went wrong!")
}

// recoverFromPanic demonstrates how to recover from a panic
func recoverFromPanic() {
	fmt.Println("    Recover from Panic Example:")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("        Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("        About to panic...")
	panic("deliberate panic!")
}

// customPanicHandler demonstrates a more sophisticated panic recovery
func customPanicHandler() {
	fmt.Println("    Custom Panic Handler Example:")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("        Stack Trace:")
			fmt.Printf("        Panic: %v\n", r)
			fmt.Println("        Stack:")
			debug.PrintStack()
		}
	}()

	deeplyNestedFunction(3)
}

func deeplyNestedFunction(depth int) {
	if depth > 0 {
		deeplyNestedFunction(depth - 1)
	} else {
		panic("reached maximum depth!")
	}
}

// deferChain demonstrates chaining multiple defers
func deferChain() {
	fmt.Println("    Defer Chain Example:")

	fmt.Println("        Starting processing")
	defer fmt.Println("        3. Cleanup step 3")

	// Resource acquisition
	defer fmt.Println("        2. Cleanup step 2")

	// More setup
	defer fmt.Println("        1. Cleanup step 1")

	fmt.Println("        Main processing")
}

// resourceManagement demonstrates proper resource management with defer
type DBConnection struct {
	id string
}

func (db *DBConnection) Close() error {
	fmt.Printf("        Closing DB connection %s\n", db.id)
	return nil
}

func resourceManagement() {
	fmt.Println("    Resource Management Example:")

	// Simulate database connections
	db1 := &DBConnection{"primary"}
	db2 := &DBConnection{"secondary"}

	defer db1.Close()
	defer db2.Close()

	fmt.Println("        Performing database operations...")
	// Simulate some work
	fmt.Println("        Operations completed")
}

// panicInDefer demonstrates panic handling in deferred functions
func panicInDefer() {
	fmt.Println("    Panic in Defer Example:")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("        Recovered from panic in defer: %v\n", r)
		}
	}()

	defer func() {
		fmt.Println("        Second defer about to panic")
		panic("panic in deferred function!")
	}()

	fmt.Println("        Main function executing normally")
}

// deferredMethodCalls demonstrates defer with method calls
type Resource struct {
	name string
}

func (r *Resource) Close() {
	fmt.Printf("        Closing resource: %s\n", r.name)
}

func (r *Resource) Process() {
	fmt.Printf("        Processing resource: %s\n", r.name)
}

func deferredMethodCalls() {
	fmt.Println("    Deferred Method Calls Example:")

	r := &Resource{"example"}
	defer r.Close()

	r.Process()
}

func main() {
	fmt.Println("=== Go Panic, Defer, and Recover Examples ===")

	fmt.Println("\n1. Basic Defer:")
	basicDefer()

	fmt.Println("\n2. Defer with Files:")
	deferWithFiles()

	fmt.Println("\n3. Defer with Arguments:")
	deferWithArguments()

	fmt.Println("\n4. Defer Chain:")
	deferChain()

	fmt.Println("\n5. Resource Management:")
	resourceManagement()

	fmt.Println("\n6. Deferred Method Calls:")
	deferredMethodCalls()

	fmt.Println("\n7. Recover from Panic:")
	recoverFromPanic()

	fmt.Println("\n8. Panic in Defer:")
	panicInDefer()

	fmt.Println("\n9. Custom Panic Handler:")
	customPanicHandler()

	// Note: basicPanic() is not called in main as it would terminate the program
	fmt.Println("\nProgram completed successfully!")
}
