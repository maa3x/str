package str

import (
	"testing"
)

func TestSplitWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "snake_case",
			input:    "hello_world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "kebab-case",
			input:    "hello-world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "camelCase",
			input:    "helloWorld",
			expected: []string{"hello", "World"},
		},
		{
			name:     "PascalCase",
			input:    "HelloWorld",
			expected: []string{"Hello", "World"},
		},
		{
			name:     "with spaces",
			input:    "HELLO to the WORLD",
			expected: []string{"HELLO", "to", "the", "WORLD"},
		},
		{
			name:     "upper snake",
			input:    "HELLO_WORLD",
			expected: []string{"HELLO", "WORLD"},
		},
		{
			name:     "mixed case",
			input:    "hello_worldAndUniverse",
			expected: []string{"hello", "world", "And", "Universe"},
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: []string{"user", "ID"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitWords(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("splitWords(%q) = %v, want %v", tt.input, got, tt.expected)
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("splitWords(%q)[%d] = %v, want %v", tt.input, i, got[i], tt.expected[i])
				}
			}
		})
	}
}
