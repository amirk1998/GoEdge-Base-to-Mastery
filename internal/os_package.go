// os_package_examples.go
package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// RunOSPackageExamples - main function to run all OS package examples
func RunOSPackageExamples() {
	commandLineArgsExample()
	environmentVariablesExample()
	fileSystemOperationsExample()
	fileInfoExample()
	processControlExample()
	signalHandlingExample()
	workingDirectoryExample()
	userInfoExample()
	pathManipulationExample()
	temporaryFilesExample()
}

// commandLineArgsExample demonstrates working with command line arguments
func commandLineArgsExample() {
	fmt.Println(SectionHeader("Command Line Arguments"))

	// os.Args contains command line arguments
	fmt.Printf("Program name: %s\n", Bold(os.Args[0]))
	fmt.Printf("Total arguments: %d\n", len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println("Arguments:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("  [%d]: %s\n", i+1, Yellow(arg))
		}
	} else {
		fmt.Println(InfoText("No additional arguments provided"))
	}

	// Advanced: Parse flags manually
	var verbose bool
	var outputFile string

	for i, arg := range os.Args[1:] {
		switch arg {
		case "-v", "--verbose":
			verbose = true
		case "-o", "--output":
			if i+1 < len(os.Args)-1 {
				outputFile = os.Args[i+2]
			}
		}
	}

	fmt.Printf("Verbose mode: %t\n", verbose)
	fmt.Printf("Output file: %s\n", outputFile)
	fmt.Println()
}

// environmentVariablesExample demonstrates environment variable operations
func environmentVariablesExample() {
	fmt.Println(SectionHeader("Environment Variables"))

	// Get environment variable
	path := os.Getenv("PATH")
	fmt.Printf("PATH environment variable (first 100 chars): %s...\n",
		Yellow(path[:min(len(path), 100)]))

	// Set environment variable
	os.Setenv("CUSTOM_VAR", "Hello, World!")
	customVar := os.Getenv("CUSTOM_VAR")
	fmt.Printf("Custom variable: %s\n", Green(customVar))

	// Check if variable exists
	home, exists := os.LookupEnv("HOME")
	if exists {
		fmt.Printf("Home directory: %s\n", Cyan(home))
	} else {
		fmt.Println(WarningText("HOME environment variable not found"))
	}

	// Get all environment variables
	envVars := os.Environ()
	fmt.Printf("Total environment variables: %d\n", len(envVars))

	// Show some system-related variables
	systemVars := []string{"OS", "USER", "USERNAME", "COMPUTERNAME", "HOSTNAME"}
	fmt.Println("System variables:")
	for _, varName := range systemVars {
		if value := os.Getenv(varName); value != "" {
			fmt.Printf("  %s: %s\n", Bold(varName), value)
		}
	}

	// Unset environment variable
	os.Unsetenv("CUSTOM_VAR")
	if os.Getenv("CUSTOM_VAR") == "" {
		fmt.Println(SuccessText("Custom variable successfully unset"))
	}
	fmt.Println()
}

// fileSystemOperationsExample demonstrates file system operations
func fileSystemOperationsExample() {
	fmt.Println(SectionHeader("File System Operations"))

	// Create a test file
	testFile := "test_file.txt"
	content := "Hello, Go File System!"

	// Create and write to file
	file, err := os.Create(testFile)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", ErrorText(err.Error()))
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("File created and written: %s\n", Green(testFile))

	// Check if file exists
	if _, err := os.Stat(testFile); err == nil {
		fmt.Printf("File exists: %s\n", SuccessText("✓"))
	} else if os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", ErrorText("✗"))
	}

	// Read file
	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("File content: %s\n", Yellow(string(data)))

	// Rename file
	newName := "renamed_file.txt"
	err = os.Rename(testFile, newName)
	if err != nil {
		fmt.Printf("Error renaming file: %s\n", ErrorText(err.Error()))
	} else {
		fmt.Printf("File renamed to: %s\n", Green(newName))
	}

	// Create directory
	dirName := "test_directory"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %s\n", ErrorText(err.Error()))
	} else {
		fmt.Printf("Directory created: %s\n", Green(dirName))
	}

	// Create nested directories
	nestedDir := "nested/deep/directory"
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("Error creating nested directories: %s\n", ErrorText(err.Error()))
	} else {
		fmt.Printf("Nested directories created: %s\n", Green(nestedDir))
	}

	// Cleanup
	os.Remove(newName)
	os.Remove(dirName)
	os.RemoveAll("nested")
	fmt.Println(InfoText("Cleanup completed"))
	fmt.Println()
}

