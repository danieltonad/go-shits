package main

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv reads a .env file and sets environment variables
func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove optional surrounding quotes from values
		value = strings.Trim(value, `"'`)
		// Set environment variable
		os.Setenv(key, value)
	}

	return scanner.Err()
}
