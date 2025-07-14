// embedding_composition.go
package internal

import (
	"fmt"
	"time"
)

func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.prefix, message)
}

// Base types for embedding examples
type AutoEngine struct {
	Horsepower int
	Fuel       string
	Running    bool
}

type VehicleWheels struct {
	Count int
	Size  string
}

type NavigationGPS struct {
	Latitude  float64
	Longitude float64
	Enabled   bool
}

// Interface definitions
type VehicleStarter interface {
	Start() error
	Stop() error
}

type VehicleNavigator interface {
	Navigate(destination string) error
	GetLocation() (float64, float64)
}

type VehicleHonker interface {
	Honk() string
}

// Methods for base types
func (e *AutoEngine) Start() error {
	if e.Running {
		return fmt.Errorf("engine is already running")
	}
	e.Running = true
	fmt.Printf("Engine started: %d HP, %s fuel\n", e.Horsepower, e.Fuel)
	return nil
}

func (e *AutoEngine) Stop() error {
	if !e.Running {
		return fmt.Errorf("engine is already stopped")
	}
	e.Running = false
	fmt.Println("Engine stopped")
	return nil
}

func (e *AutoEngine) Status() string {
	status := "stopped"
	if e.Running {
		status = "running"
	}
	return fmt.Sprintf("Engine: %d HP, %s fuel, %s", e.Horsepower, e.Fuel, status)
}

func (g *NavigationGPS) Navigate(destination string) error {
	if !g.Enabled {
		return fmt.Errorf("GPS is disabled")
	}
	fmt.Printf("Navigating to: %s\n", destination)
	return nil
}

func (g *NavigationGPS) GetLocation() (float64, float64) {
	return g.Latitude, g.Longitude
}

func (g *NavigationGPS) Enable() {
	g.Enabled = true
	fmt.Println("GPS enabled")
}

func (w *VehicleWheels) Description() string {
	return fmt.Sprintf("%d wheels, size: %s", w.Count, w.Size)
}

// Struct embedding examples
type AutoCar struct {
	AutoEngine    // Embedded struct
	VehicleWheels // Embedded struct
	NavigationGPS // Embedded struct
	Brand         string
	Model         string
	Year          int
	horn          string // private field
}

type AutoMotorcycle struct {
	AutoEngine    // Embedded struct
	VehicleWheels // Embedded struct
	Brand         string
	Model         string
	HasSidecar    bool
}

type AutoTruck struct {
	AutoEngine    // Embedded struct
	VehicleWheels // Embedded struct
	NavigationGPS // Embedded struct
	Brand         string
	Model         string
	PayloadKg     int
}

// Additional methods for embedded types
func (c *AutoCar) Honk() string {
	if c.horn == "" {
		c.horn = "Beep beep!"
	}
	return c.horn
}

func (c *AutoCar) String() string {
	return fmt.Sprintf("%d %s %s", c.Year, c.Brand, c.Model)
}

func (m *AutoMotorcycle) Honk() string {
	return "Vrooom vrooom!"
}

func (m *AutoMotorcycle) String() string {
	sidecar := ""
	if m.HasSidecar {
		sidecar = " (with sidecar)"
	}
	return fmt.Sprintf("%s %s%s", m.Brand, m.Model, sidecar)
}

func (t *AutoTruck) Honk() string {
	return "HOOOOOONK!"
}

func (t *AutoTruck) String() string {
	return fmt.Sprintf("%s %s (Payload: %d kg)", t.Brand, t.Model, t.PayloadKg)
}

func (t *AutoTruck) LoadCargo(weight int) error {
	if weight > t.PayloadKg {
		return fmt.Errorf("cargo too heavy: %d kg exceeds capacity of %d kg", weight, t.PayloadKg)
	}
	fmt.Printf("Loaded %d kg cargo into %s\n", weight, t.String())
	return nil
}

// Interface embedding example
type EnhancedNavigator interface {
	VehicleNavigator // Embedded interface
	SetDestination(string) error
	GetRoute() []string
}

type IntelligentGPS struct {
	NavigationGPS // Embedded struct
	destinations  []string
	currentRoute  []string
}

func (sg *IntelligentGPS) SetDestination(dest string) error {
	if !sg.Enabled {
		return fmt.Errorf("GPS must be enabled first")
	}
	sg.destinations = append(sg.destinations, dest)
	sg.currentRoute = []string{"Start", "Highway 1", "Exit 42", dest}
	fmt.Printf("Destination set: %s\n", dest)
	return nil
}

func (sg *IntelligentGPS) GetRoute() []string {
	return sg.currentRoute
}

// Complex embedding example
type PremiumCar struct {
	AutoCar        // Embedded struct
	IntelligentGPS // Embedded struct (overrides Car's GPS)
	leather        bool
	sunroof        bool
	heatedSeats    bool
}

