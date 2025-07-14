// context_examples.go
package internal

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// RequestData represents data passed through context
type RequestData struct {
	UserID    string
	RequestID string
	IP        string
}

// DatabaseService simulates a database service
type DatabaseService struct {
	delay time.Duration
}

// APIService simulates an API service
type APIService struct {
	delay time.Duration
}

// RunContextExamples - main function to run all context examples
func RunContextExamples() {
	basicContextExample()
	contextWithValueExample()
	contextWithTimeoutExample()
	contextWithCancelExample()
	contextWithDeadlineExample()
	contextPropagationExample()
	httpServerContextExample()
	pipelineContextExample()
	contextBestPracticesExample()
	realWorldScenarioExample()
}

// basicContextExample demonstrates basic context usage
func basicContextExample() {
	fmt.Println(Subtitle("1. Basic Context Example"))

	// Background context - never canceled, has no values, has no deadline
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)

	// TODO context - used when you're not sure what context to use
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)

	// Check context properties
	select {
	case <-ctx.Done():
		fmt.Println("Context is done")
	default:
		fmt.Println("Context is not done")
	}

	fmt.Printf("Context error: %v\n", ctx.Err())
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("Context deadline: %v\n", deadline)
	} else {
		fmt.Println("Context has no deadline")
	}
	fmt.Println()
}

// contextWithValueExample demonstrates context with values
func contextWithValueExample() {
	fmt.Println(Subtitle("2. Context with Values Example"))

	// Create context with values
	ctx := context.Background()

	// Add user ID to context
	ctx = context.WithValue(ctx, "userID", "user123")

	// Add request ID to context
	ctx = context.WithValue(ctx, "requestID", "req456")

	// Add more structured data
	requestData := RequestData{
		UserID:    "user123",
		RequestID: "req456",
		IP:        "192.168.1.1",
	}
	ctx = context.WithValue(ctx, "requestData", requestData)

	// Pass context to functions
	processRequest(ctx)

	fmt.Println()
}

// processRequest demonstrates reading values from context
func processRequest(ctx context.Context) {
	fmt.Println("Processing request...")

	// Extract values from context
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	requestData := ctx.Value("requestData")

	if userID != nil {
		fmt.Printf("User ID: %s\n", userID)
	}

	if requestID != nil {
		fmt.Printf("Request ID: %s\n", requestID)
	}

	if requestData != nil {
		if data, ok := requestData.(RequestData); ok {
			fmt.Printf("Request Data: %+v\n", data)
		}
	}

	// Simulate some work
	performDatabaseOperation(ctx)
	callExternalAPI(ctx)
}

// performDatabaseOperation simulates database operation with context
func performDatabaseOperation(ctx context.Context) {
	fmt.Println("Performing database operation...")

	// Get user ID from context
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Printf("Database query for user: %s\n", userID)
	}

	// Simulate database delay
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Database operation completed")
}

// callExternalAPI simulates external API call with context
func callExternalAPI(ctx context.Context) {
	fmt.Println("Calling external API...")

	// Get request ID from context
	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Printf("API call with request ID: %s\n", requestID)
	}

	// Simulate API call delay
	time.Sleep(30 * time.Millisecond)
	fmt.Println("API call completed")
}

// contextWithTimeoutExample demonstrates context with timeout
func contextWithTimeoutExample() {
	fmt.Println(Subtitle("3. Context with Timeout Example"))

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to release resources

	fmt.Println("Starting operation with 2-second timeout...")

	// Start multiple operations
	var wg sync.WaitGroup

	// Fast operation (should complete)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fastOperation(ctx, "Operation 1")
	}()

	// Slow operation (should timeout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		slowOperation(ctx, "Operation 2")
	}()

	wg.Wait()
	fmt.Println("All operations completed or timed out")
	fmt.Println()
}

