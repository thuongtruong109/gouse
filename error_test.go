package gouse

import (
	"fmt"
	"testing"
)

func TestDetectErr(t *testing.T) {
	err := fmt.Errorf("this is an error")
	result := DetectErr(err)
	if result != "this is an error" {
		t.Errorf("expected 'this is an error', got '%s'", result)
	}

	str := "this is a string"
	result = DetectErr(str)
	if result != "this is a string" {
		t.Errorf("expected 'this is a string', got '%s'", result)
	}

	num := 123
	result = DetectErr(num)
	if result != "123" {
		t.Errorf("expected '123', got '%s'", result)
	}

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	result = DetectErr(p)
	if result != "{Alice 30}" {
		t.Errorf("expected '{Alice 30}', got '%s'", result)
	}

	var nilVal any = nil
	result = DetectErr(nilVal)
	if result != "<nil>" {
		t.Errorf("expected '<nil>', got '%s'", result)
	}
}
