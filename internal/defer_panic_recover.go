// defer_panic_recover.go
package internal

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// RunDeferPanicRecoverExamples - main function to run all defer, panic, and recover examples
func RunDeferPanicRecoverExamples() {
	fmt.Println(SectionTitle("⏰ Defer Examples:"))
	basicDeferExample()
	deferOrderExample()
	deferWithLoopsExample()
	deferWithResourcesExample()

	fmt.Println(SectionTitle("🚨 Panic Examples:"))
	basicPanicExample()
	panicWithDeferExample()

	fmt.Println(SectionTitle("🛡️ Recover Examples:"))
	basicRecoverExample()
	recoverWithCleanupExample()
	recoverPatternExample()
	advancedErrorHandlingExample()
}

// basicDeferExample - demonstrates basic defer usage
func basicDeferExample() {
	fmt.Println(BoldText("1. Basic Defer Usage:"))

	func() {
		fmt.Println("  Function start")
		defer fmt.Println("  Deferred: This runs at function end")
		fmt.Println("  Function middle")
		defer fmt.Println("  Deferred: This runs second-to-last")
		fmt.Println("  Function end")
	}()

	// Defer with variables (captured at defer time)
	func() {
		x := 10
		defer fmt.Printf("  Deferred: x = %d (captured at defer time)\n", x)
		x = 20
		fmt.Printf("  Current: x = %d\n", x)
	}()

	// Defer with function return
	result := deferReturnExample()
	fmt.Printf("  Function returned: %d\n", result)

	fmt.Println()
}

// deferReturnExample - demonstrates defer with return values
func deferReturnExample() (result int) {
	result = 10
	defer func() {
		fmt.Printf("  Deferred: result = %d (before modification)\n", result)
		result = 20 // This modifies the return value
		fmt.Printf("  Deferred: result = %d (after modification)\n", result)
	}()
	return result
}

// deferOrderExample - demonstrates defer execution order (LIFO)
func deferOrderExample() {
	fmt.Println(BoldText("2. Defer Execution Order (LIFO):"))

	func() {
		fmt.Println("  Function start")
		defer fmt.Println("  Defer 1: First deferred")
		defer fmt.Println("  Defer 2: Second deferred")
		defer fmt.Println("  Defer 3: Third deferred")
		fmt.Println("  Function end")
	}()

	// Defer with loop variables
	fmt.Println("  Defer in loop (common mistake):")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("    Loop defer (wrong): %d\n", i) // All will print 3
	}

	fmt.Println("  Defer in loop (correct way):")
	for i := 0; i < 3; i++ {
		func(val int) {
			defer fmt.Printf("    Loop defer (correct): %d\n", val)
		}(i)
	}

	fmt.Println()
}

// deferWithLoopsExample - demonstrates defer with loops and closures
func deferWithLoopsExample() {
	fmt.Println(BoldText("3. Defer with Loops and Closures:"))

	// Problem: defer in loop
	fmt.Println("  Problem - defer in loop:")
	slice := []int{1, 2, 3}
	for i, v := range slice {
		defer fmt.Printf("    Index: %d, Value: %d\n", i, v) // Will print final values
	}

	// Solution 1: Use closure with immediate invocation
	fmt.Println("  Solution 1 - closure with immediate invocation:")
	for i, v := range slice {
		func(index, value int) {
			defer fmt.Printf("    Index: %d, Value: %d\n", index, value)
		}(i, v)
	}

	// Solution 2: Use closure that captures current values
	fmt.Println("  Solution 2 - closure capturing current values:")
	for i, v := range slice {
		defer func(index, value int) {
			fmt.Printf("    Index: %d, Value: %d\n", index, value)
		}(i, v)
	}

	fmt.Println()
}