// fastOperation simulates a fast operation
func fastOperation(ctx context.Context, name string) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("%s completed successfully\n", name)
	case <-ctx.Done():
		fmt.Printf("%s canceled: %v\n", name, ctx.Err())
	}
}

// slowOperation simulates a slow operation
func slowOperation(ctx context.Context, name string) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("%s completed successfully\n", name)
	case <-ctx.Done():
		fmt.Printf("%s canceled: %v\n", name, ctx.Err())
	}
}

// contextWithCancelExample demonstrates context with cancellation
func contextWithCancelExample() {
	fmt.Println(Subtitle("4. Context with Cancel Example"))

	// Create cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	// Start multiple workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			contextWorker(ctx, workerID)
		}(i)
	}

	// Let workers run for a bit
	time.Sleep(1 * time.Second)

	// Cancel all workers
	fmt.Println("Canceling all workers...")
	cancel()

	wg.Wait()
	fmt.Println("All workers stopped")
	fmt.Println()
}

// contextWorker simulates a worker that respects context cancellation
func contextWorker(ctx context.Context, id int) {
	fmt.Printf("Worker %d started\n", id)

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Worker %d is working...\n", id)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped: %v\n", id, ctx.Err())
			return
		}
	}
}

// contextWithDeadlineExample demonstrates context with deadline
func contextWithDeadlineExample() {
	fmt.Println(Subtitle("5. Context with Deadline Example"))

	// Create context with deadline
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Deadline set for: %v\n", deadline.Format("15:04:05.000"))

	// Start operation
	result := make(chan string, 1)
	go func() {
		result <- performLongOperation(ctx)
	}()

	// Wait for result or timeout
	select {
	case res := <-result:
		fmt.Printf("Operation result: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("Operation deadline exceeded: %v\n", ctx.Err())
	}

	fmt.Println()
}

// performLongOperation simulates a long-running operation
func performLongOperation(ctx context.Context) string {
	// Simulate work that takes random time
	workTime := time.Duration(rand.Intn(3000)) * time.Millisecond

	select {
	case <-time.After(workTime):
		return "Operation completed successfully"
	case <-ctx.Done():
		return "Operation was canceled"
	}
}

// contextPropagationExample demonstrates context propagation through call chain
func contextPropagationExample() {
	fmt.Println(Subtitle("6. Context Propagation Example"))

	// Create root context with values and timeout
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, "traceID", "trace123")
	rootCtx = context.WithValue(rootCtx, "userID", "user456")

	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	// Start request processing
	handleRequest(ctx)

	fmt.Println()
}

// handleRequest simulates request handling with context propagation
func handleRequest(ctx context.Context) {
	fmt.Println("Handling request...")

	// Extract trace ID for logging
	traceID := ctx.Value("traceID")
	if traceID != nil {
		fmt.Printf("Trace ID: %s\n", traceID)
	}

	// Pass context to service layer
	processBusinessLogic(ctx)
}

// processBusinessLogic simulates business logic processing
func processBusinessLogic(ctx context.Context) {
	fmt.Println("Processing business logic...")

	// Create child context with additional values
	childCtx := context.WithValue(ctx, "operationID", "op789")

	// Pass to data access layer
	accessDatabase(childCtx)
	callExternalService(childCtx)
}

// accessDatabase simulates database access
func accessDatabase(ctx context.Context) {
	fmt.Println("Accessing database...")

	// Check for cancellation before expensive operation
	select {
	case <-ctx.Done():
		fmt.Printf("Database access canceled: %v\n", ctx.Err())
		return
	default:
	}

	// Simulate database query
	time.Sleep(300 * time.Millisecond)

	// Extract values for logging
	traceID := ctx.Value("traceID")
	userID := ctx.Value("userID")
	operationID := ctx.Value("operationID")

	fmt.Printf("Database query completed - Trace: %v, User: %v, Op: %v\n",
		traceID, userID, operationID)
}

