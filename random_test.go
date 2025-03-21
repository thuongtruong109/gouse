package gouse

import (
	"fmt"
	"reflect"
	"testing"
	"time"
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

	for b.Loop() {
		RandNum(min, max)
	}
}

func TestRandID(t *testing.T) {
	t.Run("NonEmptyString", func(t *testing.T) {
		id := RandID()
		if len(id) == 0 {
			t.Errorf("Expected non-empty string, got an empty string")
		}
	})

	t.Run("StringRepresentationOfNumber", func(t *testing.T) {
		id := RandID()
		_, err := fmt.Sscanf(id, "%d", new(int64))
		if err != nil {
			t.Errorf("Expected a valid integer string, got error: %v", err)
		}
	})

	t.Run("UniqueIDs", func(t *testing.T) {
		id1 := RandID()
		time.Sleep(1 * time.Millisecond)
		id2 := RandID()

		if id1 == id2 {
			t.Errorf("Expected different IDs for consecutive calls, got %s and %s", id1, id2)
		}
	})
}

func BenchmarkRandID(b *testing.B) {
	for b.Loop() {
		RandID()
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

func BenchmarkRandStr(b *testing.B) {
	expectedLength := 10

	for b.Loop() {
		RandStr(expectedLength)
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

func BenchmarkRandDigit(b *testing.B) {
	expectedLength := 6

	for b.Loop() {
		RandDigit(expectedLength)
	}
}