// fileInfoExample demonstrates file information operations
func fileInfoExample() {
	fmt.Println(SectionHeader("File Information"))

	// Create a test file with some content
	testFile := "info_test.txt"
	content := "This is a test file for information demo.\nIt has multiple lines.\nAnd some content."

	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %s\n", ErrorText(err.Error()))
		return
	}
	defer os.Remove(testFile)

	// Get file information
	fileInfo, err := os.Stat(testFile)
	if err != nil {
		fmt.Printf("Error getting file info: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("File name: %s\n", Bold(fileInfo.Name()))
	fmt.Printf("File size: %s bytes\n", Yellow(fmt.Sprintf("%d", fileInfo.Size())))
	fmt.Printf("File mode: %s\n", Cyan(fileInfo.Mode().String()))
	fmt.Printf("Modification time: %s\n", Green(fileInfo.ModTime().Format("2006-01-02 15:04:05")))
	fmt.Printf("Is directory: %t\n", fileInfo.IsDir())

	// Check file permissions
	mode := fileInfo.Mode()
	fmt.Printf("Permissions: %s\n", mode.Perm().String())
	fmt.Printf("Is regular file: %t\n", mode.IsRegular())

	// Get current directory info
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %s\n", ErrorText(err.Error()))
		return
	}

	dirInfo, err := os.Stat(currentDir)
	if err != nil {
		fmt.Printf("Error getting directory info: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Current directory: %s\n", Bold(currentDir))
	fmt.Printf("Directory modification time: %s\n",
		Green(dirInfo.ModTime().Format("2006-01-02 15:04:05")))
	fmt.Println()
}

// processControlExample demonstrates process control operations
func processControlExample() {
	fmt.Println(SectionHeader("Process Control"))

	// Get process information
	fmt.Printf("Process ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getpid())))
	fmt.Printf("Parent Process ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getppid())))

	// Get user and group IDs (Unix-like systems)
	fmt.Printf("User ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getuid())))
	fmt.Printf("Group ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getgid())))

	// Execute external command
	fmt.Println("\nExecuting external command...")

	// Cross-platform command execution
	var cmd *exec.Cmd
	switch os.Getenv("OS") {
	case "Windows_NT":
		cmd = exec.Command("cmd", "/C", "echo", "Hello from Windows!")
	default:
		cmd = exec.Command("echo", "Hello from Unix-like system!")
	}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", ErrorText(err.Error()))
	} else {
		fmt.Printf("Command output: %s", Green(string(output)))
	}

	// Demonstrate process exit
	fmt.Println(InfoText("Process control example completed"))
	// Note: os.Exit() would terminate the program, so we don't use it here
	fmt.Println()
}

// signalHandlingExample demonstrates signal handling (basic example)
func signalHandlingExample() {
	fmt.Println(SectionHeader("Signal Handling"))

	// Note: Full signal handling requires the os/signal package
	// This is a basic demonstration of signal-related concepts

	fmt.Println(InfoText("Signal handling typically requires os/signal package"))
	fmt.Println(InfoText("Common signals: SIGINT (Ctrl+C), SIGTERM, SIGKILL"))

	// Demonstrate process termination concepts
	fmt.Println("Process termination methods:")
	fmt.Println("  - os.Exit(code) - terminates immediately")
	fmt.Println("  - return from main() - normal termination")
	fmt.Println("  - panic() - abnormal termination")

	fmt.Println(InfoText("For full signal handling, use os/signal package"))
	fmt.Println()
}

// workingDirectoryExample demonstrates working directory operations
func workingDirectoryExample() {
	fmt.Println(SectionHeader("Working Directory Operations"))

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Current directory: %s\n", Bold(currentDir))

	// Create a test directory
	testDir := "test_workdir"
	err = os.Mkdir(testDir, 0755)
	if err != nil {
		fmt.Printf("Error creating test directory: %s\n", ErrorText(err.Error()))
		return
	}
	defer os.Remove(testDir)

	// Change to test directory
	err = os.Chdir(testDir)
	if err != nil {
		fmt.Printf("Error changing directory: %s\n", ErrorText(err.Error()))
		return
	}

	// Verify directory change
	newDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting new directory: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Changed to directory: %s\n", Green(newDir))

	// Change back to original directory
	err = os.Chdir(currentDir)
	if err != nil {
		fmt.Printf("Error changing back to original directory: %s\n", ErrorText(err.Error()))
		return
	}

	// Verify we're back
	backDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting directory after change back: %s\n", ErrorText(err.Error()))
		return
	}
	fmt.Printf("Back to original directory: %s\n", Green(backDir))
	fmt.Println()
}

// userInfoExample demonstrates user information operations
func userInfoExample() {
	fmt.Println(SectionHeader("User Information"))

	// Get user information
	fmt.Printf("User ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getuid())))
	fmt.Printf("Effective User ID: %s\n", Yellow(fmt.Sprintf("%d", os.Geteuid())))
	fmt.Printf("Group ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getgid())))
	fmt.Printf("Effective Group ID: %s\n", Yellow(fmt.Sprintf("%d", os.Getegid())))

	// Get user-related environment variables
	userVars := map[string]string{
		"USER":        os.Getenv("USER"),
		"USERNAME":    os.Getenv("USERNAME"),
		"HOME":        os.Getenv("HOME"),
		"USERPROFILE": os.Getenv("USERPROFILE"),
	}

	fmt.Println("User environment variables:")
	for key, value := range userVars {
		if value != "" {
			fmt.Printf("  %s: %s\n", Bold(key), value)
		}
	}

	// Get groups (requires additional packages for full functionality)
	fmt.Println(InfoText("For detailed user/group info, use os/user package"))
	fmt.Println()
}

// pathManipulationExample demonstrates path manipulation
func pathManipulationExample() {
	fmt.Println(SectionHeader("Path Manipulation"))

	// Basic path operations
	testPath := "/home/user/documents/file.txt"
	fmt.Printf("Original path: %s\n", Bold(testPath))

	// Extract directory and filename
	dir := filepath.Dir(testPath)
	base := filepath.Base(testPath)
	ext := filepath.Ext(testPath)

	fmt.Printf("Directory: %s\n", Yellow(dir))
	fmt.Printf("Base name: %s\n", Green(base))
	fmt.Printf("Extension: %s\n", Cyan(ext))

	// Join paths
	joinedPath := filepath.Join("home", "user", "documents", "newfile.txt")
	fmt.Printf("Joined path: %s\n", Bold(joinedPath))

	// Split path
	pathParts := strings.Split(testPath, string(filepath.Separator))
	fmt.Printf("Path parts: %v\n", pathParts)

	// Absolute path
	relPath := "./test.txt"
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %s\n", ErrorText(err.Error()))
	} else {
		fmt.Printf("Relative path: %s\n", Yellow(relPath))
		fmt.Printf("Absolute path: %s\n", Green(absPath))
	}

	// Path separator
	fmt.Printf("Path separator: %s\n", Bold(string(filepath.Separator)))
	fmt.Printf("List separator: %s\n", Bold(string(filepath.ListSeparator)))
	fmt.Println()
}

// temporaryFilesExample demonstrates temporary file operations
func temporaryFilesExample() {
	fmt.Println(SectionHeader("Temporary Files"))

	// Get temporary directory
	tempDir := os.TempDir()
	fmt.Printf("System temporary directory: %s\n", Bold(tempDir))

	// Create temporary file
	tempFile, err := os.CreateTemp(tempDir, "example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temporary file: %s\n", ErrorText(err.Error()))
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up
	defer tempFile.Close()

	fmt.Printf("Temporary file created: %s\n", Green(tempFile.Name()))

	// Write to temporary file
	content := "This is temporary content with timestamp: " + time.Now().Format("2006-01-02 15:04:05")
	_, err = tempFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to temporary file: %s\n", ErrorText(err.Error()))
		return
	}

	// Read from temporary file
	tempFile.Seek(0, 0) // Reset file position
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Printf("Error reading temporary file: %s\n", ErrorText(err.Error()))
		return
	}

	fmt.Printf("Temporary file content: %s\n", Yellow(string(data)))

	// Create temporary directory
	tempSubDir, err := os.MkdirTemp(tempDir, "example_dir_*")
	if err != nil {
		fmt.Printf("Error creating temporary directory: %s\n", ErrorText(err.Error()))
		return
	}
	defer os.RemoveAll(tempSubDir) // Clean up

	fmt.Printf("Temporary directory created: %s\n", Green(tempSubDir))

	fmt.Println(SuccessText("Temporary files example completed"))
	fmt.Println()
}

// Helper function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
