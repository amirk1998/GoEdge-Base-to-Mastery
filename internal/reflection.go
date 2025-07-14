// reflection_examples.go
package internal

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// User represents a sample user struct for reflection examples
type AccountUser struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"min=0,max=120"`
	IsActive bool   `json:"is_active"`
}

// Product represents a sample product struct
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// Validator interface for custom validation
type Validator interface {
	Validate() error
}

// RunReflectionExamples - main function to run all reflection examples
func RunReflectionExamples() {
	basicReflectionExample()
	typeAndValueExample()
	structFieldReflectionExample()
	methodReflectionExample()
	sliceReflectionExample()
	interfaceReflectionExample()
	tagReflectionExample()
	dynamicFunctionCallExample()
	jsonMarshallingExample()
	validationFrameworkExample()
}

// basicReflectionExample demonstrates basic reflection concepts
func basicReflectionExample() {
	fmt.Println(Subtitle("1. Basic Reflection Example"))

	// Different types of values
	values := []interface{}{
		42,
		"hello world",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1, "b": 2},
	}

	for _, v := range values {
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)

		fmt.Printf("Value: %v, Type: %v, Kind: %v\n",
			v, rt, rv.Kind())
	}
	fmt.Println()
}

// typeAndValueExample demonstrates the difference between Type and Value
func typeAndValueExample() {
	fmt.Println(Subtitle("2. Type vs Value Example"))

	user := AccountUser{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      25,
		IsActive: true,
	}

	// Get Type information
	userType := reflect.TypeOf(user)
	fmt.Printf("Type Name: %s\n", userType.Name())
	fmt.Printf("Type Kind: %s\n", userType.Kind())
	fmt.Printf("Package Path: %s\n", userType.PkgPath())

	// Get Value information
	userValue := reflect.ValueOf(user)
	fmt.Printf("Value Kind: %s\n", userValue.Kind())
	fmt.Printf("Value Type: %s\n", userValue.Type())
	fmt.Printf("Is Valid: %t\n", userValue.IsValid())
	fmt.Printf("Can Set: %t\n", userValue.CanSet())

	// Working with pointer to enable setting
	userPtr := reflect.ValueOf(&user)
	userElem := userPtr.Elem()
	fmt.Printf("Pointer Elem Can Set: %t\n", userElem.CanSet())
	fmt.Println()
}

// structFieldReflectionExample demonstrates struct field reflection
func structFieldReflectionExample() {
	fmt.Println(Subtitle("3. Struct Field Reflection Example"))

	user := AccountUser{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30}
	userValue := reflect.ValueOf(&user).Elem()
	userType := reflect.TypeOf(user)

	fmt.Printf("Struct has %d fields:\n", userType.NumField())

	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fieldValue := userValue.Field(i)

		fmt.Printf("Field %d: %s (Type: %s, Value: %v)\n",
			i, field.Name, field.Type, fieldValue.Interface())

		// Demonstrate field modification
		if fieldValue.CanSet() {
			switch field.Type.Kind() {
			case reflect.String:
				if field.Name == "Name" {
					fieldValue.SetString("Modified " + fieldValue.String())
				}
			case reflect.Int:
				if field.Name == "Age" {
					fieldValue.SetInt(fieldValue.Int() + 1)
				}
			}
		}
	}

	fmt.Printf("Modified user: %+v\n", user)
	fmt.Println()
}

// methodReflectionExample demonstrates method reflection
func methodReflectionExample() {
	fmt.Println(Subtitle("4. Method Reflection Example"))

	user := AccountUser{ID: 1, Name: "Bob", Email: "bob@example.com", Age: 28}
	userValue := reflect.ValueOf(&user)
	userType := reflect.TypeOf(&user)

	fmt.Printf("Type has %d methods:\n", userType.NumMethod())

	// Add a method to User type (this would be defined elsewhere)
	// For demonstration, we'll show method discovery
	for i := 0; i < userType.NumMethod(); i++ {
		method := userType.Method(i)
		fmt.Printf("Method %d: %s (Type: %s)\n",
			i, method.Name, method.Type)
	}

	// Demonstrate method calling by name
	methodName := "String" // This would be a method you've defined
	method := userValue.MethodByName(methodName)
	if method.IsValid() {
		fmt.Printf("Method %s exists and is callable\n", methodName)
	} else {
		fmt.Printf("Method %s not found\n", methodName)
	}
	fmt.Println()
}

// sliceReflectionExample demonstrates slice reflection
func sliceReflectionExample() {
	fmt.Println(Subtitle("5. Slice Reflection Example"))

	numbers := []int{1, 2, 3, 4, 5}
	sliceValue := reflect.ValueOf(numbers)
	sliceType := reflect.TypeOf(numbers)

	fmt.Printf("Slice Type: %s, Kind: %s\n", sliceType, sliceValue.Kind())
	fmt.Printf("Slice Length: %d, Capacity: %d\n",
		sliceValue.Len(), sliceValue.Cap())
	fmt.Printf("Element Type: %s\n", sliceType.Elem())

	// Iterate through slice elements
	fmt.Print("Elements: ")
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		fmt.Printf("%v ", element.Interface())
	}
	fmt.Println()

	// Create new slice dynamically
	newSliceType := reflect.SliceOf(reflect.TypeOf(0))
	newSlice := reflect.MakeSlice(newSliceType, 0, 5)

	// Append elements to new slice
	for i := 10; i < 15; i++ {
		newSlice = reflect.Append(newSlice, reflect.ValueOf(i))
	}

	fmt.Printf("Dynamic slice: %v\n", newSlice.Interface())
	fmt.Println()
}

// interfaceReflectionExample demonstrates interface reflection
func interfaceReflectionExample() {
	fmt.Println(Subtitle("6. Interface Reflection Example"))

	var values []interface{} = []interface{}{
		42,
		"hello",
		AccountUser{ID: 1, Name: "Charlie"},
		[]string{"a", "b", "c"},
	}

	for i, v := range values {
		value := reflect.ValueOf(v)
		typ := reflect.TypeOf(v)

		fmt.Printf("Item %d:\n", i)
		fmt.Printf("  Concrete Type: %s\n", typ)
		fmt.Printf("  Kind: %s\n", value.Kind())
		fmt.Printf("  Value: %v\n", v)

		// Check if value implements specific interfaces
		errorType := reflect.TypeOf((*error)(nil)).Elem()
		stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

		if typ.Implements(errorType) {
			fmt.Printf("  Implements error interface\n")
		}
		if typ.Implements(stringerType) {
			fmt.Printf("  Implements fmt.Stringer interface\n")
		}
		fmt.Println()
	}
}

// tagReflectionExample demonstrates struct tag reflection
func tagReflectionExample() {
	fmt.Println(Subtitle("7. Struct Tag Reflection Example"))

	userType := reflect.TypeOf(AccountUser{})

	fmt.Println("Field tags:")
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")

		fmt.Printf("Field: %s\n", field.Name)
		fmt.Printf("  JSON tag: %s\n", jsonTag)
		fmt.Printf("  Validate tag: %s\n", validateTag)
		fmt.Printf("  Full tag: %s\n", field.Tag)
		fmt.Println()
	}
}

// dynamicFunctionCallExample demonstrates dynamic function calling
func dynamicFunctionCallExample() {
	fmt.Println(Subtitle("8. Dynamic Function Call Example"))

	// Function to be called dynamically
	multiply := func(a, b int) int {
		return a * b
	}

	funcValue := reflect.ValueOf(multiply)
	funcType := reflect.TypeOf(multiply)

	fmt.Printf("Function Type: %s\n", funcType)
	fmt.Printf("Number of input parameters: %d\n", funcType.NumIn())
	fmt.Printf("Number of output parameters: %d\n", funcType.NumOut())

	// Prepare arguments
	args := []reflect.Value{
		reflect.ValueOf(5),
		reflect.ValueOf(3),
	}

	// Call function dynamically
	results := funcValue.Call(args)
	fmt.Printf("Result: %v\n", results[0].Interface())

	// Function registry pattern
	functionRegistry := map[string]interface{}{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
	}

	callFunction := func(name string, args ...int) interface{} {
		fn, exists := functionRegistry[name]
		if !exists {
			return "Function not found"
		}

		fnValue := reflect.ValueOf(fn)
		fnArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			fnArgs[i] = reflect.ValueOf(arg)
		}

		results := fnValue.Call(fnArgs)
		return results[0].Interface()
	}

	fmt.Printf("Dynamic add(10, 5): %v\n", callFunction("add", 10, 5))
	fmt.Printf("Dynamic mul(4, 3): %v\n", callFunction("mul", 4, 3))
	fmt.Println()
}

// jsonMarshallingExample demonstrates JSON marshalling using reflection
func jsonMarshallingExample() {
	fmt.Println(Subtitle("9. JSON Marshalling with Reflection Example"))

	user := AccountUser{
		ID:       1,
		Name:     "David",
		Email:    "david@example.com",
		Age:      32,
		IsActive: true,
	}

	// Simple JSON marshalling using reflection
	jsonStr := structToJSON(user)
	fmt.Printf("JSON representation: %s\n", jsonStr)
	fmt.Println()
}

// structToJSON converts struct to JSON string using reflection
func structToJSON(v interface{}) string {
	value := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return ""
	}

	var fields []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)

		if !fieldValue.CanInterface() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = strings.ToLower(field.Name)
		}

		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fmt.Sprintf(`"%s"`, fieldValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueStr = fmt.Sprintf("%d", fieldValue.Int())
		case reflect.Bool:
			valueStr = fmt.Sprintf("%t", fieldValue.Bool())
		case reflect.Float32, reflect.Float64:
			valueStr = fmt.Sprintf("%f", fieldValue.Float())
		default:
			valueStr = fmt.Sprintf(`"%v"`, fieldValue.Interface())
		}

		fields = append(fields, fmt.Sprintf(`"%s": %s`, jsonTag, valueStr))
	}

	return "{" + strings.Join(fields, ", ") + "}"
}

// validationFrameworkExample demonstrates a simple validation framework using reflection
func validationFrameworkExample() {
	fmt.Println(Subtitle("10. Validation Framework Example"))

	// Test valid user
	validUser := AccountUser{
		ID:       1,
		Name:     "Eve",
		Email:    "eve@example.com",
		Age:      25,
		IsActive: true,
	}

	// Test invalid user
	invalidUser := AccountUser{
		ID:    0,
		Name:  "X", // Too short
		Email: "invalid-email",
		Age:   150, // Too old
	}

	fmt.Println("Validating valid user:")
	if errors := validateStruct(validUser); len(errors) == 0 {
		fmt.Println("✓ Valid user")
	} else {
		fmt.Printf("✗ Validation errors: %v\n", errors)
	}

	fmt.Println("\nValidating invalid user:")
	if errors := validateStruct(invalidUser); len(errors) == 0 {
		fmt.Println("✓ Valid user")
	} else {
		fmt.Printf("✗ Validation errors:\n")
		for _, err := range errors {
			fmt.Printf("  - %s\n", err)
		}
	}
	fmt.Println()
}

// validateStruct validates a struct using reflection and tags
func validateStruct(v interface{}) []string {
	var errors []string

	value := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return []string{"Value is not a struct"}
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)
		validateTag := field.Tag.Get("validate")

		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)

			if err := validateField(field.Name, fieldValue, rule); err != "" {
				errors = append(errors, err)
			}
		}
	}

	return errors
}

// validateField validates individual field based on validation rule
func validateField(fieldName string, fieldValue reflect.Value, rule string) string {
	switch {
	case rule == "required":
		if isZeroValue(fieldValue) {
			return fmt.Sprintf("%s is required", fieldName)
		}
	case strings.HasPrefix(rule, "min="):
		minStr := strings.TrimPrefix(rule, "min=")
		min, _ := strconv.Atoi(minStr)

		switch fieldValue.Kind() {
		case reflect.String:
			if len(fieldValue.String()) < min {
				return fmt.Sprintf("%s must be at least %d characters", fieldName, min)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldValue.Int() < int64(min) {
				return fmt.Sprintf("%s must be at least %d", fieldName, min)
			}
		}
	case strings.HasPrefix(rule, "max="):
		maxStr := strings.TrimPrefix(rule, "max=")
		max, _ := strconv.Atoi(maxStr)

		switch fieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldValue.Int() > int64(max) {
				return fmt.Sprintf("%s must be at most %d", fieldName, max)
			}
		}
	case rule == "email":
		if fieldValue.Kind() == reflect.String {
			email := fieldValue.String()
			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				return fmt.Sprintf("%s must be a valid email", fieldName)
			}
		}
	}

	return ""
}

// isZeroValue checks if a reflect.Value is zero
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
