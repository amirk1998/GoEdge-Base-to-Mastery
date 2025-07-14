# GoEdge: Base to Mastery

A comprehensive Go learning project that covers fundamental to advanced concepts with practical examples.

## 🚀 Project Structure

```
GoEdge-Base-to-Mastery/
├── cmd/
│   └── goedge/
│       └── main.go          # Main application entry point
├── internal/                # Internal packages with examples
│   ├── arrays_slices.go
│   ├── channels.go
│   ├── colors.go
│   ├── context.go
│   ├── defer_panic_recover.go
│   ├── embedding_composition.go
│   ├── errors.go
│   ├── file_io.go
│   ├── functions.go
│   ├── goroutines.go
│   ├── interfaces.go
│   ├── ioutil_examples.go
│   ├── io_examples.go
│   ├── io_package.go
│   ├── json_serialization.go
│   ├── maps.go
│   ├── methods.go
│   ├── os_examples.go
│   ├── os_package.go
│   ├── package_system.go
│   ├── pointers.go
│   ├── reflection.go
│   ├── string_formatting.go
│   ├── structs.go
│   └── type_system.go
├── examples/                # Additional examples
├── test/                    # Test files
├── .github/
│   └── workflows/
│       └── go.yml          # GitHub Actions CI/CD
├── go.mod                  # Go module definition
├── Makefile               # Build automation
└── README.md              # This file
```

## 📚 Topics Covered

- **🔗 Pointers** - Memory addresses and pointer operations
- **🔧 Functions** - Function definitions, parameters, and return values
- **📊 Arrays & Slices** - Data structures and manipulation
- **🗺️ Maps** - Key-value data structures
- **🔄 Defer/Panic/Recover** - Error handling and resource management
- **📝 String Formatting** - Text processing and formatting
- **📦 Structs** - Custom data types and structures
- **📦 Methods** - Method definitions and receivers
- **🔌 Interfaces** - Interface definitions and implementations
- **🔌 Errors** - Error handling patterns
- **🚀 Goroutines** - Concurrent programming
- **📺 Channels** - Communication between goroutines
- **📦 Package System** - Code organization and imports
- **🧩 Embedding & Composition** - Code reuse patterns
- **🔍 Reflection** - Runtime type inspection
- **🌐 Context** - Request-scoped values and cancellation
- **📋 JSON & Serialization** - Data serialization and deserialization
- **📁 File I/O** - File operations and readers/writers
- **🖥️ OS Package** - Operating system interactions
- **📄 IO Package** - Input/output operations
- **📁 IO/ioutil Package** - I/O utility functions
- **🖥️ System Interaction** - System-level programming
- **📄 I/O Streams** - Stream processing

## 🛠️ Installation

1. **Prerequisites**
   - Go 1.21+ installed
   - Git installed

2. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/GoEdge-Base-to-Mastery.git
   cd GoEdge-Base-to-Mastery
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

## 🎯 Usage

### Run specific topic examples:
```bash
go run ./cmd/goedge <topic>
```

### Available topics:
- `pointers` - Pointer examples
- `functions` - Function examples
- `arrays` - Array & Slice examples
- `maps` - Map examples
- `defer` - Defer/Panic/Recover examples
- `strings` - String formatting examples
- `structs` - Structs examples
- `methods` - Method examples
- `interfaces` - Interface examples
- `errors` - Errors examples
- `goroutines` - Goroutine examples
- `channels` - Channel examples
- `packages` - Package System & Imports examples
- `embedding` - Embedding & Composition examples
- `reflection` - Reflection examples
- `context` - Context Package examples
- `json` - JSON & Serialization examples
- `fileio` - File I/O & Readers/Writers examples
- `os` - OS Package examples
- `io` - IO Package examples
- `ioutil` - IO/ioutil Package examples
- `system` - System Interaction examples
- `streams` - I/O Streams examples
- `colors` - Color examples
- `all` - Run all examples

### Examples:
```bash
# Run pointer examples
go run ./cmd/goedge pointers

# Run all examples
go run ./cmd/goedge all

# Show help
go run ./cmd/goedge
```

## 🔧 Development

### Using Makefile:
```bash
# Build the project
make build

# Run tests
make test

# Clean build artifacts
make clean

# Run specific example
make run TOPIC=pointers
```

### Manual commands:
```bash
# Build
go build -o bin/goedge ./cmd/goedge

# Test
go test ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

## 🚀 CI/CD

This project uses GitHub Actions for continuous integration. The workflow:
- Runs on Go 1.21+
- Executes tests
- Performs code formatting checks
- Runs go vet for static analysis

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎓 Learning Path

1. **Start with basics**: `pointers`, `functions`, `arrays`
2. **Data structures**: `maps`, `structs`
3. **Advanced concepts**: `interfaces`, `methods`, `embedding`
4. **Concurrency**: `goroutines`, `channels`, `context`
5. **I/O and system**: `fileio`, `os`, `io`
6. **Advanced topics**: `reflection`, `json`

## 🎯 Goals

This project aims to:
- Provide practical, runnable examples for each Go concept
- Demonstrate best practices and common patterns
- Serve as a reference for Go developers
- Facilitate learning through hands-on examples

## 📞 Support

If you have any questions or issues, please:
1. Check the existing examples
2. Open an issue on GitHub
3. Refer to the official Go documentation

---

**Happy Learning! 🎉**