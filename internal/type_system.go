// type_system.go
package internal

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Custom types based on built-in types
type AccountID int
type EmailAddr string
type TempValue float64
type AccountStatus string

// Type aliases (Go 1.9+)
type MyInteger = int
type StringType = string

// Constants with custom types
const (
	StatusActive   AccountStatus = "active"
	StatusInactive AccountStatus = "inactive"
	StatusPending  AccountStatus = "pending"
)

// Struct types
type Account struct {
	ID    AccountID
	Email EmailAddr
	Name  string
}

type Item struct {
	ID    int
	Name  string
	Price float64
}

// Interface types
type StringRenderer interface {
	String() string
}

type DataValidator interface {
	Validate() bool
}

// Method on custom type
func (a AccountID) String() string {
	return fmt.Sprintf("Account-%d", int(a))
}

func (e EmailAddr) String() string {
	return string(e)
}

func (e EmailAddr) Validate() bool {
	return len(e) > 0 && strings.Contains(string(e), "@")
}

func (t TempValue) Celsius() float64 {
	return float64(t)
}

func (t TempValue) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t TempValue) String() string {
	return fmt.Sprintf("%.2f°C", t.Celsius())
}

func (a Account) String() string {
	return fmt.Sprintf("Account{ID: %s, Email: %s, Name: %s}",
		a.ID.String(), a.Email.String(), a.Name)
}

func (a Account) Validate() bool {
	return a.ID > 0 && a.Email.Validate() && len(a.Name) > 0
}

// RunTypeSystemDemo - main function to run all type system examples
func RunTypeSystemDemo() {
	customTypeDemo()
	typeAliasDemo()
	typeConversionDemo()
	typeAssertionDemo()
	typeSwitchDemo()
	underlyingTypeDemo()
	methodOnCustomTypeDemo()
	typeEmbeddingDemo()
	interfaceTypeDemo()
	reflectionTypeDemo()
}

// Example 1: Custom types and their benefits
func customTypeDemo() {
	fmt.Println(createHeader("1. Custom Types"))

	// Using custom types prevents mixing different concepts
	var accountID AccountID = 123
	var email EmailAddr = "john@example.com"
	var temp TempValue = 25.5

	fmt.Printf("AccountID: %s\n", accountID)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Temperature: %s (%.2f°F)\n", temp, temp.Fahrenheit())

	// Type safety - this would cause compilation error:
	// var anotherID AccountID = 456
	// if accountID == anotherID { } // OK - same types
	// if accountID == 456 { } // ERROR - different types

	fmt.Printf("Email is valid: %t\n", email.Validate())
	fmt.Println()
}

// Example 2: Type aliases vs custom types
func typeAliasDemo() {
	fmt.Println(createHeader("2. Type Aliases vs Custom Types"))

	// Type alias - same underlying type
	var aliasInt MyInteger = 42
	var normalInt int = 42

	// These are compatible (same type)
	aliasInt = normalInt
	normalInt = aliasInt

	fmt.Printf("Alias int: %d, Normal int: %d\n", aliasInt, normalInt)

	// Custom type - different type
	var customAccountID AccountID = 123
	var normalInt2 int = 123

	// These require explicit conversion
	customAccountID = AccountID(normalInt2)
	normalInt2 = int(customAccountID)

	fmt.Printf("Custom AccountID: %s, Normal int: %d\n", customAccountID, normalInt2)
	fmt.Println()
}

// Example 3: Type conversion and casting
func typeConversionDemo() {
	fmt.Println(createHeader("3. Type Conversion"))

	var i int = 42
	var f float64 = 3.14
	var s string = "123"

	// Basic conversions
	fmt.Printf("int to float64: %f\n", float64(i))
	fmt.Printf("float64 to int: %d\n", int(f))

	// String conversions
	intFromString, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("string to int: %d\n", intFromString)
	}

	stringFromInt := strconv.Itoa(i)
	fmt.Printf("int to string: %s\n", stringFromInt)

	// Custom type conversions
	var accountID AccountID = AccountID(i)
	var email EmailAddr = EmailAddr("user@example.com")

	fmt.Printf("Converted AccountID: %s\n", accountID)
	fmt.Printf("Converted Email: %s\n", email)
	fmt.Println()
}

// Example 4: Type assertions with interfaces
func typeAssertionDemo() {
	fmt.Println(createHeader("4. Type Assertions"))

	var val interface{} = "hello world"

	// Type assertion (unsafe)
	str := val.(string)
	fmt.Printf("Asserted string: %s\n", str)

	// Safe type assertion
	if str, ok := val.(string); ok {
		fmt.Printf("Safe assertion - string: %s\n", str)
	}

	if num, ok := val.(int); ok {
		fmt.Printf("Safe assertion - int: %d\n", num)
	} else {
		fmt.Println("Value is not an int")
	}

	// Type assertion with custom types
	var accountInterface interface{} = Account{ID: 1, Email: "test@example.com", Name: "John"}

	if account, ok := accountInterface.(Account); ok {
		fmt.Printf("Asserted Account: %s\n", account)
	}
	fmt.Println()
}

// Example 5: Type switches
func typeSwitchDemo() {
	fmt.Println(createHeader("5. Type Switches"))

	values := []interface{}{
		42,
		"hello",
		3.14,
		AccountID(123),
		EmailAddr("test@example.com"),
		Account{ID: 1, Name: "John", Email: "john@example.com"},
	}

	for i, v := range values {
		fmt.Printf("Value %d: ", i+1)

		switch val := v.(type) {
		case int:
			fmt.Printf("Integer: %d\n", val)
		case string:
			fmt.Printf("String: %s\n", val)
		case float64:
			fmt.Printf("Float: %.2f\n", val)
		case AccountID:
			fmt.Printf("AccountID: %s\n", val)
		case EmailAddr:
			fmt.Printf("Email: %s (valid: %t)\n", val, val.Validate())
		case Account:
			fmt.Printf("Account: %s (valid: %t)\n", val, val.Validate())
		default:
			fmt.Printf("Unknown type: %T\n", val)
		}
	}
	fmt.Println()
}

