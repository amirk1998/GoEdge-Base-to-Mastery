// internal/value_reference_concepts.go
package internal

import (
	"fmt"
	"strings"
)

// RunValueReferenceExamples - main function to run all value vs reference examples
func RunValueReferenceExamples() {
	fmt.Println(Subtitle("🔢 Primitive Types Examples:"))
	primitiveTypesExample()

	fmt.Println(Subtitle("👤 Struct Examples:"))
	structExample()

	fmt.Println(Subtitle("🔀 Slice Examples:"))
	sliceExample()

	fmt.Println(Subtitle("🗺️ Map Examples:"))
	mapReferenceExample()

	fmt.Println(Subtitle("📊 Array Passing Examples:"))
	arrayPassingValueDemo()

	fmt.Println(Subtitle("🏦 Real-world Banking Examples:"))
	bankingExample()
}

// ===== ARRAY PASSING EXAMPLES =====

// arrayPassingValueDemo - demonstrates how arrays are passed to functions
func arrayPassingValueDemo() {
	fmt.Println(Bold("Array Passing Demonstration:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	// Create original array
	originalArr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Original before function call: %v (address: %p)\n", originalArr, &originalArr)

	// Arrays are passed by value (copy)
	modifyArrayValueDemo(originalArr)
	fmt.Printf("Original after function call: %v\n", originalArr)

	// To modify original, pass pointer to array
	modifyArrayPointerDemo(&originalArr)
	fmt.Printf("Original after pointer modification: %v\n", originalArr)

	fmt.Println()
}

// modifyArrayValueDemo - demonstrates array passed by value
func modifyArrayValueDemo(arr [5]int) {
	fmt.Printf("Address in function (by value): %p\n", &arr)
	arr[0] = 999
	fmt.Printf("Inside function (by value): %v\n", arr)
}

// modifyArrayPointerDemo - demonstrates array passed by pointer
func modifyArrayPointerDemo(arr *[5]int) {
	fmt.Printf("Address in function (by pointer): %p\n", arr)
	arr[0] = 777
	fmt.Printf("Inside function (by pointer): %v\n", *arr)
}

// ===== PRIMITIVE TYPES EXAMPLES =====

// incrementByValue - demonstrates primitive types are always passed by value
func incrementByValue(x int) {
	x++
	fmt.Printf("Inside incrementByValue: x = %d (address: %p)\n", x, &x)
}

// incrementByPointer - demonstrates passing pointer to primitive
func incrementByPointer(x *int) {
	*x++
	fmt.Printf("Inside incrementByPointer: *x = %d (address: %p)\n", *x, x)
}

func primitiveTypesExample() {
	fmt.Println(Bold("Primitive Types Demonstration:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	num := 10
	fmt.Printf("Original: num = %d (address: %p)\n", num, &num)

	incrementByValue(num)
	fmt.Printf("After incrementByValue: num = %d\n", num)

	incrementByPointer(&num)
	fmt.Printf("After incrementByPointer: num = %d\n", num)
	fmt.Println()
}

// ===== STRUCT EXAMPLES =====

type DemoPerson struct {
	Name string
	Age  int
}

// updateDemoPersonByValue - struct passed by value (copy)
func updateDemoPersonByValue(p DemoPerson) {
	p.Age++
	fmt.Printf("Inside updateDemoPersonByValue: %+v (address: %p)\n", p, &p)
}

// updateDemoPersonByPointer - struct passed by pointer
func updateDemoPersonByPointer(p *DemoPerson) {
	p.Age++
	fmt.Printf("Inside updateDemoPersonByPointer: %+v (address: %p)\n", *p, p)
}

func structExample() {
	fmt.Println(Bold("Struct Demonstration:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	person := DemoPerson{Name: "Alice", Age: 25}
	fmt.Printf("Original: %+v (address: %p)\n", person, &person)

	updateDemoPersonByValue(person)
	fmt.Printf("After updateDemoPersonByValue: %+v\n", person)

	updateDemoPersonByPointer(&person)
	fmt.Printf("After updateDemoPersonByPointer: %+v\n", person)
	fmt.Println()
}

// ===== SLICE EXAMPLES =====

// modifySliceElementsDemo - slice header passed by value, but points to same array
func modifySliceElementsDemo(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Printf("Inside modifySliceElementsDemo: %v (address: %p)\n", s, &s)
}

// appendToIntSliceDemo - demonstrates slice growth
func appendToIntSliceDemo(s *[]int, val int) {
	*s = append(*s, val)
	fmt.Printf("Inside appendToIntSliceDemo: %v (address: %p)\n", *s, s)
}

func sliceExample() {
	fmt.Println(Bold("Slice Demonstration:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v (address: %p)\n", slice, &slice)

	modifySliceElementsDemo(slice)
	fmt.Printf("After modifySliceElementsDemo: %v\n", slice)

	appendToIntSliceDemo(&slice, 100)
	fmt.Printf("After appendToIntSliceDemo: %v\n", slice)
	fmt.Println()
}

// ===== MAP EXAMPLES =====

// modifyMapDemo - maps are reference types
func modifyMapDemo(m map[string]int) {
	m["new_key"] = 100
	fmt.Printf("Inside modifyMapDemo: %v (address: %p)\n", m, &m)
}

// replaceMapDemo - replaces entire map
func replaceMapDemo(m *map[string]int) {
	*m = make(map[string]int)
	(*m)["replaced"] = 999
	fmt.Printf("Inside replaceMapDemo: %v (address: %p)\n", *m, m)
}

func mapReferenceExample() {
	fmt.Println(Bold("Map Demonstration:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	myMap := map[string]int{"key1": 1, "key2": 2}
	fmt.Printf("Original: %v (address: %p)\n", myMap, &myMap)

	modifyMapDemo(myMap)
	fmt.Printf("After modifyMapDemo: %v\n", myMap)

	replaceMapDemo(&myMap)
	fmt.Printf("After replaceMapDemo: %v\n", myMap)
	fmt.Println()
}

// ===== REAL-WORLD BANKING SYSTEM =====

type DemoAccount struct {
	AccountID string
	Funds     float64
}

type DemoBank struct {
	Name     string
	Accounts [3]DemoAccount // Fixed-size array
}

// depositByValueDemo - won't modify original account
func depositByValueDemo(acc DemoAccount, amount float64) DemoAccount {
	acc.Funds += amount
	fmt.Printf("Inside depositByValueDemo: Account %s, Funds: %.2f\n", acc.AccountID, acc.Funds)
	return acc
}

// depositByPointerDemo - modifies original account
func depositByPointerDemo(acc *DemoAccount, amount float64) {
	acc.Funds += amount
	fmt.Printf("Inside depositByPointerDemo: Account %s, Funds: %.2f\n", acc.AccountID, acc.Funds)
}

// processBankByValueDemo - won't modify original bank
func processBankByValueDemo(bank DemoBank) {
	bank.Accounts[0].Funds += 1000
	fmt.Printf("Inside processBankByValueDemo: %s, Account[0] Funds: %.2f\n",
		bank.Name, bank.Accounts[0].Funds)
}

// processBankByPointerDemo - modifies original bank
func processBankByPointerDemo(bank *DemoBank) {
	bank.Accounts[0].Funds += 1000
	fmt.Printf("Inside processBankByPointerDemo: %s, Account[0] Funds: %.2f\n",
		bank.Name, bank.Accounts[0].Funds)
}

func bankingExample() {
	fmt.Println(Bold("Real-world Banking System:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	// Create accounts
	account1 := DemoAccount{AccountID: "ACC001", Funds: 1000.0}
	account2 := DemoAccount{AccountID: "ACC002", Funds: 2000.0}
	account3 := DemoAccount{AccountID: "ACC003", Funds: 3000.0}

	// Create bank with fixed-size array of accounts
	bank := DemoBank{
		Name:     "Go Bank",
		Accounts: [3]DemoAccount{account1, account2, account3},
	}

	fmt.Printf("Original Bank: %s\n", bank.Name)
	fmt.Printf("Account[0] Funds: %.2f\n", bank.Accounts[0].Funds)

	// Try to deposit by value (won't work)
	updatedAccount := depositByValueDemo(bank.Accounts[0], 500)
	fmt.Printf("After depositByValueDemo - Original: %.2f, Returned: %.2f\n",
		bank.Accounts[0].Funds, updatedAccount.Funds)

	// Deposit by pointer (will work)
	depositByPointerDemo(&bank.Accounts[0], 500)
	fmt.Printf("After depositByPointerDemo: %.2f\n", bank.Accounts[0].Funds)

	// Process bank by value (won't work)
	processBankByValueDemo(bank)
	fmt.Printf("After processBankByValueDemo: %.2f\n", bank.Accounts[0].Funds)

	// Process bank by pointer (will work)
	processBankByPointerDemo(&bank)
	fmt.Printf("After processBankByPointerDemo: %.2f\n", bank.Accounts[0].Funds)

	fmt.Println()
}

// ===== STUDENT GRADE MANAGEMENT SYSTEM =====

type Student struct {
	Name   string
	Grades [5]int
}

// calculateAverageByValue - calculates average without modifying original
func calculateAverageByValue(grades [5]int) float64 {
	// This function works with a copy, so original is safe
	total := 0
	for _, grade := range grades {
		total += grade
	}
	return float64(total) / float64(len(grades))
}

// normalizeGradesByPointer - normalizes grades by adding curve points
func normalizeGradesByPointer(grades *[5]int, curvePoints int) {
	// This function modifies the original array
	for i := range grades {
		grades[i] += curvePoints
		if grades[i] > 100 {
			grades[i] = 100
		}
	}
}

// RunStudentGradeExample - demonstrates real-world application
func RunStudentGradeExample() {
	fmt.Println(Bold("Student Grade Management System:"))
	fmt.Println(Cyan(strings.Repeat("=", 50)))

	student := Student{
		Name:   "John Doe",
		Grades: [5]int{78, 85, 92, 88, 76},
	}

	fmt.Printf("Student: %s\n", student.Name)
	fmt.Printf("Original grades: %v\n", student.Grades)

	// Calculate average (safe operation - by value)
	average := calculateAverageByValue(student.Grades)
	fmt.Printf("Average grade: %.2f\n", average)
	fmt.Printf("Grades after average calculation: %v\n", student.Grades)

	// Apply curve (modifies original - by pointer)
	fmt.Println(InfoText("Applying curve of 5 points..."))
	normalizeGradesByPointer(&student.Grades, 5)
	fmt.Printf("Grades after curve: %v\n", student.Grades)

	// New average after curve
	newAverage := calculateAverageByValue(student.Grades)
	fmt.Printf("New average: %.2f\n", newAverage)
	fmt.Println()
}
