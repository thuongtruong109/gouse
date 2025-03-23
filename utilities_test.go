package gouse

import (
	"fmt"
	"testing"
)

func TestDetectError(t *testing.T) {
	err := fmt.Errorf("this is an error")
	result := DetectError(err)
	if result != "this is an error" {
		t.Errorf("expected 'this is an error', got '%s'", result)
	}

	str := "this is a string"
	result = DetectError(str)
	if result != "this is a string" {
		t.Errorf("expected 'this is a string', got '%s'", result)
	}

	num := 123
	result = DetectError(num)
	if result != "123" {
		t.Errorf("expected '123', got '%s'", result)
	}

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	result = DetectError(p)
	if result != "{Alice 30}" {
		t.Errorf("expected '{Alice 30}', got '%s'", result)
	}

	var nilVal any = nil
	result = DetectError(nilVal)
	if result != "<nil>" {
		t.Errorf("expected '<nil>', got '%s'", result)
	}
}
