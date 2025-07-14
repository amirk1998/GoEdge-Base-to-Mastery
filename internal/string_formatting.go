// string_formatting.go
package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// RunStringFormattingExamples - main function to run all string formatting examples
func RunStringFormattingExamples() {
	fmt.Println(Subtitle("📝 String Formatting Examples:"))
	basicFormattingExample()
	numericFormattingExample()
	stringManipulationExample()
	advancedFormattingExample()
	stringConversionExample()
	unicodeStringExample()
	stringBuilderExample()
	stringTemplateExample()
}

func basicFormattingExample() {
	fmt.Println(InfoText("1. Basic String Formatting:"))

	name := "Alice"
	age := 30
	height := 5.6
	isStudent := false

	// Basic Printf formatting
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Height: %.1f feet\n", height)
	fmt.Printf("Is student: %t\n", isStudent)

	// Sprintf for string creation
	description := fmt.Sprintf("%s is %d years old and %.1f feet tall", name, age, height)
	fmt.Printf("Description: %s\n", description)

	// Different verb demonstrations
	value := 42
	fmt.Printf("Decimal: %d\n", value)
	fmt.Printf("Binary: %b\n", value)
	fmt.Printf("Octal: %o\n", value)
	fmt.Printf("Hexadecimal: %x\n", value)
	fmt.Printf("Hexadecimal (upper): %X\n", value)
	fmt.Printf("Character: %c\n", value)
	fmt.Printf("Quoted: %q\n", value)

	// Type and value information
	fmt.Printf("Type: %T, Value: %v\n", value, value)
	fmt.Printf("Go syntax: %#v\n", value)
}

func numericFormattingExample() {
	fmt.Println(InfoText("2. Numeric Formatting:"))

	// Integer formatting
	number := 12345
	fmt.Printf("Default: %d\n", number)
	fmt.Printf("With width: %8d\n", number)
	fmt.Printf("Left-aligned: %-8d|\n", number)
	fmt.Printf("Zero-padded: %08d\n", number)
	fmt.Printf("With plus sign: %+d\n", number)
	fmt.Printf("With space: % d\n", number)

	// Float formatting
	pi := 3.14159265359
	fmt.Printf("Default float: %f\n", pi)
	fmt.Printf("2 decimal places: %.2f\n", pi)
	fmt.Printf("Scientific notation: %e\n", pi)
	fmt.Printf("Scientific (upper): %E\n", pi)
	fmt.Printf("Compact format: %g\n", pi)
	fmt.Printf("Width and precision: %10.2f\n", pi)
	fmt.Printf("Zero-padded float: %08.2f\n", pi)

	// Currency formatting simulation
	price := 1234.56
	fmt.Printf("Price: $%.2f\n", price)
	fmt.Printf("Price with commas: $%,.2f\n", price) // Note: Go doesn't have built-in comma formatting

	// Percentage formatting
	ratio := 0.85
	fmt.Printf("Success rate: %.1f%%\n", ratio*100)

	// Large numbers
	bigNumber := 1234567890
	fmt.Printf("Big number: %d\n", bigNumber)
	fmt.Printf("Big number with separators: %,d\n", bigNumber) // Custom implementation needed
}

func stringManipulationExample() {
	fmt.Println(InfoText("3. String Manipulation:"))

	text := "  Hello, World!  "
	fmt.Printf("Original: '%s'\n", text)

	// Basic string operations
	fmt.Printf("Length: %d\n", len(text))
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("Upper case: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower case: %s\n", strings.ToLower(text))
	fmt.Printf("Title case: %s\n", strings.Title(strings.ToLower(text)))

	// String replacement
	message := "Hello, World! Welcome to the World of Go!"
	fmt.Printf("Original: %s\n", message)
	fmt.Printf("Replace first: %s\n", strings.Replace(message, "World", "Universe", 1))
	fmt.Printf("Replace all: %s\n", strings.ReplaceAll(message, "World", "Universe"))

	// String splitting and joining
	csv := "apple,banana,orange,grape"
	fruits := strings.Split(csv, ",")
	fmt.Printf("Split: %v\n", fruits)
	fmt.Printf("Joined with ' | ': %s\n", strings.Join(fruits, " | "))

	// String contains and searching
	text2 := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("Contains 'fox': %t\n", strings.Contains(text2, "fox"))
	fmt.Printf("Starts with 'The': %t\n", strings.HasPrefix(text2, "The"))
	fmt.Printf("Ends with 'dog': %t\n", strings.HasSuffix(text2, "dog"))
	fmt.Printf("Index of 'fox': %d\n", strings.Index(text2, "fox"))
	fmt.Printf("Last index of 'o': %d\n", strings.LastIndex(text2, "o"))

	// String repetition
	pattern := "Go! "
	fmt.Printf("Repeated: %s\n", strings.Repeat(pattern, 5))
}