// callExternalService simulates external service call
func callExternalService(ctx context.Context) {
	fmt.Println("Calling external service...")

	// Check for cancellation
	select {
	case <-ctx.Done():
		fmt.Printf("External service call canceled: %v\n", ctx.Err())
		return
	default:
	}

	// Simulate service call
	time.Sleep(400 * time.Millisecond)

	// Extract trace ID for correlation
	traceID := ctx.Value("traceID")
	fmt.Printf("External service call completed - Trace: %v\n", traceID)
}

// httpServerContextExample demonstrates context in HTTP server
func httpServerContextExample() {
	fmt.Println(Subtitle("7. HTTP Server Context Example"))

	// Create HTTP server with context-aware handlers
	mux := http.NewServeMux()

	// Add middleware for request context
	mux.HandleFunc("/api/users", withContext(userHandler))
	mux.HandleFunc("/api/orders", withContext(orderHandler))

	// Simulate HTTP requests
	fmt.Println("Simulating HTTP requests...")
	simulateHTTPRequest("/api/users")
	simulateHTTPRequest("/api/orders")

	fmt.Println()
}

// withContext middleware adds context to HTTP requests
func withContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Add request metadata to context
		requestID := fmt.Sprintf("req-%d", rand.Intn(10000))
		ctx = context.WithValue(ctx, "requestID", requestID)
		ctx = context.WithValue(ctx, "startTime", time.Now())

		// Create new request with context
		r = r.WithContext(ctx)

		// Call next handler
		next(w, r)
	}
}

// userHandler handles user-related requests
func userHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value("requestID")

	fmt.Printf("User handler called - Request ID: %v\n", requestID)

	// Simulate user service call
	users := getUsersFromService(ctx)
	fmt.Printf("Retrieved %d users\n", len(users))
}

// orderHandler handles order-related requests
func orderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value("requestID")

	fmt.Printf("Order handler called - Request ID: %v\n", requestID)

	// Simulate order service call
	orders := getOrdersFromService(ctx)
	fmt.Printf("Retrieved %d orders\n", len(orders))
}

// getUsersFromService simulates user service call
func getUsersFromService(ctx context.Context) []string {
	// Check for cancellation
	select {
	case <-ctx.Done():
		fmt.Printf("User service call canceled: %v\n", ctx.Err())
		return nil
	default:
	}

	// Simulate service delay
	time.Sleep(100 * time.Millisecond)

	return []string{"user1", "user2", "user3"}
}

// getOrdersFromService simulates order service call
func getOrdersFromService(ctx context.Context) []string {
	// Check for cancellation
	select {
	case <-ctx.Done():
		fmt.Printf("Order service call canceled: %v\n", ctx.Err())
		return nil
	default:
	}

	// Simulate service delay
	time.Sleep(150 * time.Millisecond)

	return []string{"order1", "order2"}
}

// simulateHTTPRequest simulates an HTTP request
func simulateHTTPRequest(path string) {
	fmt.Printf("Simulating request to %s\n", path)

	// In real scenario, this would be handled by HTTP server
	switch path {
	case "/api/users":
		userHandler(nil, &http.Request{})
	case "/api/orders":
		orderHandler(nil, &http.Request{})
	}
}

