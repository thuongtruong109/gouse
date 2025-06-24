package gouse

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test capitalize", "Test Capitalize"},
		{"test capitalize", "Test Capitalize"},
		{"test capitalize", "Test Capitalize"},
		{"capitalize123", "Capitalize123"},
	}

	for _, test := range tests {
		actual := Capitalize(test.input)
		if actual != test.expected {
			t.Errorf("Capitalize(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test snake", "test_snake"},
		{"test-snake", "test_snake"},
		{"testSnake", "test_snake"},
		{"snake123", "snake_123"},
	}

	for _, test := range tests {
		actual := SnakeCase(test.input)
		if actual != test.expected {
			t.Errorf("SnakeCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test kebab", "test-kebab"},
		{"test_kebab", "test-kebab"},
		{"testKebab", "test-kebab"},
		{"kebab123", "kebab-123"},
	}

	for _, test := range tests {
		actual := KebabCase(test.input)
		if actual != test.expected {
			t.Errorf("KebabCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test camel", "testCamel"},
		{"test-camel", "testCamel"},
		{"test_camel", "testCamel"},
		{"camel123", "camel123"},
	}

	for _, test := range tests {
		actual := CamelCase(test.input)
		if actual != test.expected {
			t.Errorf("CamelCase(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', true},
		{'A', true},
		{'z', true},
		{'Z', true},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsLetter(test.input)
		if actual != test.expected {
			t.Errorf("IsLetter(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'1', true},
		{'9', true},
		{'a', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsDigit(test.input)
		if actual != test.expected {
			t.Errorf("IsDigit(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", true},
		{"test", "test", true},
		{"test", "st", true},
		{"test", "t", true},
		{"test", "e", true},
		{"test", "s", true},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := Includes(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("Includes(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", false},
		{"test", "test", true},
		{"test", "st", false},
		{"test", "t", true},
		{"test", "e", false},
		{"test", "s", false},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := StartsWith(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("StartsWith(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"test", "es", false},
		{"test", "test", true},
		{"test", "st", true},
		{"test", "t", true},
		{"test", "e", false},
		{"test", "s", false},
		{"test", "z", false},
		{"test", "tt", false},
		{"test", "ttt", false},
		{"test", "tttt", false},
		{"test", "ttttt", false},
		{"test", "tttttt", false},
	}

	for _, test := range tests {
		actual := EndsWith(test.input, test.substr)
		if actual != test.expected {
			t.Errorf("EndsWith(%q, %q): expected %t, actual %t", test.input, test.substr, test.expected, actual)
		}
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', true},
		{'A', false},
		{'z', true},
		{'Z', false},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsLower(test.input)
		if actual != test.expected {
			t.Errorf("IsLower(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestIsUpper(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'a', false},
		{'A', true},
		{'z', false},
		{'Z', true},
		{'1', false},
		{'!', false},
		{' ', false},
	}

	for _, test := range tests {
		actual := IsUpper(test.input)
		if actual != test.expected {
			t.Errorf("IsUpper(%q): expected %t, actual %t", test.input, test.expected, actual)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"test", "e", []string{"t", "st"}},
		{"test", "t", []string{"", "es", ""}},
		{"test", "s", []string{"te", "t"}},
		{"test", "x", []string{"test"}},
		{"test", "", []string{"t", "e", "s", "t"}},
	}

	for _, test := range tests {
		actual := Split(test.input, test.sep)
		if len(actual) != len(test.expected) {
			t.Errorf("Split(%q, %q): length expected %d, length actual %d", test.input, test.sep, len(test.expected), len(actual))
		} else {
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("Split(%q, %q): expected %q, actual %q", test.input, test.sep, test.expected, actual)
				}
			}
		}
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"test", []string{"t", "e", "s", "t"}},
		{"test test", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test\ntest", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test\ttest", []string{"t", "e", "s", "t", "t", "e", "s", "t"}},
		{"test test test", []string{"t", "e", "s", "t", "t", "e", "s", "t", "t", "e", "s", "t"}},
	}

	for _, test := range tests {
		actual := Words(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("Words(%q): length expected %d, length actual %d", test.input, len(test.expected), len(actual))
		} else {
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("Words(%q): expected %q, actual %q", test.input, test.expected, actual)
				}
			}
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "tset"},
		{"test test", "tset tset"},
		{"test\ntest", "tset\ntset"},
		{"test\ttest", "tset\ttset"},
		{"test test test", "tset tset tset"},
	}

	for _, test := range tests {
		actual := Reverse(test.input)
		if actual != test.expected {
			t.Errorf("Reverse(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLowers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"TEST", "test"},
		{"test TEST", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\tTEST", "test\ttest"},
		{"test TEST test", "test test test"},
	}

	for _, test := range tests {
		actual := Lowers(test.input)
		if actual != test.expected {
			t.Errorf("Lowers(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'t', 't'},
		{'T', 't'},
		{'\n', '\n'},
		{'\t', '\t'},
		{' ', ' '},
	}

	for _, test := range tests {
		actual := Lower(test.input)
		if actual != test.expected {
			t.Errorf("Lower(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"TEST", "tEST"},
		{"test TEST", "test TEST"},
		{"test\ntest", "test\ntest"},
		{"test\tTEST", "test\tTEST"},
		{"test TEST test", "test TEST test"},
	}

	for _, test := range tests {
		actual := LowerFirst(test.input)
		if actual != test.expected {
			t.Errorf("LowerFirst(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUpper(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'t', 'T'},
		{'T', 'T'},
		{'\n', '\n'},
		{'\t', '\t'},
		{' ', ' '},
	}

	for _, test := range tests {
		actual := Upper(test.input)
		if actual != test.expected {
			t.Errorf("Upper(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUppers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "TEST"},
		{"TEST", "TEST"},
		{"test TEST", "TEST TEST"},
		{"test\ntest", "TEST\nTEST"},
		{"test\tTEST", "TEST\tTEST"},
		{"test TEST test", "TEST TEST TEST"},
	}

	for _, test := range tests {
		actual := Uppers(test.input)
		if actual != test.expected {
			t.Errorf("Uppers(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"TEST", "TEST"},
		{"test TEST", "Test TEST"},
		{"test\ntest", "Test\ntest"},
		{"test\tTEST", "Test\tTEST"},
		{"test TEST test", "Test TEST test"},
	}

	for _, test := range tests {
		actual := UpperFirst(test.input)
		if actual != test.expected {
			t.Errorf("UpperFirst(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected string
	}{
		{"test", 1, "test"},
		{"test", 2, "testtest"},
		{"test", 3, "testtesttest"},
		{"test", 4, "testtesttesttest"},
		{"test", 5, "testtesttesttesttest"},
	}

	for _, test := range tests {
		actual := Repeat(test.input, test.count)
		if actual != test.expected {
			t.Errorf("Repeat(%q, %q): expected %q, actual %q", test.input, test.count, test.expected, actual)
		}
	}
}

func TestTruncate(t *testing.T) {
	testsDefault := []struct {
		input    string
		length   int
		expected string
	}{
		{"test", 1, "t..."},
		{"test", 2, "te..."},
		{"test", 3, "tes..."},
		{"test", 4, "test"},
		{"test", 5, "test"},
	}

	for _, test := range testsDefault {
		actualDefault := Truncate(test.input, test.length)
		if actualDefault != test.expected {
			t.Errorf("Truncate(%q, %q): expected %q, actual %q", test.input, test.length, test.expected, actualDefault)
		}
	}

	testsOmission := []struct {
		input    string
		length   int
		omission string
		expected string
	}{
		{"test", 1, "...", "t..."},
		{"test", 2, "@@@", "te@@@"},
		{"test", 3, "$$$", "tes$$$"},
		{"test", 4, "!!!", "test"},
		{"test", 5, "---", "test"},
	}

	for _, test := range testsOmission {
		actualOmission := Truncate(test.input, test.length, test.omission)
		if actualOmission != test.expected {
			t.Errorf("Truncate(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.omission, test.expected, actualOmission)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		input    string
		old      string
		new      string
		expected string
	}{
		{"test", "e", "a", "tast"},
		{"test", "t", "a", "aesa"},
		{"test", "s", "a", "teat"},
		{"test", "x", "a", "test"},
		{"test", "", "a", "test"},
	}

	for _, test := range tests {
		actual := Replace(test.input, test.old, test.new)
		if actual != test.expected {
			t.Errorf("Replace(%q, %q, %q): expected %q, actual %q", test.input, test.old, test.new, test.expected, actual)
		} else {
			if len(actual) != len(test.input) {
				t.Errorf("Replace(%q, %q, %q): length expected %d, length actual %d", test.input, test.old, test.new, len(test.input), len(actual))
			}
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test"},
		{" test ", "test"},
		{" test test ", "test test"},
	}

	for _, test := range tests {
		actual := Trim(test.input)
		if actual != test.expected {
			t.Errorf("Trim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestLTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test "},
		{" test ", "test "},
		{" test test ", "test test "},
	}

	for _, test := range tests {
		actual := LTrim(test.input)
		if actual != test.expected {
			t.Errorf("LTrim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestRTrim(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", " test"},
		{"test ", "test"},
		{" test ", " test"},
		{" test test ", " test test"},
	}

	for _, test := range tests {
		actual := RTrim(test.input)
		if actual != test.expected {
			t.Errorf("RTrim(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestTrimBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" test", "test"},
		{"test ", "test"},
		{" test ", "test"},
		{" test test ", "test test"},
		{"\ntest", "test"},
		{"test\n", "test"},
		{"\ntest\n", "test"},
		{"\ntest\ntest\n", "test\ntest"},
		{"\ttest", "test"},
		{"test\t", "test"},
		{"\ttest\t", "test"},
		{"\ttest\ttest\t", "test\ttest"},
		{"\ttest\ntest\t", "test\ntest"},
		{"\ntest\ttest\n", "test\ttest"},
	}

	for _, test := range tests {
		actual := TrimBlank(test.input)
		if actual != test.expected {
			t.Errorf("TrimBlank(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestTrimPrefix(t *testing.T) {
	tests := []struct {
		input    string
		prefix   string
		expected string
	}{
		{"test", "te", "st"},
		{"test", "st", "test"},
		{"test", "x", "test"},
		{"test", "", "test"},
	}

	for _, test := range tests {
		actual := TrimPre(test.input, test.prefix)
		if actual != test.expected {
			t.Errorf("TrimPrefix(%q, %q): expected %q, actual %q", test.input, test.prefix, test.expected, actual)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		input    string
		suffix   string
		expected string
	}{
		{"test", "st", "te"},
		{"test", "te", "test"},
		{"test", "x", "test"},
		{"test", "", "test"},
	}

	for _, test := range tests {
		actual := TrimSuf(test.input, test.suffix)
		if actual != test.expected {
			t.Errorf("TrimSuffix(%q, %q): expected %q, actual %q", test.input, test.suffix, test.expected, actual)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input    []string
		sep      string
		expected string
	}{
		{[]string{"test", "test"}, " ", "test test"},
		{[]string{"test", "test"}, "\n", "test\ntest"},
		{[]string{"test", "test"}, "\t", "test\ttest"},
		{[]string{"test", "test"}, "x", "testxtest"},
		{[]string{"test", "test"}, "", "testtest"},
	}

	for _, test := range tests {
		actual := Join(test.input, test.sep)
		if actual != test.expected {
			t.Errorf("Join(%q, %q): expected %q, actual %q", test.input, test.sep, test.expected, actual)
		}
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"test", "test"}, "testtest"},
		{[]string{"test", "test", "test"}, "testtesttest"},
		{[]string{"test", "test", "test", "test"}, "testtesttesttest"},
		{[]string{"test", "test", "test", "test", "test"}, "testtesttesttesttest"},
	}

	for _, test := range tests {
		actual := Concat(test.input...)
		if actual != test.expected {
			t.Errorf("Concat(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestSubStr(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		end      int
		expected string
	}{
		{"test", 0, 1, "t"},
		{"test", 0, 2, "te"},
		{"test", 0, 3, "tes"},
		{"test", 0, 4, "test"},
		{"test", 0, 5, "test"},
		{"test", 1, 1, ""},
		{"test", 1, 2, "e"},
		{"test", 1, 3, "es"},
		{"test", 1, 4, "est"},
		{"test", 1, 5, "est"},
		{"test", 2, 2, ""},
		{"test", 2, 3, "s"},
		{"test", 2, 4, "st"},
		{"test", 2, 5, "st"},
		{"test", 3, 3, ""},
		{"test", 3, 4, "t"},
		{"test", 3, 5, "t"},
		{"test", 4, 4, ""},
		{"test", 4, 5, ""},
		{"test", 5, 5, ""},
	}

	for _, test := range tests {
		actual := SubStr(test.input, test.start, test.end)
		if actual != test.expected {
			t.Errorf("SubStr(%q, %q, %q): expected %q, actual %q", test.input, test.start, test.end, test.expected, actual)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		end      int
		expected string
	}{
		{"test", 0, 1, "t"},
		{"test", 0, 2, "te"},
		{"test", 0, 3, "tes"},
		{"test", 0, 4, "test"},
		{"test", 0, 5, "test"},
		{"test", 1, 1, ""},
		{"test", 1, 2, "e"},
		{"test", 1, 3, "es"},
		{"test", 1, 4, "est"},
		{"test", 1, 5, "est"},
		{"test", 2, 2, ""},
		{"test", 2, 3, "s"},
		{"test", 2, 4, "st"},
		{"test", 2, 5, "st"},
		{"test", 3, 3, ""},
		{"test", 3, 4, "t"},
		{"test", 3, 5, "t"},
		{"test", 4, 4, ""},
		{"test", 4, 5, ""},
		{"test", 5, 5, ""},
	}

	for _, test := range tests {
		actual := Slice(test.input, test.start, test.end)
		if actual != test.expected {
			t.Errorf("Slice(%q, %q, %q): expected %q, actual %q", test.input, test.start, test.end, test.expected, actual)
		}
	}
}

func TestSplice(t *testing.T) {
	tests := []struct {
		input       string
		start       int
		deleteCount int
		items       []string
		expected    string
	}{
		{"test", 0, 1, []string{"a"}, "aest"},
		{"test", 0, 2, []string{"a"}, "atest"},
		{"test", 1, 1, []string{"a"}, "tast"},
		// {"test", 1, 2, []string{"a"}, "taste"},
		{"test", 2, 1, []string{"a", "b"}, "teabt"},
		{"test", 3, 1, []string{"a"}, "tesa"},
		{"test", 4, 1, []string{"a"}, "testa"},
		{"test", 4, 2, []string{"a"}, "testa"},
		{"test", 4, 3, []string{"a"}, "testa"},
		{"test", 4, 4, []string{"a"}, "testa"},
	}

	for _, test := range tests {
		actual := Splice(test.input, test.start, test.deleteCount, test.items...)
		if actual != test.expected {
			t.Errorf("Splice(%q, %q, %q, %q): expected %q, actual %q", test.input, test.start, test.deleteCount, test.items, test.expected, actual)
		}
	}
}

func TestEscape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test test", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\ttest", "test\ttest"},
		{"test test test", "test test test"},
	}

	for _, test := range tests {
		actual := Esc(test.input)
		if actual != test.expected {
			t.Errorf("Escape(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestUnescape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test test", "test test"},
		{"test\ntest", "test\ntest"},
		{"test\ttest", "test\ttest"},
		{"test test test", "test test test"},
	}

	for _, test := range tests {
		actual := Unesc(test.input)
		if actual != test.expected {
			t.Errorf("Unescape(%q): expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestPadStart(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		pad      byte
		expected string
	}{
		{"test", 1, 'a', "test"},
		{"test", 2, 'a', "test"},
		{"test", 3, 'a', "test"},
		{"test", 4, 'a', "test"},
		{"test", 5, 'a', "atest"},
		{"test", 6, 'a', "aatest"},
	}

	for _, test := range tests {
		actual := LPad(test.input, test.length, test.pad)
		if actual != test.expected {
			t.Errorf("PadStart(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.pad, test.expected, actual)
		}
	}
}

func TestPadEnd(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		pad      byte
		expected string
	}{
		{"test", 1, 'a', "test"},
		{"test", 2, 'a', "test"},
		{"test", 3, 'a', "test"},
		{"test", 4, 'a', "test"},
		{"test", 5, 'a', "testa"},
		{"test", 6, 'a', "testaa"},
	}

	for _, test := range tests {
		actual := RPad(test.input, test.length, test.pad)
		if actual != test.expected {
			t.Errorf("PadEnd(%q, %q, %q): expected %q, actual %q", test.input, test.length, test.pad, test.expected, actual)
		}
	}
}

func TestCount(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Count 1",
			args: args{
				s:      "Hello World",
				substr: "",
			},
			want: 11,
		},
		{
			name: "Count 2",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 3,
		},
		{
			name: "Count 3",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		var got int

		if tt.args.substr == "" {
			got = Count(tt.args.s)
		} else {
			got = Count(tt.args.s, tt.args.substr)
		}

		if got != tt.want {
			t.Errorf("Count() = %v, want %v", got, tt.want)
		}
	}
}

func TestLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Lines 1",
			args: args{
				s: "Hello World",
			},
			want: 1,
		},
		{
			name: "Lines 2",
			args: args{
				s: "Hello\nWorld",
			},
			want: 2,
		},
		{
			name: "Lines 3",
			args: args{
				s: "Hello\nWorld\n",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := Lines(tt.args.s); got != tt.want {
			t.Errorf("Lines() = %v, want %v", got, tt.want)
		}
	}
}

func TestIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name      string
		args      args
		wantFirst int
		wantLast  int
	}{
		{
			name: "Index 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			wantFirst: 2,
			wantLast:  9,
		},
		{
			name: "Index 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			wantFirst: 2,
			wantLast:  2,
		},
		{
			name: "Index 3",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			wantFirst: -1,
			wantLast:  -1,
		},
		{
			name: "Index 4",
			args: args{
				s:      "Hello World",
				substr: " ",
			},
			wantFirst: 5,
			wantLast:  5,
		},
		{
			name: "Index 5",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			wantFirst: 6,
			wantLast:  6,
		},
	}

	for _, tt := range tests {
		gotFirst, gotLast := IdxSubStr(tt.args.s, tt.args.substr)

		if gotFirst != tt.wantFirst {
			t.Errorf("Index() gotFirst = %v, want %v", gotFirst, tt.wantFirst)
		}

		if gotLast != tt.wantLast {
			t.Errorf("Index() gotLast = %v, want %v", gotLast, tt.wantLast)
		}
	}
}

func TestFIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests1 := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FIndex 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 2,
		},
		{
			name: "FIndex 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 2,
		},
		{
			name: "FIndex 4",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			want: -1,
		},
		{
			name: "FIndex 5",
			args: args{
				s:      "Hello World",
				substr: " ",
			},
			want: 5,
		},
		{
			name: "FIndex 9",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			want: 6,
		},
	}

	for _, tt := range tests1 {
		if got := FirstIdx(tt.args.s, tt.args.substr); got != tt.want {
			t.Errorf("FIndex() = %v, want %v", got, tt.want)
		}
	}
}

func TestLIndex(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests2 := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LIndex 1",
			args: args{
				s:      "Hello World",
				substr: "l",
			},
			want: 9,
		},
		{
			name: "LIndex 2",
			args: args{
				s:      "Hello World",
				substr: "ll",
			},
			want: 2,
		},
		{
			name: "LIndex 4",
			args: args{
				s:      "Hello World",
				substr: "oo",
			},
			want: -1,
		},
		{
			name: "LIndex 9",
			args: args{
				s:      "Hello World",
				substr: "Wor",
			},
			want: 6,
		},
	}

	for _, tt := range tests2 {
		got := LastIdx(tt.args.s, tt.args.substr)

		if got != tt.want {
			t.Errorf("LIndex() = %v, want %v", got, tt.want)
		}
	}
}

func TestAt(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "At 1",
			args: args{
				s: "Hello World",
				i: 0,
			},
			want: "H",
		},
		{
			name: "At 2",
			args: args{
				s: "Hello World",
				i: 1,
			},
			want: "e",
		},
		{
			name: "At 3",
			args: args{
				s: "Hello World",
				i: 2,
			},
			want: "l",
		},
		{
			name: "At 4",
			args: args{
				s: "Hello World",
				i: 3,
			},
			want: "l",
		},
		{
			name: "At 5",
			args: args{
				s: "Hello World",
				i: 4,
			},
			want: "o",
		},
		{
			name: "At 6",
			args: args{
				s: "Hello World",
				i: 5,
			},
			want: " ",
		},
		{
			name: "At 7",
			args: args{
				s: "Hello World",
				i: 6,
			},
			want: "W",
		},
		{
			name: "At 8",
			args: args{
				s: "Hello World",
				i: 7,
			},
			want: "o",
		},
		{
			name: "At 9",
			args: args{
				s: "Hello World",
				i: 8,
			},
			want: "r",
		},
		{
			name: "At 10",
			args: args{
				s: "Hello World",
				i: 9,
			},
			want: "l",
		},
		{
			name: "At 11",
			args: args{
				s: "Hello World",
				i: 10,
			},
			want: "d",
		},
	}

	for _, tt := range tests {
		if got := At(tt.args.s, tt.args.i); got != tt.want {
			t.Errorf("At() = %v, want %v", got, tt.want)
		}
	}
}

func TestCodePointAt(t *testing.T) {
	type args struct {
		s string
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CodePointAt 1",
			args: args{
				s: "Hello World",
				i: 0,
			},
			want: 72,
		},
		{
			name: "CodePointAt 2",
			args: args{
				s: "Hello World",
				i: 1,
			},
			want: 101,
		},
		{
			name: "CodePointAt 3",
			args: args{
				s: "Hello World",
				i: 2,
			},
			want: 108,
		},
		{
			name: "CodePointAt 4",
			args: args{
				s: "Hello World",
				i: 3,
			},
			want: 108,
		},
		{
			name: "CodePointAt 5",
			args: args{
				s: "Hello World",
				i: 4,
			},
			want: 111,
		},
		{
			name: "CodePointAt 6",
			args: args{
				s: "Hello World",
				i: 5,
			},
			want: 32,
		},
		{
			name: "CodePointAt 7",
			args: args{
				s: "Hello World",
				i: 6,
			},
			want: 87,
		},
		{
			name: "CodePointAt 8",
			args: args{
				s: "Hello World",
				i: 7,
			},
			want: 111,
		},
		{
			name: "CodePointAt 9",
			args: args{
				s: "Hello World",
				i: 8,
			},
			want: 114,
		},
		{
			name: "CodePointAt 10",
			args: args{
				s: "Hello World",
				i: 9,
			},
			want: 108,
		},
		{
			name: "CodePointAt 11",
			args: args{
				s: "Hello World",
				i: 10,
			},
			want: 100,
		},
	}

	for _, tt := range tests {
		if got := CodePointAt(tt.args.s, tt.args.i); got != tt.want {
			t.Errorf("CodePointAt() = %v, want %v", got, tt.want)
		}
	}
}

func TestCodePoint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CodePoint 1",
			args: args{
				s: "Hello World",
			},
			want: []int{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
		},
	}

	for _, tt := range tests {
		got := CodePoint(tt.args.s)
		if len(got) != len(tt.want) {
			t.Errorf("CodePoint() = %v, want %v", len(got), len(tt.want))
		} else if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("CodePoint() = %v, want %v", got, tt.want)
		}
	}
}

func TestFromCodePointAt(t *testing.T) {
	type args struct {
		codePoint int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FromCodePointAt 1",
			args: args{
				codePoint: 72,
			},
			want: "H",
		},
		{
			name: "FromCodePointAt 2",
			args: args{
				codePoint: 101,
			},
			want: "e",
		},
		{
			name: "FromCodePointAt 3",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 4",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 5",
			args: args{
				codePoint: 111,
			},
			want: "o",
		},
		{
			name: "FromCodePointAt 6",
			args: args{
				codePoint: 32,
			},
			want: " ",
		},
		{
			name: "FromCodePointAt 7",
			args: args{
				codePoint: 87,
			},
			want: "W",
		},
		{
			name: "FromCodePointAt 8",
			args: args{
				codePoint: 111,
			},
			want: "o",
		},
		{
			name: "FromCodePointAt 9",
			args: args{
				codePoint: 114,
			},
			want: "r",
		},
		{
			name: "FromCodePointAt 10",
			args: args{
				codePoint: 108,
			},
			want: "l",
		},
		{
			name: "FromCodePointAt 11",
			args: args{
				codePoint: 100,
			},
			want: "d",
		},
	}

	for _, tt := range tests {
		if got := FromCodePointAt(tt.args.codePoint); got != tt.want {
			t.Errorf("FromCodePointAt() = %v, want %v", got, tt.want)
		}
	}
}

func TestFromCodePoint(t *testing.T) {
	type args struct {
		codePoint []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FromCodePoint 1",
			args: args{
				codePoint: []int{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100},
			},
			want: "Hello World",
		},
	}

	for _, tt := range tests {
		if got := FromCodePoint(tt.args.codePoint...); got != tt.want {
			t.Errorf("FromCodePoint() = %v, want %v", got, tt.want)
		}
	}
}

func TestMatchStr(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected int
	}{
		{"hello world", "world", 6},
		{"abababc", "ababc", 2},
		{"abcdabcabcd", "abcabcd", 4},
		{"aaaaa", "bba", -1},
		{"mississippi", "issi", 1},
		{"", "", 0},
		{"a", "", 0},
		{"", "a", -1},
	}

	for _, test := range tests {
		result := MatchStr(test.text, test.pattern)
		if result != test.expected {
			t.Errorf("MatchStr(%q, %q) = %d; want %d", test.text, test.pattern, result, test.expected)
		}
	}
}
