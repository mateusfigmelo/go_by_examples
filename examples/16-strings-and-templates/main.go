package main

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// stringFunctions demonstrates various string manipulation functions
func stringFunctions() {
	fmt.Println("    String Functions Example:")

	// Basic string operations
	str := "  Hello, World!  "
	fmt.Printf("        Original: %q\n", str)
	fmt.Printf("        Trimmed: %q\n", strings.TrimSpace(str))
	fmt.Printf("        Upper: %q\n", strings.ToUpper(str))
	fmt.Printf("        Lower: %q\n", strings.ToLower(str))

	// String splitting and joining
	words := strings.Split("apple,banana,cherry", ",")
	fmt.Printf("        Split: %#v\n", words)
	joined := strings.Join(words, " | ")
	fmt.Printf("        Joined: %q\n", joined)

	// String replacement
	s := "hello hello hello"
	replaced := strings.Replace(s, "hello", "hi", 2) // Replace first 2 occurrences
	fmt.Printf("        Replace: %q\n", replaced)
	replacedAll := strings.ReplaceAll(s, "hello", "hi")
	fmt.Printf("        ReplaceAll: %q\n", replacedAll)

	// String contains and index functions
	str = "Go programming is fun"
	fmt.Printf("        Contains 'Go': %v\n", strings.Contains(str, "Go"))
	fmt.Printf("        Index of 'programming': %d\n", strings.Index(str, "programming"))
	fmt.Printf("        HasPrefix 'Go': %v\n", strings.HasPrefix(str, "Go"))
	fmt.Printf("        HasSuffix 'fun': %v\n", strings.HasSuffix(str, "fun"))
}

// stringFormatting demonstrates various string formatting options
func stringFormatting() {
	fmt.Println("    String Formatting Example:")

	// Basic formatting
	name := "Alice"
	age := 30
	height := 1.75
	fmt.Printf("        Basic: Name=%s, Age=%d, Height=%.2f\n", name, age, height)

	// Width and alignment
	fmt.Printf("        Right-aligned: |%10s|\n", name)
	fmt.Printf("        Left-aligned:  |%-10s|\n", name)
	fmt.Printf("        Zero-padded:   |%05d|\n", age)

	// Different bases
	num := 42
	fmt.Printf("        Different bases: Decimal=%d, Binary=%b, Octal=%o, Hex=%x\n",
		num, num, num, num)

	// Struct formatting
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 25}
	fmt.Printf("        Struct default: %v\n", p)
	fmt.Printf("        Struct with fields: %+v\n", p)
	fmt.Printf("        Struct with syntax: %#v\n", p)

	// String and quote formatting
	special := "Hello\tWorld\n"
	fmt.Printf("        Quoted string: %q\n", special)
	fmt.Printf("        Backquoted string: %#q\n", special)
}

// User represents a user for template examples
type User struct {
	Name    string
	Age     int
	Hobbies []string
}

// textTemplates demonstrates text/template package usage
func textTemplates() {
	fmt.Println("    Text Templates Example:")

	// Simple template
	tmpl := `User Profile:
    Name: {{.Name}}
    Age: {{.Age}}
    Hobbies: {{range .Hobbies}}
        - {{.}}{{end}}
`

	user := User{
		Name:    "Alice",
		Age:     30,
		Hobbies: []string{"Reading", "Hiking", "Photography"},
	}

	// Parse and execute template
	t := template.Must(template.New("profile").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, user); err != nil {
		fmt.Printf("        Template error: %v\n", err)
		return
	}
	fmt.Print("        ", strings.ReplaceAll(buf.String(), "\n", "\n        "))

	// Template with functions
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"join":  strings.Join,
	}

	tmplWithFuncs := `
Modified Profile:
    NAME: {{upper .Name}}
    Hobbies (joined): {{join .Hobbies ", "}}`

	t = template.Must(template.New("profile_funcs").Funcs(funcMap).Parse(tmplWithFuncs))
	buf.Reset()
	if err := t.Execute(&buf, user); err != nil {
		fmt.Printf("        Template error: %v\n", err)
		return
	}
	fmt.Print("        ", strings.ReplaceAll(buf.String(), "\n", "\n        "))
}

// regularExpressions demonstrates regexp package usage
func regularExpressions() {
	fmt.Println("    Regular Expressions Example:")

	// Email validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emails := []string{
		"user@example.com",
		"invalid.email@",
		"another.user123@company.co.uk",
	}

	re := regexp.MustCompile(emailPattern)
	fmt.Println("        Email Validation:")
	for _, email := range emails {
		fmt.Printf("        %s: %v\n", email, re.MatchString(email))
	}

	// Finding matches
	text := "The dates are 2024-03-15 and 2024-04-01"
	datePattern := `\d{4}-\d{2}-\d{2}`
	re = regexp.MustCompile(datePattern)

	fmt.Println("\n        Date Extraction:")
	dates := re.FindAllString(text, -1)
	for _, date := range dates {
		fmt.Printf("        Found date: %s\n", date)
	}

	// String replacement with regex
	text = "The color is #FF0000 and the background is #00FF00"
	colorPattern := `#[0-9A-Fa-f]{6}`
	re = regexp.MustCompile(colorPattern)

	result := re.ReplaceAllStringFunc(text, func(s string) string {
		return "COLOR(" + s + ")"
	})
	fmt.Printf("\n        Color Replacement:\n        %s\n", result)

	// Named capture groups
	text = "John Doe <john@example.com>"
	pattern := `(?P<name>[^<]+)\s*<(?P<email>[^>]+)>`
	re = regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(text)
	if match != nil {
		fmt.Println("\n        Named Captures:")
		for i, name := range re.SubexpNames() {
			if i > 0 && name != "" { // Skip the first empty subexp name
				fmt.Printf("        %s: %s\n", name, match[i])
			}
		}
	}
}

// unicodeHandling demonstrates Unicode and UTF-8 handling
func unicodeHandling() {
	fmt.Println("    Unicode Handling Example:")

	text := "Hello, 世界! 👋"

	// Length in bytes vs runes
	fmt.Printf("        Bytes length: %d\n", len(text))
	fmt.Printf("        Runes length: %d\n", utf8.RuneCountInString(text))

	// Iterate over runes
	fmt.Println("        Runes analysis:")
	for i, r := range text {
		fmt.Printf("        Position %d: %c (Unicode: %U)\n", i, r, r)

		// Character properties
		fmt.Printf("            Is letter: %v\n", unicode.IsLetter(r))
		fmt.Printf("            Is digit: %v\n", unicode.IsDigit(r))
		fmt.Printf("            Is symbol: %v\n", unicode.IsSymbol(r))
		fmt.Printf("            Is space: %v\n", unicode.IsSpace(r))
	}

	// Unicode normalization and case folding
	mixed := "é"
	fmt.Printf("\n        Original: %+q\n", mixed)
	fmt.Printf("        Is letter: %v\n", unicode.IsLetter([]rune(mixed)[0]))
}

func main() {
	fmt.Println("=== Go String Functions, Formatting, Templates, and Regex Examples ===")

	fmt.Println("\n1. String Functions:")
	stringFunctions()

	fmt.Println("\n2. String Formatting:")
	stringFormatting()

	fmt.Println("\n3. Text Templates:")
	textTemplates()

	fmt.Println("\n4. Regular Expressions:")
	regularExpressions()

	fmt.Println("\n5. Unicode Handling:")
	unicodeHandling()
}