func advancedFormattingExample() {
	fmt.Println(InfoText("4. Advanced Formatting:"))

	// Struct formatting
	type Person struct {
		Name   string
		Age    int
		Email  string
		Active bool
	}

	person := Person{
		Name:   "John Doe",
		Age:    35,
		Email:  "john@example.com",
		Active: true,
	}

	fmt.Printf("Struct default: %v\n", person)
	fmt.Printf("Struct detailed: %+v\n", person)
	fmt.Printf("Struct Go syntax: %#v\n", person)

	// Pointer formatting
	ptr := &person
	fmt.Printf("Pointer: %p\n", ptr)
	fmt.Printf("Pointer value: %v\n", *ptr)

	// Slice and map formatting
	numbers := []int{1, 2, 3, 4, 5}
	grades := map[string]int{"Alice": 90, "Bob": 85}

	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("Slice detailed: %#v\n", numbers)
	fmt.Printf("Map: %v\n", grades)
	fmt.Printf("Map detailed: %#v\n", grades)

	// Custom formatting with width and alignment
	fmt.Printf("%-20s | %5s | %10s\n", "Name", "Age", "Status")
	fmt.Printf("%-20s | %5d | %10s\n", person.Name, person.Age, getStatus(person.Active))
	fmt.Printf("%-20s | %5s | %10s\n", strings.Repeat("-", 20), strings.Repeat("-", 5), strings.Repeat("-", 10))

	// Time formatting
	now := time.Now()
	fmt.Printf("Time default: %v\n", now)
	fmt.Printf("Time custom: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Date only: %s\n", now.Format("January 2, 2006"))
	fmt.Printf("Time only: %s\n", now.Format("3:04 PM"))
}

func stringConversionExample() {
	fmt.Println(InfoText("5. String Conversions:"))

	// Number to string conversions
	intVal := 42
	floatVal := 3.14159
	boolVal := true

	fmt.Printf("Int to string: %s\n", strconv.Itoa(intVal))
	fmt.Printf("Float to string: %s\n", strconv.FormatFloat(floatVal, 'f', 2, 64))
	fmt.Printf("Bool to string: %s\n", strconv.FormatBool(boolVal))

	// String to number conversions
	strInt := "123"
	strFloat := "3.14"
	strBool := "true"

	if intResult, err := strconv.Atoi(strInt); err == nil {
		fmt.Printf("String to int: %d\n", intResult)
	}

	if floatResult, err := strconv.ParseFloat(strFloat, 64); err == nil {
		fmt.Printf("String to float: %.2f\n", floatResult)
	}

	if boolResult, err := strconv.ParseBool(strBool); err == nil {
		fmt.Printf("String to bool: %t\n", boolResult)
	}

	// Quote and unquote
	original := `Hello "World" with 'quotes'`
	quoted := strconv.Quote(original)
	fmt.Printf("Quoted: %s\n", quoted)

	if unquoted, err := strconv.Unquote(quoted); err == nil {
		fmt.Printf("Unquoted: %s\n", unquoted)
	}

	// Base conversions
	number := 255
	fmt.Printf("Binary: %s\n", strconv.FormatInt(int64(number), 2))
	fmt.Printf("Octal: %s\n", strconv.FormatInt(int64(number), 8))
	fmt.Printf("Hex: %s\n", strconv.FormatInt(int64(number), 16))
}

