package gouse

import "testing"

func TestIsMatchReg(t *testing.T) {
	var arr = []struct {
		regex    string
		input    string
		expected bool
	}{
		{
			regex:    "^[a-zA-Z0-9]{3,20}$",
			input:    "zoomer",
			expected: true,
		},
		{
			regex:    "^[a-zA-Z0-9]{3,20}$",
			input:    "zoomer123",
			expected: true,
		},
	}

	for _, item := range arr {
		actual := IsMatchReg(item.regex, item.input)
		if actual != item.expected {
			t.Errorf("IsMatch(%q, %q): expected %t, actual %t", item.regex, item.input, item.expected, actual)
		}
	}
}

func BenchmarkIsMatchReg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMatchReg("^[a-zA-Z0-9]{3,20}$", "zoomer")
	}
}

func TestMatchReg(t *testing.T) {
	arr := []struct {
		regex    string
		input    string
		expected []string
	}{
		{
			regex:    "[A-Z]",
			input:    "Hello World 123",
			expected: []string{"H", "W"},
		},
	}

	for _, item := range arr {
		actual := MatchReg(item.regex, item.input)
		if len(actual) != len(item.expected) {
			t.Errorf("Match(%q, %q): expected %q, actual %q", item.regex, item.input, item.expected, actual)
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MatchReg("[A-Z]", "Hello World 123")
	}
}
