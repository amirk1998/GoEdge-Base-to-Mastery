// package_system.go
package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	// Import examples with different techniques

	// Standard library imports

	// Import with alias
	mrand "math/rand"

	// Blank import (for side effects only)
	_ "time/tzdata"
)

// Package-level variables (exported)
var (
	PackageVersion    = "1.0.0"
	PackageAuthor     = "Golang Developer"
	defaultConfigFile = "config.json" // unexported
)

// Package-level constants
const (
	MaxRetries     = 3
	DefaultTimeout = 30 * time.Second
	apiEndpoint    = "https://api.example.com" // unexported
)

// Exported types
type Config struct {
	APIKey    string   `json:"api_key"`
	Timeout   int      `json:"timeout"`
	Debug     bool     `json:"debug"`
	endpoints []string // unexported field
}

type Logger struct {
	prefix string
	debug  bool
}

// Exported functions
func NewConfig() *Config {
	return &Config{
		Timeout:   30,
		Debug:     false,
		endpoints: []string{apiEndpoint},
	}
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		debug:  false,
	}
}

// Exported methods
func (c *Config) SetAPIKey(key string) {
	c.APIKey = key
}

func (c *Config) GetEndpoint() string {
	if len(c.endpoints) > 0 {
		return c.endpoints[0]
	}
	return ""
}

func (l *Logger) Debug(message string) {
	if l.debug {
		fmt.Printf("[%s] DEBUG: %s\n", l.prefix, message)
	}
}

// unexported functions
func validateConfig(c *Config) error {
	if c.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	return nil
}

func loadConfigFromFile(filename string) (*Config, error) {
	// Simulate loading config from file
	return &Config{
		APIKey:  "demo-key",
		Timeout: 60,
		Debug:   true,
	}, nil
}

// Package initialization
func init() {
	fmt.Println(InfoText("Package system initialized"))
	mrand.Seed(time.Now().UnixNano())
}

// RunPackageSystemExamples - main function to run all package system examples
func RunPackageSystemExamples() {
	basicPackageExample()
	importAliasExample()
	visibilityExample()
	packageVariablesExample()
	initFunctionExample()
	packageDocumentationExample()
	packageOrganizationExample()
	importPathExample()
	blankImportExample()
	packageTestingExample()
}

// Example 1: Basic package usage
func basicPackageExample() {
	fmt.Println(Header("1. Basic Package Usage"))

	// Using standard library packages
	fmt.Println("Current time:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("Working directory:", getCurrentDir())

	// Using package-level variables and constants
	fmt.Printf("Package version: %s\n", PackageVersion)
	fmt.Printf("Package author: %s\n", PackageAuthor)
	fmt.Printf("Max retries: %d\n", MaxRetries)
	fmt.Printf("Default timeout: %v\n", DefaultTimeout)
	fmt.Println()
}

// Example 2: Import aliases
func importAliasExample() {
	fmt.Println(Header("2. Import Aliases"))

	// Using math/rand with alias 'mrand'
	fmt.Println("Random number with alias:", mrand.Intn(100))

	// Standard import without alias
	jsonData := `{"name": "John", "age": 30}`
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &data); err == nil {
		fmt.Printf("Parsed JSON: %+v\n", data)
	}

	// Multiple ways to import and use
	num := "42"
	if parsed, err := strconv.Atoi(num); err == nil {
		fmt.Printf("Parsed number: %d\n", parsed)
	}
	fmt.Println()
}

// Example 3: Exported vs unexported (visibility)
func visibilityExample() {
	fmt.Println(Header("3. Visibility (Exported vs Unexported)"))

	// Exported - can be accessed from other packages
	config := NewConfig()
	config.SetAPIKey("my-secret-key")

	fmt.Printf("Config API Key: %s\n", config.APIKey)
	fmt.Printf("Config Timeout: %d\n", config.Timeout)
	fmt.Printf("Config Debug: %t\n", config.Debug)
	fmt.Printf("Config Endpoint: %s\n", config.GetEndpoint())

	// Unexported - can only be accessed within the same package
	if err := validateConfig(config); err != nil {
		fmt.Printf("Config validation error: %v\n", err)
	} else {
		fmt.Println("Config is valid")
	}

	// Accessing unexported package variable
	fmt.Printf("Default config file: %s\n", defaultConfigFile)
	fmt.Println()
}

// Example 4: Package-level variables and constants
func packageVariablesExample() {
	fmt.Println(Header("4. Package-level Variables and Constants"))

	// Package variables can be modified
	fmt.Printf("Original version: %s\n", PackageVersion)
	PackageVersion = "1.1.0"
	fmt.Printf("Updated version: %s\n", PackageVersion)

	// Constants cannot be modified
	fmt.Printf("Max retries (constant): %d\n", MaxRetries)
	fmt.Printf("Default timeout (constant): %v\n", DefaultTimeout)

	// Package variables are shared across the package
	logger1 := NewLogger("APP")
	logger2 := NewLogger("DB")

	logger1.Log("Application started")
	logger2.Log("Database connected")

	fmt.Println()
}

