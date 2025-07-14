// cmd/goedge/main.go
package main

import (
	"fmt"
	"github.com/amirk1998/GoEdge-Base-to-Mastery/internal"
	"os"
)

//// Version information (set by build flags)
//var (
//	version = "dev"
//	commit  = "unknown"
//	date    = "unknown"
//)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	topic := os.Args[1]

	switch topic {
	//case "version", "-v", "--version":
	//	fmt.Printf("GoEdge v%s\n", version)
	//	fmt.Printf("Commit: %s\n", commit)
	//	fmt.Printf("Built: %s\n", date)
	//	return
	case "pointers":
		fmt.Println(internal.Header("🔗 Running Pointer Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunPointerExamples()
	case "functions":
		fmt.Println(internal.Header("🔧 Running Function Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunFunctionExamples()
	case "arrays":
		fmt.Println(internal.Header("📊 Running Array & Slice Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunArraySliceExamples()
	case "maps":
		fmt.Println(internal.Header("🗺️ Running Map Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunMapExamples()
	case "defer":
		fmt.Println(internal.Header("🔄 Running Defer/Panic/Recover Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunDeferPanicRecoverExamples()
	case "strings":
		fmt.Println(internal.Header("📝 Running String Formatting Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunStringFormattingExamples()
	case "methods":
		fmt.Println(internal.Header("📦 Running Method Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunMethodExamples()
	case "structs":
		fmt.Println(internal.Header("📦 Running Structs Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunStructureExamples()
	case "interfaces":
		fmt.Println(internal.Header("🔌 Running Interface Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunInterfaceExamples()
	case "errors":
		fmt.Println(internal.Header("🔌 Running Errors Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunErrorHandlingExamples()
	case "goroutines":
		fmt.Println(internal.Header("🚀 Running Goroutine Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunGoroutineExamples()
	case "channels":
		fmt.Println(internal.Header("📺 Running Channel Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunChannelExamples()
	case "packages":
		fmt.Println(internal.Header("📦 Running Package System Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunPackageSystemExamples()
	case "embedding":
		fmt.Println(internal.Header("🧩 Running Embedding & Composition Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunEmbeddingCompositionExamples()
	case "reflection":
		fmt.Println(internal.Header("🔍 Running Reflection Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunReflectionExamples()
	case "context":
		fmt.Println(internal.Header("🌐 Running Context Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunContextExamples()
	case "json":
		fmt.Println(internal.Header("📋 Running JSON & Serialization Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunJSONSerializationExamples()
	case "fileio":
		fmt.Println(internal.Header("📁 Running File I/O & Readers/Writers Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunFileIOExamples()
	case "os":
		fmt.Println(internal.Header("🖥️ Running OS Package Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunOSExamples()
	case "io":
		fmt.Println(internal.Header("📄 Running IO Package Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunIOExamples()
	case "ioutil":
		fmt.Println(internal.Header("📁 Running IO/ioutil Package Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunIOUtilExamples()
	case "system":
		fmt.Println(internal.Header("🖥️ Running System Interaction Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunOSPackageExamples()
	case "streams":
		fmt.Println(internal.Header("📄 Running I/O Streams Examples:"))
		fmt.Println(internal.Cyan("=" + repeat("=", 40)))
		internal.RunIOPackageExamples()
	case "colors":
		internal.ColorExamples()
	case "all":
		runAllExamples()
	default:
		fmt.Println(internal.ErrorText(fmt.Sprintf("Unknown topic: %s", topic)))
		showHelp()
	}
}

func showHelp() {
	fmt.Println(internal.Header("🐹 Golang Review Project"))
	fmt.Printf("Version: %s (commit: %s)\n", version, commit)
	fmt.Println(internal.Cyan("=" + repeat("=", 40)))
	fmt.Println(internal.Bold("Usage:"), "go run ./cmd/goedge <topic>")
	fmt.Println("\n" + internal.Subtitle("Available topics:"))

	topics := []struct {
		name, desc string
	}{
		{"pointers", "Pointer examples"},
		{"functions", "Function examples"},
		{"arrays", "Array & Slice examples"},
		{"maps", "Map examples"},
		{"defer", "Defer/Panic/Recover examples"},
		{"strings", "String formatting examples"},
		{"structs", "Structs examples"},
		{"methods", "Method examples"},
		{"interfaces", "Interface examples"},
		{"errors", "Errors examples"},
		{"goroutines", "Goroutine examples"},
		{"channels", "Channel examples"},
		{"packages", "Package System & Imports examples"},
		{"embedding", "Embedding & Composition examples"},
		{"reflection", "Reflection examples"},
		{"context", "Context Package examples"},
		{"json", "JSON & Serialization examples"},
		{"fileio", "File I/O & Readers/Writers examples"},
		{"os", "OS Package examples"},
		{"io", "IO Package examples"},
		{"ioutil", "IO/ioutil Package examples"},
		{"system", "System Interaction examples"},
		{"streams", "I/O Streams examples"},
		{"colors", "Color examples"},
		{"all", "Run all examples"},
	}

	for _, topic := range topics {
		fmt.Printf("  %s - %s\n",
			internal.Yellow(topic.name),
			topic.desc)
	}

	fmt.Println("\n" + internal.InfoText("Example: go run ./cmd/goedge json"))
}

func runAllExamples() {
	topics := []struct {
		name string
		fn   func()
	}{
		{"🔗 Pointers", internal.RunPointerExamples},
		{"🔧 Functions", internal.RunFunctionExamples},
		{"📊 Arrays & Slices", internal.RunArraySliceExamples},
		{"🗺️ Maps", internal.RunMapExamples},
		{"🔄 Defer/Panic/Recover", internal.RunDeferPanicRecoverExamples},
		{"📝 String Formatting", internal.RunStringFormattingExamples},
		{"📦 Methods", internal.RunMethodExamples},
		{"📦 Structs", internal.RunStructureExamples},
		{"🔌 Interfaces", internal.RunInterfaceExamples},
		{"🔌 Errors", internal.RunErrorHandlingExamples},
		{"🚀 Goroutines", internal.RunGoroutineExamples},
		{"📺 Channels", internal.RunChannelExamples},
		{"📦 Package System", internal.RunPackageSystemExamples},
		{"🧩 Embedding & Composition", internal.RunEmbeddingCompositionExamples},
		{"🔍 Reflection", internal.RunReflectionExamples},
		{"🌐 Context", internal.RunContextExamples},
		{"📋 JSON & Serialization", internal.RunJSONSerializationExamples},
		{"📁 File I/O & Readers/Writers", internal.RunFileIOExamples},
		{"🖥️ OS Package", internal.RunOSExamples},
		{"📄 IO Package", internal.RunIOExamples},
		{"📁 IO/ioutil Package", internal.RunIOUtilExamples},
		{"🖥️ System Interaction", internal.RunOSPackageExamples},
		{"📄 I/O Streams", internal.RunIOPackageExamples},
	}

	for i, topic := range topics {
		fmt.Printf("\n%s Examples:\n", internal.Header(topic.name))
		fmt.Println(internal.Cyan("=" + repeat("=", 50)))
		topic.fn()

		if i < len(topics)-1 {
			fmt.Println("\n" + internal.Dim(repeat("-", 50)))
		}
	}
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
