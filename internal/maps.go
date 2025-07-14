// maps.go
package internal

import (
	"fmt"
	"sort"
)

// RunMapExamples - main function to run all map examples
func RunMapExamples() {
	fmt.Println(Subtitle("🗺️ Maps Examples:"))
	basicMapExample()
	mapOperationsExample()
	mapIterationExample()
	mapAdvancedExample()
	mapPerformanceExample()
	nestedMapsExample()
	mapWithStructsExample()
	mapConcurrencyExample()
}

// basicMapExample - demonstrates basic map operations
func basicMapExample() {
	fmt.Println(Bold("1. Basic Map Operations:"))

	// Map declaration and initialization
	var ages map[string]int
	fmt.Printf("Nil map: %v (len: %d)\n", ages, len(ages))

	// Initialize with make
	ages = make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	// Map literal
	scores := map[string]int{
		"Math":    95,
		"Science": 87,
		"English": 92,
	}

	fmt.Printf("Ages map: %v\n", ages)
	fmt.Printf("Scores map: %v\n", scores)

	// Accessing values
	fmt.Printf("Alice's age: %d\n", ages["Alice"])
	fmt.Printf("Math score: %d\n", scores["Math"])

	// Zero value for missing key
	fmt.Printf("Missing key (David): %d\n", ages["David"])

	fmt.Println()
}

// mapOperationsExample - demonstrates map operations
func mapOperationsExample() {
	fmt.Println(Bold("2. Map Operations:"))

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Printf("Original map: %v\n", colors)

	// Check if key exists
	if value, exists := colors["red"]; exists {
		fmt.Printf("Red color code: %s\n", value)
	}

	if value, exists := colors["yellow"]; exists {
		fmt.Printf("Yellow color code: %s\n", value)
	} else {
		fmt.Println("Yellow color not found")
	}

	// Add new key-value pair
	colors["yellow"] = "#FFFF00"
	fmt.Printf("After adding yellow: %v\n", colors)

	// Update existing value
	colors["red"] = "#CC0000"
	fmt.Printf("After updating red: %v\n", colors)

	// Delete key
	delete(colors, "blue")
	fmt.Printf("After deleting blue: %v\n", colors)

	// Length of map
	fmt.Printf("Map length: %d\n", len(colors))

	fmt.Println()
}

// mapIterationExample - demonstrates map iteration
func mapIterationExample() {
	fmt.Println(Bold("3. Map Iteration:"))

	inventory := map[string]int{
		"apples":  50,
		"bananas": 30,
		"oranges": 25,
		"grapes":  40,
	}

	fmt.Printf("Inventory: %v\n", inventory)

	// Iterate over key-value pairs
	fmt.Println("Iterating over key-value pairs:")
	for item, quantity := range inventory {
		fmt.Printf("  %s: %d\n", item, quantity)
	}

	// Iterate over keys only
	fmt.Println("Iterating over keys only:")
	for item := range inventory {
		fmt.Printf("  %s\n", item)
	}

	// Iterate over values only
	fmt.Println("Iterating over values only:")
	for _, quantity := range inventory {
		fmt.Printf("  %d\n", quantity)
	}

	// Sorted iteration (maps are unordered)
	fmt.Println("Sorted iteration:")
	var keys []string
	for key := range inventory {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("  %s: %d\n", key, inventory[key])
	}

	fmt.Println()
}

// mapAdvancedExample - demonstrates advanced map techniques
func mapAdvancedExample() {
	fmt.Println(Bold("4. Advanced Map Techniques:"))

	// Map with slice values
	groups := map[string][]string{
		"fruits":     {"apple", "banana", "orange"},
		"vegetables": {"carrot", "broccoli", "spinach"},
		"grains":     {"rice", "wheat", "oats"},
	}

	fmt.Println("Map with slice values:")
	for category, items := range groups {
		fmt.Printf("  %s: %v\n", category, items)
	}

	// Map with map values (nested maps)
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 87,
			"English": 92,
		},
		"Bob": {
			"Math":    78,
			"Science": 82,
			"English": 85,
		},
	}

	fmt.Println("Nested maps:")
	for student, subjects := range students {
		fmt.Printf("  %s:\n", student)
		for subject, score := range subjects {
			fmt.Printf("    %s: %d\n", subject, score)
		}
	}

	// Map with function values
	operations := map[string]func(int, int) int{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
		"multiply": func(a, b int) int { return a * b },
		"divide":   func(a, b int) int { return a / b },
	}

	fmt.Println("Map with function values:")
	for operation, fn := range operations {
		if operation != "divide" { // avoid division by zero in example
			result := fn(10, 5)
			fmt.Printf("  %s(10, 5) = %d\n", operation, result)
		}
	}

	fmt.Println()
}

