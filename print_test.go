package gouse

import (
	"bytes"
	"os"
	"testing"
)

func TestSprint(t *testing.T) {
	tests := []struct {
		args     []any
		expected string
	}{
		{
			args:     []any{"Hello", "world"},
			expected: "Hello world",
		},
		{
			args:     []any{1, 2, 3},
			expected: "1 2 3",
		},
		{
			args:     []any{"Go", 3.14, true},
			expected: "Go 3.14 true",
		},
		{
			args:     []any{},
			expected: "",
		},
		{
			args:     []any{nil},
			expected: "nil",
		},
	}

	for _, test := range tests {
		t.Run("Sprint Test", func(t *testing.T) {
			actual := Sprint(test.args...)
			if actual != test.expected {
				t.Errorf("Sprint(%v) = %v; expected %v", test.args, actual, test.expected)
			}
		})
	}
}

func TestSprintln(t *testing.T) {
	tests := []struct {
		args     []any
		expected string
	}{
		{
			args:     []any{"Hello", "world"},
			expected: "Hello world\n",
		},
		{
			args:     []any{1, 2, 3},
			expected: "1 2 3\n",
		},
		{
			args:     []any{"Go", 3.14, true},
			expected: "Go 3.14 true\n",
		},
		{
			args:     []any{},
			expected: "\n",
		},
		{
			args:     []any{nil},
			expected: "nil\n",
		},
	}

	for _, test := range tests {
		t.Run("Sprintln Test", func(t *testing.T) {
			actual := Sprintln(test.args...)
			if actual != test.expected {
				t.Errorf("Sprintln(%v) = %v; expected %v", test.args, actual, test.expected)
			}
		})
	}
}

func TestSprintf(t *testing.T) {
	tests := []struct {
		format   string
		args     []any
		expected string
	}{
		{
			format:   "Hello %s!",
			args:     []any{"world"},
			expected: "Hello world!",
		},
		{
			format:   "Value: %d",
			args:     []any{123},
			expected: "Value: 123",
		},
		{
			format:   "Pi is approximately %f",
			args:     []any{3.14159},
			expected: "Pi is approximately 3.14",
		},
		{
			format:   "Boolean: %t",
			args:     []any{true},
			expected: "Boolean: true",
		},
		{
			format:   "Character: %c",
			args:     []any{65}, // ASCII value of 'A'
			expected: "Character: A",
		},
		{
			format:   "Invalid: %s %d",
			args:     []any{true, "hello"},
			expected: "Invalid: invalid invalid",
		},
		{
			format:   "Mix: %s, %d, %f",
			args:     []any{"test", 123, 3.14159},
			expected: "Mix: test, 123, 3.14",
		},
		{
			format:   "No argument: %d",
			args:     []any{},
			expected: "No argument: invalid",
		},
		{
			format:   "Percent sign: %% should appear",
			args:     []any{},
			expected: "Percent sign: % should appear",
		},
		{
			format:   "%T",
			args:     []any{true},
			expected: "true",
		},
		{
			format:   "%T",
			args:     []any{false},
			expected: "false",
		},
		{
			format:   "Invalid type: %f",
			args:     []any{"not a float"},
			expected: "Invalid type: invalid",
		},
	}

	for _, test := range tests {
		t.Run("Sprintf Test", func(t *testing.T) {
			actual := Sprintf(test.format, test.args...)
			if actual != test.expected {
				t.Errorf("Sprintf(%q, %v) = %v; expected %v", test.format, test.args, actual, test.expected)
			}
		})
	}
}

func captureOutput(f func()) string {
	// Create a pipe to capture output
	r, w, _ := os.Pipe()

	// Save the original stdout (so we can restore it later)
	stdout := os.Stdout
	// Redirect os.Stdout to the pipe's write end
	os.Stdout = w

	// Call the function that writes to stdout
	f()

	// Close the write end of the pipe
	w.Close()

	// Read the captured output from the read end of the pipe
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	// Restore the original stdout
	os.Stdout = stdout

	// Return the captured output as a string
	return buf.String()
}

func TestPrintln(t *testing.T) {
	tests := []struct {
		args     []any
		expected string
	}{
		{
			args:     []any{"Hello", "World"},
			expected: "Hello World\n",
		},
		{
			args:     []any{"SingleLine"},
			expected: "SingleLine\n",
		},
		{
			args:     []any{123, "is a number"},
			expected: "123 is a number\n",
		},
	}

	for _, test := range tests {
		t.Run("Println Test", func(t *testing.T) {
			actual := captureOutput(func() {
				Println(test.args...)
			})
			if actual != test.expected {
				t.Errorf("Println(%v) = %v; expected %v", test.args, actual, test.expected)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		args     []any
		expected string
	}{
		{
			args:     []any{"Hello", "World"},
			expected: "Hello World",
		},
		{
			args:     []any{"SingleLine"},
			expected: "SingleLine",
		},
		{
			args:     []any{123, "is a number"},
			expected: "123 is a number",
		},
	}

	for _, test := range tests {
		t.Run("Print Test", func(t *testing.T) {
			actual := captureOutput(func() {
				Print(test.args...)
			})
			if actual != test.expected {
				t.Errorf("Print(%v) = %v; expected %v", test.args, actual, test.expected)
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	tests := []struct {
		format   string
		args     []any
		expected string
	}{
		{
			format:   "Hello %s!",
			args:     []any{"World"},
			expected: "Hello World!",
		},
		{
			format:   "The number is: %d",
			args:     []any{123},
			expected: "The number is: 123",
		},
		{
			format:   "Pi: %f",
			args:     []any{3.14159},
			expected: "Pi: 3.14",
		},
		{
			format:   "Boolean: %t",
			args:     []any{true},
			expected: "Boolean: true",
		},
		{
			format:   "Character: %c",
			args:     []any{65},
			expected: "Character: A",
		},
		{
			format:   "No argument: %d",
			args:     []any{},
			expected: "No argument: invalid",
		},
	}

	for _, test := range tests {
		t.Run("Printf Test", func(t *testing.T) {
			actual := captureOutput(func() {
				Printf(test.format, test.args...)
			})
			if actual != test.expected {
				t.Errorf("Printf(%q, %v) = %v; expected %v", test.format, test.args, actual, test.expected)
			}
		})
	}
}
