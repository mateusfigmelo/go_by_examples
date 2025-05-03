package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Basic command-line arguments example
func demoBasicArgs() {
	fmt.Println("\n=== Basic Arguments ===")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [arguments]")
		fmt.Println("Example: go run main.go arg1 arg2 arg3")
		return
	}

	fmt.Println("All arguments:", os.Args)
	fmt.Println("Program name:", os.Args[0])
	fmt.Println("Arguments:", os.Args[1:])
}

// Standard flag package example
func demoStandardFlags() {
	fmt.Println("\n=== Standard Flags ===")

	// Reset the flag package to allow multiple runs
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Define flags
	name := flag.String("name", "World", "name to greet")
	age := flag.Int("age", 0, "age of the person")
	verbose := flag.Bool("verbose", false, "enable verbose output")
	var color string
	flag.StringVar(&color, "color", "blue", "favorite color")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: go run main.go -name Alice -age 25 -verbose\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Use the flags
	fmt.Printf("Hello, %s!\n", *name)
	if *age > 0 {
		fmt.Printf("Age: %d\n", *age)
	}
	if *verbose {
		fmt.Printf("Verbose mode enabled\n")
		fmt.Printf("Favorite color: %s\n", color)
		fmt.Printf("Remaining args: %v\n", flag.Args())
	}
}

// Cobra library example for subcommands
func demoCobraCommands() {
	fmt.Println("\n=== Cobra Subcommands ===")

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A sample CLI application",
		Long:  "A sample CLI application demonstrating subcommands and flags",
	}

	var getCmd = &cobra.Command{
		Use:   "get [resource]",
		Short: "Get a resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Getting resource: %s\n", args[0])
			all, _ := cmd.Flags().GetBool("all")
			if all {
				fmt.Println("Fetching all items")
			}
		},
	}

	var createCmd = &cobra.Command{
		Use:   "create [resource]",
		Short: "Create a resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Creating resource: %s\n", args[0])
			force, _ := cmd.Flags().GetBool("force")
			if force {
				fmt.Println("Force flag enabled")
			}
		},
	}

	// Add flags to commands
	getCmd.Flags().BoolP("all", "a", false, "get all resources")
	createCmd.Flags().BoolP("force", "f", false, "force creation")

	// Add commands to root
	rootCmd.AddCommand(getCmd, createCmd)

	// Example usage instructions
	fmt.Println("Example usage:")
	fmt.Println("  go run main.go get users --all")
	fmt.Println("  go run main.go create user --force")

	// Execute the cobra command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("=== Command-Line Interface Examples ===")
	fmt.Println("\nThis example demonstrates three ways to handle command-line input in Go:")
	fmt.Println("1. Basic arguments (os.Args)")
	fmt.Println("2. Standard flags (flag package)")
	fmt.Println("3. Subcommands (cobra library)")

	// Demo all three approaches
	demoBasicArgs()
	demoStandardFlags()
	demoCobraCommands()

	fmt.Println("\nTry running the program with different arguments and flags!")
}