func (lc *PremiumCar) EnableLuxuryFeatures() {
	lc.leather = true
	lc.sunroof = true
	lc.heatedSeats = true
	fmt.Println("Luxury features enabled: leather, sunroof, heated seats")
}

func (lc *PremiumCar) String() string {
	return fmt.Sprintf("Luxury %s", lc.AutoCar.String())
}

// Method shadowing example
type PerformanceCar struct {
	AutoCar
	turbo bool
}

func (sc *PerformanceCar) Start() error {
	if sc.turbo {
		fmt.Println("Turbo activated!")
	}
	return sc.AutoCar.Start() // Call embedded method
}

func (sc *PerformanceCar) EnableTurbo() {
	sc.turbo = true
	fmt.Println("Turbo enabled")
}

// Composition vs Embedding
type VehicleFleet struct {
	vehicles []AutoVehicle // Composition - has-a relationship
	manager  string
}

type AutoVehicle interface {
	Start() error
	Stop() error
	String() string
}

func (f *VehicleFleet) AddVehicle(v AutoVehicle) {
	f.vehicles = append(f.vehicles, v)
	fmt.Printf("Added vehicle to fleet: %s\n", v.String())
}

func (f *VehicleFleet) StartAll() {
	fmt.Printf("Fleet manager %s starting all vehicles:\n", f.manager)
	for _, v := range f.vehicles {
		if err := v.Start(); err != nil {
			fmt.Printf("Failed to start %s: %v\n", v.String(), err)
		}
	}
}

// RunEmbeddingCompositionExamples - main function to run all embedding examples
func RunEmbeddingCompositionExamples() {
	basicEmbeddingExample()
	methodPromotionExample()
	interfaceEmbeddingExample()
	embeddingVsCompositionExample()
	methodShadowingExample()
	complexEmbeddingExample()
	embeddingWithInterfacesExample()
	embeddingConflictExample()
	embeddingBestPracticesExample()
	realWorldExample()
}

// Example 1: Basic struct embedding
func basicEmbeddingExample() {
	fmt.Println(Header("1. Basic Struct Embedding"))

	// Create a car with embedded structs
	car := AutoCar{
		AutoEngine: AutoEngine{
			Horsepower: 200,
			Fuel:       "gasoline",
		},
		VehicleWheels: VehicleWheels{
			Count: 4,
			Size:  "18 inch",
		},
		NavigationGPS: NavigationGPS{
			Latitude:  40.7128,
			Longitude: -74.0060,
			Enabled:   true,
		},
		Brand: "Toyota",
		Model: "Camry",
		Year:  2023,
		horn:  "Beep beep!",
	}

	// Access embedded fields directly
	fmt.Printf("Car: %s\n", car.String())
	fmt.Printf("Horsepower: %d\n", car.Horsepower)
	fmt.Printf("Wheel count: %d\n", car.Count)
	fmt.Printf("GPS enabled: %t\n", car.Enabled)

	// Access embedded fields through struct names
	fmt.Printf("Engine status: %s\n", car.AutoEngine.Status())
	fmt.Printf("Wheels: %s\n", car.VehicleWheels.Description())
	fmt.Println()
}

// Example 2: Method promotion
func methodPromotionExample() {
	fmt.Println(Header("2. Method Promotion"))

	car := AutoCar{
		AutoEngine:    AutoEngine{Horsepower: 250, Fuel: "premium"},
		VehicleWheels: VehicleWheels{Count: 4, Size: "19 inch"},
		NavigationGPS: NavigationGPS{Latitude: 34.0522, Longitude: -118.2437, Enabled: true},
		Brand:         "BMW",
		Model:         "M3",
		Year:          2023,
	}

	// Promoted methods from embedded structs
	fmt.Printf("Starting %s:\n", car.String())
	car.Start() // Promoted from AutoEngine

	car.Enable()             // Promoted from NavigationGPS
	car.Navigate("Downtown") // Promoted from NavigationGPS

	lat, lng := car.GetLocation() // Promoted from NavigationGPS
	fmt.Printf("Current location: %.4f, %.4f\n", lat, lng)

	fmt.Printf("Wheels description: %s\n", car.Description()) // Promoted from VehicleWheels

	car.Stop() // Promoted from AutoEngine
	fmt.Println()
}

// Example 3: Interface embedding
func interfaceEmbeddingExample() {
	fmt.Println(Header("3. Interface Embedding"))

	smartGPS := &IntelligentGPS{
		NavigationGPS: NavigationGPS{
			Latitude:  37.7749,
			Longitude: -122.4194,
			Enabled:   true,
		},
	}

	// IntelligentGPS implements EnhancedNavigator (which embeds VehicleNavigator)
	var navigator EnhancedNavigator = smartGPS

	// Use embedded interface methods
	navigator.Navigate("Golden Gate Bridge")
	lat, lng := navigator.GetLocation()
	fmt.Printf("Current position: %.4f, %.4f\n", lat, lng)

	// Use additional methods
	navigator.SetDestination("Fisherman's Wharf")
	route := navigator.GetRoute()
	fmt.Printf("Route: %v\n", route)
	fmt.Println()
}

