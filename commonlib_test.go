package gomodlibs

import (
	"os"
	"testing"
)

// Test that generated strings have the correct length
func TestGenerateRandomStringLength(t *testing.T) {
	lengths := []int{5, 10, 20, 50} // Different lengths to test
	for _, length := range lengths {
		randomStr, err := GenerateRandomString(length)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if len(randomStr) != length {
			t.Errorf("Expected length %d, but got %d", length, len(randomStr))
		}
	}
}

// Test that multiple calls generate different results
func TestGenerateRandomStringUniqueness(t *testing.T) {
	randomStr1, err1 := GenerateRandomString(10)
	randomStr2, err2 := GenerateRandomString(10)

	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error: %v, %v", err1, err2)
	}

	if randomStr1 == randomStr2 {
		t.Errorf("Expected different strings, but got identical: %s and %s", randomStr1, randomStr2)
	}
}

// Test that it handles zero-length requests correctly
func TestGenerateRandomStringZeroLength(t *testing.T) {
	randomStr, err := GenerateRandomString(0)
	if err != nil {
		t.Fatalf("Unexpected error for length 0: %v", err)
	}
	if len(randomStr) != 0 {
		t.Errorf("Expected empty string, but got: %s", randomStr)
	}
}

// TestCreateDirIfNotExists verifies that the function correctly creates a directory when it does not exist.
func TestCreateDirIfNotExists(t *testing.T) {
	// Define a test directory name
	testDir := "testdir"

	// Ensure test cleanup after the test finishes
	defer os.RemoveAll(testDir) // Deletes the directory after the test

	// Step 1: Check if directory does not exist initially
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Fatalf("Test directory %s already exists before test", testDir)
	}

	// Step 2: Call the function to create the directory
	err := CreateDirIfNotExists(testDir)
	if err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	// Step 3: Verify that the directory was created
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Fatalf("Directory %s was not created", testDir)
	}

	// Step 4: Call the function again and ensure it does not return an error
	err = CreateDirIfNotExists(testDir)
	if err != nil {
		t.Fatalf("Unexpected error when calling CreateDirIfNotExists on an existing directory: %v", err)
	}
}