// mapPerformanceExample - demonstrates map performance considerations
func mapPerformanceExample() {
	fmt.Println(Bold("5. Map Performance Considerations:"))

	// Map with initial capacity
	largeMap := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		largeMap[i] = fmt.Sprintf("value_%d", i)
	}

	fmt.Printf("Large map created with %d elements\n", len(largeMap))

	// Map key types performance
	fmt.Println("Different key types:")

	// String keys
	stringMap := make(map[string]int)
	stringMap["key1"] = 1
	stringMap["key2"] = 2

	// Integer keys (generally faster)
	intMap := make(map[int]string)
	intMap[1] = "value1"
	intMap[2] = "value2"

	fmt.Printf("String key map: %v\n", stringMap)
	fmt.Printf("Integer key map: %v\n", intMap)

	// Struct keys (must be comparable)
	type Point struct {
		X, Y int
	}

	pointMap := make(map[Point]string)
	pointMap[Point{1, 2}] = "point1"
	pointMap[Point{3, 4}] = "point2"

	fmt.Printf("Struct key map: %v\n", pointMap)

	fmt.Println()
}

// nestedMapsExample - demonstrates complex nested map structures
func nestedMapsExample() {
	fmt.Println(Bold("6. Nested Maps and Complex Structures:"))

	// Company organizational structure
	company := map[string]map[string]map[string]interface{}{
		"Engineering": {
			"Backend": {
				"lead":     "Alice",
				"members":  []string{"Bob", "Charlie", "David"},
				"projects": 3,
				"budget":   100000,
			},
			"Frontend": {
				"lead":     "Eve",
				"members":  []string{"Frank", "Grace"},
				"projects": 2,
				"budget":   75000,
			},
		},
		"Marketing": {
			"Digital": {
				"lead":     "Henry",
				"members":  []string{"Ivy", "Jack"},
				"projects": 4,
				"budget":   50000,
			},
		},
	}

	fmt.Println("Company structure:")
	for department, teams := range company {
		fmt.Printf("  %s:\n", department)
		for team, details := range teams {
			fmt.Printf("    %s:\n", team)
			for key, value := range details {
				fmt.Printf("      %s: %v\n", key, value)
			}
		}
	}

	// Safe nested access
	if engineering, exists := company["Engineering"]; exists {
		if backend, exists := engineering["Backend"]; exists {
			if lead, exists := backend["lead"]; exists {
				fmt.Printf("Backend lead: %s\n", lead)
			}
		}
	}

	fmt.Println()
}

// Employee - struct for map examples
type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

// mapWithStructsExample - demonstrates maps with structs
func mapWithStructsExample() {
	fmt.Println(Bold("7. Maps with Structs:"))

	// Map with struct values
	employees := map[int]Employee{
		1: {ID: 1, Name: "Alice", Position: "Engineer", Salary: 75000},
		2: {ID: 2, Name: "Bob", Position: "Designer", Salary: 65000},
		3: {ID: 3, Name: "Charlie", Position: "Manager", Salary: 85000},
	}

	fmt.Println("Employees map:")
	for id, emp := range employees {
		fmt.Printf("  ID %d: %s (%s) - $%.2f\n", id, emp.Name, emp.Position, emp.Salary)
	}

	// Update struct in map
	emp := employees[1]
	emp.Salary = 80000
	employees[1] = emp
	fmt.Printf("Updated Alice's salary: $%.2f\n", employees[1].Salary)

	// Map with struct pointers (for easier updates)
	employeePtrs := map[int]*Employee{
		1: {ID: 1, Name: "Alice", Position: "Engineer", Salary: 75000},
		2: {ID: 2, Name: "Bob", Position: "Designer", Salary: 65000},
	}

	// Direct update through pointer
	employeePtrs[1].Salary = 82000
	fmt.Printf("Updated Alice's salary (via pointer): $%.2f\n", employeePtrs[1].Salary)

	// Index by different fields
	employeesByName := make(map[string]*Employee)
	for _, emp := range employeePtrs {
		employeesByName[emp.Name] = emp
	}

	fmt.Println("Employees by name:")
	for name, emp := range employeesByName {
		fmt.Printf("  %s: %s (ID: %d)\n", name, emp.Position, emp.ID)
	}

	fmt.Println()
}

// mapConcurrencyExample - demonstrates map concurrency considerations
func mapConcurrencyExample() {
	fmt.Println(Bold("8. Map Concurrency Considerations:"))

	// Note: This is a demonstration of concepts, not actual concurrent code
	// Maps are NOT safe for concurrent access

	fmt.Println("Map concurrency notes:")
	fmt.Println("  - Maps are NOT thread-safe")
	fmt.Println("  - Concurrent read/write operations cause panic")
	fmt.Println("  - Use sync.RWMutex for concurrent access")
	fmt.Println("  - Consider sync.Map for high-concurrency scenarios")

	// Thread-safe map wrapper example (conceptual)
	type SafeMap struct {
		data map[string]int
		// In real implementation, add sync.RWMutex here
	}

	safeMap := SafeMap{
		data: make(map[string]int),
	}

	// In real implementation, these would be protected by mutex
	safeMap.data["key1"] = 1
	safeMap.data["key2"] = 2

	fmt.Printf("Safe map data: %v\n", safeMap.data)

	// Example of map copying for safe concurrent read
	original := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Create copy for safe concurrent access
	copy := make(map[string]int)
	for k, v := range original {
		copy[k] = v
	}

	fmt.Printf("Original map: %v\n", original)
	fmt.Printf("Copy for concurrent access: %v\n", copy)

	fmt.Println()
}
