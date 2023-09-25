// encrypt_test.go
package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name          string
		plainText      string
		key           string
		iv            string
		expectedRegex string
		expectError   bool
	}{
		{
			name:          "Valid encryption",
			plainText:      "Hello, World!",
			key:           "0123456789abcdef0123456789abcdef",
			iv:            "1234567812345678",
			expectedRegex: `[0-9a-f]{32}`, // Regex for a 32-character hexadecimal string
			expectError:   false,
		},
		{
			name:          "Invalid encryption (wrong key)",
			plainText:      "Hello, World!",
			key:           "wrongkey",
			iv:            "1234567812345678",
			expectedRegex: ``,
			expectError:   true,
		},
		{
			name:          "Invalid encryption (wrong iv)",
			plainText:      "Hello, World!",
			key:           "0123456789abcdef0123456789abcdef",
			iv:            "wrongiv",
			expectedRegex: ``,
			expectError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encrypted := encrypt(test.plainText, test.key, test.iv)

			if test.expectError && encrypted != "" {
				t.Errorf("Expected an error but got a result: %s", encrypted)
			}

			if !test.expectError {
				// Check if the encrypted result matches the expected regex pattern
				if matched, _ := regexp.MatchString(test.expectedRegex, encrypted); !matched {
					t.Errorf("Expected a result matching pattern %s, Got: %s", test.expectedRegex, encrypted)
				}
			}
		})
	}
}

