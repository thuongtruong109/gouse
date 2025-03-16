package gouse

import (
	"testing"
	"time"
)

func TestGetLocal(t *testing.T) {
	c := NewLocalCache()
	c.SetLocalCache("test", "test")
	if result, err := c.GetLocalCache("test"); result != "test" {
		t.Error("GetLocal error", err)
	}
}

func TestSetLocal(t *testing.T) {
	c := NewLocalCache()
	c.SetLocalCache("test", "test")
	if result, err := c.GetLocalCache("test"); result != "test" {
		t.Error("SetLocal error", err)
	}
}

func TestDelLocal(t *testing.T) {
	c := NewLocalCache()
	c.SetLocalCache("test", "test")
	c.DelLocalCache("test")
	if result, err := c.GetLocalCache("test"); result != "" {
		t.Error("DelLocal error", err)
	}
}

func TestFlushLocal(t *testing.T) {
	c := NewLocalCache()
	c.SetLocalCache("test", "test")
	c.FlushLocalCache()
	if result, err := c.GetLocalCache("test"); result != "" {
		t.Error("FlushLocal error", err)
	}
}

func TestAllLocalCache(t *testing.T) {
	c := NewLocalCache()
	c.SetLocalCache("test", "test")
	if c.AllLocalCache()["test"] != "test" {
		t.Error("AllLocal error")
	}
}

func TestGetTmpCache(t *testing.T) {
	tmp := NewTmpCache(1 * time.Hour)
	tmp.SetTmpCache("test", "test", 0)
	if tmp.GetTmpCache("test") != "test" {
		t.Error("TestGetTmp failed")
	}
}

func TestSetTmpCache(t *testing.T) {
	tmp := NewTmpCache(1 * time.Hour)
	tmp.SetTmpCache("test", "test", 0)
	if tmp.GetTmpCache("test") != "test" {
		t.Error("TestSetTmp failed")
	}
}

func TestDelTmpCache(t *testing.T) {
	tmp := NewTmpCache(1 * time.Hour)
	tmp.SetTmpCache("test", "test", 0)
	tmp.DelTmpCache("test")
	if tmp.GetTmpCache("test") != nil {
		t.Error("TestDelTmp failed")
	}
}

func TestFlushTmp(t *testing.T) {
	tmp := NewTmpCache(1 * time.Hour)
	tmp.SetTmpCache("test", "test", 0)
	tmp.FlushTmpCache()
	if tmp.GetTmpCache("test") != nil {
		t.Error("TestFlushTmp failed")
	}
}

func TestAllTmpCache(t *testing.T) {
	tmp := NewTmpCache(1 * time.Hour)
	tmp.SetTmpCache("test", "test", 0)
	if tmp.AllTmpCache()["test"] != "test" {
		t.Error("TestAllTmp failed")
	}
}
