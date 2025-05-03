# Command-Line Interface in Go

This example demonstrates three different approaches to handling command-line input in Go:
1. Basic command-line arguments using `os.Args`
2. Command-line flags using the standard `flag` package
3. Command-line subcommands using the `cobra` library

## Project Structure
```
24-cli/
├── main.go    # Main program with CLI examples
├── go.mod     # Module definition and dependencies
└── README.md  # Documentation
```

## Key Components

### 1. Basic Arguments (`os.Args`)
- Direct access to command-line arguments
- Simple string slice handling
- No parsing or validation
- Useful for simple scripts and tools

Example:
```bash
go run main.go arg1 arg2 arg3
```

### 2. Standard Flags (`flag` package)
- Built-in flag parsing
- Type safety for arguments
- Default values
- Usage documentation
- Support for custom types

Example:
```bash
go run main.go -name Alice -age 25 -verbose
```

### 3. Subcommands (`cobra` library)
- Hierarchical command structure
- Nested subcommands
- Command-specific flags
- Automatic help generation
- Middleware support

Example:
```bash
go run main.go get users --all
go run main.go create user --force
```

## Best Practices

1. **Command Structure**
   - Use verbs for commands (get, create, update, delete)
   - Use nouns for resources (user, file, config)
   - Keep command hierarchy logical
   - Provide clear help messages

2. **Flag Naming**
   - Use both short (-v) and long (--verbose) forms
   - Follow common conventions (-h for help)
   - Group related flags
   - Provide meaningful defaults

3. **Error Handling**
   - Validate input early
   - Provide clear error messages
   - Exit with appropriate status codes
   - Handle edge cases

4. **Documentation**
   - Include examples in help text
   - Document all flags and commands
   - Provide usage patterns
   - Include version information

## Common Patterns

1. **Configuration Loading**
   ```go
   config := flag.String("config", "config.yaml", "path to config file")
   flag.Parse()
   ```

2. **Version Flag**
   ```go
   version := flag.Bool("version", false, "print version information")
   if *version {
       fmt.Println("v1.0.0")
       os.Exit(0)
   }
   ```

3. **Required Flags**
   ```go
   name := flag.String("name", "", "name (required)")
   flag.Parse()
   if *name == "" {
       flag.Usage()
       os.Exit(1)
   }
   ```

4. **Custom Types**
   ```go
   type interval []time.Duration
   
   func (i *interval) String() string {
       return fmt.Sprint(*i)
   }
   
   func (i *interval) Set(value string) error {
       duration, err := time.ParseDuration(value)
       if err != nil {
           return err
       }
       *i = append(*i, duration)
       return nil
   }
   ```

## Advanced Topics

1. **Environment Variables**
   - Reading environment variables
   - Overriding flags with env vars
   - Configuration precedence

2. **Interactive Input**
   - Password prompts
   - Confirmation dialogs
   - Progress bars
   - Interactive menus

3. **Output Formatting**
   - Table output
   - JSON/YAML output
   - Color and styling
   - Progress indicators

4. **Shell Completion**
   - Bash completion
   - Zsh completion
   - Fish completion
   - PowerShell completion

## Running the Example

1. **Basic Arguments**
   ```bash
   go run main.go arg1 arg2 arg3
   ```

2. **Standard Flags**
   ```bash
   go run main.go -name Alice -age 25 -verbose
   go run main.go -help
   ```

3. **Cobra Subcommands**
   ```bash
   go run main.go get users --all
   go run main.go create user --force
   go run main.go help get
   ```

## Additional Resources
- [flag package documentation](https://golang.org/pkg/flag/)
- [cobra library documentation](https://github.com/spf13/cobra)
- [12 Factor CLI Apps](https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46)
- [Command Line Interface Guidelines](https://clig.dev/) 
