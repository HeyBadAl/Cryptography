// decrypt_test.go
package main

import (
	"testing"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name          string
		cipherText    string
		key           string
		iv            string
		expected      string
		expectError   bool
	}{
		{
			name:       "Valid decryption",
			cipherText: "5f4b292cfe9a16bcb5837a60622c8d3a",
			key:        "0123456789abcdef0123456789abcdef",
			iv:         "1234567812345678",
			expected:   "Hello, World!",
			expectError: false,
		},
		{
			name:       "Invalid decryption (wrong key)",
			cipherText: "5f4b292cfe9a16bcb5837a60622c8d3a",
			key:        "wrongkey",
			iv:         "1234567812345678",
			expected:   "",
			expectError: true,
		},
		{
			name:       "Invalid decryption (wrong iv)",
			cipherText: "5f4b292cfe9a16bcb5837a60622c8d3a",
			key:        "0123456789abcdef0123456789abcdef",
			iv:         "wrongiv",
			expected:   "",
			expectError: true,
		},
		{
			name:       "Invalid decryption (corrupted ciphertext)",
			cipherText: "invalidciphertext",
			key:        "0123456789abcdef0123456789abcdef",
			iv:         "1234567812345678",
			expected:   "",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			decrypted := decrypt(test.cipherText, test.key, test.iv)

			if test.expectError && decrypted != "" {
				t.Errorf("Expected an error but got a result: %s", decrypted)
			}

			if !test.expectError && decrypted != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, decrypted)
			}
		})
	}
}

