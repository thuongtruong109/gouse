package gouse

import (
	"testing"
	"time"
)

func TestGetLocal(t *testing.T) {
	c := NewLCache()
	c.SetLCache("test", "test")
	if result, err := c.GetLCache("test"); result != "test" {
		t.Error("GetLocal error", err)
	}
}

func TestSetLocal(t *testing.T) {
	c := NewLCache()
	c.SetLCache("test", "test")
	if result, err := c.GetLCache("test"); result != "test" {
		t.Error("SetLocal error", err)
	}
}

func TestDelLocal(t *testing.T) {
	c := NewLCache()
	c.SetLCache("test", "test")
	c.DelLCache("test")
	if result, err := c.GetLCache("test"); result != "" {
		t.Error("DelLocal error", err)
	}
}

func TestFlushLocal(t *testing.T) {
	c := NewLCache()
	c.SetLCache("test", "test")
	c.FlushLCache()
	if result, err := c.GetLCache("test"); result != "" {
		t.Error("FlushLocal error", err)
	}
}

func TestAllLCache(t *testing.T) {
	c := NewLCache()
	c.SetLCache("test", "test")
	if c.AllLCache()["test"] != "test" {
		t.Error("AllLocal error")
	}
}

func TestGetTCache(t *testing.T) {
	tmp := NewTCache(1 * time.Hour)
	tmp.SetTCache("test", "test", 0)
	if tmp.GetTCache("test") != "test" {
		t.Error("TestGetTmp failed")
	}
}

func TestSetTCache(t *testing.T) {
	tmp := NewTCache(1 * time.Hour)
	tmp.SetTCache("test", "test", 0)
	if tmp.GetTCache("test") != "test" {
		t.Error("TestSetTmp failed")
	}
}

func TestDelTCache(t *testing.T) {
	tmp := NewTCache(1 * time.Hour)
	tmp.SetTCache("test", "test", 0)
	tmp.DelTCache("test")
	if tmp.GetTCache("test") != nil {
		t.Error("TestDelTmp failed")
	}
}

func TestFlushTmp(t *testing.T) {
	tmp := NewTCache(1 * time.Hour)
	tmp.SetTCache("test", "test", 0)
	tmp.FlushTCache()
	if tmp.GetTCache("test") != nil {
		t.Error("TestFlushTmp failed")
	}
}

func TestAllTCache(t *testing.T) {
	tmp := NewTCache(1 * time.Hour)
	tmp.SetTCache("test", "test", 0)
	if tmp.AllTCache()["test"] != "test" {
		t.Error("TestAllTmp failed")
	}
}