func unicodeStringExample() {
	fmt.Println(InfoText("6. Unicode String Handling:"))

	// Unicode strings
	text := "Hello, 世界! 🌍 Golang"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Byte length: %d\n", len(text))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))

	// Iterate over runes
	fmt.Print("Runes: ")
	for i, r := range text {
		fmt.Printf("[%d:%c] ", i, r)
	}
	fmt.Println()

	// Iterate over bytes
	fmt.Print("Bytes: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%02x ", text[i])
	}
	fmt.Println()

	// Unicode normalization and validation
	fmt.Printf("Valid UTF-8: %t\n", utf8.ValidString(text))

	// Substring with unicode awareness
	runes := []rune(text)
	if len(runes) >= 7 {
		substring := string(runes[0:7])
		fmt.Printf("First 7 runes: %s\n", substring)
	}

	// Working with emojis
	emojis := "🚀🎯🔥💯⭐"
	fmt.Printf("Emoji string: %s\n", emojis)
	fmt.Printf("Emoji byte length: %d\n", len(emojis))
	fmt.Printf("Emoji rune count: %d\n", utf8.RuneCountInString(emojis))
}

func stringBuilderExample() {
	fmt.Println(InfoText("7. Efficient String Building:"))

	// Using strings.Builder for efficient string concatenation
	var builder strings.Builder

	// Build a large string efficiently
	builder.WriteString("Building a string efficiently:\n")

	for i := 1; i <= 5; i++ {
		builder.WriteString(fmt.Sprintf("Line %d: ", i))
		builder.WriteString("Some content here")
		if i < 5 {
			builder.WriteString("\n")
		}
	}

	result := builder.String()
	fmt.Printf("Built string:\n%s\n", result)
	fmt.Printf("Builder length: %d\n", builder.Len())

	// Performance comparison demonstration (conceptual)
	fmt.Println("\nString concatenation methods:")

	// Method 1: Using += (inefficient for large strings)
	str1 := ""
	for i := 0; i < 3; i++ {
		str1 += fmt.Sprintf("Part %d ", i)
	}
	fmt.Printf("Method 1 (+=): %s\n", str1)

	// Method 2: Using strings.Builder (efficient)
	var builder2 strings.Builder
	for i := 0; i < 3; i++ {
		builder2.WriteString(fmt.Sprintf("Part %d ", i))
	}
	fmt.Printf("Method 2 (Builder): %s\n", builder2.String())

	// Method 3: Using slice and Join (efficient for known size)
	parts := make([]string, 3)
	for i := 0; i < 3; i++ {
		parts[i] = fmt.Sprintf("Part %d", i)
	}
	str3 := strings.Join(parts, " ")
	fmt.Printf("Method 3 (Join): %s\n", str3)
}

func stringTemplateExample() {
	fmt.Println(InfoText("8. String Templates and Patterns:"))

	// Simple template replacement
	template := "Hello, {name}! Welcome to {place}."

	replacements := map[string]string{
		"{name}":  "Alice",
		"{place}": "Golang World",
	}

	result := template
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	fmt.Printf("Template: %s\n", template)
	fmt.Printf("Result: %s\n", result)

	// Email template example
	emailTemplate := `
Subject: Welcome {name}!

Dear {name},

Thank you for joining {company}. Your account has been created successfully.

Best regards,
{company} Team
`

	emailData := map[string]string{
		"{name}":    "John Doe",
		"{company}": "TechCorp",
	}

	email := emailTemplate
	for placeholder, value := range emailData {
		email = strings.ReplaceAll(email, placeholder, value)
	}

	fmt.Printf("Generated email:%s\n", email)

	// URL building
	baseURL := "https://api.example.com"
	endpoint := "/users"
	userID := "123"

	fullURL := fmt.Sprintf("%s%s/%s", baseURL, endpoint, userID)
	fmt.Printf("Built URL: %s\n", fullURL)

	// Query parameter building
	params := map[string]string{
		"format": "json",
		"limit":  "10",
		"offset": "0",
	}

	var queryParts []string
	for key, value := range params {
		queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
	}

	queryString := strings.Join(queryParts, "&")
	fullURLWithParams := fmt.Sprintf("%s?%s", fullURL, queryString)
	fmt.Printf("URL with params: %s\n", fullURLWithParams)
}

// Helper function for status formatting
func getStatus(active bool) string {
	if active {
		return "Active"
	}
	return "Inactive"
}
