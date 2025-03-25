package main

import (
	"bufio"
	"os"
	"strings"
)

// Load reads a .env file and sets environment variables.
// It does not override existing variables.
func Load(filenames ...string) error {
	for _, filename := range filenames {
		err := loadFile(filename, false)
		if err != nil {
			return err
		}
	}
	return nil
}

// Overload reads a .env file and sets environment variables,
// overriding existing ones.
func Overload(filenames ...string) error {
	for _, filename := range filenames {
		err := loadFile(filename, true)
		if err != nil {
			return err
		}
	}
	return nil
}

// loadFile processes the .env file and sets variables
func loadFile(filename string, overwrite bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore empty lines and full-line comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Remove inline comments
		line = removeInlineComment(line)

		// Extract key-value pair
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		// Set the environment variable
		if overwrite || os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

// removeInlineComment removes inline comments from a line
func removeInlineComment(line string) string {
	parts := strings.SplitN(line, "#", 2)
	return strings.TrimSpace(parts[0])
}
