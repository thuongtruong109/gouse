package gouse

import (
	"reflect"
	"testing"
)

func TestRandNum(t *testing.T) {
	min := 1
	max := 10

	result := RandNum(min, max)

	if result < min || result > max {
		t.Errorf("Random() = %d; want >= %d and <= %d", result, min, max)
	}
}

func BenchmarkRandNum(b *testing.B) {
	min := 1
	max := 10

	for i := 0; i < b.N; i++ {
		RandNum(min, max)
	}
}

func TestRandStr(t *testing.T) {
	expectedLength := 10
	got := RandStr(expectedLength)

	if len(got) != expectedLength {
		t.Errorf("RandomStr() = %v, want %v", len(got), expectedLength)
	}

	if reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("RandomStr() = %v, want %v", reflect.TypeOf(got).Kind(), reflect.String)
	}
}

func TestRandDigit(t *testing.T) {
	expectedLength := 6
	got := RandDigit(expectedLength)

	if len(got) != expectedLength {
		t.Errorf("RandomNum() = %v, want %v", len(got), expectedLength)
	}

	if reflect.TypeOf(got).Kind() != reflect.String {
		t.Errorf("RandomNum() = %v, want %v", reflect.TypeOf(got).Kind(), reflect.String)
	}
}
