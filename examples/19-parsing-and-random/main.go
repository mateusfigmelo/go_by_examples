package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"
)

// randomNumbers demonstrates various ways to generate random numbers
func randomNumbers() {
	fmt.Println("    Random Number Generation:")

	// Secure random number using crypto/rand
	max := big.NewInt(100)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		fmt.Printf("        Error generating crypto random: %v\n", err)
	} else {
		fmt.Printf("        Crypto random (0-99): %d\n", n)
	}

	// Generate random bytes
	bytes := make([]byte, 10)
	_, err = rand.Read(bytes)
	if err != nil {
		fmt.Printf("        Error generating random bytes: %v\n", err)
	} else {
		fmt.Printf("        Random bytes: %v\n", bytes)
	}

	// Generate random string
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, 10)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Printf("        Error generating random string: %v\n", err)
			return
		}
		result[i] = letters[n.Int64()]
	}
	fmt.Printf("        Random string: %s\n", string(result))
}

// numberParsing demonstrates parsing numbers from strings
func numberParsing() {
	fmt.Println("    Number Parsing:")

	// Integer parsing
	intStr := "42"
	if num, err := strconv.Atoi(intStr); err != nil {
		fmt.Printf("        Error parsing int: %v\n", err)
	} else {
		fmt.Printf("        Parsed int: %d\n", num)
	}

	// Parse int with base
	hexStr := "2A"
	if num, err := strconv.ParseInt(hexStr, 16, 64); err != nil {
		fmt.Printf("        Error parsing hex: %v\n", err)
	} else {
		fmt.Printf("        Parsed hex: %d\n", num)
	}

	// Float parsing
	floatStr := "3.14159"
	if num, err := strconv.ParseFloat(floatStr, 64); err != nil {
		fmt.Printf("        Error parsing float: %v\n", err)
	} else {
		fmt.Printf("        Parsed float: %f\n", num)
	}

	// Boolean parsing
	boolStr := "true"
	if val, err := strconv.ParseBool(boolStr); err != nil {
		fmt.Printf("        Error parsing bool: %v\n", err)
	} else {
		fmt.Printf("        Parsed bool: %v\n", val)
	}

	// Error handling examples
	invalidInt := "not a number"
	if _, err := strconv.Atoi(invalidInt); err != nil {
		fmt.Printf("        Expected error parsing invalid int: %v\n", err)
	}

	// Number formatting
	num := 42
	binary := strconv.FormatInt(int64(num), 2)
	hex := strconv.FormatInt(int64(num), 16)
	fmt.Printf("        Number 42 in binary: %s\n", binary)
	fmt.Printf("        Number 42 in hex: %s\n", hex)
}

// urlParsing demonstrates URL parsing and manipulation
func urlParsing() {
	fmt.Println("    URL Parsing:")

	// Parse basic URL
	rawURL := "https://example.com:8080/path?query=value#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("        Error parsing URL: %v\n", err)
		return
	}

	fmt.Printf("        Scheme: %s\n", parsedURL.Scheme)
	fmt.Printf("        Host: %s\n", parsedURL.Host)
	fmt.Printf("        Path: %s\n", parsedURL.Path)
	fmt.Printf("        RawQuery: %s\n", parsedURL.RawQuery)
	fmt.Printf("        Fragment: %s\n", parsedURL.Fragment)

	// Query parameters
	queryParams := parsedURL.Query()
	fmt.Printf("        Query params: %v\n", queryParams)

	// Build URL from components
	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/users",
		RawQuery: "id=123&name=John",
	}
	fmt.Printf("        Built URL: %s\n", u.String())

	// Escape/unescape
	path := "/path with spaces/and+symbols"
	escaped := url.PathEscape(path)
	fmt.Printf("        Escaped path: %s\n", escaped)

	if unescaped, err := url.PathUnescape(escaped); err != nil {
		fmt.Printf("        Error unescaping path: %v\n", err)
	} else {
		fmt.Printf("        Unescaped path: %s\n", unescaped)
	}

	// Query parameter manipulation
	query := url.Values{}
	query.Add("name", "John Doe")
	query.Add("age", "30")
	query.Add("tags", "go")
	query.Add("tags", "programming")
	fmt.Printf("        Query string: %s\n", query.Encode())
	fmt.Printf("        Multiple values for 'tags': %v\n", query["tags"])

	// Relative URL resolution
	base, _ := url.Parse("https://example.com/base/")
	relative, _ := url.Parse("../other")
	fmt.Printf("        Resolved URL: %s\n", base.ResolveReference(relative))
}

// advancedURLParsing demonstrates more complex URL operations
func advancedURLParsing() {
	fmt.Println("    Advanced URL Parsing:")

	// Parse and normalize URLs
	urls := []string{
		"HTTPS://EXAMPLE.COM/path",
		"http://example.com:80/path",
		"https://example.com:443/path",
	}

	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			fmt.Printf("        Error parsing URL %s: %v\n", u, err)
			continue
		}

		// Normalize
		parsed.Host = strings.ToLower(parsed.Host)
		if (parsed.Scheme == "http" && strings.HasSuffix(parsed.Host, ":80")) ||
			(parsed.Scheme == "https" && strings.HasSuffix(parsed.Host, ":443")) {
			parsed.Host = strings.TrimSuffix(parsed.Host, ":80")
			parsed.Host = strings.TrimSuffix(parsed.Host, ":443")
		}

		fmt.Printf("        Normalized URL: %s\n", parsed)
	}

	// URL validation
	validateURL := func(rawURL string) bool {
		u, err := url.Parse(rawURL)
		if err != nil {
			return false
		}
		return u.Scheme != "" && u.Host != ""
	}

	testURLs := []string{
		"https://example.com",
		"not-a-url",
		"ftp://example.com",
		"//no-scheme.com",
	}

	for _, u := range testURLs {
		fmt.Printf("        Is %s valid? %v\n", u, validateURL(u))
	}

	// Working with userinfo
	urlWithAuth := "https://user:pass@example.com/path"
	parsed, _ := url.Parse(urlWithAuth)
	if userinfo := parsed.User; userinfo != nil {
		username := userinfo.Username()
		password, hasPassword := userinfo.Password()
		fmt.Printf("        Username: %s\n", username)
		fmt.Printf("        Has password: %v\n", hasPassword)
		if hasPassword {
			fmt.Printf("        Password: %s\n", password)
		}
	}
}

func main() {
	fmt.Println("=== Go Random Numbers, Number Parsing, and URL Parsing Examples ===")

	fmt.Println("\n1. Random Numbers:")
	randomNumbers()

	fmt.Println("\n2. Number Parsing:")
	numberParsing()

	fmt.Println("\n3. URL Parsing:")
	urlParsing()

	fmt.Println("\n4. Advanced URL Parsing:")
	advancedURLParsing()
}