// Example 4: Embedding vs Composition
func embeddingVsCompositionExample() {
	fmt.Println(Header("4. Embedding vs Composition"))

	// Embedding example (is-a relationship)
	car := AutoCar{
		AutoEngine:    AutoEngine{Horsepower: 180, Fuel: "gasoline"},
		VehicleWheels: VehicleWheels{Count: 4, Size: "16 inch"},
		Brand:         "Honda",
		Model:         "Civic",
		Year:          2023,
	}

	motorcycle := AutoMotorcycle{
		AutoEngine:    AutoEngine{Horsepower: 100, Fuel: "gasoline"},
		VehicleWheels: VehicleWheels{Count: 2, Size: "17 inch"},
		Brand:         "Harley-Davidson",
		Model:         "Sportster",
		HasSidecar:    false,
	}

	// Composition example (has-a relationship)
	fleet := VehicleFleet{
		manager: "John Doe",
	}

	fleet.AddVehicle(&car)
	fleet.AddVehicle(&motorcycle)
	fleet.StartAll()

	fmt.Println("\nEmbedding provides 'is-a' relationship")
	fmt.Println("Composition provides 'has-a' relationship")
	fmt.Println()
}

// Example 5: Method shadowing
func methodShadowingExample() {
	fmt.Println(Header("5. Method Shadowing"))

	sportsCar := PerformanceCar{
		AutoCar: AutoCar{
			AutoEngine: AutoEngine{Horsepower: 400, Fuel: "premium"},
			Brand:      "Ferrari",
			Model:      "488",
			Year:       2023,
		},
		turbo: false,
	}

	fmt.Printf("Starting %s:\n", sportsCar.String())

	// Method without turbo
	sportsCar.Start()
	sportsCar.Stop()

	// Enable turbo and start again
	sportsCar.EnableTurbo()
	sportsCar.Start() // This calls the overridden method

	// Can still call the original method explicitly
	fmt.Println("Calling original engine start method:")
	sportsCar.AutoCar.Start()

	sportsCar.Stop()
	fmt.Println()
}

// Example 6: Complex embedding scenarios
func complexEmbeddingExample() {
	fmt.Println(Header("6. Complex Embedding"))

	luxuryCar := PremiumCar{
		AutoCar: AutoCar{
			AutoEngine:    AutoEngine{Horsepower: 300, Fuel: "premium"},
			VehicleWheels: VehicleWheels{Count: 4, Size: "20 inch"},
			Brand:         "Mercedes",
			Model:         "S-Class",
			Year:          2023,
		},
		IntelligentGPS: IntelligentGPS{
			NavigationGPS: NavigationGPS{
				Latitude:  51.5074,
				Longitude: -0.1278,
				Enabled:   true,
			},
		},
	}

	fmt.Printf("Luxury car: %s\n", luxuryCar.String())

	// Access methods from multiple embedded types
	luxuryCar.Start()
	luxuryCar.EnableLuxuryFeatures()

	// IntelligentGPS methods override regular GPS
	luxuryCar.SetDestination("Buckingham Palace")
	route := luxuryCar.GetRoute()
	fmt.Printf("Navigation route: %v\n", route)

	luxuryCar.Stop()
	fmt.Println()
}

// Example 7: Embedding with interfaces
func embeddingWithInterfacesExample() {
	fmt.Println(Header("7. Embedding with Interfaces"))

	// Create different vehicle types
	car := &AutoCar{
		AutoEngine: AutoEngine{Horsepower: 200, Fuel: "gasoline"},
		Brand:      "Toyota",
		Model:      "Prius",
		Year:       2023,
	}

	truck := &AutoTruck{
		AutoEngine: AutoEngine{Horsepower: 400, Fuel: "diesel"},
		PayloadKg:  5000,
		Brand:      "Ford",
		Model:      "F-150",
	}

	// All implement VehicleStarter interface through embedding
	starters := []VehicleStarter{car, truck}

	fmt.Println("Starting all vehicles:")
	for _, starter := range starters {
		starter.Start()
	}

	// All implement VehicleHonker interface
	honkers := []VehicleHonker{car, truck}

	fmt.Println("\nHonking all vehicles:")
	for _, honker := range honkers {
		fmt.Println(honker.Honk())
	}

	fmt.Println("\nStopping all vehicles:")
	for _, starter := range starters {
		starter.Stop()
	}
	fmt.Println()
}

// Move these types to package level

type ComponentA struct {
	Name string
}

type ComponentB struct {
	Name string
}