// Example 6: Underlying types and type definitions
func underlyingTypeDemo() {
	fmt.Println(createHeader("6. Underlying Types"))

	// Show underlying types
	var accountID AccountID = 123
	var email EmailAddr = "test@example.com"
	var temp TempValue = 25.5

	fmt.Printf("AccountID underlying type: %T\n", int(accountID))
	fmt.Printf("Email underlying type: %T\n", string(email))
	fmt.Printf("Temperature underlying type: %T\n", float64(temp))

	// Type alias has same underlying type
	var aliasInt MyInteger = 42
	fmt.Printf("MyInteger (alias) underlying type: %T\n", aliasInt)

	// Demonstrate type identity
	fmt.Printf("AccountID == int: %t\n", reflect.TypeOf(accountID) == reflect.TypeOf(int(0)))
	fmt.Printf("MyInteger == int: %t\n", reflect.TypeOf(aliasInt) == reflect.TypeOf(int(0)))
	fmt.Println()
}

// Example 7: Methods on custom types
func methodOnCustomTypeDemo() {
	fmt.Println(createHeader("7. Methods on Custom Types"))

	temp := TempValue(25.5)
	fmt.Printf("Temperature: %s\n", temp)
	fmt.Printf("In Fahrenheit: %.2f°F\n", temp.Fahrenheit())

	accountID := AccountID(123)
	fmt.Printf("AccountID: %s\n", accountID)

	email := EmailAddr("john.doe@example.com")
	fmt.Printf("Email: %s (valid: %t)\n", email, email.Validate())

	invalidEmail := EmailAddr("invalid-email")
	fmt.Printf("Invalid email: %s (valid: %t)\n", invalidEmail, invalidEmail.Validate())
	fmt.Println()
}

// Example 8: Type embedding (composition)
func typeEmbeddingDemo() {
	fmt.Println(createHeader("8. Type Embedding"))

	// Embedded struct
	type Location struct {
		Street string
		City   string
		ZIP    string
	}

	type DetailedAccount struct {
		Account  // Embedded type
		Location // Embedded type
		Age      int
	}

	detailedAccount := DetailedAccount{
		Account: Account{
			ID:    AccountID(456),
			Email: EmailAddr("jane@example.com"),
			Name:  "Jane Doe",
		},
		Location: Location{
			Street: "123 Main St",
			City:   "New York",
			ZIP:    "10001",
		},
		Age: 30,
	}

	// Access embedded fields directly
	fmt.Printf("Account ID: %s\n", detailedAccount.ID)
	fmt.Printf("Account Email: %s\n", detailedAccount.Email)
	fmt.Printf("Account Name: %s\n", detailedAccount.Name)
	fmt.Printf("Street: %s\n", detailedAccount.Street)
	fmt.Printf("City: %s\n", detailedAccount.City)
	fmt.Printf("Age: %d\n", detailedAccount.Age)

	// Embedded methods are promoted
	fmt.Printf("Account valid: %t\n", detailedAccount.Validate())
	fmt.Printf("Account string: %s\n", detailedAccount.Account.String())
	fmt.Println()
}

// Example 9: Interface types and implementations
func interfaceTypeDemo() {
	fmt.Println(createHeader("9. Interface Types"))

	// Multiple types implementing same interface
	var stringers []StringRenderer = []StringRenderer{
		AccountID(123),
		EmailAddr("test@example.com"),
		TempValue(25.5),
		Account{ID: 1, Name: "John", Email: "john@example.com"},
	}

	fmt.Println("StringRenderer implementations:")
	for i, s := range stringers {
		fmt.Printf("%d. %s\n", i+1, s.String())
	}

	// DataValidator interface
	var validators []DataValidator = []DataValidator{
		EmailAddr("valid@example.com"),
		EmailAddr("invalid-email"),
		Account{ID: 1, Name: "John", Email: "john@example.com"},
		Account{ID: 0, Name: "", Email: ""},
	}

	fmt.Println("\nDataValidator implementations:")
	for i, v := range validators {
		fmt.Printf("%d. Valid: %t\n", i+1, v.Validate())
	}
	fmt.Println()
}

// Example 10: Reflection with types
func reflectionTypeDemo() {
	fmt.Println(createHeader("10. Reflection and Types"))

	values := []interface{}{
		AccountID(123),
		EmailAddr("test@example.com"),
		TempValue(25.5),
		Account{ID: 1, Name: "John", Email: "john@example.com"},
	}

	for i, v := range values {
		t := reflect.TypeOf(v)
		val := reflect.ValueOf(v)

		fmt.Printf("Value %d:\n", i+1)
		fmt.Printf("  Type: %s\n", t.Name())
		fmt.Printf("  Kind: %s\n", t.Kind())
		fmt.Printf("  Package: %s\n", t.PkgPath())
		fmt.Printf("  String: %s\n", val.String())

		// Check if it implements interfaces
		stringerType := reflect.TypeOf((*StringRenderer)(nil)).Elem()
		validatorType := reflect.TypeOf((*DataValidator)(nil)).Elem()

		fmt.Printf("  Implements StringRenderer: %t\n", t.Implements(stringerType))
		fmt.Printf("  Implements DataValidator: %t\n", t.Implements(validatorType))
		fmt.Println()
	}
}

// createHeader helper function
func createHeader(title string) string {
	return "==== " + title + " ===="
}
