package gouse

import (
	"context"
	"fmt"
	"testing"
)

type contextKey string

func TestCtx(t *testing.T) {
	key := contextKey("key")
	newCtx := context.WithValue(context.Background(), key, "value")

	SetCtx(newCtx)

	ctx := GetCtx()

	if ctx != newCtx {
		t.Errorf("expected context %v, got %v", newCtx, ctx)
	}

	if value := ctx.Value(key); value != "value" {
		t.Errorf("expected value 'value', got %v", value)
	}
}

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

	var nilVal interface{} = nil
	result = DetectError(nilVal)
	if result != "<nil>" {
		t.Errorf("expected '<nil>', got '%s'", result)
	}
}