// deferWithResourcesExample - demonstrates defer for resource management
func deferWithResourcesExample() {
	fmt.Println(BoldText("4. Defer for Resource Management:"))

	// File handling example
	func() {
		fmt.Println("  File handling example:")
		file, err := os.Create("temp_example.txt")
		if err != nil {
			fmt.Printf("    Error creating file: %v\n", err)
			return
		}
		defer func() {
			file.Close()
			os.Remove("temp_example.txt") // Clean up
			fmt.Println("    File closed and removed")
		}()

		file.WriteString("Hello, defer!")
		fmt.Println("    File written successfully")
	}()

	// Timer example
	func() {
		fmt.Println("  Timer example:")
		start := time.Now()
		defer func() {
			duration := time.Since(start)
			fmt.Printf("    Function took: %v\n", duration)
		}()

		time.Sleep(10 * time.Millisecond)
		fmt.Println("    Some work done")
	}()

	// Mutex example (conceptual)
	fmt.Println("  Mutex pattern (conceptual):")
	fmt.Println("    // mutex.Lock()")
	fmt.Println("    // defer mutex.Unlock()")
	fmt.Println("    // Critical section code here")

	fmt.Println()
}

// basicPanicExample - demonstrates basic panic usage
func basicPanicExample() {
	fmt.Println(BoldText("5. Basic Panic Usage:"))

	// Panic with string
	fmt.Println("  Example 1 - Panic with string:")
	func() {
		defer fmt.Println("    Deferred: This runs even during panic")
		fmt.Println("    Before panic")
		// Uncomment next line to see panic
		// panic("Something went wrong!")
		fmt.Println("    This would run if no panic")
	}()

	// Panic with custom error
	fmt.Println("  Example 2 - Panic with custom error:")
	func() {
		defer fmt.Println("    Deferred: Cleanup during panic")
		fmt.Println("    Doing some work...")
		// Simulate conditional panic
		condition := false
		if condition {
			panic(fmt.Errorf("custom error: invalid condition"))
		}
		fmt.Println("    Work completed successfully")
	}()

	// Runtime panic example
	fmt.Println("  Example 3 - Runtime panic (slice out of bounds):")
	func() {
		defer fmt.Println("    Deferred: Handling runtime panic")
		slice := []int{1, 2, 3}
		fmt.Printf("    Slice: %v\n", slice)
		// Uncomment next line to trigger panic
		// fmt.Printf("    Element at index 10: %d\n", slice[10])
		fmt.Println("    No panic occurred")
	}()

	fmt.Println()
}

// panicWithDeferExample - demonstrates panic with defer
func panicWithDeferExample() {
	fmt.Println(BoldText("6. Panic with Defer:"))

	func() {
		defer fmt.Println("    Defer 1: First deferred")
		defer fmt.Println("    Defer 2: Second deferred")
		defer fmt.Println("    Defer 3: Third deferred")

		fmt.Println("    Function start")
		// All defers will execute in reverse order during panic
		// Uncomment next line to see panic with defers
		// panic("Panic with multiple defers!")
		fmt.Println("    Function end (no panic)")
	}()

	fmt.Println()
}

// basicRecoverExample - demonstrates basic recover usage
func basicRecoverExample() {
	fmt.Println(BoldText("7. Basic Recover Usage:"))

	// Recover from panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("    Recovered from panic: %v\n", r)
			}
		}()

		fmt.Println("    Before panic")
		panic("This is a test panic!")
		fmt.Println("    This line won't execute")
	}()

	fmt.Println("    Execution continues after recovered panic")

	// Recover only works in defer
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("    Panic recovered: %v\n", r)
				// Can get stack trace
				fmt.Printf("    Stack trace:\n")
				stack := make([]byte, 1024)
				runtime.Stack(stack, false)
				fmt.Printf("    %s\n", stack[:200]) // Show first 200 bytes
			}
		}()

		fmt.Println("    About to panic...")
		panic("Another test panic!")
	}()

	fmt.Println()
}

