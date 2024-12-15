package gouse

import (
	"testing"
	"time"
)

func TestISODate(t *testing.T) {
	now := time.Now()
	if ISODate() != now.Format("2006-01-02T15:04:05.999Z") {
		t.Error("ISO() should return today's date in ISO format")
	}
}

func TestNormalDate(t *testing.T) {
	if NormalDate() != time.Now().Format("01/02/2006") {
		t.Error("ShortNormal() should return today's date in ShortNormal format")
	}
}

func TestReverseDate(t *testing.T) {
	if ReverseDate() != time.Now().Format("2006/01/02") {
		t.Error("ShortReverse() should return today's date in ShortReverse format")
	}
}

func TestDashDate(t *testing.T) {
	if DashDate() != time.Now().Format("2006-01-02") {
		t.Error("ShortDash() should return today's date in ShortDash format")
	}
}

func TestDotDash(t *testing.T) {
	if DotDate() != time.Now().Format("2006.01.02") {
		t.Error("ShortDot() should return today's date in ShortDot format")
	}
}

func TestUnderlineDate(t *testing.T) {
	if UnderlineDate() != time.Now().Format("2006_01_02") {
		t.Error("ShortUnderscore() should return today's date in ShortUnderscore format")
	}
}

func TestSpaceDate(t *testing.T) {
	if SpaceDate() != time.Now().Format("2006 01 02") {
		t.Error("ShortSpace() should return today's date in ShortSpace format")
	}
}

func TestMonthDate(t *testing.T) {
	if MonthDate() != time.Now().Format("01/2006") {
		t.Error("ShortMonth() should return today's date in ShortMonth format")
	}
}

func TestLongDate(t *testing.T) {
	if LongDate() != time.Now().Format("Monday, January 2, 2006") {
		t.Error("Long() should return today's date in Long format")
	}
}

func TestUTCDate(t *testing.T) {
	if UTCDate() != time.Now().Format("Jan 2, 2006 at 3:04pm (MST)") {
		t.Error("UTC() should return today's date in UTC format")
	}
}
