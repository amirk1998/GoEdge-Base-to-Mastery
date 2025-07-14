// json_serialization.go
package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// User struct with various JSON tags
type JSONUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`             // Exclude from JSON
	Age       int       `json:"age,omitempty"` // Omit if empty
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	Profile   *Profile  `json:"profile,omitempty"` // Pointer to nested struct
}

// Profile nested struct
type Profile struct {
	Bio       string   `json:"bio"`
	Website   string   `json:"website,omitempty"`
	Interests []string `json:"interests"`
}

// Product struct for custom marshaling example
type JSONProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	InStock     bool    `json:"in_stock"`
	Description string  `json:"description"`
}

// Custom JSON marshaling for Product
func (p JSONProduct) MarshalJSON() ([]byte, error) {
	// Create a custom representation
	type Alias JSONProduct
	return json.Marshal(&struct {
		*Alias
		PriceFormatted string `json:"price_formatted"`
		Status         string `json:"status"`
	}{
		Alias:          (*Alias)(&p),
		PriceFormatted: fmt.Sprintf("%.2f %s", p.Price, p.Currency),
		Status:         map[bool]string{true: "Available", false: "Out of Stock"}[p.InStock],
	})
}

// Custom JSON unmarshaling for Product
func (p *JSONProduct) UnmarshalJSON(data []byte) error {
	type Alias JSONProduct
	aux := &struct {
		*Alias
		PriceFormatted string `json:"price_formatted"`
		Status         string `json:"status"`
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse custom fields if needed
	if aux.Status == "Available" {
		p.InStock = true
	} else {
		p.InStock = false
	}

	return nil
}

// Config struct with different tag options
type JSONConfig struct {
	AppName  string            `json:"app_name"`
	Version  string            `json:"version"`
	Debug    bool              `json:"debug,omitempty"`
	Database DatabaseConfig    `json:"database"`
	Features map[string]bool   `json:"features"`
	Servers  []ServerConfig    `json:"servers"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"-"` // Never include in JSON
	SSL      bool   `json:"ssl"`
}

type ServerConfig struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Weight int    `json:"weight,omitempty"`
}

// Custom time format example
type Event struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	StartTime CustomTime `json:"start_time"`
	EndTime   CustomTime `json:"end_time"`
}

// CustomTime wrapper for custom time marshaling
type CustomTime struct {
	time.Time
}

const CustomTimeFormat = "2006-01-02 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(CustomTimeFormat))
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	t, err := time.Parse(CustomTimeFormat, s)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// RunJSONSerializationExamples - main function to run all JSON serialization examples
func RunJSONSerializationExamples() {
	basicMarshalingExample()
	structTagsExample()
	customMarshalingExample()
	nestedStructExample()
	arraySliceExample()
	mapExample()
	customTimeExample()
	jsonStreamingExample()
	errorHandlingExample()
	configFileExample()
}

// Basic marshaling and unmarshaling
func basicMarshalingExample() {
	fmt.Println(Subtitle("📝 Basic JSON Marshaling/Unmarshaling"))

	// Create a user
	user := JSONUser{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "secret123", // This will be excluded
		Age:       30,
		IsActive:  true,
		CreatedAt: time.Now(),
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling: %v", err)
		return
	}

	fmt.Printf("Marshaled JSON: %s\n", string(jsonData))

	// Unmarshal back to struct
	var newUser JSONUser
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return
	}

	fmt.Printf("Unmarshaled User: %+v\n", newUser)
	fmt.Printf("Password field (excluded): '%s'\n", newUser.Password)
	fmt.Println()
}

// Struct tags demonstration
func structTagsExample() {
	fmt.Println(Subtitle("🏷️ Struct Tags Examples"))

	// User with all fields
	fullUser := JSONUser{
		ID:        1,
		Name:      "Alice Smith",
		Email:     "alice@example.com",
		Password:  "topsecret",
		Age:       25,
		IsActive:  true,
		CreatedAt: time.Now(),
		Profile: &Profile{
			Bio:       "Software Developer",
			Website:   "https://alice.dev",
			Interests: []string{"coding", "reading", "traveling"},
		},
	}

	// User with minimal fields (omitempty demonstration)
	minimalUser := JSONUser{
		ID:        2,
		Name:      "Bob Johnson",
		Email:     "bob@example.com",
		IsActive:  false,
		CreatedAt: time.Now(),
		// Age is 0 (will be omitted)
		// Profile is nil (will be omitted)
	}

	fmt.Println(Bold("Full User JSON:"))
	printJSON(fullUser)

	fmt.Println(Bold("Minimal User JSON (omitempty demo):"))
	printJSON(minimalUser)
}

// Custom marshaling example
func customMarshalingExample() {
	fmt.Println(Subtitle("🎨 Custom Marshaling Example"))

	product := JSONProduct{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Currency:    "USD",
		InStock:     true,
		Description: "High-performance laptop",
	}

	fmt.Println(Bold("Product with custom marshaling:"))
	printJSON(product)

	// Unmarshal the custom JSON
	jsonStr := `{
		"id": 2,
		"name": "Phone",
		"price": 599.99,
		"currency": "EUR",
		"description": "Smartphone",
		"price_formatted": "599.99 EUR",
		"status": "Out of Stock"
	}`

	var newProduct JSONProduct
	err := json.Unmarshal([]byte(jsonStr), &newProduct)
	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return
	}

	fmt.Printf("Unmarshaled product: %+v\n", newProduct)
	fmt.Println()
}

// Nested struct example
func nestedStructExample() {
	fmt.Println(Subtitle("🏗️ Nested Structures"))

	config := JSONConfig{
		AppName: "MyApp",
		Version: "1.0.0",
		Debug:   true,
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "admin",
			Password: "secret", // Won't appear in JSON
			SSL:      true,
		},
		Features: map[string]bool{
			"authentication": true,
			"logging":        true,
			"metrics":        false,
		},
		Servers: []ServerConfig{
			{Name: "web-1", Host: "192.168.1.10", Port: 8080, Weight: 100},
			{Name: "web-2", Host: "192.168.1.11", Port: 8080, Weight: 100},
			{Name: "api-1", Host: "192.168.1.20", Port: 3000},
		},
		Metadata: map[string]string{
			"environment": "production",
			"region":      "us-east-1",
		},
	}

	fmt.Println(Bold("Complex nested configuration:"))
	printJSON(config)
}

// Array and slice examples
func arraySliceExample() {
	fmt.Println(Subtitle("📊 Arrays and Slices"))

	type DataSet struct {
		Numbers  []int      `json:"numbers"`
		Strings  []string   `json:"strings"`
		Objects  []JSONUser `json:"objects"`
		Matrix   [][]int    `json:"matrix"`
		Empty    []string   `json:"empty,omitempty"`
		NilSlice []string   `json:"nil_slice,omitempty"`
	}

	data := DataSet{
		Numbers: []int{1, 2, 3, 4, 5},
		Strings: []string{"apple", "banana", "cherry"},
		Objects: []JSONUser{
			{ID: 1, Name: "User 1", Email: "user1@example.com", IsActive: true, CreatedAt: time.Now()},
			{ID: 2, Name: "User 2", Email: "user2@example.com", IsActive: false, CreatedAt: time.Now()},
		},
		Matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		Empty:    []string{},
		NilSlice: nil,
	}

	fmt.Println(Bold("Array and slice marshaling:"))
	printJSON(data)
}

// Map example
func mapExample() {
	fmt.Println(Subtitle("🗺️ Map Examples"))

	type APIResponse struct {
		Status   string                 `json:"status"`
		Data     map[string]interface{} `json:"data"`
		Metadata map[string]string      `json:"metadata"`
		Counts   map[string]int         `json:"counts"`
	}

	response := APIResponse{
		Status: "success",
		Data: map[string]interface{}{
			"user_id":    123,
			"username":   "john_doe",
			"balance":    1250.50,
			"active":     true,
			"last_login": time.Now().Format(time.RFC3339),
			"preferences": map[string]interface{}{
				"theme":         "dark",
				"language":      "en",
				"notifications": true,
			},
		},
		Metadata: map[string]string{
			"version":   "1.0",
			"timestamp": time.Now().Format(time.RFC3339),
			"source":    "api",
		},
		Counts: map[string]int{
			"total_users":    1000,
			"active_users":   850,
			"inactive_users": 150,
		},
	}

	fmt.Println(Bold("Map marshaling:"))
	printJSON(response)
}

// Custom time example
func customTimeExample() {
	fmt.Println(Subtitle("⏰ Custom Time Formatting"))

	event := Event{
		ID:        1,
		Title:     "Team Meeting",
		StartTime: CustomTime{time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)},
		EndTime:   CustomTime{time.Date(2024, 1, 15, 11, 30, 0, 0, time.UTC)},
	}

	fmt.Println(Bold("Event with custom time format:"))
	printJSON(event)

	// Unmarshal custom time
	jsonStr := `{
		"id": 2,
		"title": "Code Review",
		"start_time": "2024-01-16 14:00:00",
		"end_time": "2024-01-16 15:30:00"
	}`

	var newEvent Event
	err := json.Unmarshal([]byte(jsonStr), &newEvent)
	if err != nil {
		log.Printf("Error unmarshaling: %v", err)
		return
	}

	fmt.Printf("Unmarshaled event: %+v\n", newEvent)
	fmt.Println()
}

// JSON streaming example
func jsonStreamingExample() {
	fmt.Println(Subtitle("🌊 JSON Streaming"))

	// Create a JSON array string
	jsonArray := `[
		{"id": 1, "name": "Item 1"},
		{"id": 2, "name": "Item 2"},
		{"id": 3, "name": "Item 3"}
	]`

	// Use decoder for streaming
	decoder := json.NewDecoder(strings.NewReader(jsonArray))

	// Read opening bracket
	token, err := decoder.Token()
	if err != nil {
		log.Printf("Error reading token: %v", err)
		return
	}
	fmt.Printf("Opening token: %v\n", token)

	// Read array elements
	for decoder.More() {
		var item map[string]interface{}
		err := decoder.Decode(&item)
		if err != nil {
			log.Printf("Error decoding item: %v", err)
			continue
		}
		fmt.Printf("Decoded item: %+v\n", item)
	}

	// Read closing bracket
	token, err = decoder.Token()
	if err != nil {
		log.Printf("Error reading closing token: %v", err)
		return
	}
	fmt.Printf("Closing token: %v\n", token)
	fmt.Println()
}

// Error handling example
func errorHandlingExample() {
	fmt.Println(Subtitle("🚨 Error Handling"))

	// Invalid JSON
	invalidJSON := `{"name": "John", "age": "not a number"}`

	var user JSONUser
	err := json.Unmarshal([]byte(invalidJSON), &user)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)

		// Check for specific error type
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("Syntax error at position %d\n", syntaxErr.Offset)
		}
		if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
			fmt.Printf("Type error: cannot unmarshal %v into Go struct field %s of type %v\n",
				typeErr.Value, typeErr.Field, typeErr.Type)
		}
	}

	// Marshal error example (circular reference)
	a := &CircularA{}
	b := &CircularB{}
	a.B = b
	b.A = a

	_, err = json.Marshal(a)
	if err != nil {
		fmt.Printf("Marshal error (circular reference): %v\n", err)
	}
	fmt.Println()
}

// Config file example
func configFileExample() {
	fmt.Println(Subtitle("⚙️ Configuration File Example"))

	// Simulate reading a config file
	configJSON := `{
		"app_name": "WebService",
		"version": "2.1.0",
		"debug": false,
		"database": {
			"host": "db.example.com",
			"port": 5432,
			"username": "webapp",
			"ssl": true
		},
		"features": {
			"authentication": true,
			"logging": true,
			"metrics": true,
			"caching": false
		},
		"servers": [
			{
				"name": "primary",
				"host": "web1.example.com",
				"port": 80,
				"weight": 100
			},
			{
				"name": "secondary",
				"host": "web2.example.com",
				"port": 80,
				"weight": 50
			}
		],
		"metadata": {
			"environment": "staging",
			"region": "eu-west-1",
			"deployment_id": "dep-123456"
		}
	}`

	var config JSONConfig
	err := json.Unmarshal([]byte(configJSON), &config)
	if err != nil {
		log.Printf("Error parsing config: %v", err)
		return
	}

	fmt.Printf("Loaded configuration:\n")
	fmt.Printf("  App: %s v%s\n", config.AppName, config.Version)
	fmt.Printf("  Debug: %v\n", config.Debug)
	fmt.Printf("  Database: %s:%d (SSL: %v)\n",
		config.Database.Host, config.Database.Port, config.Database.SSL)
	fmt.Printf("  Features enabled: ")
	for feature, enabled := range config.Features {
		if enabled {
			fmt.Printf("%s ", feature)
		}
	}
	fmt.Println()
	fmt.Printf("  Servers: %d configured\n", len(config.Servers))

	// Pretty print the entire config
	fmt.Println(Bold("Full configuration:"))
	printJSON(config)
}

// Helper function to print JSON with proper formatting
func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Error marshaling: %v", err)
		return
	}
	fmt.Printf("%s\n\n", string(jsonData))
}

// Place these at the top level, outside any function

type CircularA struct {
	B *CircularB `json:"b"`
}

type CircularB struct {
	A *CircularA `json:"a"`
}
