package str

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "string",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "byte",
			input:    byte('A'),
			expected: "A",
		},
		{
			name:     "bytes slice",
			input:    []byte("hello"),
			expected: "hello",
		},
		{
			name:     "rune",
			input:    'B',
			expected: "B",
		},
		{
			name:     "runes slice",
			input:    []rune("hello"),
			expected: "hello",
		},
		{
			name:     "string slice",
			input:    []string{"hello", "world"},
			expected: "hello world",
		},
		{
			name:     "int",
			input:    42,
			expected: "42",
		},
		{
			name:     "int8",
			input:    int8(8),
			expected: "8",
		},
		{
			name:     "int16",
			input:    int16(16),
			expected: "16",
		},
		{
			name:     "int64",
			input:    int64(64),
			expected: "64",
		},
		{
			name:     "uint",
			input:    uint(42),
			expected: "42",
		},
		{
			name:     "uint16",
			input:    uint16(16),
			expected: "16",
		},
		{
			name:     "uint32",
			input:    uint32(32),
			expected: "32",
		},
		{
			name:     "uint64",
			input:    uint64(64),
			expected: "64",
		},
		{
			name:     "float32",
			input:    float32(3.14),
			expected: "3.14",
		},
		{
			name:     "float64",
			input:    3.14159,
			expected: "3.14159",
		},
		{
			name:     "bool true",
			input:    true,
			expected: "true",
		},
		{
			name:     "bool false",
			input:    false,
			expected: "false",
		},
		{
			name:     "nil",
			input:    nil,
			expected: "",
		},
		{
			name:     "pointer",
			input:    ptr("test"),
			expected: "test",
		},
		{
			name:     "nil pointer",
			input:    (*string)(nil),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.input); got != Str(tt.expected) {
				t.Errorf("New() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to create pointer
func ptr[T any](v T) *T {
	return &v
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "helloWorld",
		},
		{
			name:     "with special characters",
			input:    "hello_world-test",
			expected: "helloWorldTest",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "hello123World456",
		},
		{
			name:     "already camel case",
			input:    "helloWorld",
			expected: "helloWorld",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "helloWorld",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "helloWorldTest",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "userId",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToCamel(); got != Str(tt.expected) {
				t.Errorf("Str.ToCamel() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToGoCamel(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "helloWorld",
		},
		{
			name:     "with special characters",
			input:    "hello_world-test",
			expected: "helloWorldTest",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "hello123World456",
		},
		{
			name:     "already camel case",
			input:    "helloWorld",
			expected: "helloWorld",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "helloWorld",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "helloWorldTest",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "userID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToGoCamel(); got != Str(tt.expected) {
				t.Errorf("Str.ToGoCamel() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToGoPascal(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "HelloWorld",
		},
		{
			name:     "with special characters",
			input:    "hello_world-test",
			expected: "HelloWorldTest",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "Hello123World456",
		},
		{
			name:     "already pascal case",
			input:    "HelloWorld",
			expected: "HelloWorld",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "HelloWorld",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "HelloWorldTest",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "UserID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToGoPascal(); got != Str(tt.expected) {
				t.Errorf("Str.ToGoPascal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToPascal(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "HelloWorld",
		},
		{
			name:     "with special characters",
			input:    "hello_world-test",
			expected: "HelloWorldTest",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "Hello123World456",
		},
		{
			name:     "already pascal case",
			input:    "HelloWorld",
			expected: "HelloWorld",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "HelloWorld",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "HelloWorldTest",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "UserId",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToPascal(); got != Str(tt.expected) {
				t.Errorf("Str.ToPascal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToKebab(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "hello-world",
		},
		{
			name:     "with special characters",
			input:    "hello_world.test",
			expected: "hello-world-test",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "hello123-world456",
		},
		{
			name:     "camel case",
			input:    "helloWorld",
			expected: "hello-world",
		},
		{
			name:     "pascal case",
			input:    "HelloWorld",
			expected: "hello-world",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "hello-world",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "hello-world-test",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "user-id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToKebab(); got != Str(tt.expected) {
				t.Errorf("Str.ToKebab() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "hello_world",
		},
		{
			name:     "with special characters",
			input:    "hello-world.test",
			expected: "hello_world_test",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "hello123_world456",
		},
		{
			name:     "camel case",
			input:    "helloWorld",
			expected: "hello_world",
		},
		{
			name:     "pascal case",
			input:    "HelloWorld",
			expected: "hello_world",
		},
		{
			name:     "uppercase words",
			input:    "HELLO WORLD",
			expected: "hello_world",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "hello_world_test",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "user_id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToSnake(); got != Str(tt.expected) {
				t.Errorf("Str.ToSnake() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestToUpperSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected string
	}{
		{
			name:     "simple two words",
			input:    "hello world",
			expected: "HELLO_WORLD",
		},
		{
			name:     "with special characters",
			input:    "hello-world.test",
			expected: "HELLO_WORLD_TEST",
		},
		{
			name:     "with numbers",
			input:    "hello123_world456",
			expected: "HELLO123_WORLD456",
		},
		{
			name:     "camel case",
			input:    "helloWorld",
			expected: "HELLO_WORLD",
		},
		{
			name:     "pascal case",
			input:    "HelloWorld",
			expected: "HELLO_WORLD",
		},
		{
			name:     "already uppercase",
			input:    "HELLO WORLD",
			expected: "HELLO_WORLD",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "HELLO",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world   test",
			expected: "HELLO_WORLD_TEST",
		},
		{
			name:     "sequential uppercase",
			input:    "userID",
			expected: "USER_ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ToUpperSnake(); got != Str(tt.expected) {
				t.Errorf("Str.ToUpperSnake() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		start    int
		end      int
		expected string
	}{
		{
			name:     "normal slice",
			input:    "hello",
			start:    1,
			end:      4,
			expected: "ell",
		},
		{
			name:     "slice whole string",
			input:    "hello",
			start:    0,
			end:      5,
			expected: "hello",
		},
		{
			name:     "negative start index",
			input:    "hello",
			start:    -3,
			end:      5,
			expected: "llo",
		},
		{
			name:     "negative end index",
			input:    "hello",
			start:    1,
			end:      -1,
			expected: "ell",
		},
		{
			name:     "both negative indices",
			input:    "hello",
			start:    -4,
			end:      -1,
			expected: "ell",
		},
		{
			name:     "start equals end",
			input:    "hello",
			start:    2,
			end:      2,
			expected: "",
		},
		{
			name:     "start greater than end",
			input:    "hello",
			start:    3,
			end:      2,
			expected: "",
		},
		{
			name:     "start out of bounds",
			input:    "hello",
			start:    10,
			end:      12,
			expected: "",
		},
		{
			name:     "end out of bounds",
			input:    "hello",
			start:    0,
			end:      10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			start:    0,
			end:      0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Slice(tt.start, tt.end); got != Str(tt.expected) {
				t.Errorf("Str.Slice() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSliceFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		index    int
		expected string
	}{
		{
			name:     "slice from positive index",
			input:    "hello",
			index:    1,
			expected: "ello",
		},
		{
			name:     "slice from start",
			input:    "hello",
			index:    0,
			expected: "hello",
		},
		{
			name:     "slice from negative index",
			input:    "hello",
			index:    -3,
			expected: "llo",
		},
		{
			name:     "index equals length",
			input:    "hello",
			index:    5,
			expected: "",
		},
		{
			name:     "index out of bounds",
			input:    "hello",
			index:    10,
			expected: "",
		},
		{
			name:     "large negative index",
			input:    "hello",
			index:    -10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			index:    0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SliceFrom(tt.index); got != Str(tt.expected) {
				t.Errorf("Str.SliceFrom() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSliceRunes(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		start    int
		end      int
		expected string
	}{
		{
			name:     "ascii string slice",
			input:    "hello",
			start:    1,
			end:      4,
			expected: "ell",
		},
		{
			name:     "unicode string slice",
			input:    "你好世界",
			start:    1,
			end:      3,
			expected: "好世",
		},
		{
			name:     "slice whole string",
			input:    "hello世界",
			start:    0,
			end:      7,
			expected: "hello世界",
		},
		{
			name:     "negative start index",
			input:    "hello世界",
			start:    -3,
			end:      6,
			expected: "o世",
		},
		{
			name:     "negative end index",
			input:    "hello世界",
			start:    1,
			end:      -1,
			expected: "ello世",
		},
		{
			name:     "both negative indices",
			input:    "hello世界",
			start:    -4,
			end:      -1,
			expected: "lo世",
		},
		{
			name:     "start equals end",
			input:    "hello世界",
			start:    2,
			end:      2,
			expected: "",
		},
		{
			name:     "start greater than end",
			input:    "hello世界",
			start:    3,
			end:      2,
			expected: "",
		},
		{
			name:     "start out of bounds",
			input:    "hello世界",
			start:    10,
			end:      12,
			expected: "",
		},
		{
			name:     "end out of bounds",
			input:    "hello世界",
			start:    0,
			end:      10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			start:    0,
			end:      0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SliceRunes(tt.start, tt.end); got != Str(tt.expected) {
				t.Errorf("Str.SliceRunes() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSliceRunesFrom(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		index    int
		expected string
	}{
		{
			name:     "ascii string slice from positive index",
			input:    "hello",
			index:    1,
			expected: "ello",
		},
		{
			name:     "unicode string slice from positive index",
			input:    "你好世界",
			index:    1,
			expected: "好世界",
		},
		{
			name:     "slice from start",
			input:    "hello世界",
			index:    0,
			expected: "hello世界",
		},
		{
			name:     "slice from negative index",
			input:    "hello世界",
			index:    -3,
			expected: "o世界",
		},
		{
			name:     "index equals length",
			input:    "hello世界",
			index:    7,
			expected: "",
		},
		{
			name:     "index out of bounds",
			input:    "hello世界",
			index:    10,
			expected: "",
		},
		{
			name:     "large negative index",
			input:    "hello世界",
			index:    -10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			index:    0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SliceRunesFrom(tt.index); got != Str(tt.expected) {
				t.Errorf("Str.SliceRunesFrom() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSliceRunesTo(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		index    int
		expected string
	}{
		{
			name:     "ascii string slice to positive index",
			input:    "hello",
			index:    3,
			expected: "hel",
		},
		{
			name:     "unicode string slice to positive index",
			input:    "你好世界",
			index:    2,
			expected: "你好",
		},
		{
			name:     "slice to start",
			input:    "hello世界",
			index:    0,
			expected: "",
		},
		{
			name:     "slice to negative index",
			input:    "hello世界",
			index:    -1,
			expected: "hello世",
		},
		{
			name:     "index equals length",
			input:    "hello世界",
			index:    7,
			expected: "hello世界",
		},
		{
			name:     "index out of bounds",
			input:    "hello世界",
			index:    10,
			expected: "",
		},
		{
			name:     "large negative index",
			input:    "hello世界",
			index:    -10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			index:    0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SliceRunesTo(tt.index); got != Str(tt.expected) {
				t.Errorf("Str.SliceRunesTo() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSliceTo(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		index    int
		expected string
	}{
		{
			name:     "slice to positive index",
			input:    "hello",
			index:    3,
			expected: "hel",
		},
		{
			name:     "slice to start",
			input:    "hello",
			index:    0,
			expected: "",
		},
		{
			name:     "slice to negative index",
			input:    "hello",
			index:    -2,
			expected: "hel",
		},
		{
			name:     "index equals length",
			input:    "hello",
			index:    5,
			expected: "",
		},
		{
			name:     "index out of bounds",
			input:    "hello",
			index:    10,
			expected: "",
		},
		{
			name:     "large negative index",
			input:    "hello",
			index:    -10,
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			index:    0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SliceTo(tt.index); got != Str(tt.expected) {
				t.Errorf("Str.SliceTo() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		delim    string
		expected []string
	}{
		{
			name:     "split on space",
			input:    "hello world test",
			delim:    " ",
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "split on comma",
			input:    "a,b,c",
			delim:    ",",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with empty delimiter",
			input:    "hello",
			delim:    "",
			expected: []string{"h", "e", "l", "l", "o"},
		},
		{
			name:     "split empty string",
			input:    "",
			delim:    ",",
			expected: []string{""},
		},
		{
			name:     "split with delimiter not in string",
			input:    "hello",
			delim:    ",",
			expected: []string{"hello"},
		},
		{
			name:     "split with multi-character delimiter",
			input:    "hello::world::test",
			delim:    "::",
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "split with leading delimiter",
			input:    ",a,b,c",
			delim:    ",",
			expected: []string{"", "a", "b", "c"},
		},
		{
			name:     "split with trailing delimiter",
			input:    "a,b,c,",
			delim:    ",",
			expected: []string{"a", "b", "c", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Split(tt.delim)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Str.Split() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		delim    string
		count    int
		expected []string
	}{
		{
			name:     "split on space with count 2",
			input:    "hello world test",
			delim:    " ",
			count:    2,
			expected: []string{"hello", "world test"},
		},
		{
			name:     "split on comma with count 3",
			input:    "a,b,c,d",
			delim:    ",",
			count:    3,
			expected: []string{"a", "b", "c,d"},
		},
		{
			name:     "split with count larger than segments",
			input:    "a,b,c",
			delim:    ",",
			count:    5,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with empty delimiter",
			input:    "hello",
			delim:    "",
			count:    3,
			expected: []string{"h", "e", "llo"},
		},
		{
			name:     "split empty string",
			input:    "",
			delim:    ",",
			count:    2,
			expected: []string{""},
		},
		{
			name:     "split with negative count",
			input:    "a,b,c",
			delim:    ",",
			count:    -1,
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "split with count 1",
			input:    "a,b,c",
			delim:    ",",
			count:    1,
			expected: []string{"a,b,c"},
		},
		{
			name:     "split with delimiter not in string",
			input:    "hello",
			delim:    ",",
			count:    2,
			expected: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.SplitN(tt.delim, tt.count)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Str.SplitN() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSplitPattern(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		pattern  string
		expected []string
	}{
		{
			name:     "split on word boundaries",
			input:    "hello world test",
			pattern:  `\s+`,
			expected: []string{"hello", "world", "test"},
		},
		{
			name:     "split on numbers",
			input:    "abc123def456ghi",
			pattern:  `\d+`,
			expected: []string{"abc", "def", "ghi"},
		},
		{
			name:     "split with capturing groups",
			input:    "a=b,c=d",
			pattern:  `[,=]`,
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "split with invalid pattern",
			input:    "hello world",
			pattern:  "[",
			expected: []string{"hello world"},
		},
		{
			name:     "split empty string",
			input:    "",
			pattern:  `\s+`,
			expected: []string{""},
		},
		{
			name:     "split with pattern not in string",
			input:    "hello",
			pattern:  `\d`,
			expected: []string{"hello"},
		},
		{
			name:     "split with complex pattern",
			input:    "hello:world;test,example",
			pattern:  `[;,:]`,
			expected: []string{"hello", "world", "test", "example"},
		},
		{
			name:     "split with consecutive delimiters",
			input:    "a,,b,,c",
			pattern:  `,+`,
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.SplitPattern(tt.pattern)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Str.SplitPattern() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReplaceLast(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		old      string
		new      string
		expected string
	}{
		{
			name:     "replace last occurrence",
			input:    "hello hello world",
			old:      "hello",
			new:      "hi",
			expected: "hello hi world",
		},
		{
			name:     "replace when only one occurrence exists",
			input:    "hello world",
			old:      "world",
			new:      "earth",
			expected: "hello earth",
		},
		{
			name:     "replace with empty string",
			input:    "hello world",
			old:      "world",
			new:      "",
			expected: "hello ",
		},
		{
			name:     "no replacement when old string not found",
			input:    "hello world",
			old:      "xyz",
			new:      "abc",
			expected: "hello world",
		},
		{
			name:     "empty input string",
			input:    "",
			old:      "hello",
			new:      "hi",
			expected: "",
		},
		{
			name:     "empty old string",
			input:    "hello world",
			old:      "",
			new:      "test",
			expected: "hello world",
		},
		{
			name:     "replace with longer string",
			input:    "hello world",
			old:      "world",
			new:      "beautiful world",
			expected: "hello beautiful world",
		},
		{
			name:     "multiple character old string",
			input:    "hello hello hello",
			old:      "hello",
			new:      "hi",
			expected: "hello hello hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ReplaceLast(tt.old, tt.new); got != Str(tt.expected) {
				t.Errorf("Str.ReplaceLast() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReplacePattern(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		pattern  string
		new      string
		expected string
	}{
		{
			name:     "replace word pattern",
			input:    "hello world test",
			pattern:  `\w+`,
			new:      "foo",
			expected: "foo foo foo",
		},
		{
			name:     "replace digit pattern",
			input:    "abc123def456",
			pattern:  `\d+`,
			new:      "NUM",
			expected: "abcNUMdefNUM",
		},
		{
			name:     "replace with capturing groups",
			input:    "hello world",
			pattern:  `(\w+)\s+(\w+)`,
			new:      "$2 $1",
			expected: "world hello",
		},
		{
			name:     "invalid pattern returns original string",
			input:    "hello world",
			pattern:  "[",
			new:      "test",
			expected: "hello world",
		},
		{
			name:     "empty input string",
			input:    "",
			pattern:  `\w+`,
			new:      "test",
			expected: "",
		},
		{
			name:     "pattern not found",
			input:    "hello world",
			pattern:  `\d+`,
			new:      "num",
			expected: "hello world",
		},
		{
			name:     "replace with empty string",
			input:    "hello world",
			pattern:  `\s+`,
			new:      "",
			expected: "helloworld",
		},
		{
			name:     "multiple replacements",
			input:    "a=b,c=d",
			pattern:  `[,=]`,
			new:      ":",
			expected: "a:b:c:d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ReplacePattern(tt.pattern, tt.new); got != Str(tt.expected) {
				t.Errorf("Str.ReplacePattern() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPopEnd(t *testing.T) {
	tests := []struct {
		name          string
		input         Str
		expectedRune  rune
		expectedSlice string
	}{
		{
			name:          "basic ascii string",
			input:         "hello",
			expectedRune:  'o',
			expectedSlice: "hell",
		},
		{
			name:          "single character",
			input:         "a",
			expectedRune:  'a',
			expectedSlice: "",
		},
		{
			name:          "empty string",
			input:         "",
			expectedRune:  -1,
			expectedSlice: "",
		},
		{
			name:          "unicode string",
			input:         "你好世界",
			expectedRune:  '界',
			expectedSlice: "你好世",
		},
		{
			name:          "string with numbers",
			input:         "abc123",
			expectedRune:  '3',
			expectedSlice: "abc12",
		},
		{
			name:          "string with spaces",
			input:         "hello ",
			expectedRune:  ' ',
			expectedSlice: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRune, gotSlice := tt.input.PopEnd()
			if gotRune != tt.expectedRune {
				t.Errorf("Str.PopEnd() rune = %v, want %v", gotRune, tt.expectedRune)
			}
			if gotSlice != Str(tt.expectedSlice) {
				t.Errorf("Str.PopEnd() slice = %v, want %v", gotSlice, tt.expectedSlice)
			}
		})
	}
}

func TestPopStart(t *testing.T) {
	tests := []struct {
		name          string
		input         Str
		expectedRune  rune
		expectedSlice string
	}{
		{
			name:          "basic ascii string",
			input:         "hello",
			expectedRune:  'h',
			expectedSlice: "ello",
		},
		{
			name:          "single character",
			input:         "a",
			expectedRune:  'a',
			expectedSlice: "",
		},
		{
			name:          "empty string",
			input:         "",
			expectedRune:  -1,
			expectedSlice: "",
		},
		{
			name:          "unicode string",
			input:         "你好世界",
			expectedRune:  '你',
			expectedSlice: "好世界",
		},
		{
			name:          "string with numbers",
			input:         "123abc",
			expectedRune:  '1',
			expectedSlice: "23abc",
		},
		{
			name:          "string with spaces",
			input:         " hello",
			expectedRune:  ' ',
			expectedSlice: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRune, gotSlice := tt.input.PopStart()
			if gotRune != tt.expectedRune {
				t.Errorf("Str.PopStart() rune = %v, want %v", gotRune, tt.expectedRune)
			}
			if gotSlice != Str(tt.expectedSlice) {
				t.Errorf("Str.PopStart() slice = %v, want %v", gotSlice, tt.expectedSlice)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		want    uint64
		wantErr bool
	}{
		{
			name:    "valid uint",
			input:   "42",
			want:    42,
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "large number",
			input:   "18446744073709551615", // max uint64
			want:    18446744073709551615,
			wantErr: false,
		},
		{
			name:    "negative number",
			input:   "-42",
			want:    0,
			wantErr: true,
		},
		{
			name:    "decimal number",
			input:   "42.5",
			want:    0,
			wantErr: true,
		},
		{
			name:    "non-numeric string",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "overflow",
			input:   "184416744073709551616", // max uint64 + 1
			want:    18446744073709551615,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.ParseUint()
			if (err != nil) != tt.wantErr {
				t.Errorf("Str.ParseUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Str.ParseUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		want    int64
		wantErr bool
	}{
		{
			name:    "positive number",
			input:   "42",
			want:    42,
			wantErr: false,
		},
		{
			name:    "negative number",
			input:   "-42",
			want:    -42,
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "max int64",
			input:   "9223372036854775807",
			want:    9223372036854775807,
			wantErr: false,
		},
		{
			name:    "min int64",
			input:   "-9223372036854775808",
			want:    -9223372036854775808,
			wantErr: false,
		},
		{
			name:    "overflow positive",
			input:   "9223372036854775808", // max int64 + 1
			want:    9223372036854775807,
			wantErr: true,
		},
		{
			name:    "overflow negative",
			input:   "-9223372036854775809", // min int64 - 1
			want:    -9223372036854775808,
			wantErr: true,
		},
		{
			name:    "decimal number",
			input:   "42.5",
			want:    0,
			wantErr: true,
		},
		{
			name:    "non-numeric string",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.ParseInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("Str.ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Str.ParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		want    float64
		wantErr bool
	}{
		{
			name:    "positive integer",
			input:   "42",
			want:    42.0,
			wantErr: false,
		},
		{
			name:    "negative integer",
			input:   "-42",
			want:    -42.0,
			wantErr: false,
		},
		{
			name:    "positive decimal",
			input:   "3.14",
			want:    3.14,
			wantErr: false,
		},
		{
			name:    "negative decimal",
			input:   "-3.14",
			want:    -3.14,
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "scientific notation",
			input:   "1.23e-4",
			want:    0.000123,
			wantErr: false,
		},
		{
			name:    "invalid number",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "multiple decimal points",
			input:   "1.2.3",
			want:    0,
			wantErr: true,
		},
		{
			name:    "max float64",
			input:   "1.7976931348623157e+308",
			want:    1.7976931348623157e+308,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.ParseFloat()
			if (err != nil) != tt.wantErr {
				t.Errorf("Str.ParseFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Str.ParseFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadStart(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		length   int
		pad      string
		expected string
	}{
		{
			name:     "pad with single character",
			input:    "hello",
			length:   8,
			pad:      "*",
			expected: "***hello",
		},
		{
			name:     "pad with multiple characters",
			input:    "world",
			length:   10,
			pad:      "ab",
			expected: "ababaworld",
		},
		{
			name:     "length less than string length",
			input:    "hello",
			length:   3,
			pad:      "*",
			expected: "hello",
		},
		{
			name:     "empty string input",
			input:    "",
			length:   3,
			pad:      "*",
			expected: "***",
		},
		{
			name:     "empty padding string",
			input:    "test",
			length:   8,
			pad:      "",
			expected: "test",
		},
		{
			name:     "length equals string length",
			input:    "test",
			length:   4,
			pad:      "*",
			expected: "test",
		},
		{
			name:     "unicode padding",
			input:    "hello",
			length:   8,
			pad:      "世界",
			expected: "世界世hello",
		},
		{
			name:     "length requiring partial pad",
			input:    "test",
			length:   7,
			pad:      "ab",
			expected: "abatest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.PadStart(tt.length, tt.pad); got != Str(tt.expected) {
				t.Errorf("Str.PadStart() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPadEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		length   int
		pad      string
		expected string
	}{
		{
			name:     "pad with single character",
			input:    "hello",
			length:   8,
			pad:      "*",
			expected: "hello***",
		},
		{
			name:     "pad with multiple characters",
			input:    "world",
			length:   10,
			pad:      "ab",
			expected: "worldababa",
		},
		{
			name:     "length less than string length",
			input:    "hello",
			length:   3,
			pad:      "*",
			expected: "hello",
		},
		{
			name:     "empty string input",
			input:    "",
			length:   3,
			pad:      "*",
			expected: "***",
		},
		{
			name:     "empty padding string",
			input:    "test",
			length:   8,
			pad:      "",
			expected: "test",
		},
		{
			name:     "length equals string length",
			input:    "test",
			length:   4,
			pad:      "*",
			expected: "test",
		},
		{
			name:     "unicode padding",
			input:    "hello",
			length:   8,
			pad:      "世界",
			expected: "hello世界世",
		},
		{
			name:     "length requiring partial pad",
			input:    "test",
			length:   7,
			pad:      "ab",
			expected: "testaba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.PadEnd(tt.length, tt.pad); got != Str(tt.expected) {
				t.Errorf("Str.PadEnd() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		pattern string
		want    bool
	}{
		{
			name:    "simple pattern match",
			input:   "hello",
			pattern: "^hello$",
			want:    true,
		},
		{
			name:    "pattern with wildcard",
			input:   "hello world",
			pattern: "hello.*",
			want:    true,
		},
		{
			name:    "no match",
			input:   "hello",
			pattern: "world",
			want:    false,
		},
		{
			name:    "pattern with character class",
			input:   "abc123",
			pattern: `\w+\d+`,
			want:    true,
		},
		{
			name:    "case sensitive match",
			input:   "Hello",
			pattern: "hello",
			want:    false,
		},
		{
			name:    "empty string match",
			input:   "",
			pattern: "^$",
			want:    true,
		},
		{
			name:    "invalid pattern",
			input:   "test",
			pattern: "[",
			want:    false,
		},
		{
			name:    "pattern with groups",
			input:   "abc123def",
			pattern: `(\w+)(\d+)(\w+)`,
			want:    true,
		},
		{
			name:    "pattern with alternation",
			input:   "cat",
			pattern: "cat|dog",
			want:    true,
		},
		{
			name:    "unicode string match",
			input:   "你好世界",
			pattern: "^你好.*",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Match(tt.pattern); got != tt.want {
				t.Errorf("Str.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		fn       func(rune) rune
		expected string
	}{
		{
			name:  "to upper case",
			input: "hello",
			fn: func(r rune) rune {
				return []rune(strings.ToUpper(string(r)))[0]
			},
			expected: "HELLO",
		},
		{
			name:  "remove vowels",
			input: "hello world",
			fn: func(r rune) rune {
				if strings.ContainsRune("aeiou", r) {
					return -1
				}
				return r
			},
			expected: "hll wrld",
		},
		{
			name:  "replace spaces",
			input: "hello world",
			fn: func(r rune) rune {
				if r == ' ' {
					return '-'
				}
				return r
			},
			expected: "hello-world",
		},
		{
			name:  "empty string",
			input: "",
			fn: func(r rune) rune {
				return r
			},
			expected: "",
		},
		{
			name:  "unicode string",
			input: "你好世界",
			fn: func(r rune) rune {
				if r == '好' {
					return '很'
				}
				return r
			},
			expected: "你很世界",
		},
		{
			name:  "remove all characters",
			input: "hello",
			fn: func(r rune) rune {
				return -1
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.Map(tt.fn); got != Str(tt.expected) {
				t.Errorf("Str.Map() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLastIndexAny(t *testing.T) {
	tests := []struct {
		name  string
		input Str
		chars string
		want  int
	}{
		{
			name:  "single character match at end",
			input: "hello",
			chars: "o",
			want:  4,
		},
		{
			name:  "single character match in middle",
			input: "hello",
			chars: "l",
			want:  3,
		},
		{
			name:  "multiple characters, last match",
			input: "hello world",
			chars: "od",
			want:  10,
		},
		{
			name:  "no match",
			input: "hello",
			chars: "xyz",
			want:  -1,
		},
		{
			name:  "empty input string",
			input: "",
			chars: "abc",
			want:  -1,
		},
		{
			name:  "empty chars string",
			input: "hello",
			chars: "",
			want:  -1,
		},
		{
			name:  "unicode characters",
			input: "hello世界世",
			chars: "世",
			want:  11,
		},
		{
			name:  "multiple matches same character",
			input: "hello",
			chars: "l",
			want:  3,
		},
		{
			name:  "special characters",
			input: "hello!@#$",
			chars: "@#$",
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.LastIndexAny(tt.chars); got != tt.want {
				t.Errorf("Str.LastIndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		name   string
		input  Str
		search string
		want   int
	}{
		{
			name:   "match at end",
			input:  "hello world",
			search: "world",
			want:   6,
		},
		{
			name:   "match in middle",
			input:  "hello hello world",
			search: "hello",
			want:   6,
		},
		{
			name:   "no match",
			input:  "hello world",
			search: "goodbye",
			want:   -1,
		},
		{
			name:   "empty search string",
			input:  "hello",
			search: "",
			want:   5,
		},
		{
			name:   "empty input string",
			input:  "",
			search: "hello",
			want:   -1,
		},
		{
			name:   "unicode string",
			input:  "你好世界世界",
			search: "世界",
			want:   12,
		},
		{
			name:   "substring match",
			input:  "hello hello hello",
			search: "hell",
			want:   12,
		},
		{
			name:   "case sensitive",
			input:  "hello HELLO",
			search: "hello",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.LastIndex(tt.search); got != tt.want {
				t.Errorf("Str.LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		pattern string
		want    []int
	}{
		{
			name:    "simple pattern match",
			input:   "hello world",
			pattern: "world",
			want:    []int{6, 11},
		},
		{
			name:    "pattern with wildcard",
			input:   "hello world",
			pattern: "h.*o",
			want:    []int{0, 8},
		},
		{
			name:    "no match",
			input:   "hello world",
			pattern: "xyz",
			want:    nil,
		},
		{
			name:    "pattern with character class",
			input:   "abc123def",
			pattern: `\d+`,
			want:    []int{3, 6},
		},
		{
			name:    "empty string",
			input:   "",
			pattern: "test",
			want:    nil,
		},
		{
			name:    "invalid pattern",
			input:   "test",
			pattern: "[",
			want:    nil,
		},
		{
			name:    "pattern with groups",
			input:   "test123test",
			pattern: `(\d+)`,
			want:    []int{4, 7},
		},
		{
			name:    "unicode string match",
			input:   "hello 世界",
			pattern: "世界",
			want:    []int{6, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.FindIndex(tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str.FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllIndex(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		pattern string
		want    [][]int
	}{
		{
			name:    "simple pattern multiple matches",
			input:   "hello world hello world",
			pattern: "hello",
			want:    [][]int{{0, 5}, {12, 17}},
		},
		{
			name:    "pattern with wildcard",
			input:   "test123test456test",
			pattern: `\d+`,
			want:    [][]int{{4, 7}, {11, 14}},
		},
		{
			name:    "no matches",
			input:   "hello world",
			pattern: "xyz",
			want:    nil,
		},
		{
			name:    "pattern with character class",
			input:   "a1b2c3d4",
			pattern: `[0-9]`,
			want:    [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			name:    "empty string",
			input:   "",
			pattern: "test",
			want:    nil,
		},
		{
			name:    "invalid pattern",
			input:   "test",
			pattern: "[",
			want:    nil,
		},
		{
			name:    "pattern with groups",
			input:   "test123test456",
			pattern: `(\d+)`,
			want:    [][]int{{4, 7}, {11, 14}},
		},
		{
			name:    "unicode string multiple matches",
			input:   "你好世界你好世界",
			pattern: "你好",
			want:    [][]int{{0, 6}, {12, 18}},
		},
		{
			name:    "overlapping matches not allowed",
			input:   "aaa",
			pattern: "aa",
			want:    [][]int{{0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.FindAllIndex(tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str.FindAllIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		pattern string
		want    []string
	}{
		{
			name:    "simple pattern multiple matches",
			input:   "hello world hello world",
			pattern: "hello",
			want:    []string{"hello", "hello"},
		},
		{
			name:    "pattern with wildcard",
			input:   "test123test456test",
			pattern: `\d+`,
			want:    []string{"123", "456"},
		},
		{
			name:    "no matches",
			input:   "hello world",
			pattern: "xyz",
			want:    nil,
		},
		{
			name:    "pattern with character class",
			input:   "a1b2c3d4",
			pattern: `[0-9]`,
			want:    []string{"1", "2", "3", "4"},
		},
		{
			name:    "empty string",
			input:   "",
			pattern: "test",
			want:    nil,
		},
		{
			name:    "invalid pattern",
			input:   "test",
			pattern: "[",
			want:    nil,
		},
		{
			name:    "pattern with groups",
			input:   "test123test456",
			pattern: `(\d+)`,
			want:    []string{"123", "456"},
		},
		{
			name:    "unicode string multiple matches",
			input:   "你好世界你好世界",
			pattern: "你好",
			want:    []string{"你好", "你好"},
		},
		{
			name:    "overlapping matches not allowed",
			input:   "aaa",
			pattern: "aa",
			want:    []string{"aa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.FindAll(tt.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Str.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		input   Str
		pattern string
		want    string
	}{
		{
			name:    "simple pattern match",
			input:   "hello world",
			pattern: "world",
			want:    "world",
		},
		{
			name:    "pattern with wildcard",
			input:   "hello world",
			pattern: "h.*o",
			want:    "hello wo",
		},
		{
			name:    "no match",
			input:   "hello world",
			pattern: "xyz",
			want:    "",
		},
		{
			name:    "pattern with character class",
			input:   "abc123def",
			pattern: `\d+`,
			want:    "123",
		},
		{
			name:    "empty string",
			input:   "",
			pattern: "test",
			want:    "",
		},
		{
			name:    "invalid pattern",
			input:   "test",
			pattern: "[",
			want:    "",
		},
		{
			name:    "pattern with groups",
			input:   "test123test",
			pattern: `(\d+)`,
			want:    "123",
		},
		{
			name:    "unicode string match",
			input:   "hello 世界",
			pattern: "世界",
			want:    "世界",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Find(tt.pattern)
			if got != tt.want {
				t.Errorf("Str.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name   string
		input  Str
		chars  string
		wanted bool
	}{
		{
			name:   "single character match",
			input:  "hello",
			chars:  "h",
			wanted: true,
		},
		{
			name:   "multiple characters - one matches",
			input:  "hello",
			chars:  "xyz h",
			wanted: true,
		},
		{
			name:   "no match",
			input:  "hello",
			chars:  "xyz",
			wanted: false,
		},
		{
			name:   "empty input string",
			input:  "",
			chars:  "xyz",
			wanted: false,
		},
		{
			name:   "empty chars string",
			input:  "hello",
			chars:  "",
			wanted: false,
		},
		{
			name:   "unicode string match",
			input:  "hello世界",
			chars:  "界",
			wanted: true,
		},
		{
			name:   "special characters match",
			input:  "hello!@#",
			chars:  "@#$",
			wanted: true,
		},
		{
			name:   "space match",
			input:  "hello world",
			chars:  " ",
			wanted: true,
		},
		{
			name:   "case sensitive",
			input:  "hello",
			chars:  "H",
			wanted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ContainsAny(tt.chars); got != tt.wanted {
				t.Errorf("Str.ContainsAny() = %v, want %v", got, tt.wanted)
			}
		})
	}
}

func TestRuneCount(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "ascii string",
			input:    "hello",
			expected: 5,
		},
		{
			name:     "string with spaces",
			input:    "hello world",
			expected: 11,
		},
		{
			name:     "unicode string",
			input:    "你好世界",
			expected: 4,
		},
		{
			name:     "mixed ascii and unicode",
			input:    "hello世界",
			expected: 7,
		},
		{
			name:     "special characters",
			input:    "!@#$%^&*()",
			expected: 10,
		},
		{
			name:     "numbers",
			input:    "12345",
			expected: 5,
		},
		{
			name:     "emojis",
			input:    "👋🌍",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.RuneCount(); got != tt.expected {
				t.Errorf("Str.RuneCount() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRuneIndex(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		r        rune
		expected int
	}{
		{
			name:     "ascii character match",
			input:    "hello",
			r:        'h',
			expected: 0,
		},
		{
			name:     "ascii character match in middle",
			input:    "hello",
			r:        'l',
			expected: 2,
		},
		{
			name:     "ascii character match at end",
			input:    "hello",
			r:        'o',
			expected: 4,
		},
		{
			name:     "ascii character no match",
			input:    "hello",
			r:        'x',
			expected: -1,
		},
		{
			name:     "empty string",
			input:    "",
			r:        'a',
			expected: -1,
		},
		{
			name:     "unicode character match",
			input:    "你好世界",
			r:        '好',
			expected: 3,
		},
		{
			name:     "unicode character no match",
			input:    "你好世界",
			r:        '嗨',
			expected: -1,
		},
		{
			name:     "mixed ascii and unicode match ascii",
			input:    "hello世界",
			r:        'o',
			expected: 4,
		},
		{
			name:     "mixed ascii and unicode match unicode",
			input:    "hello世界",
			r:        '界',
			expected: 8,
		},
		{
			name:     "space character match",
			input:    "hello world",
			r:        ' ',
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.RuneIndex(tt.r); got != tt.expected {
				t.Errorf("Str.RuneIndex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReplaceEach(t *testing.T) {
	tests := []struct {
		name     string
		input    Str
		old      string
		new      string
		expected string
	}{
		{
			name:     "simple character replacements",
			input:    "hello",
			old:      "hlo",
			new:      "HLO",
			expected: "HeLLO",
		},
		{
			name:     "multiple occurrences",
			input:    "test test",
			old:      "te",
			new:      "TE",
			expected: "TEsT TEsT",
		},
		{
			name:     "no matches",
			input:    "hello",
			old:      "xyz",
			new:      "ABC",
			expected: "hello",
		},
		{
			name:     "empty input string",
			input:    "",
			old:      "abc",
			new:      "123",
			expected: "",
		},
		{
			name:     "different length strings",
			input:    "hello",
			old:      "abc",
			new:      "abcd",
			expected: "hello",
		},
		{
			name:     "empty old and new strings",
			input:    "hello",
			old:      "",
			new:      "",
			expected: "hello",
		},
		{
			name:     "unicode character replacements",
			input:    "你好世界",
			old:      "好世",
			new:      "很可",
			expected: "你很可界",
		},
		{
			name:     "special characters",
			input:    "!@#$%",
			old:      "!@#",
			new:      "123",
			expected: "123$%",
		},
		{
			name:     "overlapping replacements",
			input:    "abba",
			old:      "ab",
			new:      "ba",
			expected: "aaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.ReplaceEach(tt.old, tt.new); got != Str(tt.expected) {
				t.Errorf("Str.ReplaceEach() = %v, want %v", got, tt.expected)
			}
		})
	}
}
