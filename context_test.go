package gouse

import (
	"context"
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
