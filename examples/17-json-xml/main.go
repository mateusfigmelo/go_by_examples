package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"time"
)

// Address represents a physical address
type Address struct {
	Street  string `json:"street" xml:"street"`
	City    string `json:"city" xml:"city"`
	Country string `json:"country" xml:"country"`
}

// Person represents a person with various attributes
type Person struct {
	ID        int                    `json:"id" xml:"id,attr"`
	FirstName string                 `json:"first_name" xml:"firstName"`
	LastName  string                 `json:"last_name" xml:"lastName"`
	Age       int                    `json:"age,omitempty" xml:"age,omitempty"`
	Address   Address                `json:"address" xml:"address"`
	Birthday  time.Time              `json:"birthday" xml:"birthday"`
	Hobbies   []string               `json:"hobbies,omitempty" xml:"hobbies>hobby,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty" xml:"-"`
}

// CustomTime wraps time.Time for custom JSON marshaling
type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Format("2006-01-02"))), nil
}

// basicJSON demonstrates basic JSON marshaling and unmarshaling
func basicJSON() {
	fmt.Println("    Basic JSON Example:")

	// Create a person
	person := Person{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Birthday: time.Date(1993, time.April, 15, 0, 0, 0, 0, time.UTC),
		Hobbies:  []string{"reading", "hiking"},
		Metadata: map[string]interface{}{
			"verified": true,
			"score":    4.5,
		},
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Printf("        Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Printf("        Marshaled JSON:\n%s\n", jsonData)

	// Unmarshal from JSON
	var decodedPerson Person
	if err := json.Unmarshal(jsonData, &decodedPerson); err != nil {
		fmt.Printf("        Error unmarshaling JSON: %v\n", err)
		return
	}
	fmt.Printf("\n        Unmarshaled Person: %+v\n", decodedPerson)
}

// advancedJSON demonstrates advanced JSON features
func advancedJSON() {
	fmt.Println("    Advanced JSON Example:")

	// JSON streaming
	fmt.Println("        JSON Streaming:")
	jsonStream := `
		{"name": "Alice", "age": 25}
		{"name": "Bob", "age": 30}
		{"name": "Charlie", "age": 35}
	`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var m map[string]interface{}
		if err := decoder.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("        Error decoding JSON: %v\n", err)
			return
		}
		fmt.Printf("        Decoded: %v\n", m)
	}

	// Custom JSON marshaling example
	type Event struct {
		Name string     `json:"name"`
		Date CustomTime `json:"date"`
	}

	event := Event{
		Name: "Conference",
		Date: CustomTime{time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)},
	}

	fmt.Println("\n        Custom Marshaling:")
	jsonData, _ := json.MarshalIndent(event, "", "    ")
	fmt.Printf("        %s\n", jsonData)
}

// basicXML demonstrates basic XML marshaling and unmarshaling
func basicXML() {
	fmt.Println("    Basic XML Example:")

	// Create a person
	person := Person{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Smith",
		Age:       28,
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "London",
			Country: "UK",
		},
		Birthday: time.Date(1995, time.June, 20, 0, 0, 0, 0, time.UTC),
		Hobbies:  []string{"painting", "music"},
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Printf("        Error marshaling XML: %v\n", err)
		return
	}
	// Add XML header
	xmlData = append([]byte(xml.Header), xmlData...)
	fmt.Printf("        Marshaled XML:\n%s\n", xmlData)

	// Unmarshal from XML
	var decodedPerson Person
	if err := xml.Unmarshal(xmlData, &decodedPerson); err != nil {
		fmt.Printf("        Error unmarshaling XML: %v\n", err)
		return
	}
	fmt.Printf("\n        Unmarshaled Person: %+v\n", decodedPerson)
}

// advancedXML demonstrates advanced XML features
func advancedXML() {
	fmt.Println("    Advanced XML Example:")

	// Define a complex XML structure
	type Item struct {
		XMLName xml.Name `xml:"item"`
		ID      int      `xml:"id,attr"`
		Name    string   `xml:"name"`
		Price   float64  `xml:"price"`
		Tags    []string `xml:"tags>tag"`
	}

	type Inventory struct {
		XMLName xml.Name `xml:"inventory"`
		Items   []Item   `xml:"item"`
	}

	// Create sample data
	inventory := Inventory{
		Items: []Item{
			{
				ID:    1,
				Name:  "Laptop",
				Price: 999.99,
				Tags:  []string{"electronics", "computers"},
			},
			{
				ID:    2,
				Name:  "Mouse",
				Price: 24.99,
				Tags:  []string{"electronics", "accessories"},
			},
		},
	}

	// Marshal to XML with custom formatting
	xmlData, err := xml.MarshalIndent(inventory, "", "    ")
	if err != nil {
		fmt.Printf("        Error marshaling XML: %v\n", err)
		return
	}
	xmlData = append([]byte(xml.Header), xmlData...)
	fmt.Printf("        Marshaled XML:\n%s\n", xmlData)

	// XML Decoder example
	fmt.Println("\n        XML Decoder Example:")
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("        Error decoding XML: %v\n", err)
			return
		}

		switch t := token.(type) {
		case xml.StartElement:
			fmt.Printf("        Start Element: %s\n", t.Name.Local)
			for _, attr := range t.Attr {
				fmt.Printf("            Attribute: %s=%s\n", attr.Name.Local, attr.Value)
			}
		case xml.EndElement:
			fmt.Printf("        End Element: %s\n", t.Name.Local)
		case xml.CharData:
			content := strings.TrimSpace(string(t))
			if content != "" {
				fmt.Printf("        Character Data: %s\n", content)
			}
		}
	}
}

func main() {
	fmt.Println("=== Go JSON and XML Examples ===")

	fmt.Println("\n1. Basic JSON:")
	basicJSON()

	fmt.Println("\n2. Advanced JSON:")
	advancedJSON()

	fmt.Println("\n3. Basic XML:")
	basicXML()

	fmt.Println("\n4. Advanced XML:")
	advancedXML()
}