// pipelineContextExample demonstrates context in processing pipeline
func pipelineContextExample() {
	fmt.Println(Subtitle("8. Pipeline Context Example"))

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create pipeline stages
	input := make(chan int, 5)
	stage1 := make(chan int, 5)
	stage2 := make(chan int, 5)
	output := make(chan int, 5)

	// Start pipeline stages
	go pipelineStage1(ctx, input, stage1)
	go pipelineStage2(ctx, stage1, stage2)
	go pipelineStage3(ctx, stage2, output)

	// Send input data
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			select {
			case input <- i:
				fmt.Printf("Sent input: %d\n", i)
			case <-ctx.Done():
				fmt.Printf("Input canceled: %v\n", ctx.Err())
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Read output
	go func() {
		for {
			select {
			case result, ok := <-output:
				if !ok {
					fmt.Println("Output channel closed")
					return
				}
				fmt.Printf("Received output: %d\n", result)
			case <-ctx.Done():
				fmt.Printf("Output reading canceled: %v\n", ctx.Err())
				return
			}
		}
	}()

	// Wait for pipeline to complete or timeout
	time.Sleep(2 * time.Second)
	fmt.Println("Pipeline processing completed")
	fmt.Println()
}

// pipelineStage1 processes input and multiplies by 2
func pipelineStage1(ctx context.Context, input <-chan int, output chan<- int) {
	defer close(output)

	for {
		select {
		case value, ok := <-input:
			if !ok {
				return
			}
			// Process value
			result := value * 2
			select {
			case output <- result:
				fmt.Printf("Stage 1: %d -> %d\n", value, result)
			case <-ctx.Done():
				fmt.Printf("Stage 1 canceled: %v\n", ctx.Err())
				return
			}
		case <-ctx.Done():
			fmt.Printf("Stage 1 canceled: %v\n", ctx.Err())
			return
		}
	}
}

// pipelineStage2 processes input and adds 10
func pipelineStage2(ctx context.Context, input <-chan int, output chan<- int) {
	defer close(output)

	for {
		select {
		case value, ok := <-input:
			if !ok {
				return
			}
			// Process value
			result := value + 10
			select {
			case output <- result:
				fmt.Printf("Stage 2: %d -> %d\n", value, result)
			case <-ctx.Done():
				fmt.Printf("Stage 2 canceled: %v\n", ctx.Err())
				return
			}
		case <-ctx.Done():
			fmt.Printf("Stage 2 canceled: %v\n", ctx.Err())
			return
		}
	}
}

// pipelineStage3 processes input and divides by 2
func pipelineStage3(ctx context.Context, input <-chan int, output chan<- int) {
	defer close(output)

	for {
		select {
		case value, ok := <-input:
			if !ok {
				return
			}
			// Process value
			result := value / 2
			select {
			case output <- result:
				fmt.Printf("Stage 3: %d -> %d\n", value, result)
			case <-ctx.Done():
				fmt.Printf("Stage 3 canceled: %v\n", ctx.Err())
				return
			}
		case <-ctx.Done():
			fmt.Printf("Stage 3 canceled: %v\n", ctx.Err())
			return
		}
	}
}

// contextBestPracticesExample demonstrates context best practices
func contextBestPracticesExample() {
	fmt.Println(Subtitle("9. Context Best Practices Example"))

	// ✅ DO: Pass context as first parameter
	goodFunction(context.Background(), "data")

	// ✅ DO: Use context.WithTimeout for operations with time limits
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// ✅ DO: Check for context cancellation in loops
	demonstrateContextChecking(ctx)

	// ✅ DO: Use context.WithValue sparingly and with proper keys
	demonstrateContextValues()

	// ✅ DO: Always call cancel function
	demonstrateCancelUsage()

	fmt.Println()
}

// goodFunction demonstrates proper context usage as first parameter
func goodFunction(ctx context.Context, data string) {
	fmt.Printf("Processing data: %s\n", data)

	// Always check for cancellation before expensive operations
	select {
	case <-ctx.Done():
		fmt.Printf("Operation canceled: %v\n", ctx.Err())
		return
	default:
	}

	// Simulate work
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Operation completed")
}

// demonstrateContextChecking shows how to check context in loops
func demonstrateContextChecking(ctx context.Context) {
	fmt.Println("Demonstrating context checking in loops...")

	for i := 0; i < 5; i++ {
		// Check for cancellation
		select {
		case <-ctx.Done():
			fmt.Printf("Loop canceled at iteration %d: %v\n", i, ctx.Err())
			return
		default:
		}

		// Simulate work
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Loop iteration %d completed\n", i)
	}
}

// Custom key type for context values
type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

// demonstrateContextValues shows proper context value usage
func demonstrateContextValues() {
	fmt.Println("Demonstrating context values...")

	// Use typed keys instead of strings
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, "user123")
	ctx = context.WithValue(ctx, requestIDKey, "req456")

	// Extract values with type safety
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Printf("User ID: %s\n", userID)
	}

	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Printf("Request ID: %s\n", requestID)
	}
}

