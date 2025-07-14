// functions.go
package internal

import "fmt"

// RunFunctionExamples - main function to run all function examples
func RunFunctionExamples() {
	basicFunctionExample()
	multiplePReturnExample()
	variableArgumentsExample()
	closureExample()
	higherOrderFunctionExample()
	anonymousFunctionExample()
	recursionExample()
	deferExample()
	panicRecoverExample()
}

// Example 1: Basic function
func add(a, b int) int {
	return a + b
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func basicFunctionExample() {
	fmt.Println("\n=== Basic Function Example ===")

	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	message := greet("Ali")
	fmt.Printf("Greeting: %s\n", message)
}

// Example 2: Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func getNameAndAge() (string, int) {
	return "Sara", 25
}

func multiplePReturnExample() {
	fmt.Println("\n=== Multiple Return Example ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	name, age := getNameAndAge()
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

// Example 3: Variable arguments (variadic functions)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func printInfo(name string, scores ...int) {
	fmt.Printf("Student: %s\n", name)
	if len(scores) > 0 {
		fmt.Printf("Scores: %v\n", scores)
		fmt.Printf("Average: %.2f\n", float64(sum(scores...))/float64(len(scores)))
	}
}

func variableArgumentsExample() {
	fmt.Println("\n=== Variable Arguments Example ===")

	result := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d\n", result)

	numbers := []int{10, 20, 30}
	result = sum(numbers...)
	fmt.Printf("Sum of slice: %d\n", result)

	printInfo("Ali", 85, 90, 78, 92)
}

// Example 4: Closures
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func adder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func closureExample() {
	fmt.Println("\n=== Closure Example ===")

	// Counter closure
	nextNumber := counter()
	fmt.Printf("First call: %d\n", nextNumber())
	fmt.Printf("Second call: %d\n", nextNumber())
	fmt.Printf("Third call: %d\n", nextNumber())

	// Adder closure
	add5 := adder(5)
	fmt.Printf("Add 5 to 3: %d\n", add5(3))
	fmt.Printf("Add 5 to 10: %d\n", add5(10))
}

// Example 5: Higher-order functions
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func multiply(a, b int) int {
	return a * b
}

func higherOrderFunctionExample() {
	fmt.Println("\n=== Higher-order Function Example ===")

	result := applyOperation(5, 3, add)
	fmt.Printf("Addition: %d\n", result)

	result = applyOperation(5, 3, multiply)
	fmt.Printf("Multiplication: %d\n", result)

	// Using anonymous function
	result = applyOperation(5, 3, func(a, b int) int {
		return a - b
	})
	fmt.Printf("Subtraction: %d\n", result)
}

// Example 6: Anonymous functions
func anonymousFunctionExample() {
	fmt.Println("\n=== Anonymous Function Example ===")

	// Immediately invoked function expression (IIFE)
	result := func(x, y int) int {
		return x * y
	}(4, 5)
	fmt.Printf("IIFE result: %d\n", result)

	// Assigning anonymous function to variable
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 7: %d\n", square(7))
}

// Example 7: Recursion
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func recursionExample() {
	fmt.Println("\n=== Recursion Example ===")

	fmt.Printf("Factorial of 5: %d\n", factorial(5))

	fmt.Print("Fibonacci sequence (first 10): ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}

// Example 8: Defer statement
func deferExample() {
	fmt.Println("\n=== Defer Example ===")

	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("Middle")

	// Defer with variables
	for i := 0; i < 3; i++ {
		defer func(x int) {
			fmt.Printf("Deferred loop: %d\n", x)
		}(i)
	}

	fmt.Println("End")
}

// Example 9: Panic and Recover
func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("Something went wrong!")
	fmt.Println("This will not be printed")
}

func panicRecoverExample() {
	fmt.Println("\n=== Panic and Recover Example ===")

	fmt.Println("Before calling risky function")
	riskyFunction()
	fmt.Println("After calling risky function")
}
