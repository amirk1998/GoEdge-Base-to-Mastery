package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Custom error types for advanced examples
type ValidationError struct {
	Field   string
	Message string
	Code    int
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s (code: %d)", e.Field, e.Message, e.Code)
}

type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on table %s: %v", e.Operation, e.Table, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// User struct for examples
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Basic error handling example
func basicErrorExample() {
	fmt.Println("=== Basic Error Handling ===")

	// Example 1: Simple error creation and handling
	result, err := divideNumbers(10, 0)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)

	// Example 2: Successful operation
	result2, err2 := divideNumbers(10, 2)
	if err2 != nil {
		fmt.Printf("Error occurred: %v\n", err2)
		return
	}
	fmt.Printf("Result: %.2f\n", result2)
}

func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Error creation methods
func errorCreationExample() {
	fmt.Println("\n=== Error Creation Methods ===")

	// Method 1: Using errors.New()
	err1 := errors.New("this is a simple error")
	fmt.Printf("errors.New(): %v\n", err1)

	// Method 2: Using fmt.Errorf()
	username := "john_doe"
	err2 := fmt.Errorf("user %s not found", username)
	fmt.Printf("fmt.Errorf(): %v\n", err2)

	// Method 3: Creating custom error
	err3 := &ValidationError{
		Field:   "email",
		Message: "invalid email format",
		Code:    400,
	}
	fmt.Printf("Custom error: %v\n", err3)
}

// Multiple return values with error
func multipleReturnExample() {
	fmt.Println("\n=== Multiple Return Values ===")

	// Example: Function that can fail
	user, err := getUserByID(123)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return
	}

	fmt.Printf("User found: %+v\n", user)

	// Example: Function that succeeds
	user2, err2 := getUserByID(1)
	if err2 != nil {
		fmt.Printf("Failed to get user: %v\n", err2)
		return
	}

	fmt.Printf("User found: %+v\n", user2)
}

func getUserByID(id int) (*User, error) {
	// Simulate database lookup
	if id == 1 {
		return &User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		}, nil
	}

	return nil, fmt.Errorf("user with ID %d not found", id)
}

// Error wrapping example
func errorWrappingExample() {
	fmt.Println("\n=== Error Wrapping ===")

	err := processUserData(999)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		// Check if it's a specific error type
		var dbErr *DatabaseError
		if errors.As(err, &dbErr) {
			fmt.Printf("Database operation failed: %s on table %s\n", dbErr.Operation, dbErr.Table)
		}
	}
}

func processUserData(userID int) error {
	err := fetchUserFromDatabase(userID)
	if err != nil {
		return fmt.Errorf("failed to process user data: %w", err)
	}
	return nil
}

func fetchUserFromDatabase(userID int) error {
	// Simulate database error
	originalErr := errors.New("connection timeout")
	return &DatabaseError{
		Operation: "SELECT",
		Table:     "users",
		Err:       originalErr,
	}
}

// Error checking patterns
func errorCheckingExample() {
	fmt.Println("\n=== Error Checking Patterns ===")

	// Pattern 1: Early return
	result, err := validateAndProcess("john@example.com")
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}
	fmt.Printf("Processing result: %s\n", result)

	// Pattern 2: Error accumulation
	errors := validateUser(&User{
		Name:  "",
		Email: "invalid-email",
		Age:   -5,
	})

	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	}
}

func validateAndProcess(email string) (string, error) {
	if email == "" {
		return "", errors.New("email cannot be empty")
	}

	if !strings.Contains(email, "@") {
		return "", errors.New("invalid email format")
	}

	return "Email processed successfully", nil
}

func validateUser(user *User) []error {
	var errors []error

	if user.Name == "" {
		errors = append(errors, &ValidationError{
			Field:   "name",
			Message: "name cannot be empty",
			Code:    400,
		})
	}

	if !strings.Contains(user.Email, "@") {
		errors = append(errors, &ValidationError{
			Field:   "email",
			Message: "invalid email format",
			Code:    400,
		})
	}

	if user.Age < 0 {
		errors = append(errors, &ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
			Code:    400,
		})
	}

	return errors
}

// File operations with error handling
func fileOperationsExample() {
	fmt.Println("\n=== File Operations with Error Handling ===")

	// Example: Reading a file
	content, err := readFileContent("example.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	fmt.Printf("File content: %s\n", content)

	// Example: Writing to a file
	err = writeFileContent("output.txt", "Hello, World!")
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return
	}

	fmt.Println("File written successfully")
}

func readFileContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file: %v\n", closeErr)
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %w", err)
	}

	return string(content), nil
}

func writeFileContent(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file: %v\n", closeErr)
		}
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to file: %w", err)
	}

	return nil
}

// Advanced error handling with panic and recover
func panicAndRecoverExample() {
	fmt.Println("\n=== Panic and Recover Example ===")

	// Example of handling panics
	err := safeOperation()
	if err != nil {
		fmt.Printf("Operation failed safely: %v\n", err)
	}

	// Example of successful operation
	err2 := safeOperation2()
	if err2 != nil {
		fmt.Printf("Operation failed: %v\n", err2)
	} else {
		fmt.Println("Operation completed successfully")
	}
}

