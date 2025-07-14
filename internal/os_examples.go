// os_examples.go
package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RunOSExamples - main function to run all OS package examples
func RunOSExamples() {
	fmt.Println(Subtitle("🖥️  OS Package Examples"))
	fmt.Println()

	environmentVariablesDemo()
	commandLineArgumentsDemo()
	fileSystemOperationsDemo()
	fileInfoDemo()
	workingDirectoryDemo()
	processInfoDemo()
	filePermissionsDemo()
	temporaryFilesDemo()
}

// Environment Variables Operations
func environmentVariablesDemo() {
	fmt.Println(Yellow("📌 Environment Variables:"))

	// Set environment variable
	os.Setenv("MY_APP_NAME", "GoLang Tutorial")
	os.Setenv("DEBUG_MODE", "true")

	// Get environment variable
	appName := os.Getenv("MY_APP_NAME")
	debugMode := os.Getenv("DEBUG_MODE")
	nonExistent := os.Getenv("NON_EXISTENT")

	fmt.Printf("App Name: %s\n", Green(appName))
	fmt.Printf("Debug Mode: %s\n", Green(debugMode))
	fmt.Printf("Non-existent var: '%s' (empty)\n", Red(nonExistent))

	// Check if environment variable exists
	if value, exists := os.LookupEnv("HOME"); exists {
		fmt.Printf("HOME directory: %s\n", Cyan(value))
	}

	// Get all environment variables
	fmt.Println(Bold("All environment variables:"))
	envVars := os.Environ()
	for i, env := range envVars {
		if i < 3 { // Show only first 3 for brevity
			fmt.Printf("  %s\n", Dim(env))
		}
	}
	fmt.Printf("  ... and %d more\n", len(envVars)-3)

	// Unset environment variable
	os.Unsetenv("MY_APP_NAME")
	fmt.Println()
}

// Command Line Arguments Processing
func commandLineArgumentsDemo() {
	fmt.Println(Yellow("📌 Command Line Arguments:"))

	fmt.Printf("Program name: %s\n", Green(os.Args[0]))
	fmt.Printf("Number of arguments: %d\n", len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println(Bold("Arguments:"))
		for i, arg := range os.Args[1:] {
			fmt.Printf("  [%d]: %s\n", i, Cyan(arg))
		}
	} else {
		fmt.Println(Dim("No additional arguments provided"))
	}
	fmt.Println()
}

// File System Operations
func fileSystemOperationsDemo() {
	fmt.Println(Yellow("📌 File System Operations:"))

	// Create directory
	dirName := "test_directory"
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Printf("Directory created: %s\n", Green(dirName))

	// Create nested directories
	nestedDir := filepath.Join(dirName, "nested", "deep")
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("Error creating nested directory: %v\n", err)
		return
	}
	fmt.Printf("Nested directory created: %s\n", Green(nestedDir))

	// Create a file
	fileName := filepath.Join(dirName, "test_file.txt")
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	file.WriteString("Hello, Go OS package!")
	file.Close()
	fmt.Printf("File created: %s\n", Green(fileName))

	// Check if file/directory exists
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File exists: %s\n", Cyan(fileName))
	} else if os.IsNotExist(err) {
		fmt.Printf("File does not exist: %s\n", Red(fileName))
	}

	// Rename file
	newFileName := filepath.Join(dirName, "renamed_file.txt")
	err = os.Rename(fileName, newFileName)
	if err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
	} else {
		fmt.Printf("File renamed to: %s\n", Green(newFileName))
	}

	// Remove file and directories (cleanup)
	os.Remove(newFileName)
	os.RemoveAll(dirName)
	fmt.Printf("Cleanup completed\n")
	fmt.Println()
}

// File Information and Metadata
func fileInfoDemo() {
	fmt.Println(Yellow("📌 File Information:"))

	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "fileinfo_test_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Cleanup

	// Write some content
	content := "This is a test file for demonstrating file info operations."
	tempFile.WriteString(content)
	tempFile.Close()

	// Get file information
	fileInfo, err := os.Stat(tempFile.Name())
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("File name: %s\n", Green(fileInfo.Name()))
	fmt.Printf("File size: %d bytes\n", fileInfo.Size())
	fmt.Printf("File mode: %s\n", Cyan(fileInfo.Mode().String()))
	fmt.Printf("Is directory: %t\n", fileInfo.IsDir())
	fmt.Printf("Modification time: %s\n", Yellow(fileInfo.ModTime().Format(time.RFC3339)))

	// Check file permissions
	mode := fileInfo.Mode()
	fmt.Printf("Permissions:\n")
	fmt.Printf("  Owner: %s\n", getPermissionString(mode>>6&7))
	fmt.Printf("  Group: %s\n", getPermissionString(mode>>3&7))
	fmt.Printf("  Other: %s\n", getPermissionString(mode&7))
	fmt.Println()
}