// Example 5: Init function behavior
func initFunctionExample() {
	fmt.Println(Header("5. Init Function Behavior"))

	// Init function already ran when package was imported
	fmt.Println("Init function ran during package initialization")
	fmt.Println("Random number (seeded in init):", mrand.Intn(1000))

	// Multiple init functions can exist and run in declaration order
	fmt.Println("Check console output for init message")
	fmt.Println()
}

// Example 6: Package documentation patterns
func packageDocumentationExample() {
	fmt.Println(Header("6. Package Documentation"))

	// Documenting package usage
	fmt.Println("Package documentation example:")
	fmt.Println("// Package main provides examples of Go package system")
	fmt.Println("// It demonstrates imports, visibility, and organization")
	fmt.Println()

	// Documenting types and functions
	fmt.Println("Type documentation:")
	fmt.Println("// Config represents application configuration")
	fmt.Println("// It provides methods for setting and getting config values")
	fmt.Println()

	// Usage examples in documentation
	config := NewConfig()
	config.SetAPIKey("example-key")
	fmt.Printf("Example usage: %+v\n", config)
	fmt.Println()
}

// Example 7: Package organization patterns
func packageOrganizationExample() {
	fmt.Println(Header("7. Package Organization"))

	// Demonstrate logical grouping
	fmt.Println("Package organization patterns:")
	fmt.Println("1. By functionality (auth, db, api)")
	fmt.Println("2. By domain (user, product, order)")
	fmt.Println("3. By layer (handler, service, repository)")

	// Example of internal package structure
	packageStructure := `
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── auth/
│   ├── db/
│   └── api/
├── pkg/
│   ├── config/
│   └── logger/
└── go.mod
`
	fmt.Println("Example structure:")
	fmt.Println(packageStructure)
	fmt.Println()
}

// Example 8: Import paths and module system
func importPathExample() {
	fmt.Println(Header("8. Import Paths"))

	// Standard library imports
	fmt.Println("Standard library imports:")
	fmt.Println(`import "fmt"`)
	fmt.Println(`import "time"`)
	fmt.Println(`import "net/http"`)
	fmt.Println()

	// Third-party imports (examples)
	fmt.Println("Third-party imports:")
	fmt.Println(`import "github.com/gin-gonic/gin"`)
	fmt.Println(`import "github.com/gorilla/mux"`)
	fmt.Println(`import "gorm.io/gorm"`)
	fmt.Println()

	// Local imports
	fmt.Println("Local imports:")
	fmt.Println(`import "myproject/internal/auth"`)
	fmt.Println(`import "myproject/pkg/config"`)
	fmt.Println()

	// Import grouping
	fmt.Println("Import grouping best practice:")
	importExample := `import (
    // Standard library
    "fmt"
    "time"
    
    // Third-party
    "github.com/gin-gonic/gin"
    "github.com/gorilla/mux"
    
    // Local
    "myproject/internal/auth"
    "myproject/pkg/config"
)`
	fmt.Println(importExample)
	fmt.Println()
}

// Example 9: Blank imports (side effects)
func blankImportExample() {
	fmt.Println(Header("9. Blank Imports"))

	// Blank import example (already done with _ "time/tzdata")
	fmt.Println("Blank import usage:")
	fmt.Println(`import _ "time/tzdata"`)
	fmt.Println("This imports tzdata for timezone support")
	fmt.Println()

	// Common blank import scenarios
	fmt.Println("Common blank import scenarios:")
	fmt.Println("1. Database drivers:")
	fmt.Println(`   import _ "github.com/go-sql-driver/mysql"`)
	fmt.Println()
	fmt.Println("2. Image format support:")
	fmt.Println(`   import _ "image/png"`)
	fmt.Println(`   import _ "image/jpeg"`)
	fmt.Println()
	fmt.Println("3. Plugin registration:")
	fmt.Println(`   import _ "myproject/plugins/auth"`)
	fmt.Println()
}

// Example 10: Package testing organization
func packageTestingExample() {
	fmt.Println(Header("10. Package Testing"))

	// Testing package organization
	fmt.Println("Testing package organization:")
	fmt.Println("1. Same package testing (package main)")
	fmt.Println("2. Black box testing (package main_test)")
	fmt.Println()

	// Example test structure
	testStructure := `
auth/
├── auth.go
├── auth_test.go          // same package testing
├── auth_integration_test.go // black box testing
└── testdata/
    └── fixtures.json
`
	fmt.Println("Test structure example:")
	fmt.Println(testStructure)

	// Demonstrate testing concepts
	config := NewConfig()
	config.SetAPIKey("test-key")

	if err := validateConfig(config); err != nil {
		fmt.Printf("Test failed: %v\n", err)
	} else {
		fmt.Println("Test passed: Config validation successful")
	}
	fmt.Println()
}

// Helper functions
func getCurrentDir() string {
	if dir, err := os.Getwd(); err == nil {
		return filepath.Base(dir)
	}
	return "unknown"
}

// Additional init function (demonstrates multiple init functions)
func init() {
	// This will run after the previous init function
	if strings.Contains(os.Args[0], "test") {
		fmt.Println(InfoText("Test mode detected"))
	}
}
