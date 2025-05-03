package main

import (
	"fmt"
	"sort"
	"time"
)

// Person represents a person with various attributes
type Person struct {
	Name string
	Age  int
}

// Track represents a music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// byAge implements sort.Interface for []Person based on Age
type byAge []Person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// byArtist implements sort.Interface for []Track based on Artist
type byArtist []Track

func (a byArtist) Len() int           { return len(a) }
func (a byArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byArtist) Less(i, j int) bool { return a[i].Artist < a[j].Artist }

// customLess represents a custom less function for Track sorting
type customLess func(p1, p2 *Track) bool

// multiSort implements multi-criteria sort for tracks
type multiSort struct {
	tracks []Track
	less   []customLess
}

func (ms multiSort) Len() int { return len(ms.tracks) }
func (ms multiSort) Swap(i, j int) {
	ms.tracks[i], ms.tracks[j] = ms.tracks[j], ms.tracks[i]
}
func (ms multiSort) Less(i, j int) bool {
	p, q := &ms.tracks[i], &ms.tracks[j]
	// Try all but the last comparison
	for _, less := range ms.less[:len(ms.less)-1] {
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	// Try the final comparison
	return ms.less[len(ms.less)-1](p, q)
}

// basicSorting demonstrates sorting of basic types
func basicSorting() {
	// Sort integers
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	sort.Ints(numbers)
	fmt.Printf("    Sorted integers: %v\n", numbers)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(words)
	fmt.Printf("    Sorted strings: %v\n", words)

	// Sort float64s
	floats := []float64{3.14, 1.41, 2.71, 1.73}
	sort.Float64s(floats)
	fmt.Printf("    Sorted floats: %v\n", floats)
}

// customSorting demonstrates sorting with custom types
func customSorting() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
		{"David", 35},
	}

	// Sort by age
	sort.Sort(byAge(people))
	fmt.Println("    Sorted by age:")
	for _, p := range people {
		fmt.Printf("        %s: %d\n", p.Name, p.Age)
	}

	// Sort by name using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("\n    Sorted by name:")
	for _, p := range people {
		fmt.Printf("        %s: %d\n", p.Name, p.Age)
	}
}

// trackSorting demonstrates complex sorting with multiple criteria
func trackSorting() {
	tracks := []Track{
		{"Go", "Moby", "Moby", 1992, 3 * time.Minute},
		{"Go", "Delilah", "From the Roots Up", 2012, 4 * time.Minute},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, 4 * time.Minute},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, 4 * time.Minute},
	}

	// Sort by artist
	sort.Sort(byArtist(tracks))
	fmt.Println("    Sorted by artist:")
	printTracks(tracks)

	// Sort by multiple fields
	sort.Sort(multiSort{
		tracks: tracks,
		less: []customLess{
			func(x, y *Track) bool { return x.Title < y.Title },   // Sort by title
			func(x, y *Track) bool { return x.Year < y.Year },     // Then by year
			func(x, y *Track) bool { return x.Length < y.Length }, // Then by length
			func(x, y *Track) bool { return x.Artist < y.Artist }, // Then by artist
		},
	})
	fmt.Println("\n    Sorted by title, year, length, artist:")
	printTracks(tracks)
}

// reverseSorting demonstrates reverse sorting
func reverseSorting() {
	// Reverse sort integers
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("    Reverse sorted integers: %v\n", numbers)

	// Reverse sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	fmt.Printf("    Reverse sorted strings: %v\n", words)
}

// stableSorting demonstrates stable sorting
func stableSorting() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 25},
		{"Charlie", 20},
		{"David", 25},
	}

	// Stable sort preserves order of equal elements
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("    Stable sorted by age (preserves order of equal ages):")
	for _, p := range people {
		fmt.Printf("        %s: %d\n", p.Name, p.Age)
	}
}

// searchSorting demonstrates searching in sorted slices
func searchSorting() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5

	// Binary search
	index := sort.SearchInts(numbers, target)
	fmt.Printf("    Found %d at index %d\n", target, index)

	// Search with custom function
	people := []Person{
		{"Alice", 20},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 35},
	}

	// Sort first
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	// Search for person with age >= 30
	index = sort.Search(len(people), func(i int) bool {
		return people[i].Age >= 30
	})
	fmt.Printf("    First person with age >= 30: %s\n", people[index].Name)
}

// Helper function to print tracks
func printTracks(tracks []Track) {
	const format = "        %-20s %-20s %-20s %d %s\n"
	for _, t := range tracks {
		fmt.Printf(format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
}

func main() {
	fmt.Println("=== Go Sorting Examples ===")

	fmt.Println("\n1. Basic Sorting:")
	basicSorting()

	fmt.Println("\n2. Custom Sorting:")
	customSorting()

	fmt.Println("\n3. Track Sorting (Multiple Criteria):")
	trackSorting()

	fmt.Println("\n4. Reverse Sorting:")
	reverseSorting()

	fmt.Println("\n5. Stable Sorting:")
	stableSorting()

	fmt.Println("\n6. Search in Sorted Slices:")
	searchSorting()
}