// recoverWithCleanupExample - demonstrates recover with cleanup
func recoverWithCleanupExample() {
	fmt.Println(BoldText("8. Recover with Cleanup:"))

	processWithCleanup := func(data []int) {
		defer func() {
			fmt.Println("    Cleanup: Closing resources")
			if r := recover(); r != nil {
				fmt.Printf("    Panic during processing: %v\n", r)
				fmt.Println("    Cleanup completed despite panic")
			}
		}()

		fmt.Println("    Processing data...")
		for i, v := range data {
			if v < 0 {
				panic(fmt.Sprintf("negative value at index %d: %d", i, v))
			}
			fmt.Printf("    Processing: %d\n", v)
		}
		fmt.Println("    Processing completed successfully")
	}

	// Success case
	fmt.Println("  Success case:")
	processWithCleanup([]int{1, 2, 3})

	// Panic case
	fmt.Println("  Panic case:")
	processWithCleanup([]int{1, -2, 3})

	fmt.Println("    Main function continues...")

	fmt.Println()
}

// recoverPatternExample - demonstrates common recover patterns
func recoverPatternExample() {
	fmt.Println(BoldText("9. Common Recover Patterns:"))

	// Pattern 1: Convert panic to error
	safeFunction := func() (result int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("function panicked: %v", r)
			}
		}()

		// Simulate work that might panic
		panic("simulated error")
		return 42, nil
	}

	if result, err := safeFunction(); err != nil {
		fmt.Printf("  Pattern 1 - Error returned: %v\n", err)
	} else {
		fmt.Printf("  Pattern 1 - Result: %d\n", result)
	}

	// Pattern 2: Selective recovery
	selectiveRecover := func() {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case string:
					fmt.Printf("  Pattern 2 - String panic: %s\n", v)
				case error:
					fmt.Printf("  Pattern 2 - Error panic: %v\n", v)
				default:
					fmt.Printf("  Pattern 2 - Unknown panic type: %v\n", v)
					// Re-panic for unknown types
					panic(r)
				}
			}
		}()

		panic("This is a string panic")
	}

	selectiveRecover()

	// Pattern 3: Graceful shutdown
	gracefulShutdown := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Pattern 3 - Graceful shutdown due to: %v\n", r)
				fmt.Println("  Pattern 3 - Performing cleanup...")
				fmt.Println("  Pattern 3 - Shutdown completed")
			}
		}()

		panic("Critical system error")
	}

	gracefulShutdown()

	fmt.Println()
}

// CustomValidationError - custom error type for validation errors
type CustomValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface for CustomValidationError
func (e CustomValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// advancedErrorHandlingExample - demonstrates advanced error handling patterns
func advancedErrorHandlingExample() {
	fmt.Println(BoldText("10. Advanced Error Handling Patterns:"))

	// Robust error handling function
	processData := func(data map[string]interface{}) (err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case CustomValidationError:
					err = v
				case error:
					err = fmt.Errorf("processing error: %w", v)
				default:
					err = fmt.Errorf("unexpected panic: %v", r)
				}
			}
		}()

		// Validate required fields
		if name, ok := data["name"]; !ok || name == "" {
			panic(CustomValidationError{Field: "name", Message: "is required"})
		}

		if age, ok := data["age"]; !ok {
			panic(CustomValidationError{Field: "age", Message: "is required"})
		} else if ageInt, ok := age.(int); !ok || ageInt < 0 {
			panic(CustomValidationError{Field: "age", Message: "must be a non-negative integer"})
		}

		fmt.Printf("  Processing data: %v\n", data)
		return nil
	}

	// Test cases
	testCases := []map[string]interface{}{
		{"name": "Alice", "age": 30},          // Valid
		{"name": "", "age": 25},               // Invalid name
		{"name": "Bob"},                       // Missing age
		{"name": "Charlie", "age": "invalid"}, // Invalid age type
		{"name": "David", "age": -5},          // Invalid age value
	}

	for i, testCase := range testCases {
		fmt.Printf("  Test case %d: %v\n", i+1, testCase)
		if err := processData(testCase); err != nil {
			fmt.Printf("    Error: %v\n", err)
		} else {
			fmt.Printf("    Success\n")
		}
	}

	fmt.Println()
}

// Helper functions (you'll need to implement these or import them)
func SectionTitle(text string) string {
	return fmt.Sprintf("\n=== %s ===\n", text)
}

func BoldText(text string) string {
	return fmt.Sprintf("** %s **", text)
}