// Working Directory Operations
func workingDirectoryDemo() {
	fmt.Println(Yellow("📌 Working Directory Operations:"))

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}
	fmt.Printf("Current directory: %s\n", Green(currentDir))

	// Create a test directory
	testDir := "temp_work_dir"
	os.Mkdir(testDir, 0755)
	defer os.RemoveAll(testDir) // Cleanup

	// Change working directory
	originalDir := currentDir
	err = os.Chdir(testDir)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return
	}

	// Verify directory change
	newDir, _ := os.Getwd()
	fmt.Printf("Changed to: %s\n", Cyan(newDir))

	// Change back to original directory
	os.Chdir(originalDir)
	restoredDir, _ := os.Getwd()
	fmt.Printf("Restored to: %s\n", Green(restoredDir))
	fmt.Println()
}

// Process Information
func processInfoDemo() {
	fmt.Println(Yellow("📌 Process Information:"))

	// Get process ID
	pid := os.Getpid()
	fmt.Printf("Process ID: %s\n", Green(fmt.Sprintf("%d", pid)))

	// Get parent process ID
	ppid := os.Getppid()
	fmt.Printf("Parent Process ID: %s\n", Cyan(fmt.Sprintf("%d", ppid)))

	// Get user ID (Unix-like systems)
	uid := os.Getuid()
	fmt.Printf("User ID: %s\n", Yellow(fmt.Sprintf("%d", uid)))

	// Get group ID (Unix-like systems)
	gid := os.Getgid()
	fmt.Printf("Group ID: %s\n", Yellow(fmt.Sprintf("%d", gid)))

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", Green(hostname))
	}
	fmt.Println()
}

// File Permissions Example
func filePermissionsDemo() {
	fmt.Println(Yellow("📌 File Permissions:"))

	// Create a test file
	testFile := "permission_test.txt"
	file, err := os.Create(testFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	file.Close()
	defer os.Remove(testFile) // Cleanup

	// Change file permissions
	err = os.Chmod(testFile, 0644) // rw-r--r--
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
		return
	}

	// Check permissions
	fileInfo, _ := os.Stat(testFile)
	fmt.Printf("File permissions: %s\n", Green(fileInfo.Mode().String()))

	// Change to different permissions
	err = os.Chmod(testFile, 0755) // rwxr-xr-x
	if err == nil {
		fileInfo, _ = os.Stat(testFile)
		fmt.Printf("Updated permissions: %s\n", Cyan(fileInfo.Mode().String()))
	}
	fmt.Println()
}

// Temporary Files and Directories
func temporaryFilesDemo() {
	fmt.Println(Yellow("📌 Temporary Files and Directories:"))

	// Get temporary directory
	tempDir := os.TempDir()
	fmt.Printf("System temp directory: %s\n", Green(tempDir))

	// Create temporary file
	tempFile, err := os.CreateTemp("", "golang_example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Cleanup

	fmt.Printf("Created temp file: %s\n", Cyan(tempFile.Name()))

	// Write to temporary file
	content := "This is temporary content"
	tempFile.WriteString(content)
	tempFile.Close()

	// Create temporary directory
	tempDirPath, err := os.MkdirTemp("", "golang_example_dir_*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDirPath) // Cleanup

	fmt.Printf("Created temp directory: %s\n", Yellow(tempDirPath))

	// Create file in temporary directory
	tempFileInDir := filepath.Join(tempDirPath, "nested_temp.txt")
	nestedFile, err := os.Create(tempFileInDir)
	if err == nil {
		nestedFile.WriteString("Nested temporary file content")
		nestedFile.Close()
		fmt.Printf("Created nested temp file: %s\n", Green(tempFileInDir))
	}
	fmt.Println()
}

// Helper function to convert permission bits to string
func getPermissionString(perm os.FileMode) string {
	permissions := ""
	if perm&4 != 0 {
		permissions += "r"
	} else {
		permissions += "-"
	}
	if perm&2 != 0 {
		permissions += "w"
	} else {
		permissions += "-"
	}
	if perm&1 != 0 {
		permissions += "x"
	} else {
		permissions += "-"
	}
	return permissions
}