type ComponentC struct {
	ComponentA
	ComponentB
	Name string // Resolves conflict
}

// Example 8: Embedding conflicts and resolution
func embeddingConflictExample() {
	fmt.Println(Header("8. Embedding Conflicts"))

	// Example of potential naming conflicts
	c := ComponentC{
		ComponentA: ComponentA{Name: "From A"},
		ComponentB: ComponentB{Name: "From B"},
		Name:       "From C",
	}

	fmt.Printf("C.Name: %s\n", c.Name)
	fmt.Printf("C.ComponentA.Name: %s\n", c.ComponentA.Name)
	fmt.Printf("C.ComponentB.Name: %s\n", c.ComponentB.Name)

	// Method calls need to be explicit when there's conflict
	fmt.Printf("A.GetName(): %s\n", c.ComponentA.GetName())
	fmt.Printf("B.GetName(): %s\n", c.ComponentB.GetName())
	fmt.Println()
}

// Move these methods to package level
func (a ComponentA) GetName() string {
	return "A: " + a.Name
}

func (b ComponentB) GetName() string {
	return "B: " + b.Name
}

// Move these types and method to package level

type TimestampedEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (te *TimestampedEntity) Touch() {
	te.UpdatedAt = time.Now()
}

type SystemUser struct {
	TimestampedEntity // Good: common behavior
	ID                int
	Name              string
	Email             string
}

type SystemProduct struct {
	TimestampedEntity // Good: reusable pattern
	ID                int
	Name              string
	Price             float64
}

// Example 9: Best practices for embedding
func embeddingBestPracticesExample() {
	fmt.Println(Header("9. Embedding Best Practices"))

	fmt.Println("Best practices for embedding:")
	fmt.Println("1. Use embedding for 'is-a' relationships")
	fmt.Println("2. Use composition for 'has-a' relationships")
	fmt.Println("3. Prefer small, focused embedded types")
	fmt.Println("4. Document method promotion clearly")
	fmt.Println("5. Be careful with naming conflicts")
	fmt.Println("6. Consider interface embedding for extensibility")
	fmt.Println()

	// Example of good embedding design
	user := SystemUser{
		TimestampedEntity: TimestampedEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	product := SystemProduct{
		TimestampedEntity: TimestampedEntity{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ID:    100,
		Name:  "Laptop",
		Price: 999.99,
	}

	fmt.Printf("User created: %s\n", user.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Product created: %s\n", product.CreatedAt.Format("2006-01-02 15:04:05"))

	// Update timestamps
	user.Touch()
	product.Touch()

	fmt.Printf("User updated: %s\n", user.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Product updated: %s\n", product.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println()
}

// Move these handler types to package level

type BaseHandler struct {
	logger    *Logger
	startTime time.Time
}

type UserHandler struct {
	BaseHandler // Embedded common functionality
	userDB      map[int]string
}

type ProductHandler struct {
	BaseHandler // Embedded common functionality
	productDB   map[int]string
}

// Move these methods to package level
func (bh *BaseHandler) LogRequest(method, path string) {
	if bh.logger != nil {
		bh.logger.Log(fmt.Sprintf("%s %s", method, path))
	}
}

func (bh *BaseHandler) Uptime() time.Duration {
	return time.Since(bh.startTime)
}

func (uh *UserHandler) GetUser(id int) string {
	uh.LogRequest("GET", fmt.Sprintf("/users/%d", id))
	if user, exists := uh.userDB[id]; exists {
		return user
	}
	return "User not found"
}

func (ph *ProductHandler) GetProduct(id int) string {
	ph.LogRequest("GET", fmt.Sprintf("/products/%d", id))
	if product, exists := ph.productDB[id]; exists {
		return product
	}
	return "Product not found"
}

// Example 10: Real-world embedding example
func realWorldExample() {
	fmt.Println(Header("10. Real-World Example"))

	// HTTP server with embedded functionality
	logger := NewLogger("SERVER")

	userHandler := &UserHandler{
		BaseHandler: BaseHandler{
			logger:    logger,
			startTime: time.Now(),
		},
		userDB: map[int]string{
			1: "John Doe",
			2: "Jane Smith",
		},
	}

	productHandler := &ProductHandler{
		BaseHandler: BaseHandler{
			logger:    logger,
			startTime: time.Now(),
		},
		productDB: map[int]string{
			100: "Laptop",
			101: "Mouse",
		},
	}

	// Use handlers
	fmt.Println("API Server Example:")
	fmt.Printf("User 1: %s\n", userHandler.GetUser(1))
	fmt.Printf("Product 100: %s\n", productHandler.GetProduct(100))
	fmt.Printf("User 999: %s\n", userHandler.GetUser(999))

	// Access embedded functionality
	fmt.Printf("Server uptime: %v\n", userHandler.Uptime())
	fmt.Println()
}