// demonstrateCancelUsage shows proper cancel function usage
func demonstrateCancelUsage() {
	fmt.Println("Demonstrating cancel usage...")

	// Always defer cancel to prevent resource leaks
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // This is important!

	// Do work
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Work completed before timeout")
	case <-ctx.Done():
		fmt.Printf("Work canceled: %v\n", ctx.Err())
	}
}

// realWorldScenarioExample demonstrates real-world context usage
func realWorldScenarioExample() {
	fmt.Println(Subtitle("10. Real-world Scenario Example"))

	// Simulate a web request with database and API calls
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "user789")
	ctx = context.WithValue(ctx, "requestID", "req123")

	// Set timeout for the entire request
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Process order
	order := processOrder(ctx)
	if order != nil {
		fmt.Printf("Order processed successfully: %+v\n", order)
	} else {
		fmt.Println("Order processing failed")
	}

	fmt.Println()
}

// Order represents an order
type Order struct {
	ID       string
	UserID   string
	Products []string
	Total    float64
}

// processOrder simulates order processing with multiple service calls
func processOrder(ctx context.Context) *Order {
	userID := ctx.Value("userID").(string)
	requestID := ctx.Value("requestID").(string)

	fmt.Printf("Processing order for user %s (Request: %s)\n", userID, requestID)

	// Create services
	dbService := &DatabaseService{delay: 300 * time.Millisecond}
	apiService := &APIService{delay: 400 * time.Millisecond}

	// Get user data
	user, err := dbService.GetUser(ctx, userID)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return nil
	}

	// Get product prices
	prices, err := apiService.GetProductPrices(ctx, []string{"product1", "product2"})
	if err != nil {
		fmt.Printf("Failed to get prices: %v\n", err)
		return nil
	}

	// Calculate total
	total := 0.0
	for _, price := range prices {
		total += price
	}

	// Create order
	order := &Order{
		ID:       fmt.Sprintf("order-%d", rand.Intn(10000)),
		UserID:   user,
		Products: []string{"product1", "product2"},
		Total:    total,
	}

	// Save order
	err = dbService.SaveOrder(ctx, order)
	if err != nil {
		fmt.Printf("Failed to save order: %v\n", err)
		return nil
	}

	return order
}

// GetUser simulates getting user from database
func (db *DatabaseService) GetUser(ctx context.Context, userID string) (string, error) {
	fmt.Printf("Getting user %s from database...\n", userID)

	select {
	case <-time.After(db.delay):
		fmt.Println("User retrieved from database")
		return userID, nil
	case <-ctx.Done():
		return "", fmt.Errorf("database operation canceled: %w", ctx.Err())
	}
}

// SaveOrder simulates saving order to database
func (db *DatabaseService) SaveOrder(ctx context.Context, order *Order) error {
	fmt.Printf("Saving order %s to database...\n", order.ID)

	select {
	case <-time.After(db.delay):
		fmt.Println("Order saved to database")
		return nil
	case <-ctx.Done():
		return fmt.Errorf("database operation canceled: %w", ctx.Err())
	}
}

// GetProductPrices simulates getting product prices from API
func (api *APIService) GetProductPrices(ctx context.Context, products []string) ([]float64, error) {
	fmt.Printf("Getting prices for products %v from API...\n", products)

	select {
	case <-time.After(api.delay):
		fmt.Println("Product prices retrieved from API")
		prices := make([]float64, len(products))
		for i := range prices {
			prices[i] = rand.Float64() * 100
		}
		return prices, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("API operation canceled: %w", ctx.Err())
	}
}