func safeOperation() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	// This will panic
	riskyOperation(0)
	return nil
}

func safeOperation2() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	// This will not panic
	riskyOperation(10)
	return nil
}

func riskyOperation(value int) {
	if value == 0 {
		panic("cannot process zero value")
	}
	fmt.Printf("Processing value: %d\n", value)
}

// Error handling with timeouts
func timeoutExample() {
	fmt.Println("\n=== Timeout Error Handling ===")

	// Example: Operation with timeout
	result, err := operationWithTimeout(2 * time.Second)
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return
	}

	fmt.Printf("Operation result: %s\n", result)
}

func operationWithTimeout(timeout time.Duration) (string, error) {
	done := make(chan string, 1)

	go func() {
		// Simulate long-running operation
		time.Sleep(3 * time.Second)
		done <- "Operation completed"
	}()

	select {
	case result := <-done:
		return result, nil
	case <-time.After(timeout):
		return "", errors.New("operation timeout")
	}
}

// Error handling with type assertions
func errorTypeAssertionExample() {
	fmt.Println("\n=== Error Type Assertion ===")

	// Example: Checking specific error types
	err := performOperation("invalid")
	if err != nil {
		handleSpecificError(err)
	}

	err2 := performOperation("valid")
	if err2 != nil {
		handleSpecificError(err2)
	} else {
		fmt.Println("Operation performed successfully")
	}
}

func performOperation(input string) error {
	if input == "invalid" {
		return &ValidationError{
			Field:   "input",
			Message: "input value is invalid",
			Code:    400,
		}
	}
	return nil
}

func handleSpecificError(err error) {
	// Method 1: Type assertion
	if validationErr, ok := err.(*ValidationError); ok {
		fmt.Printf("Validation error: Field=%s, Code=%d\n", validationErr.Field, validationErr.Code)
		return
	}

	// Method 2: Using errors.As
	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("Database error: Operation=%s, Table=%s\n", dbErr.Operation, dbErr.Table)
		return
	}

	// Default case
	fmt.Printf("Unknown error: %v\n", err)
}

// Best practices example
func bestPracticesExample() {
	fmt.Println("\n=== Best Practices Example ===")

	// Example: Proper error handling in a service
	service := &UserService{}

	user, err := service.CreateUser("John Doe", "john@example.com", 25)
	if err != nil {
		fmt.Printf("Failed to create user: %v\n", err)
		return
	}

	fmt.Printf("User created successfully: %+v\n", user)
}

type UserService struct{}

func (s *UserService) CreateUser(name, email string, age int) (*User, error) {
	// Validate input
	if err := s.validateInput(name, email, age); err != nil {
		return nil, fmt.Errorf("input validation failed: %w", err)
	}

	// Check if user exists
	exists, err := s.userExists(email)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil, &ValidationError{
			Field:   "email",
			Message: "user with this email already exists",
			Code:    409,
		}
	}

	// Create user
	user := &User{
		ID:    generateID(),
		Name:  name,
		Email: email,
		Age:   age,
	}

	// Save to database
	if err := s.saveUser(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *UserService) validateInput(name, email string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "name cannot be empty", Code: 400}
	}

	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Message: "invalid email format", Code: 400}
	}

	if age < 0 || age > 150 {
		return &ValidationError{Field: "age", Message: "age must be between 0 and 150", Code: 400}
	}

	return nil
}

func (s *UserService) userExists(email string) (bool, error) {
	// Simulate database check
	if email == "existing@example.com" {
		return true, nil
	}
	return false, nil
}

func (s *UserService) saveUser(user *User) error {
	// Simulate database save
	fmt.Printf("Saving user to database: %+v\n", user)
	return nil
}

func generateID() int {
	return int(time.Now().UnixNano() % 1000000)
}

// Convert string to int with error handling
func stringToIntExample() {
	fmt.Println("\n=== String to Int Conversion ===")

	numbers := []string{"123", "456", "abc", "789"}

	for _, numStr := range numbers {
		if num, err := strconv.Atoi(numStr); err != nil {
			fmt.Printf("Failed to convert '%s' to int: %v\n", numStr, err)
		} else {
			fmt.Printf("Converted '%s' to int: %d\n", numStr, num)
		}
	}
}

// Main function to run all examples
func RunErrorHandlingExamples() {
	fmt.Println("Go Error Handling Examples")
	fmt.Println("==========================")

	basicErrorExample()
	errorCreationExample()
	multipleReturnExample()
	errorWrappingExample()
	errorCheckingExample()
	fileOperationsExample()
	panicRecoverExample()
	timeoutExample()
	errorTypeAssertionExample()
	bestPracticesExample()
	stringToIntExample()

	fmt.Println("\n=== Error Handling Examples Completed ===")
}
