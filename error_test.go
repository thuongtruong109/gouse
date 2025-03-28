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

func TestErr(t *testing.T) {
	err := fmt.Errorf("this is an error")
	result := Err(err)
	if result == nil || result.Error() != "this is an error" {
		t.Errorf("expected error 'this is an error', got %v", result)
	}

	str := "this is a string"
	result = Err(str)
	if result == nil || result.Error() != "this is a string" {
		t.Errorf("expected error 'this is a string', got %v", result)
	}

	num := 123
	result = Err(num)
	if result == nil || result.Error() != "123" {
		t.Errorf("expected error '123', got %v", result)
	}

	type customStruct struct {
		field string
	}
	cs := customStruct{field: "test"}
	result = Err(cs)
	if result == nil || result.Error() != "{test}" {
		t.Errorf("expected error '{test}', got %v", result)
	}
}

func TestErrMsg(t *testing.T) {
	err := fmt.Errorf("an error occurred")
	message := "Test error"
	result := ErrMsg(message, err)
	expected := "Test error: an error occurred"
	if result.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, result.Error())
	}

	// Test Case 2: nil error
	result = ErrMsg(message, nil)
	expected = "Test error: <nil>"
	if result.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, result.Error())
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but did not panic")
		}
	}()

	err := fmt.Errorf("this is an error")
	Panic(fmt.Errorf("Panic message"), err)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but did not panic")
		}
	}()

	Panic(fmt.Errorf("Panic message"))
}
