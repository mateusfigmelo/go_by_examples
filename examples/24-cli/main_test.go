package main

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestBasicArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "no arguments",
			args: []string{"program"},
			want: []string{
				"Usage: go run main.go [arguments]",
				"Example: go run main.go arg1 arg2 arg3",
			},
		},
		{
			name: "with arguments",
			args: []string{"program", "arg1", "arg2", "arg3"},
			want: []string{
				"All arguments:",
				"Program name: program",
				"Arguments:",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args
			output := testutil.CaptureOutput(func() {
				// Reset flags before running
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
				demoBasicArgs()
			})

			for _, want := range tt.want {
				if !strings.Contains(output, want) {
					t.Errorf("demoBasicArgs() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestStandardFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "default values",
			args: []string{"program"},
			want: []string{
				"Hello, World!",
			},
		},
		{
			name: "custom name",
			args: []string{"program", "-name", "Alice"},
			want: []string{
				"Hello, Alice!",
			},
		},
		{
			name: "all flags",
			args: []string{"program", "-name", "Bob", "-age", "25", "-verbose", "-color", "red"},
			want: []string{
				"Hello, Bob!",
				"Age: 25",
				"Verbose mode enabled",
				"Favorite color: red",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args
			output := testutil.CaptureOutput(func() {
				// Reset flags before running
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
				demoStandardFlags()
			})

			for _, want := range tt.want {
				if !strings.Contains(output, want) {
					t.Errorf("demoStandardFlags() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestCobraCommands(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "get command",
			args: []string{"program", "get", "users"},
			want: []string{
				"Getting resource: users",
			},
		},
		{
			name: "get with all flag",
			args: []string{"program", "get", "users", "--all"},
			want: []string{
				"Getting resource: users",
				"Fetching all items",
			},
		},
		{
			name: "create command",
			args: []string{"program", "create", "user"},
			want: []string{
				"Creating resource: user",
			},
		},
		{
			name: "create with force flag",
			args: []string{"program", "create", "user", "--force"},
			want: []string{
				"Creating resource: user",
				"Force flag enabled",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = tt.args
			output := testutil.CaptureOutput(func() {
				// Reset flags before running
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
				demoCobraCommands()
			})

			for _, want := range tt.want {
				if !strings.Contains(output, want) {
					t.Errorf("demoCobraCommands() missing output %q in %q", want, output)
				}
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	// Save original args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"program"}
	output := testutil.CaptureOutput(func() {
		// Reset flags before running
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		main()
	})

	wants := []string{
		"=== Command-Line Interface Examples ===",
		"This example demonstrates three ways to handle command-line input in Go:",
		"1. Basic arguments (os.Args)",
		"2. Standard flags (flag package)",
		"3. Subcommands (cobra library)",
	}

	for _, want := range wants {
		if !strings.Contains(output, want) {
			t.Errorf("main() missing output %q in %q", want, output)
		}
	}
}
